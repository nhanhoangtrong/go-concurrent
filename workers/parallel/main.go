package parallel

import (
	"context"
	"deals-sync/entities"
	"deals-sync/pkg/config"
	"deals-sync/pkg/db"
	myKafka "deals-sync/pkg/kafka"
	"encoding/json"
	"fmt"
	"sync"
)

type ParallelWorker struct{}

func NewParallelWorker() *ParallelWorker {
	return &ParallelWorker{}
}

func (w *ParallelWorker) Run(config *config.Config, msgCount int) {

	// 1st, initialize db
	mongo, err := db.InitMongoDB(config.MongoUri, config.DbName)
	if err != nil {
		panic(err)
	}
	col := mongo.Collection("deals")

	// initialize custom context
	ctx, cancel := context.WithCancel(context.Background())

	// create reader first
	reader := myKafka.CreateKafkaReader(config.KafkaUrl, config.KafkaTopic, config.WorkerGroup)

	// then, start reading in a goroutine
	errc := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < msgCount; i++ {
			// Read message
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				errc <- err
				return
			}

			// For each message, create a go routine
			wg.Add(1)
			go func() {
				// Decode and save
				defer wg.Done()
				var deal entities.DealEntity
				if err := json.Unmarshal(m.Value, &deal); err != nil {
					errc <- err
					return
				}
				_, err = col.InsertOne(ctx, deal)
				if err != nil {
					errc <- err
					return
				}
			}()
		}
	}()

	// Start waiting until all goroutines are done
	go func() {
		defer cancel()
		wg.Wait()
	}()

	// Wait for error or done
	select {
	case err := <-errc:
		fmt.Println("Run error", err)
	case <-ctx.Done():
		return
	}
}

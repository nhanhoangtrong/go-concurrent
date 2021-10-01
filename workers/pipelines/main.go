package pipelines

import (
	"context"
	"deals-sync/entities"
	"deals-sync/pkg/config"
	"deals-sync/pkg/db"
	myKafka "deals-sync/pkg/kafka"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
)

type PipelinesWorker struct{}

func NewPipelinesWorker() *PipelinesWorker {
	return &PipelinesWorker{}
}

func (w *PipelinesWorker) Run(config *config.Config, msgCount int) {

	// 1st, initialize db
	mongo, err := db.InitMongoDB(config.MongoUri, config.DbName)
	if err != nil {
		panic(err)
	}
	col := mongo.Collection("deals")

	// initialize custom context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create reader first
	reader := myKafka.CreateKafkaReader(config.KafkaUrl, config.KafkaTopic, config.WorkerGroup)

	// then, start read in a goroutine
	errc := make(chan error)
	msgChan := make(chan *kafka.Message)
	go func() {
		defer close(msgChan)
		for i := 0; i < msgCount; i++ {
			// Read message
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				errc <- err
				return
			}
			select {
			case msgChan <- &m:
				continue
			case <-ctx.Done():
				return
			}
		}
	}()

	// then decode & save deal in goroutines
	var wg sync.WaitGroup
	for i := 0; i < config.MaxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case msg := <-msgChan:
					if msg != nil {
						var deal entities.DealEntity
						if err := json.Unmarshal(msg.Value, &deal); err != nil {
							errc <- err
							return
						}
						_, err = col.InsertOne(ctx, deal)
						if err != nil {
							errc <- err
							return
						}
					} else {
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		// Done save
		cancel()
	}()

	select {
	case err := <-errc:
		fmt.Println("Run error", err)
	case <-ctx.Done():
		return
	}
}

package single_pipeline

import (
	"context"
	"deals-sync/entities"
	"deals-sync/pkg/config"
	"deals-sync/pkg/db"
	myKafka "deals-sync/pkg/kafka"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type SinglePipelineWorker struct{}

func NewSinglePipelineWorker() *SinglePipelineWorker {
	return &SinglePipelineWorker{}
}

func (w *SinglePipelineWorker) Run(config *config.Config, msgCount int) {

	// 1st, initialize db
	mongo, err := db.InitMongoDB(config.MongoUri, config.DbName)
	if err != nil {
		panic(err)
	}
	col := mongo.Collection("deals")

	// initialize custom context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// on serial, create reader first
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

	// and decode deal in a goroutine
	dealChan := make(chan *entities.DealEntity)
	go func() {
		defer close(dealChan)
		for {
			select {
			case msg := <-msgChan:
				if msg != nil {
					var deal entities.DealEntity
					if err := json.Unmarshal(msg.Value, &deal); err != nil {
						errc <- err
					}
					dealChan <- &deal
				} else {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// Save deal in another goroutine
	go func() {
		defer cancel()
		for deal := range dealChan {
			_, err = col.InsertOne(ctx, deal)
			if err != nil {
				errc <- err
				return
			}
		}
	}()

	// In main thread, wait for either error or done context
	select {
	case err := <-errc:
		fmt.Println("Run error", err)
	case <-ctx.Done():
		return
	}
}

package serial

import (
	"context"
	"deals-sync/entities"
	"deals-sync/pkg/config"
	"deals-sync/pkg/db"
	"deals-sync/pkg/kafka"
	"encoding/json"
)

type SerialWorker struct{}

func NewSerialWorker() *SerialWorker {
	return &SerialWorker{}
}

func (w *SerialWorker) Run(config *config.Config, msgCount int) {

	// 1st, initialize db
	mongo, err := db.InitMongoDB(config.MongoUri, config.DbName)
	if err != nil {
		panic(err)
	}
	col := mongo.Collection("deals")

	// on serial, create reader first
	reader := kafka.CreateKafkaReader(config.KafkaUrl, config.KafkaTopic, config.WorkerGroup)

	for i := 0; i < msgCount; i++ {
		// Read message
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		// Decode
		var deal entities.DealEntity
		if err = json.Unmarshal(m.Value, &deal); err != nil {
			panic(err)
		}

		// Save
		_, err = col.InsertOne(context.Background(), deal)
		if err != nil {
			panic(err)
		}
	}
}

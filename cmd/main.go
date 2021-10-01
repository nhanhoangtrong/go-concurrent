package main

import (
	"context"
	"deals-sync/interfaces"
	"deals-sync/pkg/config"
	"deals-sync/pkg/db"
	"deals-sync/pkg/utils"
	"deals-sync/workers/feed"
	"deals-sync/workers/parallel"
	"deals-sync/workers/pipelines"
	"deals-sync/workers/serial"
	"deals-sync/workers/single_pipeline"
	"deals-sync/workers/with_buffer"
	"errors"
	"log"
	"strings"
	"time"
)

func main() {
	// Number of run
	count := 1_000
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Drop deals collection
	mongo, err := db.InitMongoDB(config.MongoUri, config.DbName)
	if err != nil {
		panic(err)
	}
	err = mongo.Collection("deals").Drop(context.Background())
	if err != nil {
		panic(err)
	}
	err = mongo.Client().Disconnect(context.Background())
	if err != nil {
		panic(err)
	}

	var worker interfaces.Worker
	switch {
	case strings.HasPrefix(config.WorkerGroup, "feed"):
		worker = feed.NewFeedWorker()
	case strings.HasPrefix(config.WorkerGroup, "serial"):
		worker = serial.NewSerialWorker()
	case strings.HasPrefix(config.WorkerGroup, "parallel"):
		worker = parallel.NewParallelWorker()
	case strings.HasPrefix(config.WorkerGroup, "single_pipeline"):
		worker = single_pipeline.NewSinglePipelineWorker()
	case strings.HasPrefix(config.WorkerGroup, "pipelines"):
		worker = pipelines.NewPipelinesWorker()
	case strings.HasPrefix(config.WorkerGroup, "with_buffer"):
		worker = with_buffer.NewWithBufferWorker()
	default:
		panic(errors.New("worker not supported"))
	}

	defer log.Printf("Done " + config.WorkerGroup)
	// Run worker and measure time
	defer utils.MeasureTime(time.Now(), config.WorkerGroup)
	worker.Run(config, count)
}

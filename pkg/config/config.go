package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUri     string
	DbName       string
	KafkaUrl     string
	KafkaTopic   string
	WorkerGroup  string
	MaxWorkers   int
	BufferLength int
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	maxWorkers, err := strconv.Atoi(os.Getenv("MAX_WORKERS"))
	if err != nil {
		return nil, err
	}
	bufferLength, err := strconv.Atoi(os.Getenv("BUFFER_LENGTH"))
	if err != nil {
		return nil, err
	}

	return &Config{
		MongoUri:     os.Getenv("MONGO_URI"),
		DbName:       os.Getenv("DBNAME"),
		KafkaUrl:     os.Getenv("KAFKA_URL"),
		KafkaTopic:   os.Getenv("KAFKA_TOPIC"),
		WorkerGroup:  os.Getenv("WORKER_GROUP"),
		MaxWorkers:   maxWorkers,
		BufferLength: bufferLength,
	}, nil
}

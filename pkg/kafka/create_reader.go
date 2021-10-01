package kafka

import (
	"strings"

	"github.com/segmentio/kafka-go"
)

func CreateKafkaReader(kafkaURLs, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURLs, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e2, // 1KB
		MaxBytes: 10e6, // 10MB
	})
}

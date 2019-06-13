package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	// Create Producer properties
	kafkaWriterConfig := kafka.WriterConfig{
		Brokers: []string{"localhost:9093"},
		Topic:   "first_go_topic",
	}
	// Create Producer
	kafkaWriter := kafka.NewWriter(kafkaWriterConfig)

	// Create a message and send it to Kafka
	kafkaMessage := kafka.Message{
		Value: []byte("Hello Kafka, Gopher here!"),
	}

	if err := kafkaWriter.WriteMessages(ctx, kafkaMessage); err != nil {
		panic(err)
	}

}

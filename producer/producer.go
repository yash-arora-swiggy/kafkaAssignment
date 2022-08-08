package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "example-123",
		AllowAutoTopicCreation: true,
		Completion: func(messages []kafka.Message, err error) {
			if err != nil {
				fmt.Print("Failed to publish message")
			}
		},
	}
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
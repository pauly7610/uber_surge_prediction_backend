package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConsumer(brokers string, groupID string, topics []string) {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  brokers,
		"group.id":           groupID,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	consumer.SubscribeTopics(topics, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			processMessage(msg)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func processMessage(msg *kafka.Message) {
	log.Printf("Processing message: %s", string(msg.Value))
}

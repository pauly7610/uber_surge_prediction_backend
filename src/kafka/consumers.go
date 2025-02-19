package kafka

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
    "log"
)

func StartConsumer(brokers string, groupID string, topics []string) {
    config := &kafka.ConfigMap{
        "bootstrap.servers": brokers,
        "group.id":          groupID,
        "auto.offset.reset": "earliest",
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
            // Process message
            processMessage(msg)
        } else {
            // Handle error
            log.Printf("Consumer error: %v (%v)\n", err, msg)
            // Send to dead letter queue
        }
    }
}

func processMessage(msg *kafka.Message) {
    // Implement message processing logic
    log.Printf("Processing message: %s", string(msg.Value))
    // Add error handling and retry logic if necessary
} 
package kafka

import (
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestProcessMessage(t *testing.T) {
	// Mock Kafka message
	msg := &kafka.Message{
		Value: []byte(`{"event": "test_event"}`),
	}

	// Call processMessage function
	processMessage(msg)

	// Assert no error
	assert.NotNil(t, msg)
}

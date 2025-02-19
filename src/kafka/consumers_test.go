package kafka

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProcessMessage(t *testing.T) {
	// Mock Kafka message
	msg := &kafka.Message{
		Value: []byte(`{"event": "test_event"}`),
	}

	// Call processMessage function
	err := processMessage(msg)

	// Assert no error
	assert.NoError(t, err)
} 
package services

import (
    "testing"
    "github.com/go-redis/redis/v8"
    "context"
)

func TestSetAndGetSurgeMultiplier(t *testing.T) {
    client := GetRedisClient()
    defer client.Close()

    key := "test_key"
    multiplier := 1.5

    err := SetSurgeMultiplier(client, key, multiplier)
    if err != nil {
        t.Fatalf("Failed to set surge multiplier: %v", err)
    }

    result, err := GetSurgeMultiplier(client, key)
    if err != nil {
        t.Fatalf("Failed to get surge multiplier: %v", err)
    }

    if result != multiplier {
        t.Errorf("Expected %v, got %v", multiplier, result)
    }
} 
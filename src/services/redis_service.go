package services

import (
    "github.com/go-redis/redis/v8"
    "context"
    "time"
    "strconv"
)

var ctx = context.Background()

func GetRedisClient() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
}

func SetSurgeMultiplier(client *redis.Client, key string, multiplier float64) error {
    return client.Set(ctx, key, multiplier, 5*time.Minute).Err()
}

func GetSurgeMultiplier(client *redis.Client, key string) (float64, error) {
    val, err := client.Get(ctx, key).Result()
    if err != nil {
        return 0, err
    }
    return strconv.ParseFloat(val, 64)
} 
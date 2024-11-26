package infrastructure

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	infrastructureconfiguration "panel-subs/infrastructure/configuration"

	"github.com/redis/go-redis/extra/redisotel/v9"
)

var RDS *redis.Client

func InitializeChannelEngine() {
	// Create a standard Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     infrastructureconfiguration.RedisAddr,
		Password: infrastructureconfiguration.RedisPass,
		DB:       infrastructureconfiguration.RedisDB,
	})

	if err := redisotel.InstrumentTracing(redisClient); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(redisClient); err != nil {
		panic(err)
	}

	RDS = redisClient

	fmt.Println("=== Load Cache, Pub/Sub Redis ===")
	fmt.Println("Channel Engine Is Running: ", RDS)
	fmt.Println("=== Load Cache, Pub/Sub Redis ===")
}

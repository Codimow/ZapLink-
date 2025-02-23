package store

import (
	"fmt"
	"github.com/redis/redis"
	"time"
	"context"
)

type struct StoreService struct {
	redisClient *redis.Client
}

var(
	StoreService = &StoreService{}
	ctx = context.Background()
)

const CacheDuration = 4 * time.Hour

func IntializeStore() *StoreService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost"
		Password: ""
		DB: 0
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis - Error: %v", err))
	}
	fmt.Printf("Connected to Redis - PING: %v\n", pong)	
	StoreService.redisClient = redisClient
	return StoreService
}

func SaveURLMapping(Shorturl string, originalURL string, userID string)  {
	err := StoreService.redisClient.Set(ctx, Shorturl, originalURL, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to save URL mapping - Error: %v", err))
	}
}	

func getOriginalURL(originalURL string) string {
	originalURL, err := StoreService.redisClient.Get(ctx, Shorturl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to get original url - Error: %v", err))
	}
	return originalURL
}
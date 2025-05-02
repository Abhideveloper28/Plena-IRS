package pubsub

import (
	"context"
	"encoding/json"
	"log"

	"access-key-service/internal/model"

	"github.com/go-redis/redis/v8"
)

type RedisPublisher struct {
	client *redis.Client
}

func NewRedisPublisher(client *redis.Client) *RedisPublisher {
	return &RedisPublisher{client: client}
}

func (p *RedisPublisher) PublishKeyUpdate(key model.AccessKey) {
	ctx := context.Background()
	data, _ := json.Marshal(key)
	if err := p.client.Publish(ctx, "key_updates", data).Err(); err != nil {
		log.Println("Redis publish error:", err)
	}
}

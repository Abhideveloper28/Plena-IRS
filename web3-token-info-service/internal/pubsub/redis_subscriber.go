package pubsub

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func StartSubscriber() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	sub := rdb.Subscribe(ctx, "key_event_channel")
	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("Error receiving message:", err)
			continue
		}
		fmt.Println("Received message:", msg.Payload)
	}
}

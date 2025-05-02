package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"access-key-service/internal/model"
)

var ctx = context.Background()

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("access_keys.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	db.AutoMigrate(&model.AccessKey{})
	return db
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("failed to connect redis: %v", err)
	}
	return rdb
}

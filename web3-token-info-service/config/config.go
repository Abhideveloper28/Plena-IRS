package config

import (
	"fmt"
)

var RedisAddress = "localhost:6379"

func GetRedisAddress() string {
	return RedisAddress
}

func InitConfig() {
	fmt.Println("Configuration loaded")
}

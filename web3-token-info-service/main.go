package main

import (
	"web3-token-info-service/config"
	"web3-token-info-service/internal/handler"
	"web3-token-info-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitConfig()

	go service.StartSubscriber()

	r := gin.Default()

	r.GET("/api/token/:key/:tokenName", handler.GetTokenInfo)

	r.Run(":8081")
}

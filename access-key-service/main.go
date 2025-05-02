package main

import (
	"access-key-service/config"
	"access-key-service/internal/handler"
	"access-key-service/internal/pubsub"
	"access-key-service/internal/repository"
	"access-key-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	redisClient := config.InitRedis()

	repo := repository.NewAccessKeyRepository(db)
	publisher := pubsub.NewRedisPublisher(redisClient)
	svc := service.NewAccessKeyService(repo, publisher)
	h := handler.NewHandler(svc)

	r := gin.Default()

	api := r.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.POST("/key", h.CreateKey)
			admin.GET("/keys", h.ListKeys)
			admin.PUT("/key/:key", h.UpdateKey)
			admin.DELETE("/key/:key", h.DeleteKey)
		}
		user := api.Group("/user")
		{
			user.GET("/plan/:key", h.GetPlanByKey)
			user.DELETE("/disable/:key", h.DisableKey)
		}
	}

	r.Run(":8080")
}

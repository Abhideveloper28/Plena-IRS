package handler

import (
	"net/http"
	"web3-token-info-service/internal/service"

	"github.com/gin-gonic/gin"
)

func GetTokenInfo(c *gin.Context) {
	key := c.Param("key")
	tokenName := c.Param("tokenName")

	info, err := service.FetchTokenInfo(key, tokenName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, info)
}

package utils

import "github.com/gin-gonic/gin"

func JSONResponse(c *gin.Context, statusCode int, response map[string]interface{}) {
	c.JSON(statusCode, response)
}

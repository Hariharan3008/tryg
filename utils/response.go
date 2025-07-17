package utils

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, gin.H{
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}
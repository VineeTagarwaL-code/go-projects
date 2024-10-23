package handlers

import (
	"go-pubsub-redis/libraries/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {

	redis := redis.GetInstance()
	redis.Publish("1234", "message")
	c.JSON(http.StatusOK, gin.H{
		"message": "Message Published",
	})
}

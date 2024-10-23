package handlers

import (
	"go-pubsub-redis/libraries/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishHandler(c *gin.Context) {
	event := c.Query("event")

	if event == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Event is required",
		})
		return
	}

	message := c.Query("message")

	if message == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Message is required",
		})
		return
	}

	redis := redis.GetInstance()
	redis.Publish(event, message)
	c.JSON(http.StatusOK, gin.H{
		"status":  "Message published successfully",
		"event":   event,
		"message": message,
	})

}

package routes

import (
	"go-pubsub-redis/handlers"
	"go-pubsub-redis/websocket"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/user", handlers.UserHandler)
	r.GET("/publish", handlers.PublishHandler)
	r.GET("/ws", websocket.WebSocketHandler)
}

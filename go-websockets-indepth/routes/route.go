package routes

import (
	"go-websockets-indepth/websocket"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	manager := websocket.NewManager()
	r.GET("/ws", manager.ServeWs)
}

package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Connection struct {
	Conn          *websocket.Conn
	Subscriptions map[string]bool
}

type WebSocketServer struct {
	Connections map[*Connection]bool
	mu          sync.Mutex
}

type Message struct {
	Action string `json:"action"`
	Event  string `json:"event,omitempty"`
}

var (
	PubSubServer = &WebSocketServer{
		Connections: make(map[*Connection]bool),
	}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatal("Failed to set websocket upgrade: ", err)
		return
	}

	defer conn.Close()

	wsConn := PubSubServer.AddConnection(conn)
	defer PubSubServer.RemoveConnection(wsConn)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %+v", err)
			return

		}
		var msg Message

		if err := json.Unmarshal(message, &msg); err != nil {
			log.Fatal("Failed to unmarshal message: ", err)
		}

		switch msg.Action {
		case "subscribe":
			if msg.Event == "" {
				log.Fatal("Event is required")
			}
			wsConn.Subscribe(msg.Event, conn)
			go startListeningToEvent(msg.Event, conn)
		case "unsubscribe":
			if msg.Event == "" {
				log.Fatal("Event is required")
			}
			wsConn.Unsubscribe(msg.Event, conn)

		}

	}
}

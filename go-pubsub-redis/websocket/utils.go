package websocket

import (
	"go-pubsub-redis/libraries/redis"
	"log"

	"github.com/gorilla/websocket"
)

func (s *WebSocketServer) AddConnection(conn *websocket.Conn) *Connection {
	s.mu.Lock()
	defer s.mu.Unlock()

	newConn := &Connection{
		Conn:          conn,
		Subscriptions: make(map[string]bool),
	}
	s.Connections[newConn] = true
	return newConn
}

func (s *WebSocketServer) RemoveConnection(c *Connection) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Connections, c)
}

func (c *Connection) Subscribe(event string, conn *websocket.Conn) {
	c.Subscriptions[event] = true
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Subscribed to the event")); err != nil {
		log.Printf("Failed to write message: %+v", err)
		return
	}
}

func (c *Connection) Unsubscribe(event string, conn *websocket.Conn) {
	if _, ok := c.Subscriptions[event]; !ok {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("Event doesn't exists")); err != nil {
			log.Printf("Failed to write message: %+v", err)
			return
		}
		return
	}
	delete(c.Subscriptions, event)
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Unsubscribed from the event")); err != nil {
		log.Printf("Failed to write message: %+v", err)
		return
	}
}

func startListeningToEvent(event string, conn *websocket.Conn) {

	pubsub := redis.GetInstance().Subscribe(event) // Pass the event as a string argument
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			log.Printf("Failed to receive message: %+v", err)
			return
		}

		PubSubServer.mu.Lock()
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			log.Printf("Failed to write message: %+v", err)
			return
		}
		PubSubServer.mu.Unlock()
	}

}

package websocket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pingWait     = 10 * time.Second
	pingInterval = (pingWait * 2) / 10
)

type ClientLists map[*Client]bool
type Client struct {
	connection *websocket.Conn
	Manager    *Manager

	// to avoid concurrent write we will use channel

	egress chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		Manager:    manager,
		egress:     make(chan []byte),
	}
}

func (c *Client) ReadMessages() {

	defer func() {
		c.Manager.removeClient(c)
	}()

	c.connection.SetReadLimit(512)
	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WEBSOCKET  - error reading message: %v", err)
			}
			break
		}

		for wsClient := range c.Manager.clients {
			wsClient.egress <- payload
		}
		log.Printf("Message: %v", string(payload))
	}
}

func (c *Client) WriteMessages() {
	defer func() {
		c.Manager.removeClient(c)
	}()
	// ticker := time.NewTicker(pingInterval)
	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Printf("Error writing to client: %v", err)
				}
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Error writing to client: %v", err)
			}
			// case <-ticker.C:
			// 	log.Print("Ping")
			// 	if err := c.connection.WriteMessage(websocket.TextMessage, []byte("Pinging")); err != nil {
			// 		log.Printf("Error writing to client: %v", err)
			// 	}

		}
	}
}

package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	CheckOrigin:     checkOrigin,
	WriteBufferSize: 1024,
}

type Manager struct {
	clients ClientLists
	sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientLists),
	}
}

func (m *Manager) ServeWs(c *gin.Context) {
	log.Print("WEBSOCKET | New connection registered ")

	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := NewClient(conn, m)
	m.addClient(client)
	go client.ReadMessages()
	go client.WriteMessages()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()

	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("origin")
	log.Print(origin)
	switch origin {
	case "localhost:3000":
		return true
	}
	return false
}

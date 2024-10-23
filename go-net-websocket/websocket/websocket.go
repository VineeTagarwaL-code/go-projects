package websocket

import (
	"sync"

	"golang.org/x/net/websocket"
)

var (
	instance *Server
	once     sync.Once
)

type Server struct {
	cons map[*websocket.Conn]bool
	mu   sync.Mutex
}

func GetInstance() *Server {
	once.Do(func() {
		instance = &Server{
			cons: make(map[*websocket.Conn]bool),
		}
	})
	return instance
}

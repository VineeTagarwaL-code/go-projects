package websocket

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func (s *Server) HandleReq(ws *websocket.Conn) {
	if ws == nil {
		log.Fatal("ws is nil")
		return
	}

	s.AddConn(ws)
	// how to get size of s

	fmt.Println("Number of clients connected: ", len(s.cons))
	for {
		msg := make([]byte, 512)
		n, err := ws.Read(msg)
		if err != nil {
			log.Println("Error reading from WebSocket:", err)
			s.RemoveConn(ws)
			break
		}

		s.BroadCast(msg[:n])
		log.Printf("Received message: %s", msg[:n])
	}
}

func (s *Server) AddConn(ws *websocket.Conn) {
	s.mu.Lock()

	s.cons[ws] = true

	s.mu.Unlock()
}

func (s *Server) RemoveConn(ws *websocket.Conn) {
	s.mu.Lock()

	delete(s.cons, ws)

	s.mu.Unlock()
}
func GetHandler() websocket.Server {
	websocketHandler := websocket.Server{
		Handler: websocket.Handler(func(ws *websocket.Conn) {
			server := GetInstance()
			server.HandleReq(ws)
		}),
		Handshake: func(c *websocket.Config, r *http.Request) error {
			return nil
		},
	}

	return websocketHandler
}

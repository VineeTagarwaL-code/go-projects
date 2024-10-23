package websocket

import (
	"log"

	"golang.org/x/net/websocket"
)

func (s *Server) BroadCast(msg []byte) {
	for ws := range s.cons {
		go func(ws *websocket.Conn) {

			_, err := ws.Write(msg)
			if err != nil {
				log.Printf("Failed to write to client: %v", err)
			}
		}(ws)
	}
}

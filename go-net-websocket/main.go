package main

import (
	"log"
	"net-websocket/websocket"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Use the handler function from the websocket package
	http.Handle("/ws", websocket.GetHandler())

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000" // Default port if not set
	}

	log.Print("Starting the server...")
	log.Print("Server started on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

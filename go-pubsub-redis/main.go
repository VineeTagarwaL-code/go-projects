package main

import (
	"fmt"
	"go-pubsub-redis/libraries/redis"
	"go-pubsub-redis/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	redis.GetInstance()
	log.Print("SERVER - Starting server")
	routes.SetupRouter(r)
	r.Run(":3000")
}

func HelloWorld() {
	fmt.Println("Hello World")
}

package main

import (
	"fmt"
	"go-websockets-indepth/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.SetupRouter(r)
	r.Run(":3000")
	fmt.Println("Hello World")
}

package main

import (
	"flag"
	"fmt"
	"go-jwt/route"

	"github.com/gin-gonic/gin"
)

func main() {
	var port string
	flag.StringVar(&port, "Port", ":7788", "server port")

	app := gin.Default()

	//route
	route.Setup(app)

	app.Run(port)
	fmt.Println("success")
}

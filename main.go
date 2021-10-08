package main

import (
	"flag"
	"fmt"
	"go-jwt/middleware"
	"go-jwt/route"

	"github.com/gin-gonic/gin"
)

func main() {
	var port string
	flag.StringVar(&port, "Port", ":7788", "server port")

	app := gin.Default()

	//middleware
	middleware.Setup(app)
	//route
	route.Setup(app)

	app.Run(port)
	fmt.Println("success")
}

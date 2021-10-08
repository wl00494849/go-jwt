package route

import (
	"go-jwt/controller"

	"github.com/gin-gonic/gin"
)

type Result struct {
	ResultString string
}

func Setup(app *gin.Engine) {
	userGroup(app)
}

func userGroup(app *gin.Engine) {
	router := app.Group("/User")

	router.POST("/Register", controller.Register)
}

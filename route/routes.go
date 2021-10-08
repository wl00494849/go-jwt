package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	ResultString string
}

func Setup(route *gin.Engine) {

	route.GET("/", func(c *gin.Context) {
		result := &Result{
			ResultString: "hellow",
		}
		c.JSON(http.StatusOK, result)
	})

}

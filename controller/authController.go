package controller

import (
	"go-jwt/server"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var data map[string]string

	ctx.ShouldBind(&data)
	server.UserRegister(data)

	ctx.JSON(http.StatusOK, "success")
}

func Login(ctx *gin.Context) {
	var data map[string]string

	ctx.ShouldBind(&data)
	token, err := server.GetToken(data)

	if err != nil {
		panic(err)
	}

	ctx.SetCookie("jwt", token, int((time.Hour * 24).Seconds()), "/", "localhost", false, true)
}

func UserInfo(ctx *gin.Context) {
	cookie, _ := ctx.Cookie("jwt")
	user := server.GetUserInfo(cookie)

	ctx.JSON(200, user)
}

func LoginOut(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "", "", false, true)
	ctx.JSON(200, "message:success")
}

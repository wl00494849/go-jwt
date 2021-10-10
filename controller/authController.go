package controller

import (
	"go-jwt/model"
	"go-jwt/server"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var data map[string]string

	ctx.ShouldBind(&data)

	pwd, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := &model.User{
		UserName: data["UserName"],
		Email:    data["Email"],
		Password: pwd,
	}

	server.UserRegister(user)

	ctx.JSON(http.StatusOK, user)
}

func Login(ctx *gin.Context) {
	var data map[string]string

	ctx.ShouldBind(&data)
	server.UserLogin(data)
	ctx.JSON(200, data)
}

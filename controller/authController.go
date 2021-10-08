package controller

import (
	"go-jwt/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var data map[string]string

	ctx.ShouldBindJSON(&data)

	pwd, _ := bcrypt.GenerateFromPassword([]byte(data["passward"]), 14)

	user := &model.User{
		UserName: data["userName"],
		Email:    data["email"],
		Password: pwd,
	}

	// server.UserRegister(user)

	ctx.JSON(http.StatusOK, user)
}

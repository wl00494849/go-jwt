package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var data map[string]string

	ctx.ShouldBindJSON(&data)

	ctx.JSON(http.StatusOK, data)
}

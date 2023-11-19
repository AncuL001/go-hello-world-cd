package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HelloAPI interface {
	Hello(ctx *gin.Context)
	Secret(ctx *gin.Context)
}

type HelloAPIImpl struct {}

func CreateHelloAPIImpl() *HelloAPIImpl {
	return &HelloAPIImpl{}
}

func (h *HelloAPIImpl) Hello (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func (h *HelloAPIImpl) Secret(ctx *gin.Context) {
	secret := os.Getenv("SECRET")

	ctx.JSON(http.StatusOK, gin.H{
		"message": secret,
	})
}
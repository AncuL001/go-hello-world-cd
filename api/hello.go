package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloAPI interface {
	Hello(ctx *gin.Context)
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
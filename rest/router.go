package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	d := gin.Default()

	d.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	return d
}

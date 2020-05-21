package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFile("/main.js", "/usr/share/hackine/assets/main.js")
	r.NoRoute(func(ctx *gin.Context) {
		ctx.File("/usr/share/hackine/assets/index.html")
	})
	return r
}

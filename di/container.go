package di

import "github.com/gin-gonic/gin"

type Container struct {
}

func (c Container) Router() *gin.Engine {
	return gin.Default()
}

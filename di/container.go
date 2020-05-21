package di

import (
	"github.com/gin-gonic/gin"
	"github.com/keitam913/hackine/rest"
)

type Container struct {
}

func (c Container) Router() *gin.Engine {
	return rest.NewRouter()
}

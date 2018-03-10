package router

import (
	"github.com/gin-gonic/gin"
)

type routerAble interface {
	create(engine *gin.Engine)
}

func Load() *gin.Engine {
	e := gin.Default()

	f := FunctionRouter{}
	f.create(e)

	return e
}
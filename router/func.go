package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FunctionRouter struct {
	routerAble
}

func ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": "pong",
	})
}

func detectRules(context *gin.Context) {

}

func (FunctionRouter) create(engine *gin.Engine) {
	funcGroup := engine.Group("/func")
	{
		funcGroup.GET("/ping", ping)
	}
}
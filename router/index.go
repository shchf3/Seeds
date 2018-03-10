package router

import (
	"github.com/gin-gonic/gin"
	"Seeds/utils"
	"net/http"
)

type routerAble interface {
	create(engine *gin.Engine)
}

var unauthorizedMessage = gin.H{
	"ret": 0,
	"data": "token or source is invalid",
}

func KeyVerify(context *gin.Context) {
	key, hasKey := context.GetQuery("key")
	if !hasKey {
		context.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedMessage)
		return
	}

	if utils.GetConfig().Get("verifyKey") != key {
		context.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedMessage)
		return
	}
}

func Load() *gin.Engine {
	e := gin.Default()
	e.Use(KeyVerify)

	function := FunctionRouter{}
	function.create(e)

	user := UserRouter{}
	user.create(e)

	return e
}

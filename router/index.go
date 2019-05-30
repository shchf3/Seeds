package router

import (
	"net/http"

	"github.com/CloudHammer/Seeds/utils"
	"github.com/gin-gonic/gin"
)

type routerAble interface {
	create(engine *gin.RouterGroup)
}

var unauthorizedMessage = gin.H{
	"ret":  0,
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

	routers := []routerAble{
		FunctionRouter{},
		UserRouter{},
		NodeRouter{},
	}

	var modPath *gin.RouterGroup

	if utils.GetConfig().GetBool("modURL") {
		modPath = e.Group("/mod_mu")
	} else {
		modPath = e.Group("")
	}

	for _, r := range routers {
		r.create(modPath)
	}

	return e
}

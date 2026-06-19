package router

import (
	"growthos/internal/handler"

	"github.com/gin-gonic/gin"
)

func AuthRouterInit(router *gin.Engine, authHandler handler.AuthHandler) {
	user := router.Group("/user")
	user.POST("register", authHandler.Register)
	user.POST("login", authHandler.Login)
}

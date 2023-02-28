package router

import (
	userHandler "GinProject/internal/router/handler/user"

	"github.com/gin-gonic/gin"

	mid "GinProject/internal/middlewareLib"
)

var (
	ge *gin.Engine
)

func Init() {
	ge = gin.Default()
}

func Run() {
	ge.Use(mid.ErrorWrapper)

	ge.GET("/user",
		mid.JWTValidate,
		userHandler.ListUser)

	ge.Run()
}

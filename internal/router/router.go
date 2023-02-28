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

	ge.GET("/user/:account",
		mid.JWTValidate,
		userHandler.GetUser)

	ge.POST("/user",
		userHandler.CreateUser)

	ge.DELETE("/user/:account",
		mid.JWTValidate,
		userHandler.DeleteUser)

	ge.POST("/login",
		userHandler.Login)

	ge.Run()
}

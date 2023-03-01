package router

import (
	userHandler "GinProject/internal/router/handler/user"
	"fmt"

	"github.com/gin-gonic/gin"

	mid "GinProject/internal/middlewareLib"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	ge.PATCH("/user/:account",
		mid.JWTValidate,
		userHandler.UpdateUser)

	ge.DELETE("/user/:account",
		mid.JWTValidate,
		userHandler.DeleteUser)

	ge.POST("/login",
		userHandler.Login)

	ge.NoRoute(userHandler.NoRoute)

	if mode := gin.Mode(); mode == gin.DebugMode {
		fmt.Println("in swagger")
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 8080))
		ge.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	ge.Run()
}

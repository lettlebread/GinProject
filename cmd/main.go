package main

import (
	_ "GinProject/docs"
	"GinProject/internal/db"
	"GinProject/internal/router"
	userHandler "GinProject/internal/router/handler/user"
	jwtUtil "GinProject/internal/util/jwt"
)

// @title GinProject
// @version 1.0
// @contact.name Johnny.Chen
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	db.Init()
	jwtUtil.Init()
	userHandler.Init()
	router.Init()
	router.Run()
}

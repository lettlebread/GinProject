package main

import (
	"GinProject/internal/db"
	"GinProject/internal/router"
	userHandler "GinProject/internal/router/handler/user"
	jwtUtil "GinProject/internal/util"
)

func main() {
	db.Init()
	jwtUtil.Init()
	userHandler.Init()
	router.Init()
	router.Run()
}

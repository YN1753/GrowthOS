package main

import (
	"fmt"
	"growthos/internal/config"
	"growthos/internal/database"
	"growthos/internal/handler"
	"growthos/internal/repository"
	"growthos/internal/router"
	"growthos/internal/service"
)

func main() {
	Config := config.ConfigInit()
	fmt.Println(Config)
	db := database.MysqlInit(Config.Mysql)
	fmt.Println(db)
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)
	authHandler := handler.NewAuthHandler(userService, Config.Jwt.Secret, Config.Jwt.ExpireHour)
	r := router.RouterInit(authHandler)
	r.Run(":" + Config.Server.Port)
}

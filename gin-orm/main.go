package main

import (
	"gin-orm/config"
	"gin-orm/controller"
	"gin-orm/repository"
	"gin-orm/routes"
	"gin-orm/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	db := config.Connect()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userCtrl := controller.NewUserController(userService)

	routes.UserRoutes(router, userCtrl)

	router.Run(":8000")
}

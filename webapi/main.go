package main

import (
	"webapi/controller"
	"webapi/repository"
	"webapi/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)

	r := gin.Default()

	r.GET("/users", ctrl.GetAllUsers)
	r.GET("/users/:id", ctrl.GetUserByID)
	r.POST("/users", ctrl.CreateUser)
	r.PUT("/users", ctrl.UpdateUser)
	r.DELETE("/users/:id", ctrl.DeleteUser)

	r.Run(":8080")
}

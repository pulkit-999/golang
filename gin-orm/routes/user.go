package routes

import (
	"gin-orm/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userCtrl *controller.UserController) {
	router.GET("/", userCtrl.GetUsers)
	router.GET("/:id", userCtrl.GetUser)
	router.POST("/", userCtrl.CreateUser)
	router.DELETE("/:id", userCtrl.DeleteUser)
	router.PUT("/:id", userCtrl.UpdateUser)
}

package routes

import (
	"webapi/webaApi/controller"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.Engine, employeectrl *controller.EmployeeController) {
	router.GET("/", employeectrl.GetEmployees)
	router.POST("/create", employeectrl.CreateEmployees)
}

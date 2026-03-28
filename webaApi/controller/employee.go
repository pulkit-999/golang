package controller

import (
	"net/http"
	"webapi/webaApi/models"
	"webapi/webaApi/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (r *EmployeeController) GetEmployees(ctx *gin.Context) {
	employees, err := r.service.GetEmployees()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, employees)
}

func (r *EmployeeController) CreateEmployees(ctx *gin.Context) {
	var employee models.Employee
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := r.service.CreateEmployees(&employee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, employee)
}

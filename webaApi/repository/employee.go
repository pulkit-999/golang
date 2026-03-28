package repository

import (
	"webapi/webaApi/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetEmployees() ([]models.Employee, error)
	CreateEmployees(*models.Employee) error
}
type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) CreateEmployees(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) GetEmployees() ([]models.Employee, error) {
	var employee []models.Employee
	err := r.db.Find(&employee).Error
	return employee, err
}

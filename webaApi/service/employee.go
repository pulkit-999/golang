package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"webapi/webaApi/models"
	"webapi/webaApi/repository"

	"github.com/redis/go-redis/v9"
)

type EmployeeService interface {
	GetEmployees() ([]models.Employee, error)
	CreateEmployees(*models.Employee) error
}

type employeeService struct {
	repo  repository.EmployeeRepository
	redis *redis.Client
	ctx   context.Context
}

func NewEmployeeService(repo repository.EmployeeRepository, redis *redis.Client) EmployeeService {
	return &employeeService{
		repo:  repo,
		redis: redis,
		ctx:   context.Background(),
	}
}

func (s *employeeService) GetEmployees() ([]models.Employee, error) {
	cacheKey := "employees:all"

	cached, err := s.redis.Get(s.ctx, cacheKey).Result()

	if err == nil {
		fmt.Println("✅ CACHE HIT (Redis)")

		var employees []models.Employee
		json.Unmarshal([]byte(cached), &employees)
		return employees, nil
	}

	if err == redis.Nil {
		fmt.Println("❌ CACHE MISS → Fetching from DB")
	} else {
		fmt.Println("⚠️ Redis error → fallback to DB:", err)
	}

	employees, err := s.repo.GetEmployees()
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(employees)
	s.redis.Set(s.ctx, cacheKey, data, 60*time.Second)

	return employees, nil
}

func (s *employeeService) CreateEmployees(employee *models.Employee) error {
	err := s.repo.CreateEmployees(employee)
	if err != nil {
		return err
	}

	// ❗ Clear cache
	s.redis.Del(s.ctx, "employees:all")

	return nil
}

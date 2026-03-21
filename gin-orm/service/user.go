package service

import (
	"gin-orm/models"
	"gin-orm/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id uint, user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(id uint, user *models.User) (*models.User, error) {
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	err = s.repo.Update(existingUser)
	if err != nil {
		return nil, err
	}
	return existingUser, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

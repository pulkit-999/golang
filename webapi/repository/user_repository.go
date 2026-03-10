package repository

import "webapi/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
}

type userRepository struct {
	users []models.User
	nextID int
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: []models.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		},
		nextID: 3,
	}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	return r.users, nil
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *userRepository) Create(user *models.User) error {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, *user)
	return nil
}

func (r *userRepository) Update(user *models.User) error {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = *user
			return nil
		}
	}
	return nil
}

func (r *userRepository) Delete(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return nil
}

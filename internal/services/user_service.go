package services

import (
	"api/internal/models"
	"api/internal/repositories"
)

type UserService interface {
	Create(user *models.User) error // Correct method name to 'Create'
	GetById(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(user *models.User) error { // Correct method name to 'Create'
	return s.repo.Create(user) // Call 'Create' method from repository
}

func (s *userService) GetById(id int) (*models.User, error) {
	return s.repo.GetById(id)
}

func (s *userService) Update(user *models.User) error { // Correct method name to 'Create'
	return s.repo.Update(user) // Call 'Create' method from repository
}

func (s *userService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

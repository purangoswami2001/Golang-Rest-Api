package repositories

import (
	"api/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error // Change method name to 'Create'
	GetById(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id int) error
	FindByEmail(email string, user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (s *userRepository) GetById(id int) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

// FindByEmail fetches a user by their email
func (r *userRepository) FindByEmail(email string, user *models.User) error {
	return r.db.Where("email = ?", email).First(user).Error
}

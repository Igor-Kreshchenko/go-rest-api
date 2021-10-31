package repositories

import (
	"errors"

	"github.com/Igor-Kreshchenko/go-rest-api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserById(userID uint) (*models.User, error)
	UpdateUserName(id uint, newName string) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Where("Email = ?", user.Email).First(user).Error
	if err == nil {
		return nil, errors.New("user already exists")
	}

	err = r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindUserById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUserName(id uint, newName string) error {
	var user *models.User

	res := r.db.Model(&user).Where("id = ?", id).Update("name", newName)

	return res.Error
}

func (r *userRepository) DeleteUser(id uint) error {
	res := r.db.Where(id).Delete(&models.User{})

	return res.Error
}

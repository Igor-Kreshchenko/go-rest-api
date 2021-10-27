package services

import (
	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"github.com/Igor-Kreshchenko/go-rest-api/repositories"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Name struct {
	Name string `json:"name" binding:"required"`
}

type UserService interface {
	CreateUser(user *CreateUserRequest) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
	UpdateUserName(id uint, newName string) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) CreateUser(userReq *CreateUserRequest) (*models.User, error) {
	createdUser, err := s.userRepository.CreateUser(&models.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	})
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *userService) FindUserById(id uint) (*models.User, error) {
	user, err := s.userRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) UpdateUserName(id uint, name string) error {
	var newName string

	_, err := s.userRepository.FindUserById(id)
	if err != nil {
		return err
	}

	newName = name

	err = s.userRepository.UpdateUserName(id, newName)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id uint) error {
	err := s.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

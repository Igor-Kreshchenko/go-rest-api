package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type VerifyUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
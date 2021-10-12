package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"title"`
	Email    string `json:"author"`
	Password string `json:"text"`
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
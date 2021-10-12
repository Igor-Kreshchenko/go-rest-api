package controllers

import (
	"net/http"

	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
  // Validate input
  var input models.CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create user
  user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}
  models.DB.Create(&user)

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func LoginUser(c *gin.Context) {
  // Validate input
  var input models.VerifyUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create user
  user := models.User{Email: input.Email, Password: input.Password}
  models.DB.Create(&user)

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func LogoutUser(c *gin.Context) {  // Get model if exist
  var user models.User

  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}
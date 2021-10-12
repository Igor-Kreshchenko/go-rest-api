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

func LogoutUser(c *gin.Context) {
  // Get model if exist
  var user models.User

  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUsers(c *gin.Context) {
  var users []models.User
  models.DB.Find(&users)

  c.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUser(c *gin.Context) {  // Get model if exist
  var user models.User

  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input models.UpdateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&user).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&user)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
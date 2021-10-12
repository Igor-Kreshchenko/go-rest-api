package controllers

import (
	"net/http"

  "github.com/gin-gonic/gin"
  "github.com/Igor-Kreshchenko/go-rest-api/models"
)

func FindPosts(c *gin.Context) {
  var posts []models.Post
  models.DB.Find(&posts)

  c.JSON(http.StatusOK, gin.H{"data": posts})
}

func CreatePost(c *gin.Context) {
  // Validate input
  var input models.CreatePostInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create post
  post := models.Post{Title: input.Title, Author: input.Author, Text: input.Text}
  models.DB.Create(&post)

  c.JSON(http.StatusOK, gin.H{"data": post})
}

func FindPost(c *gin.Context) {  // Get model if exist
  var post models.Post

  if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
  // Get model if exist
  var post models.Post
  if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input models.UpdatePostInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&post).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
  // Get model if exist
  var post models.Post
  if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&post)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
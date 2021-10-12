package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPosts(c *gin.Context) {
  var posts []models.Post
  models.DB.Find(&posts)

  c.JSON(http.StatusOK, gin.H{"data": posts})
}
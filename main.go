package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"github.com/Igor-Kreshchenko/go-rest-api/controllers"
)

func main() {
  r := gin.Default()

  models.ConnectDataBase()

  r.GET("/posts", controllers.FindPosts)
  r.POST("/posts", controllers.CreatePost)
  r.GET("/posts/:id", controllers.FindPost)
  r.PATCH("/posts/:id", controllers.UpdatePost)
  r.DELETE("/posts/:id", controllers.DeletePost)

  r.POST("/users/register", controllers.RegisterUser)
  r.POST("/users/login", controllers.LoginUser)
  r.GET("/users/logout/:id", controllers.LogoutUser)

  r.Run()
}
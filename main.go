package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"github.com/Igor-Kreshchenko/go-rest-api/controllers"
)

func main() {
  r := gin.Default()

  models.ConnectDataBase()

  r.GET("/api/posts", controllers.FindPosts)
  r.POST("/api/posts", controllers.CreatePost)
  r.GET("/api/posts/:id", controllers.FindPost)
  r.PATCH("/api/posts/:id", controllers.UpdatePost)
  r.DELETE("/api/posts/:id", controllers.DeletePost)

  r.POST("/users/register", controllers.RegisterUser)
  r.POST("/users/login", controllers.LoginUser)
  r.GET("/users/logout/:id", controllers.LogoutUser)

  r.GET("/api/users", controllers.FindUsers)
  r.GET("/api/users/:id", controllers.FindUser)
  r.DELETE("/api/users/:id", controllers.DeleteUser)
  r.PATCH("/api/users/:id", controllers.UpdateUser)

  r.Run()
}
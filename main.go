package main

import (
	"github.com/Igor-Kreshchenko/go-rest-api/controllers"
	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDataBase()

	router.GET("/api/posts", controllers.FindPosts)
	router.POST("/api/posts", controllers.CreatePost)
	router.GET("/api/posts/:id", controllers.FindPost)
	router.PATCH("/api/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/posts/:id", controllers.DeletePost)

	router.GET("/api/users", controllers.FindUsers)
	router.POST("/api/users", controllers.CreateUser)
	router.GET("/api/users/:id", controllers.FindUser)
	router.DELETE("/api/users/:id", controllers.DeleteUser)
	router.PATCH("/api/users/:id", controllers.UpdateUser)

	router.Run()
}

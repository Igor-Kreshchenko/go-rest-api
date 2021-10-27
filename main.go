package main

import (
	"github.com/Igor-Kreshchenko/go-rest-api/api"
	"github.com/Igor-Kreshchenko/go-rest-api/repositories"
	"github.com/Igor-Kreshchenko/go-rest-api/services"
	"github.com/Igor-Kreshchenko/go-rest-api/setup"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := setup.ConnectDataBase()
	if err != nil {
		panic(err)
	}

	gr := router.Group("v1/api")

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository)

	api.InjectUser(gr, userService)
	api.InjectPost(gr, postService)

	router.Run()
}

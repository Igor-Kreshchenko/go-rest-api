package api

import (
	"net/http"
	"strconv"

	"github.com/Igor-Kreshchenko/go-rest-api/services"
	"github.com/gin-gonic/gin"
)

func InjectUser(gr *gin.RouterGroup, userService services.UserService) {
	handler := gr.Group("users")

	handler.POST("", userCreate(userService))
	handler.GET(":id", userFind(userService))
	handler.PATCH(":id", updateUserName(userService))
	handler.DELETE(":id", userDelete(userService))
}

func userCreate(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request *services.CreateUserRequest

		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}

		user, err := userService.CreateUser(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created",
			"user":    user,
		})
	}
}

func userFind(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		uid, err := strconv.ParseUint(id, 10, 64)

		user, err := userService.FindUserById(uint(uid))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}

func updateUserName(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var name *services.Name
		id := c.Param("id")
		uid, err := strconv.ParseUint(id, 10, 64)

		err = c.BindJSON(&name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = userService.UpdateUserName(uint(uid), name.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Name updated",
		})
	}
}

func userDelete(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		uid, err := strconv.ParseUint(id, 10, 64)

		err = userService.DeleteUser(uint(uid))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted",
		})
	}
}

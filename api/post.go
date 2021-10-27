package api

import (
	"net/http"
	"strconv"

	"github.com/Igor-Kreshchenko/go-rest-api/services"
	"github.com/gin-gonic/gin"
)

func InjectPost(gr *gin.RouterGroup, postService services.PostService) {
	handler := gr.Group("posts")

	handler.GET("", postList(postService))
	handler.POST("", postCreate(postService))
	handler.GET(":id", getPostById(postService))
	handler.PATCH(":id", updatePostText(postService))
	handler.DELETE(":id", postDelete(postService))
}

func postList(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := postService.GetAllPosts()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	}
}

func postCreate(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var postReq *services.PostRequest

		err := c.BindJSON(&postReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}

		res, err := postService.CreatePost(postReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Post created",
			"post":    res,
		})
	}
}

func getPostById(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		uid, err := strconv.ParseUint(id, 10, 64)

		post, err := postService.GetPostByID(uint(uid))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"post": post,
		})
	}
}

func updatePostText(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var text *services.Text
		id := c.Param("id")
		uid, err := strconv.ParseUint(id, 10, 64)

		err = c.BindJSON(&text)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = postService.UpdatePostText(uint(uid), text.Text)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Text updated",
		})
	}
}

func postDelete(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		uid, err := strconv.ParseUint(id, 10, 64)

		err = postService.DeletePost(uint(uid))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Post deleted",
		})
	}
}

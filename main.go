package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"github.com/Igor-Kreshchenko/go-rest-api/controllers"
)

func main() {
  r := gin.Default()

  models.ConnectDataBase()

  r.Run()
}
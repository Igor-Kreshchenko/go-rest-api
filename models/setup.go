package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	godotenv.Load()

	dbuser, isName := os.LookupEnv("DB_USER")
	dbpass, isPass := os.LookupEnv("DB_PASSWORD")
	dbhost, isHost := os.LookupEnv("DB_HOST")
	port, isPort := os.LookupEnv("PORT")

	if !isName || !isPass || !isHost || !isPort {
		logrus.Error("Cant read .env file")
	}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable`, dbhost, dbuser, dbpass, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Post{}, &User{})

	DB = database
}

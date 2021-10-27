package setup

import (
	"fmt"
	"os"

	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() (*gorm.DB, error) {
	var db *gorm.DB
	godotenv.Load()

	dbuser, isName := os.LookupEnv("DB_USER")
	dbpass, isPass := os.LookupEnv("DB_PASSWORD")
	dbhost, isHost := os.LookupEnv("DB_HOST")

	if !isName || !isPass || !isHost {
		logrus.Error("Cant read .env file")
	}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=postgres sslmode=disable`, dbhost, dbuser, dbpass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.Post{}, &models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

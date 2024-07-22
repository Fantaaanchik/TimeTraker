package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time_traker/models"
)

func InitDB() *gorm.DB {
	Host := os.Getenv("DB_HOST")
	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	Port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		Host, User, Password, DBName, Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to db, %v", err.Error())
	}

	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Println("cannot migrate models, error description: ", err.Error())
	}
	return db
}

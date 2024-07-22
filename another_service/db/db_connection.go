package db

import (
	"another_service/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		"localhost", "testuser", "testpassword", "testdatabase", "5453")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to db, %v", err.Error())
	}

	err = db.AutoMigrate(&models.UserInformation{})
	if err != nil {
		log.Println("cannot migrate models, error description: ", err.Error())
	}
	return db
}

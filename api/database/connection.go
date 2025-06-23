package database

import (
	"api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connect error:", err)
	}
	log.Println("Connected to DB")

	_ = DB.Migrator().DropTable(&models.Invoice{})
	DB.AutoMigrate(&models.User{}, &models.Invoice{})
}

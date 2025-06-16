package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "log"
  "os"
  "goinvoice/models"
  "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
  _ = godotenv.Load()
  dsn := os.Getenv("DATABASE_URL")
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal("DB connect error:", err)
  }
  db.AutoMigrate(&models.Invoice{}, &models.User{}) // model-model
  DB = db
}

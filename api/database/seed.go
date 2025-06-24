package database

import (
	"api/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	// Seed user
	user := models.User{
		Email:    "test@example.com",
		Password: string(password),
	}
	if err := DB.Create(&user).Error; err != nil {
		log.Println("User seed error:", err)
		return
	}

	// Seed invoices
	invoices := []models.Invoice{
		{UserID: user.ID, Amount: 150000, Paid: false, DueDate: time.Now().AddDate(0, 0, 7)},
		{UserID: user.ID, Amount: 225000, Paid: true, DueDate: time.Now().AddDate(0, 0, -3)},
	}
	if err := DB.Create(&invoices).Error; err != nil {
		log.Println("Invoice seed error:", err)
		return
	}

	log.Println("✅ Seeder executed")
}

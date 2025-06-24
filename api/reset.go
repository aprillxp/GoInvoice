package main

import (
	"api/database"
	"api/models"
	"log"
)

func reset() {
	// Connect ke DB
	database.ConnectDB()

	// Drop semua tabel
	err := database.DB.Migrator().DropTable(&models.Invoice{}, &models.User{})
	if err != nil {
		log.Fatal("❌ Failed to drop table:", err)
	}
	log.Println("✅ All table dropped successfully")

	// Migrasi ulang
	err = database.DB.AutoMigrate(&models.User{}, &models.Invoice{})
	if err != nil {
		log.Fatal("❌ Failed to migrate:", err)
	}
	log.Println("✅ Migrating is successfully")

	// Jalankan seeder
	database.Seed()
	log.Println("✅ Seeding is successfully")
}

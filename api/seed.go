package main

import (
	"api/database"
)

func seed() {
	database.ConnectDB()
	database.Seed()
}

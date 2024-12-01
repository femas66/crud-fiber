package migrations

import (
	"femas/belajar-fiber/database"
	"femas/belajar-fiber/models"
	"fmt"
	"log"
)

func Migration() {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration successful")
}

package repository

import (
	"gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"

	"recipeAppBackend/internal/models"
)

func CreateDB()  {
	log.Println("Creating database")
	// Conectar a la base de datos SQLite
	db, err := gorm.Open(sqlite.Open("recipe.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	// Migrar el esquema
	db.AutoMigrate(&models.User{}, &models.Recipe{})
}
package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
	"recipeAppBackend/internal/models"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("recipe.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // Automigración de la estructura de User
    db.AutoMigrate(&models.User{}, &models.Recipe{})

    return db
}
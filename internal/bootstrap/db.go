package bootstrap

import (
	"log"

	"github.com/esmeraldarr/recipeAppBackend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
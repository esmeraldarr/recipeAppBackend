package config

import (
	"github.com/esmeraldarr/recipeAppBackend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBconnet() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("recipe.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

    return db
}
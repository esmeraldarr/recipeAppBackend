package main
import (
	"gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// User es una estructura que representa una tabla en la base de datos
type User struct {
    ID   uint   `gorm:"primaryKey"`
    Username string
    Email string
    Password string
}

type Recipe struct {
    ID   uint   `gorm:"primaryKey"`
    Title string
    Ingredients string
    Instructions string
    Category string
    ImageURL string
}

func main() {
    // Conectar a la base de datos SQLite
    db, err := gorm.Open(sqlite.Open("recipe.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database", err)
    }

    // Migrar el esquema
    db.AutoMigrate(&User{}, &Recipe{})

    db.Create(&User{Username: "admin", Email: "admin@admin.admin", Password: "admin"})

}

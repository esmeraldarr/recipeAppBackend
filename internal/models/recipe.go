package models

type Recipe struct {
    ID   uint   `gorm:"primaryKey"`
    Title string
    Ingredients string
    Instructions string
    Category string
    ImageURL string
}
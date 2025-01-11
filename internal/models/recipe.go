package models

type Recipe struct {
    ID int `gorm:"primaryKey"`
    Title string
    Ingredients string
    Instructions string
    Category string
    ImageURL string
}
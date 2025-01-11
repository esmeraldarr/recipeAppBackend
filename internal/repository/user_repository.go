package repository

import (
	"github.com/esmeraldarr/recipeAppBackend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
    var users []models.User
    result := r.db.Find(&users)
    return users, result.Error
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
    var user models.User
    result := r.db.First(&user, id)
    return user, result.Error
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
    result := r.db.Create(&user)
    return user, result.Error
}
package repository

import (
    "recipeAppBackend/internal/models"
    "errors"
)

type UserRepository struct {
    data map[int]models.User
    nextID int
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        data:   make(map[int]models.User),
        nextID: 1,
    }
}

func (r *UserRepository) GetAll() ([]models.User, error) {
    users := make([]models.User, 0, len(r.data))
    for _, user := range r.data {
        users = append(users, user)
    }
    return users, nil
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
    user, exists := r.data[id]
    if !exists {
        return models.User{}, errors.New("user not found")
    }
    return user, nil
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
    user.ID = r.nextID
    r.nextID++
    r.data[user.ID] = user
    return user, nil
}
package services

import (
	"errors"

	"github.com/esmeraldarr/recipeAppBackend/internal/models"
	"github.com/esmeraldarr/recipeAppBackend/internal/repository"
)

type UserService struct {
    Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) GetUsers() ([]models.User, error) {
    return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
    user, err := s.Repo.GetByID(id)
    if err != nil {
        return models.User{}, err
    }
    if user.ID == 0 {
        return models.User{}, errors.New("user not found")
    }
    return user, nil
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
    if user.Username == "" {
        return models.User{}, errors.New("user name is required")
    }
    return s.Repo.Create(user)
}
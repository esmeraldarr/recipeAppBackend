package bootstrap

import (
	"github.com/esmeraldarr/recipeAppBackend/internal/repository"
	"github.com/esmeraldarr/recipeAppBackend/internal/services"
	"gorm.io/gorm"
)

type Services struct {
    UserService *services.UserService
}

func InitServices(db *gorm.DB) *Services {
    userRepo := repository.NewUserRepository(db)
    userService := services.NewUserService(userRepo)

    return &Services{
        UserService: userService,
    }
}
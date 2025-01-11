package main

import (
    "log"
    "net/http"
    "recipeAppBackend/config"
    "recipeAppBackend/internal/repository"
    "recipeAppBackend/internal/services"
    "recipeAppBackend/api/handlers"
    "recipeAppBackend/internal/routes"
)

func main() {
    db := config.InitDB() // Inicializa la base de datos

    userRepo := repository.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    router := routes.RegisterRoutes(userHandler)
    log.Println("Server running on port 8080")
    http.ListenAndServe(":8080", router)
}
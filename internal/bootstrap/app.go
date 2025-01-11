package bootstrap

import (
	"log"
	"net/http"

	"github.com/esmeraldarr/recipeAppBackend/api/handlers"
	"github.com/esmeraldarr/recipeAppBackend/internal/routes"
)

func StartApp() {
	db := InitDB()
	services := InitServices(db)

	userHandler := handlers.NewUserHandler(services.UserService)

	router := routes.RegisterRoutes(userHandler)
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}

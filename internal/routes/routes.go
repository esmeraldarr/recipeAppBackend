package routes

import (
	"github.com/esmeraldarr/recipeAppBackend/api/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(userHandler *handlers.UserHandler) *mux.Router {

    r := mux.NewRouter()
    r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
    r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
    r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

    return r
}
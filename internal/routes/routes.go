package routes

import (
    "github.com/gorilla/mux"
    "recipeAppBackend/api/handlers"
)

func RegisterRoutes() *mux.Router {
    r := mux.NewRouter()

    // User routes
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

    // Recipe routes
    r.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")
    r.HandleFunc("/recipes/{id}", handlers.GetRecipeByID).Methods("GET")
    r.HandleFunc("/recipes", handlers.CreateRecipe).Methods("POST")

    return r
}
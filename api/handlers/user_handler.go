package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "recipeAppBackend/internal/services"
    "recipeAppBackend/internal/models"
    "github.com/gorilla/mux"
)

type UserHandler struct {
    Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{Service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.Service.GetUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    user, err := h.Service.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    createdUser, err := h.Service.CreateUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(createdUser)
}
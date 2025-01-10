package middlewares

import (
    "net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Lógica de autenticación
        next.ServeHTTP(w, r)
    })
}
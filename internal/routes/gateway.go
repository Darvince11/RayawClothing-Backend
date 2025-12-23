package routes

import (
	"net/http"
	"rayaw-api/internal/handlers"
	"rayaw-api/internal/middleware"
)

func ServerMux() http.Handler {
	mux := http.NewServeMux()

	authHandlers := handlers.NewAuthenticationHandler("Hello")
	authRoutes := NewAuthRoutes(mux, authHandlers)
	authRoutes.RegisterRoutes()

	return middleware.CorsMiddleware(mux)
}

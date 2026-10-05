package routes

import (
	"database/sql"
	"net/http"
	"rayaw-api/internal/config"
	"rayaw-api/internal/handlers"
	"rayaw-api/internal/middleware"
	"rayaw-api/internal/repositories"
	"rayaw-api/internal/services"
)

func ServerMux(config *config.Config, db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	//Handle auth
	tokenRepo := repositories.NewTokenRepository(db)
	tokenService := services.NewTokenService(config.AuthConfig.JWTSecretKey, tokenRepo)
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandlers := handlers.NewAuthenticationHandler(authService, tokenService)
	authRoutes := NewAuthRoutes(mux, authHandlers)
	authRoutes.RegisterRoutes()

	return middleware.CorsMiddleware(mux)
}

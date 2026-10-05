package routes

import (
	"net/http"
	"rayaw-api/internal/handlers"
)

type AuthRoutes struct {
	mux         *http.ServeMux
	authHandler *handlers.AuthenticationHandler
}

func NewAuthRoutes(mux *http.ServeMux, authHandler *handlers.AuthenticationHandler) *AuthRoutes {
	return &AuthRoutes{mux: mux, authHandler: authHandler}
}

func (a *AuthRoutes) RegisterRoutes() {
	a.mux.HandleFunc("POST /signup", a.authHandler.SignUpHandler)
	a.mux.HandleFunc("POST /login", a.authHandler.LoginHandler)
}

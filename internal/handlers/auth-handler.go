package handlers

import (
	"fmt"
	"net/http"
)

type AuthenticationHandler struct {
	as string
}

func NewAuthenticationHandler(as string) *AuthenticationHandler {
	return &AuthenticationHandler{as: as}
}

func (a *AuthenticationHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sign up")
}

func (a *AuthenticationHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	//Handle login
}

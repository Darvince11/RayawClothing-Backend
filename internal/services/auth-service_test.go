package services

import (
	"rayaw-api/internal/config"
	"testing"
)

var cfg = config.Config{
	Port:  "",
	DbUrl: "",
	AuthConfig: &config.AuthConfig{
		JWTSecretKey: "secret_key",
	},
}
var authService *AuthService = NewAuthService(nil, &cfg)

func TestJWTAuthentication(t *testing.T) {
	//test GenerateJWT()
	token, err := authService.GenerateJWT("Kafui", "customer", 3600)

	//Check for any errors
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	//Check for empty token
	if token == "" {
		t.Error("expected token, got an empty string")
	}

	//test validateJWT()
	_, err = authService.ValidateJWT(token)

	if err != nil {
		t.Errorf("Error validating token:%v", err)
	}
}

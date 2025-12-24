package services

import (
	"testing"
)

func TestAuthService(t *testing.T) {
	authService := NewAuthService("hello")
	token, err := authService.GenerateJWT()

	if err != nil {
		t.Errorf("Error generating token:%v", err)
	}

	_, err = authService.ValidateJWT(token)

	if err != nil {
		t.Errorf("Error validating token:%v", err)
	}
}

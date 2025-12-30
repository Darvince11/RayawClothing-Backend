package services

import (
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
	"rayaw-api/internal/tests"
	"testing"
)

func TestAuthServeice(t *testing.T) {
	db := tests.SetupTestDB(t)
	authRepo := repositories.NewAuthRepository(db)
	authService := NewAuthService(authRepo)

	newUser := models.SignUpRequest{
		First_name:    "Kafui",
		Last_name:     "Dotse",
		Email:         "kafui42@gmail.com",
		Phone_number:  "02031578",
		User_password: "123453423",
	}

	_, err := authService.Register(&newUser)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	user, err := authService.GetUserByEmail(newUser.Email)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	} else {
		t.Log(user)
	}
}

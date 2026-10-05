package repositories

import (
	"rayaw-api/internal/models"
	"rayaw-api/internal/tests"
	"testing"
)

func TestAuthRepo(t *testing.T) {
	db := tests.SetupTestDB(t)
	if db == nil {
		t.Fatal("db is nil")
	}

	authRepo := NewAuthRepository(db)

	user := models.User{
		First_name:    "Kafui",
		Last_name:     "Dotse",
		Email:         "kafui21h@gmail.com",
		Phone_number:  "0233202897809",
		User_password: "123456",
	}

	userId, err := authRepo.AddUser(&user)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	} else {
		t.Log("New user id:", userId)
	}

	result, err := authRepo.GetUserByEmail(user.Email)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	t.Log(result)

	err = authRepo.UpdateUser(userId, &user)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

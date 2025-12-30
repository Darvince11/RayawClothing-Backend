package repositories

import (
	"rayaw-api/internal/models"
	"rayaw-api/internal/tests"
	"testing"
	"time"
)

func TestTokenRepo(t *testing.T) {
	db := tests.SetupTestDB(t)
	if db == nil {
		t.Fatal("db is nil")
	}

	tokenRepo := NewTokenRepository(db)

	token := models.RefreshToken{
		UserId: 1,
		Token:  "okww2_p0KTiZY3jdYznguYhFKpIAiTypwluflhynLXs",
		Expiry: time.Now().Add(1800 * time.Second),
	}

	err := tokenRepo.AddRefreshToken(&token)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	result, err := tokenRepo.GetRefreshToken(token.Token)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	t.Log(result)
}

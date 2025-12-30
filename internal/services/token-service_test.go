package services

import (
	"os"
	"rayaw-api/internal/repositories"
	"rayaw-api/internal/tests"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestTokenService(t *testing.T) {
	db := tests.SetupTestDB(t)

	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Errorf("Expected not error, got: %v", err)
	}

	key := os.Getenv("JWT_SECRETKEY")
	if key == "" {
		t.Fatal("Invalid environment variable")
	}

	tokenRepo := repositories.NewTokenRepository(db)
	tokenService := NewTokenService(key, tokenRepo)

	accessToken, err := tokenService.GenerateAccesToken("custormer", time.Duration(1800))
	if err != nil {
		t.Errorf("Expected not error, got: %v", err)
	}
	t.Logf("access token: %v", accessToken)

	_, err = tokenService.ValidateAccessToken(accessToken)
	if err != nil {
		t.Errorf("Expected not error, got: %v", err)
	}

	refreshToken, err := tokenService.GenerateRefreshToken(1, time.Now().Add(1200*time.Second))
	if err != nil {
		t.Errorf("Expected not error, got: %v", err)
	}
	t.Logf("refresh token: %v", refreshToken)

	err = tokenService.ValidateRefreshToken("okww2_p0KTiZY3jdYznguYhFKpIAiTypwluflhynLXs")
	if err != nil {
		t.Errorf("Expected not error, got: %v", err)
	}

	err = tokenService.RevokeRefreshToken("okww2_p0KTiZY3jdYznguYhFKpIAiTypwluflhynLXs")
	if err != nil {
		t.Errorf("Expected not error, got: %v", err)
	}

}

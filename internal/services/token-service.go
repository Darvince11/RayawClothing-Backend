package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secreteKey string
	tokenRepo  repositories.TokenRepository
}

func NewTokenService(secreteKey string, tokenRepo repositories.TokenRepository) *TokenService {
	return &TokenService{secreteKey: secreteKey, tokenRepo: tokenRepo}
}

func (ts *TokenService) GenerateAccesToken(role string, exp time.Duration) (string, error) {
	//Generate JWT token logic
	// Create claims which contains user insensitve info
	claims := jwt.MapClaims{
		"role": role,
		"exp":  time.Now().Add(time.Duration(exp) * time.Second).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//sign the claim with a secret key
	signedToken, err := token.SignedString([]byte(ts.secreteKey))
	return signedToken, err
}

func (ts *TokenService) ValidateAccessToken(tokenString string) (*jwt.MapClaims, error) {
	//Validate jwt
	//Extract the claims
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (any, error) {
			return []byte(ts.secreteKey), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		return nil, err
	}
	// check expiry
	if !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

func (ts *TokenService) GenerateRefreshToken(userId int, exp time.Time) (*models.RefreshToken, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)

	if err != nil {
		return nil, err
	}

	token := base64.RawURLEncoding.EncodeToString(bytes)
	tokenModel := models.RefreshToken{
		UserId:  userId,
		Token:   token,
		Expiry:  exp,
		Revoked: false,
	}
	return &tokenModel, nil
}

func (ts *TokenService) ValidateRefreshToken(token string) error {
	storedToken, err := ts.tokenRepo.GetRefreshToken(token)
	if err != nil {
		return err
	}
	if storedToken.Token != token {
		return errors.New("refresh token is invalid")
	}
	return nil
}

func (ts *TokenService) RevokeRefreshToken(token string) error {
	oldToken, err := ts.tokenRepo.GetRefreshToken(token)
	if err != nil {
		return err
	}

	revokedToken := models.RefreshToken{
		Token:      oldToken.Token,
		Expiry:     oldToken.Expiry,
		Revoked:    true,
		Created_at: time.Now(),
	}
	err = ts.tokenRepo.UpdateRefreshToken(token, &revokedToken)
	return err
}

func (ts *TokenService) StoreRefreshToken(refreshToken *models.RefreshToken) error {
	return ts.tokenRepo.AddRefreshToken(refreshToken)
}

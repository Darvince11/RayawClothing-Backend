package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	message string //change to auth repo later
}

func NewAuthService(message string) *AuthService {
	return &AuthService{message: message}
}

func (as *AuthService) GenerateJWT() (string, error) {
	//Generate JWT token logic
	// Create claims which contains user insensitve info
	claims := jwt.MapClaims{
		"username": "Kafui Dotse",
		"role":     "customer",
		"exp":      time.Now().Add(3600 * time.Second).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//sign the claim with a secret key
	signedToken, err := token.SignedString([]byte("hello"))
	return signedToken, err
}

func (as *AuthService) ValidateJWT(signedToken string) (*jwt.Token, error) {
	//Validate jwt
	//Extract the claims
	token, err := jwt.Parse(signedToken,
		func(token *jwt.Token) (any, error) {
			return []byte("hello"), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil {
		fmt.Printf("Error validating token: %v", err)
	}
	// check expiry
	if !token.Valid {
		return nil, err
	}
	return token, nil
}

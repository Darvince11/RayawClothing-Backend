package services

import (
	"errors"
	"fmt"
	"rayaw-api/internal/config"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	authRepo repositories.AuthRepository
	config   *config.Config
}

func NewAuthService(authRepo repositories.AuthRepository, config *config.Config) *AuthService {
	return &AuthService{authRepo: authRepo, config: config}
}

func (as *AuthService) GenerateJWT(username, role string, exp int64) (string, error) {
	//Generate JWT token logic
	// Create claims which contains user insensitve info
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Duration(exp) * time.Second).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//sign the claim with a secret key
	signedToken, err := token.SignedString([]byte(as.config.AuthConfig.JWTSecretKey))
	return signedToken, err
}

func (as *AuthService) ValidateJWT(signedToken string) (*jwt.Token, error) {
	//Validate jwt
	//Extract the claims
	token, err := jwt.Parse(signedToken,
		func(token *jwt.Token) (any, error) {
			return []byte(as.config.AuthConfig.JWTSecretKey), nil
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

func (as *AuthService) SignUp(signUpRequest *models.SignUpRequest) (int, error) {
	var newUser models.User
	//change the user data into database ready model
	newUser.First_name = signUpRequest.First_name
	newUser.Last_name = signUpRequest.Last_name
	newUser.Email = signUpRequest.Email
	newUser.Phone_number = signUpRequest.Phone_number
	newUser.User_password = signUpRequest.User_password

	//check if user already exits
	result, err := as.authRepo.GetUserByEmail(newUser.Email)

	if result != nil || err == nil {
		return result.Id, errors.New("user already exists")
	}

	newUserId, err := as.authRepo.AddUser(&newUser)
	return newUserId, err
}

func (as *AuthService) SaveRefreshToken(token string, userId int) error {
	return as.authRepo.AddRefreshToken(token, userId)
}

func (as *AuthService) GetRefreshToken(token string, userId int) *models.RefreshToken {
	return as.authRepo.GetRefreshTokenById(userId)
}

func (as *AuthService) GetUserByEmail(email string) (*models.User, error) {
	return as.authRepo.GetUserByEmail(email)
}

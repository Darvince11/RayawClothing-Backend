package services

import (
	"errors"
	"rayaw-api/internal/models"
	"rayaw-api/internal/repositories"
)

type AuthService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (as *AuthService) Register(signUpRequest *models.SignUpRequest) (int, error) {
	var newUser models.User
	//change the user data into database ready model
	newUser.First_name = signUpRequest.First_name
	newUser.Last_name = signUpRequest.Last_name
	newUser.Email = signUpRequest.Email
	newUser.Phone_number = signUpRequest.Phone_number
	newUser.User_password = signUpRequest.User_password

	//check if user already exits
	result, err := as.authRepo.GetUserByEmail(newUser.Email)

	if err == nil {
		return result.Id, errors.New("user already exists")
	}

	newUserId, err := as.authRepo.AddUser(&newUser)
	return newUserId, err
}

func (as *AuthService) GetUserByEmail(email string) (*models.User, error) {
	return as.authRepo.GetUserByEmail(email)
}

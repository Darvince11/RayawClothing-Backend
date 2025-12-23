package services

type AuthService struct {
	message string //change to auth repo later
}

func NewAuthService(message string) *AuthService {
	return &AuthService{message: message}
}

func (as *AuthService) GenerateJWT() {
	//Generate JWT token logic
}

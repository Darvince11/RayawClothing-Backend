package interfaces

import "github.com/golang-jwt/jwt/v5"

type AuthInterface interface {
	//returns a jwt claim if the token is valid
	//
	//returns error it's invalid
	ValidateAccessToken(tokenString string) (*jwt.MapClaims, error)
}

package models

import "time"

type RefreshToken struct {
	Id         int
	UserId     int
	Token      string
	Expiry     time.Time
	Revoked    bool
	Created_at time.Time
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

package models

import "time"

type User struct {
	Id            int
	First_name    string
	Last_name     string
	Email         string
	Phone_number  string
	User_password string
	Created_at    time.Time
}

type SignUpRequest struct {
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Email         string `json:"email"`
	Phone_number  string `json:"phone_number"`
	User_password string `json:"user_password"`
}

type LoginRequest struct {
	Email         string `json:"email"`
	User_password string `json:"user_password"`
}

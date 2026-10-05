package models

import "time"

type User struct {
	Id            int       `json:"id"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Email         string    `json:"email"`
	Phone_number  string    `json:"phone_number"`
	User_password string    `json:"user_password"`
	Created_at    time.Time `json:"created_at"`
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

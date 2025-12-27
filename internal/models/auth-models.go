package models

type User struct {
	Id            int
	First_name    string
	Last_name     string
	Email         string
	Phone_number  string
	User_password string
	Created_at    int
}

type SignUpRequest struct {
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Email         string `json:"email"`
	Phone_number  string `json:"phone_number"`
	User_password string `json:"user_password"`
}

type LoginResponse struct {
	Id            int    `json:"id"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Email         string `json:"email"`
	Phone_number  string `json:"phone_number"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_tokens"`
}

type RefreshToken struct {
	Id            int
	Refresh_Token string
	User_Id       int
}

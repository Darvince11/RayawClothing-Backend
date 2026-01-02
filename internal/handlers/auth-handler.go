package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rayaw-api/internal/models"
	"rayaw-api/internal/services"
	"rayaw-api/utils"
	"time"
)

type AuthenticationHandler struct {
	authService  *services.AuthService
	tokenService *services.TokenService
}

func NewAuthenticationHandler(authService *services.AuthService, tokenService *services.TokenService) *AuthenticationHandler {
	return &AuthenticationHandler{authService: authService, tokenService: tokenService}
}

func (ah *AuthenticationHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	type responseData struct {
		User_info     models.User `json:"user_info"`
		Access_token  string      `json:"access_token"`
		Refresh_token string      `json:"refresh_token"`
	}
	//extract the data
	var newUser models.SignUpRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Fprintf(w, "Error decoding body:%v", err)
		return
	}

	//Validate data
	if newUser.First_name == "" || newUser.Last_name == "" || newUser.Email == "" || newUser.User_password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if newUser.Phone_number == "" || len(newUser.Phone_number) != 10 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//sign up
	_, err = ah.authService.Register(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error signing user up:%v", err)
		return
	}
	//on error send error code
	user, err := ah.authService.GetUserByEmail(newUser.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error getting user:", err)
		return
	}
	//on success return generate tokens and json response
	accessToken, err := ah.tokenService.GenerateAccesToken("customer", time.Duration(utils.AccessTokenExpiry))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error generating access token:", err)
		return
	}
	refreshToken, err := ah.tokenService.GenerateRefreshToken(user.Id, time.Now().Add(time.Duration(utils.RefreshTokenExpiry)*time.Second))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error generating referesh token:", err)
		return
	}

	err = ah.tokenService.StoreRefreshToken(refreshToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error storing refresh token:", err)
		return
	}

	//response
	res := models.Response[responseData]{
		Success: true,
		Data: responseData{
			User_info:     *user,
			Access_token:  accessToken,
			Refresh_token: refreshToken.Token,
		},
		Message: "user signup in successfully",
		Error:   map[string]any{},
		Meta:    map[string]any{},
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
}

func (ah *AuthenticationHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	//Handle login
	//extract the data
	type responseData struct {
		User_info     models.User `json:"user_info"`
		Access_token  string      `json:"access_token"`
		Refresh_token string      `json:"refresh_token"`
	}

	var userLoginRequest models.LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userLoginRequest)
	if err != nil {
		fmt.Fprintf(w, "Error decoding body:%v", err)
	}

	//validate data
	if userLoginRequest.Email == "" || userLoginRequest.User_password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//login user
	user, err := ah.authService.Login(&userLoginRequest)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error loging user in:%v", err)
		return
	}

	//generate tokens
	accessToken, err := ah.tokenService.GenerateAccesToken("customer", time.Duration(utils.AccessTokenExpiry))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error generating access token:", err)
		return
	}

	refreshToken, err := ah.tokenService.GetRefreshTokenById(user.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error getting refresh token:", err)
		return
	}
	//return response
	res := models.Response[responseData]{
		Success: true,
		Data: responseData{
			User_info:     *user,
			Access_token:  accessToken,
			Refresh_token: refreshToken.Token,
		},
		Message: "user loged in successfully",
		Error:   map[string]any{},
		Meta:    map[string]any{},
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
}

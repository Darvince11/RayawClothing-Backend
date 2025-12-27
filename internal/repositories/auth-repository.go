package repositories

import (
	"database/sql"
	"rayaw-api/internal/models"
)

type AuthRepository interface {
	AddUser(user *models.User) (int, error)
	GetUserByEmail(email string) (*models.User, error)
	AddRefreshToken(token string, userId int) error
	GetRefreshTokenById(userId int) *models.RefreshToken
}

type ImplAuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &ImplAuthRepository{db: db}
}

func (ar *ImplAuthRepository) AddUser(user *models.User) (int, error) {
	query := `INSERT INTO users (first_name, last_name, email, phone_number, user_password) 
	VALUES ($1, $2, $3, $4, $5)`
	result, err := ar.db.Exec(query, user.First_name, user.Last_name, user.Email, user.Phone_number, user.User_password)
	userID, _ := result.LastInsertId()
	return int(userID), err
}

func (ar *ImplAuthRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email=$1`
	var user models.User
	err := ar.db.QueryRow(query, email).Scan(&user.Id, &user.First_name, &user.Last_name, &user.Email, &user.Phone_number, &user.User_password, &user.Created_at)
	return &user, err
}

func (ar *ImplAuthRepository) AddRefreshToken(token string, userId int) error {
	query := `INSERT INTO refresh_token (refresh_token,user_id) VALUES ($1,$2)`
	_, err := ar.db.Exec(query, token, userId)
	return err
}

func (ar *ImplAuthRepository) GetRefreshTokenById(userId int) *models.RefreshToken {
	query := `SELECT * FROM refresh_token WHERE user_id=$1`
	var refreshToken models.RefreshToken
	ar.db.QueryRow(query, userId).Scan(&refreshToken.Id, &refreshToken.Refresh_Token, &refreshToken.User_Id)
	return &refreshToken
}

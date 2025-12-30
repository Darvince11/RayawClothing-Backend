package repositories

import (
	"database/sql"
	"rayaw-api/internal/models"
)

type AuthRepository interface {
	AddUser(user *models.User) (int, error)
	GetUserByEmail(email string) (*models.User, error)
}

type ImplAuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &ImplAuthRepository{db: db}
}

func (ar *ImplAuthRepository) AddUser(user *models.User) (int, error) {
	query := `INSERT INTO users (first_name, last_name, email, phone_number, user_password) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var userID int
	err := ar.db.QueryRow(query, user.First_name, user.Last_name, user.Email, user.Phone_number, user.User_password).Scan(&userID)
	return userID, err
}

func (ar *ImplAuthRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email=$1`
	var user models.User
	err := ar.db.QueryRow(query, email).Scan(&user.Id, &user.First_name, &user.Last_name, &user.Email, &user.Phone_number, &user.User_password, &user.Created_at)
	return &user, err
}

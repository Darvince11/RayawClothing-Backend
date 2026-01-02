package repositories

import (
	"database/sql"
	"rayaw-api/internal/models"
)

type TokenRepository interface {
	AddRefreshToken(refreshToken *models.RefreshToken) error
	GetRefreshToken(refreshTokenString string) (*models.RefreshToken, error)
	UpdateRefreshToken(oldRefreshTokenString string, newRefreshToken *models.RefreshToken) error
	GetRefreshTokenByUserId(userId int) (*models.RefreshToken, error)
}

type ImplTokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) TokenRepository {
	return &ImplTokenRepository{db: db}
}

func (tr *ImplTokenRepository) AddRefreshToken(refreshToken *models.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expiry, revoked)
	VALUES ($1, $2, $3, $4)
	`
	_, err := tr.db.Exec(query, refreshToken.UserId, refreshToken.Token, refreshToken.Expiry, refreshToken.Revoked)
	return err
}

func (tr *ImplTokenRepository) GetRefreshToken(refreshTokenString string) (*models.RefreshToken, error) {
	query := `SELECT * FROM refresh_tokens WHERE refresh_token=$1`
	var refreshToken models.RefreshToken
	err := tr.db.QueryRow(query, refreshTokenString).Scan(&refreshToken.Id, &refreshToken.UserId, &refreshToken.Token, &refreshToken.Expiry, &refreshToken.Revoked, &refreshToken.Created_at)
	return &refreshToken, err
}

func (tr *ImplTokenRepository) UpdateRefreshToken(oldRefreshTokenString string, newRefreshToken *models.RefreshToken) error {
	query := `UPDATE refresh_tokens
	SET refresh_token=$1, expiry=$2, revoked=$3 created_at=$4
	WHERE refresh_token=$5
	`
	_, err := tr.db.Exec(query, newRefreshToken.Token, newRefreshToken.Expiry, newRefreshToken.Revoked, newRefreshToken.Created_at, oldRefreshTokenString)
	return err
}

func (tr *ImplTokenRepository) GetRefreshTokenByUserId(userId int) (*models.RefreshToken, error) {
	query := `SELECT * FROM refresh_tokens WHERE user_id=$1`
	var refreshToken models.RefreshToken
	err := tr.db.QueryRow(query, userId).Scan(&refreshToken.Id, &refreshToken.UserId, &refreshToken.Token, &refreshToken.Expiry, &refreshToken.Revoked, &refreshToken.Created_at)
	return &refreshToken, err
}

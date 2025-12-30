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

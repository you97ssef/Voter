package models

import "time"

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	VerifiedAt *time.Time `json:"verified_at"`
}

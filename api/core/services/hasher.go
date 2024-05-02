package services

import (
	"golang.org/x/crypto/bcrypt"
)

func NewHasher(cost int) *Hasher {
	return &Hasher{
		cost: cost,
	}
}

type Hasher struct {
	cost int
}

func (hs *Hasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hs.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (hs *Hasher) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

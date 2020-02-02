package security

import (
	"golang.org/x/crypto/bcrypt"
)

// hash password before save
func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword( []byte(pass), bcrypt.DefaultCost )
}

// when loggin in, verify password
func VerifyPassword(hashedPassword, pass string) error{
	return bcrypt.CompareHashAndPassword( []byte(hashedPassword), []byte(pass) )
}
package security

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword( []byte(pass), bcrypt.DefaultCost )
}


func VerifyPassword(hashedPassword, pass string) error{
	return bcrypt.CompareHashAndPassword( []byte(hashedPassword), []byte(pass) )
}
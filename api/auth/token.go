package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	config "github.com/vonmutinda/crafted/config"
)

// CreateToken - create a token
func CreateToken(userID uint64)(string, error){

	claims := jwt.MapClaims{
		"user_id" 		: userID,
		"exp"			: time.Now().Add(1 * time.Hour).Unix(),
		"authorized"	: true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.SECRET_KEY)
}
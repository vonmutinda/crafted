package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

// ExtractToken - 
func ExtractToken(r *http.Request)string{

	token := r.URL.Query().Get("token")

	if token != "" {
		return token
	}

	bearerToken := r.Header.Get("Authorization")
	
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

// TokenValid - is it?
func TokenValid(r *http.Request) error {

	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){ 
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { 
			return nil, fmt.Errorf("Unexpected Signing method: %v", token.Header["alg"])
		} 
		return config.SECRET_KEY, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { 
		p, _ := json.MarshalIndent(claims, "", " ") 
		fmt.Printf("claims : %v", string(p)) 
	}

	return nil
} 
package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/vonmutinda/crafted/api/auth"
	"github.com/vonmutinda/crafted/api/responses"
)

// SetUpLoggerMiddleware -log details of every request
func SetUpLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){
		log.Println( fmt.Sprintf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto) )
		next(w,r)
	}
}

// SetJSONMiddleware - header to return json response
func SetJSONMiddleware(next http.HandlerFunc) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		next(w, r)
	}
}

// SetAuthMiddleware - enforce login
func SetAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){

		err := auth.TokenValid(r)  

		if err != nil { 
			fmt.Printf("token not valid : %v",err) 
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorised"))
			return
		} 

		next(w, r)
	}
}
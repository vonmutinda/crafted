package middlewares

import (
	"fmt"
	"log"
	"net/http"

)

// log details of every request
func SetUpLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		log.Println( fmt.Sprintf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto) )
		next(w,r)
	}
}

// set header to return json response
func SetJsonMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		next(w, r)
	}
}
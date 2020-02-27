package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/vonmutinda/crafted/api/auth"
	"github.com/vonmutinda/crafted/api/responses"
	logger "github.com/vonmutinda/crafted/api/log"
)

// SetUpLoggerMiddleware -log details of every request
func SetUpLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){

		entry := fmt.Sprintf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		
		log.Println(entry)
		logger.GetLogger().Info(entry)

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
			logger.GetLogger().Errorf("submitted token not valid :%v\n",err)  
			responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
			return
		}  
		next(w, r)
	}
}
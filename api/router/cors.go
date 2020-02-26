package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// LoadCORS - 
func LoadCORS(r *mux.Router) http.Handler{

	headers := handlers.AllowedHeaders([]string{
		"X-Request",
		"Content-Type",
		"Entity",
		"Location",
		"Authorization",
	})

	methods := handlers.AllowedMethods([]string{
		"GET", "POST", "PUT", "DELETE",
	})

	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(headers, methods, origins)(r)
}
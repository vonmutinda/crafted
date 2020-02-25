package routes

import (
	"net/http"

	"github.com/vonmutinda/crafted/api/controllers"
)

var authRoutes = []Route{ 
	Route{
		URI: "/login",
		Method: http.MethodPost,
		Handler: controllers.Login,
	},
}
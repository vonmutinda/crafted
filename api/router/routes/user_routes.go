package routes

import (
	"net/http"
	"github.com/vonmutinda/crafted/api/controllers"
)

var UserRoutes = []Route{
	Route{
		Uri: "/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		Uri: "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		Uri: "/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		Uri: "/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
}
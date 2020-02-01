package routes

import (
	"net/http"
	"github.com/vonmutinda/crafted/api/controllers"

)

var articleRoutes = []Route{
	Route{
		Uri: "/",
		Method: http.MethodGet,
		Handler: controllers.GetArticles,
	},
	Route{
		Uri: "/delete",
		Method: http.MethodDelete,
		Handler: controllers.DeleteAll,
	},
}
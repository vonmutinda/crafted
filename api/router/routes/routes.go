package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/vonmutinda/crafted/api/middlewares"

)

type Route struct {
	Uri 		string
	Method 		string
	Handler		func(w http.ResponseWriter, r *http.Request)
}
 
var routes = [][]Route{ 
	articleRoutes,
	UserRoutes,
}

func Load() []Route{ 

	var AppRoutes []Route 

	for _, k := range routes{
		AppRoutes = append(AppRoutes, k...) 
	}
	
	return AppRoutes
}

// NORMALLY ;In gollira mux 
// m := mux.NewRouter()  <-- m is of type *mux.Router
// m.HandleFunc("/", Handler).Methods("GET")

func SetUpRoutesWithMiddlewares(r *mux.Router) *mux.Router{

	for _, route := range Load(){
		r.HandleFunc(
			route.Uri,
			middlewares.SetUpLoggerMiddleware(
				middlewares.SetJsonMiddleware(
					route.Handler,
				),
			), 
		).Methods(route.Method)
	}

	return r

}
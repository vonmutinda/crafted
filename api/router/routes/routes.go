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

func Load() []Route{
	routes := articleRoutes
	return routes
}

// NORMALLY ;In gollira mux 
// m := mux.NewRouter()  <-- m is of type *mux.Router
// m.HandleFunc("/", Handler).Methods("GET")

func SetUpRoutes(r *mux.Router) *mux.Router{
	for _, route := range Load(){
		r.HandleFunc(
			route.Uri,
			route.Handler,
		).Methods(route.Method)
	}
	return r
}

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
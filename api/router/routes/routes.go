package routes

import (
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/vonmutinda/crafted/api/middlewares"

)

// Route struct
type Route struct {
	URI 		string
	Method 		string
	Handler		func(w http.ResponseWriter, r *http.Request)
}
 
var routes = [][]Route{ 
	articleRoutes,
	userRoutes,
	authRoutes,
}

func load() []Route{ 

	var AppRoutes []Route 

	for _, k := range routes{
		AppRoutes = append(AppRoutes, k...) 
	}

	return AppRoutes
}


// SetUpRoutesWithMiddlewares NORMALLY ;In gollira mux 
// m := mux.NewRouter()  <-- m is of type *mux.Router
// m.HandleFunc("/", Handler).Methods("GET")
func SetUpRoutesWithMiddlewares(r *mux.Router) *mux.Router{

	for _, route := range load(){
		r.HandleFunc(
			route.URI,
			middlewares.SetUpLoggerMiddleware(
				middlewares.SetJsonMiddleware(
					route.Handler,
				),
			), 
		).Methods(route.Method)
	}

	return r

}
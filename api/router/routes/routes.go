package routes

import (
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/vonmutinda/crafted/api/middlewares"

)

// Route struct
type Route struct {
	URI 			string
	Method 			string
	Handler			func(w http.ResponseWriter, r *http.Request)
	AuthRequired 	bool
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

		// if you know of a better way to Use the auth middleware
		// kindly let me know
		if route.AuthRequired {

			r.HandleFunc(
				route.URI,
				middlewares.SetUpLoggerMiddleware(
					middlewares.SetJSONMiddleware(
						middlewares.SetAuthMiddleware(route.Handler),
					),
				), 
			).Methods(route.Method)

		}else {

			r.HandleFunc(
				route.URI,
				middlewares.SetUpLoggerMiddleware(
					middlewares.SetJSONMiddleware(route.Handler),
				),
			).Methods(route.Method)
		}
 
	}

	return r
}
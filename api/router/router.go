package router

import (
	"github.com/gorilla/mux" 
	"github.com/vonmutinda/crafted/api/router/routes"

)

func New() *mux.Router{
	r := mux.NewRouter().StrictSlash(true)  
	return routes.SetUpRoutesWithMiddlewares(r)
}
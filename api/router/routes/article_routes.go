package routes

import (
	"net/http"
	"github.com/vonmutinda/crafted/api/controllers"
)

var articleRoutes = []Route{
	Route{
		Uri: "/articles",
		Method: http.MethodGet,
		Handler: controllers.GetArticles,
	},
	Route{
		Uri: "/articles",
		Method: http.MethodPost,
		Handler: controllers.CreateArticle,
	},
	Route{
		Uri: "/articles/{id}",
		Method: http.MethodGet,
		Handler: controllers.FetchArticleByID,
	}, 
	Route{
		Uri: "/articles/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteArticleByID,
	},
	// Route{
	// 	Uri: "/delete",
	// 	Method: http.MethodDelete,
	// 	Handler: controllers.DeleteAll,
	// },
}
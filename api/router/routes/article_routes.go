package routes

import (
	"net/http"
	"github.com/vonmutinda/crafted/api/controllers"
)

var articleRoutes = []Route{
	Route{
		URI: "/articles",
		Method: http.MethodGet,
		Handler: controllers.GetArticles,
	},
	Route{
		URI: "/articles",
		Method: http.MethodPost,
		Handler: controllers.CreateArticle,
	},
	Route{
		URI: "/articles/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateArticle,
	},
	Route{
		URI: "/articles/{id}",
		Method: http.MethodGet,
		Handler: controllers.FetchArticleByID,
	}, 
	Route{
		URI: "/articles/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteArticleByID,
	},
	Route{
		URI: "/delete",
		Method: http.MethodDelete,
		Handler: controllers.DeleteAll,
	},

	Route{
		URI: "/delete/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteArticleByID,
	},
}
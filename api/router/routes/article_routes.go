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
		AuthRequired: false,
	},
	Route{
		URI: "/articles",
		Method: http.MethodPost,
		Handler: controllers.CreateArticle,
		AuthRequired: true,
	},
	Route{
		URI: "/articles/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateArticle,
		AuthRequired: true,
	},
	Route{
		URI: "/articles/{id}",
		Method: http.MethodGet,
		Handler: controllers.FetchArticleByID,
		AuthRequired: false,
	}, 
	Route{
		URI: "/articles/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteArticleByID,
		AuthRequired: true,
	},
	Route{
		URI: "/delete",
		Method: http.MethodDelete,
		Handler: controllers.DeleteAll,
		AuthRequired: true,
	},

	Route{
		URI: "/delete/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteArticleByID,
		AuthRequired: false,
	},
}
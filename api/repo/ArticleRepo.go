package repo

import (
	"github.com/vonmutinda/crafted/api/models"

)

type ArticlesRepo interface {
	GetAllArticles()([]models.Article, error)
	SaveArticle(models.Article) (models.Article, error)
	// UpdateArticle(models.Article) (models.Article, error)
	// DeleteArticle() error
	// DeleteAll() error
}


package repo

import (
	"github.com/vonmutinda/crafted/api/models"

)

type ArticlesRepo interface {
	GetAllArticles()([]models.Article, error)
	SaveArticle(models.Article) (models.Article, error)
	FindByID(id uint64) (models.Article, error)
	DeleteByID(id uint64) (int64, error)
	// UpdateArticle(models.Article) (models.Article, error) 
	DeleteAllArticles() (int64, error)
}


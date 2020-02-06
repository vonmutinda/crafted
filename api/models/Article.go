package models

import (
	"errors"
	"html"
	"strings"
	"time" 

)


type Article struct {
	ID				uint64  	`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title			string		`gorm:"size:250;not_null;unique" json:"title"`
	Body			string		`gorm:"size:500;" json:"body"`
	AuthorID		uint64		`gorm:"not_null" json:"author_id"`
	Author 			User		`gorm:"foreignkey:AuthorID" json:"author"`
	CreatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` 
	UpdatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// prepare 
func (a *Article) Prepare(){ 
	a.ID = 0
	a.Title = html.EscapeString(strings.TrimSpace(a.Title))
	a.Body = html.EscapeString(strings.TrimSpace(a.Body))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

// validate article
func (a *Article) Validate() error {
	if a.Title == "" {
		return errors.New("Title Required")
	}
	if a.AuthorID < 0 {
		return errors.New("Author required")
	}
	if a.Body == ""{
		return errors.New("Article body required")
	}
	// HOWEVER, i think body can be null.
	// incomplete article == draft
	return nil
}


type ArticlesRepo interface {
	GetAllArticles()([]Article, error)
	SaveArticle(Article) (Article, error)
	FetchArticleByID(id uint64) (Article, error)
	DeleteByID(id uint64) (int64, error)
	// UpdateArticle(models.Article) (models.Article, error) 
	DeleteAllArticles() (int64, error)
}

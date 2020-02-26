package models

import (
	"html"
	"strings"
	"time"
	
	"github.com/go-playground/validator/v10"
)

// Article struct
type Article struct {
	ID				uint64  	`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title			string		`gorm:"size:250;not_null;unique" json:"title" validate:"required"`
	Body			string		`gorm:"size:500;" json:"body"`
	AuthorID		uint64		`gorm:"not_null" json:"author_id" validate:"required"`
	Author 			User		`gorm:"foreignkey:AuthorID" json:"author"`
	CreatedAt		time.Time	`json:"created_at"` 
	UpdatedAt		time.Time	`json:"updated_at"` 
	DeletedAt		*time.Time	`json:"deleted_at,omitempty" sql:"index"`
}

// Prepare func
func (a *Article) Prepare(){  
	a.Title = html.EscapeString(strings.TrimSpace(a.Title))
	a.Body = html.EscapeString(strings.TrimSpace(a.Body))  
}

// Validate cooler one
func (a *Article) Validate() error {
	v := validator.New();  
	return v.StructPartial(a,"Title", "AuthorID")
}

// ArticleInterface -
type ArticleInterface interface {
	GetAllArticles()([]Article, error)
	SaveArticle(a *Article) (*Article, error)
	FetchArticleByID(id uint64) (Article, error)
	DeleteByID(id uint64, uid uint64) (int64, error)
	UpdateArticle(a *Article, id int64, uid uint64)(int64, error) 
	DeleteAllArticles() (int64, error)
}

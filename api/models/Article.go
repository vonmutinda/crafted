package models

import ( 
	"html"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v9"
)


type Article struct {
	ID				uint64  	`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title			string		`gorm:"size:250;not_null;unique" json:"title" validate:"required"`
	Body			string		`gorm:"size:500;" json:"body"`
	AuthorID		uint64		`gorm:"not_null" json:"author_id" validate:"required`
	Author 			User		`gorm:"foreignkey:AuthorID" json:"author"`
	CreatedAt		time.Time	`json:"created_at"` 
	UpdatedAt		time.Time	`json:"updated_at"` 
	DeletedAt		*time.Time	`json:"deleted_at,omitempty" sql:"index"`
}

// prepare 
func (a *Article) Prepare(){ 
	a.ID = 0
	a.Title = html.EscapeString(strings.TrimSpace(a.Title))
	a.Body = html.EscapeString(strings.TrimSpace(a.Body))
	a.CreatedAt = time.Now() 
}

// cooler validator
func (a *Article) Validate() error {
	v := validator.New()  
	// return v.Struct(a)
	return v.StructPartial(a,"title", "author_id")
}



type ArticlesRepo interface {
	GetAllArticles()([]Article, error)
	SaveArticle(Article) (Article, error)
	FetchArticleByID(id uint64) (Article, error)
	DeleteByID(id uint64) (int64, error)
	UpdateArticle(a Article, id int64)(int64, error) 
	DeleteAllArticles() (int64, error)
}

package models

import ( 
	"time"
	"errors"
	"strings"
	"html"
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

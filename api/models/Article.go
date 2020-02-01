package models

import (
	"fmt"
	"time"
)


type Article struct {
	ID				uint32  	`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title			string		`gorm:"size:250;not_null;unique" json:"title"`
	Body			string		`gorm:"size:500;" json:"body"`
	Author			User		`gorm:"foreignkey:ID" json:"author"`
	CreatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` 
	UpdatedAt		time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

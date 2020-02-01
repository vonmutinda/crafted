package models

import (
	"time"
	"log"
	"github.com/vonmutinda/crafted/api/security" 
)

type User struct {
	ID 			uint32 		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Nickname 	string 		`gorm:"size:20;not null,unique" json:"nickname"`
	Email		string 		`gorm:"size:20;not null,unique" json:"email"`
	Password	string 		`gorm:"size:60;not null" json:"password"`
	CreatedAt	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// hash the password 
func (u *User) BeforeSave() error {
	hashedPass, err := security.Hash(u.Password)
	if err != nil {
		log.Println("Error hashing password :", err)
		return err
	}

	u.Password = string(hashedPass)
	return nil
}
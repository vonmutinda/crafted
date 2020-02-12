package models

import ( 
	"html"
	"log"
	"strings"
	"time"
	
	"github.com/go-playground/validator/v10"
	"github.com/vonmutinda/crafted/api/security"
)

type User struct {
	ID 			uint64 		`gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Nickname 	string 		`gorm:"size:20;not null,unique" json:"nickname" validate:"required"`
	Email		string 		`gorm:"size:20;not null,unique" json:"email" validate:"required,email"`
	Password	string 		`gorm:"size:60;not null" json:"-" validate:"required"`
	CreatedAt	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Articles	[]Article	`json:"articles,omitempty"`
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

// prepare 
func (u *User) Prepare(){
	u.ID = 0
	u.Nickname = html.EscapeString( strings.TrimSpace(u.Nickname) )
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// cool validator
func (u *User) Validate() error{
	v := validator.New()
	return v.Struct(u)
}

// validate before save
// func (u *User) Validate(action string) error{

// 	switch strings.ToLower(action){
// 		case "update":
// 			if u.Nickname == ""{
// 				return errors.New("Required Nickname")
// 			} 		
// 			if u.Email == ""{
// 				return errors.New("Required Email")
// 			}
		
// 			if err := checkmail.ValidateFormat(u.Email); err != nil {
// 				return errors.New("Invalid Email")
// 			}
// 			return nil
// 		default:
// 			if u.Nickname == ""{
// 				return errors.New("Required Nickname")
// 			}
		
// 			if u.Password == ""{
// 				return errors.New("Required Password")
// 			}
		
// 			if u.Email == ""{
// 				return errors.New("Required Email")
// 			}
		
// 			if err := checkmail.ValidateFormat(u.Email); err != nil {
// 				return errors.New("Invalid Email")
// 			}
// 			return nil

// 	}
	 
// }



type UsersRepo interface {
	Save(User) (User, error)
	FindAll() ([]User, error)
	FindById(uint64)(User, error)
	// Update(uint32, models.User) (uint64, error)
	// Delete(uint32) (uint64, error)
	// DeleteAll() (error)
}
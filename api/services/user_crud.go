package services 

import (
	"errors" 

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/vonmutinda/crafted/api/channels"
	"github.com/vonmutinda/crafted/api/models"
)

// UserService struct
type UserService struct {
	Logger 	*logrus.Logger
	DB 		*gorm.DB
}

// Save method of UserRepository interface
func (u *UserService) Save(user *models.User) (*models.User, error){
	
	var err error
	done := make( chan bool)   

	go func(ch chan<- bool){    
		
		gor := u.DB.Save(user)

		if err = gor.Error; err != nil { 
			u.Logger.Errorf("cannot insert new user record : %v", gor.Error)
			ch<- false
			return
		}  
		ch<- true

	}(done)

	if channels.OK(done){ 
		return user, nil
	} 
	return &models.User{}, err
}

// FindAll Fetch all the Users
func (u *UserService) FindAll() ([]models.User, error){ 
	
	var err error 

	users := []models.User{}
	// a goroutine (channel) for fetching records
	done := make( chan bool) 
	go func(ch chan<- bool){

		gor := u.DB.Raw(`
			SELECT * FROM users LIMIT 50
		`).Scan(&users)

		if err = gor.Error; err != nil {
			u.Logger.Errorf("cannot find all users : %v", err) 
			ch<- false
			return
		} 

		ch<- true 
	}(done)

	if channels.OK(done){
		return users, nil
	}
	return nil, err
}

// FindUserByID Fetch user by id
func (u *UserService)FindUserByID(uid uint64) ( models.User, error){

	var err error  
	user := models.User{}

	// a goroutine (channel) for fetching records
	done := make(chan bool) 
	go func(ch chan<- bool){ 

		gor := u.DB.Raw(`
			SELECT * FROM users WHERE id=?
			`, uid,
		).Scan(&user) 

		if err = gor.Error; err != nil {
			u.Logger.Errorf("cannot fetch user by id %d : %v", uid, gor.Error)
			ch<- false
			return  
		}
		ch<- true

	}(done)

	if channels.OK(done) {
		return user, nil
	}

	if gorm.IsRecordNotFoundError(err){
		return models.User{}, errors.New("User not found")
	}

	return models.User{}, err
}
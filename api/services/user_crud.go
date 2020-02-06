package services 

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/vonmutinda/crafted/api/channels"
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/models"
)

type UserCRUD struct {}

// Save method of UserRepository interface
func (r *UserCRUD) Save(user models.User) (models.User, error){
	var err error 

	done := make( chan bool)

	go func(ch chan<- bool){
		if err = database.GetDB().Debug().Model(&models.User{}).Create(&user).Error; err != nil {
			ch<- false
			return 
		}
		ch<- true
	}(done)

	if channels.OK(done){ 
		return user, nil
	} 
	return models.User{}, err
}

// Fetch all the Users
func (r *UserCRUD) FindAll() ([]models.User, error){
	var err error

	users := []models.User{}
	// a goroutine (channel) for fetching records
	done := make( chan bool) 
	go func(ch chan<- bool){
		if err = database.GetDB().Model(&models.User{}).Limit(50).Find(&users).Error; err != nil {
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

// Fetch all the Users
func (r *UserCRUD) FindById(uid uint64) ( models.User, error){
	var err error

	user := models.User{}
	// a goroutine (channel) for fetching records
	done := make(chan bool) 
	go func(ch chan<- bool){
		if err = database.GetDB().Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error; err != nil {
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
package crud 

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/vonmutinda/crafted/api/models" 
)

type repoUsersCrud struct {
	db *gorm.DB
}

// returns pointer to a db connection pool/repository
func NewUserCrud(db *gorm.DB) *repoUsersCrud {
	return &repoUsersCrud{db}
}

// Save method of UserRepository interface
func (r *repoUsersCrud) Save(user models.User) (models.User, error){
	var err error

	// a goroutine (channel) for saving to db
	done := make( chan bool) 
	go func(ch chan<- bool){
		if err = r.db.Debug().Model(&models.User{}).Create(&user).Error; err != nil {
			ch<- false
			return 
		}
		ch<- true
	}(done)

	if <-done == true {
		return user, nil
	}
	return models.User{}, err
}

// Fetch all the Users
func (r *repoUsersCrud) FindAll() ([]models.User, error){
	var err error

	users := []models.User{}
	// a goroutine (channel) for fetching records
	done := make( chan bool) 
	go func(ch chan<- bool){
		if err = r.db.Debug().Model(&models.User{}).Limit(50).Find(&users).Error; err != nil {
			ch<- false
			return 
		}
		ch<- true
	}(done)

	if <-done == true{
		return users, nil
	}
	return nil, err
}

// Fetch all the Users
func (r *repoUsersCrud) FindById(uid uint64) ( models.User, error){
	var err error

	user := models.User{}
	// a goroutine (channel) for fetching records
	done := make( chan bool) 
	go func(ch chan<- bool){
		if err = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error; err != nil {
			ch<- false
			return 
		}
		ch<- true
	}(done)

	if <-done == true {
		return user, nil
	}

	if gorm.IsRecordNotFoundError(err){
		return models.User{}, errors.New("User not found")
	}

	return models.User{}, err
}
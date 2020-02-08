package database

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/vonmutinda/crafted/config"
	_"github.com/jinzhu/gorm/dialects/postgres"

)

var (
	err error
	
	db *gorm.DB
)

// connect to db
func Connect() error  {

	db, err = gorm.Open(config.DB_DRIVER, config.DB_URL)

	if err != nil {
		log.Println("Error connecting to db")
		return err
	}

	return nil
}

func GetDB() *gorm.DB{
	return db
}
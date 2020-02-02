package database

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/vonmutinda/crafted/config"
	_"github.com/jinzhu/gorm/dialects/postgres"

)

// connect to db
func Connect() (*gorm.DB, error)  {
	db, err := gorm.Open(config.DB_DRIVER, config.DB_URL)

	if err != nil {
		log.Println("Error connecting to db")
		return nil, err
	}

	return db, nil
}
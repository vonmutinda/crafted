package database

import (
	"github.com/jinzhu/gorm"
	"github.com/vonmutinda/crafted/config"
	"github.com/vonmutinda/crafted/api/log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/go-sql-driver/mysql"
)

var (
	err error
	
	db *gorm.DB
)

// Connect to db
func init()  {

	db, err = gorm.Open(config.DB_DRIVER, config.DB_URL)

	if err != nil { 
		 log.GetLogger().Errorf("cannot connect to db :%v",err)
	}
 
}

// GetDB returns db conn
func GetDB() *gorm.DB{ 
	return db
}
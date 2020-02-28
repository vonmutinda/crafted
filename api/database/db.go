package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    //mssql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql dialect
	_ "github.com/jinzhu/gorm/dialects/postgres" //psql dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //sqlite dialect

	"github.com/vonmutinda/crafted/api/log"
	"github.com/vonmutinda/crafted/config"
)

var (
	err error

	db *gorm.DB
)

// Connect to db
func init() {
	db, err = gorm.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		log.GetLogger().Errorf("cannot connect to db :%v", err)
	} else {

	}

}

// GetDB returns db conn
func GetDB() *gorm.DB {
	return db
}

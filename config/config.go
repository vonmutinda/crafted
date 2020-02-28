package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonmutinda/crafted/api/log"
)

var (
	PORT       = ""
	DB_URL     = ""
	DB_DRIVER  = ""
	SECRET_KEY []byte
	DB_HOST    string
)

// load necessary configurations
func init() {

	if err := godotenv.Load(); err != nil {
		log.GetLogger().Errorf("cannot load .env file : %v", err)
		PORT = ":9000"
	}
	SECRET_KEY = []byte(os.Getenv("API_SECRET"))
	PORT = os.Getenv("PORT")
	DB_DRIVER = os.Getenv("DB_DRIVER")
	if os.Getenv("DB_HOST") == "" {
		DB_HOST = "127.0.0.1"
	}

	log.GetLogger().Info(DB_DRIVER)
	switch driver := DB_DRIVER; driver {
	case "postgres":
		DB_URL = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		)

	case "mysql":
		DB_URL = fmt.Sprintf(
			"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	}

}

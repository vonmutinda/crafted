package config 

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonmutinda/crafted/api/log"
)

var ( 
	PORT = "" 
	DB_URL = "" 
	DB_DRIVER = ""
	SECRET_KEY []byte
	DB_HOST string
)

// load necessary configurations
func init(){ 

	if err := godotenv.Load(); err != nil {
		log.GetLogger().Errorf("cannot load .env file : %v", err) 
		PORT = ":9000"
	}

	PORT = os.Getenv("PORT")

	if os.Getenv("DB_HOST") == ""{
		DB_HOST = "127.0.0.1"
	}

	// postgresql 
	// DB_URL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PASSWORD"),
	// ) 

	// mysql
	DB_URL = fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_HOST"),os.Getenv("DB_NAME")) 

		
	
	DB_DRIVER = os.Getenv("DB_DRIVER")

	SECRET_KEY = []byte(os.Getenv("API_SECRET")) 	
}
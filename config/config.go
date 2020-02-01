package config 

import (
	"github.com/joho/godotenv"
	"os"
	"fmt"
)

var (
	PORT =""
	DB_URL =""
	DB_DRIVER =""
)

func Load(){ 
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error %s",err)
		PORT = ":9000"
	}
	PORT = os.Getenv("PORT")

	DB_URL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)
	DB_DRIVER=os.Getenv("DB_DRIVER")
}
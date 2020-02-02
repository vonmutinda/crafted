package auto

import (
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/models"
	"log"

)

// Auto-migrate models
func Load(){
	db , err := database.Connect() 

	if err != nil {
		log.Fatal(err)
	}

	// Article
	if err := db.Debug().AutoMigrate(&models.Article{}).Error; err !=nil {
		log.Println("error migrating Article:", err)
	}

	// User
	if err := db.Debug().AutoMigrate(&models.User{}).Error; err !=nil {
		log.Println("error migrating User:", err)
	}

	defer db.Close()
}
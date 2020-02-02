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

	// drop exisiting tables first
	if err := db.Debug().DropTableIfExists(&models.Article{}, &models.User{}).Error; err != nil{
		log.Println(err)
	}

	// Article
	if err := db.Debug().AutoMigrate(&models.User{}, &models.Article{}).Error; err !=nil {
		log.Println("error migrating Article:", err)
	}

	// relationship
	err = db.Debug().Model(&models.Article{}).AddForeignKey("author_id","users(id)", "cascade", "cascade").Error 

	if err !=nil{
		log.Println("error creating relations:", err)
	}

	// load dummy data
	for i, _ := range users{
		if err = db.Debug().Model(&models.User{}).Create(&users[i]).Error; err != nil {
			log.Println("error adding dummy user", err)
		}

		articles[i].AuthorID = users[i].ID
		if err = db.Debug().Model(&models.Article{}).Create(&articles[i]).Error; err != nil {
			log.Println("error adding dummy article", err)
		}

		if err = db.Debug().Model(&articles[i]).Related(&articles[i].Author).Error; err != nil {
			log.Println("error relating ", err)
		}
	} 

	defer db.Close()
}
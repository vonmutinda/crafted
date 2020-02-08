package auto

import (
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/models"
	"log"

)

// Auto-migrate models
func Load(){  

	if err := database.Connect(); err != nil{
		log.Println(err)
	}

	if err := database.GetDB().DropTableIfExists(&models.Article{}, &models.User{}).Error; err != nil{
		log.Println(err)
	}

	// migrate
	if err := database.GetDB().AutoMigrate(&models.User{}, &models.Article{}).Error; err !=nil {
		log.Println("error migrating Article:", err)
	}

	// relationship
	database.GetDB().Model(&models.Article{}).AddForeignKey("author_id","users(id)", "cascade", "cascade")

	// load dummy data
	for i, _ := range users{
		database.GetDB().Model(&models.User{}).Create(&users[i])

		articles[i].AuthorID = users[i].ID
		database.GetDB().Model(&models.Article{}).Create(&articles[i])
	}  
}
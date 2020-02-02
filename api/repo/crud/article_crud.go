package crud

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/vonmutinda/crafted/api/models"

)

type repoArticleCrud struct {
	db *gorm.DB
}

func NewArticleCrud(db *gorm.DB) *repoArticleCrud{
	return &repoArticleCrud{db:db}
}

func (repo *repoArticleCrud) GetAllArticles() ([]models.Article, error){
	var err error

	articles := []models.Article{}

	done := make(chan bool)

	go func(c<-chan bool){
		err := repo.db.Debug().Model(&models.Article{}).Find(&articles).Error
		if err != nil {
			log.Println("Error fetching records :", err)
			done <- false
			return
		}
		done<- true 
	}(done)

	if <-done == true {
		return articles,nil
	}
	return nil, err
}

// save a new article 
func (repo *repoArticleCrud) SaveArticle(article models.Article) (models.Article, error){
	var err error
	// goroutine for saving
	done := make(chan bool)
	go func(c <-chan bool){
		if err := repo.db.Debug().Model(&models.Article{}).Create(&article).Error; err != nil { 
			log.Println(err) 
			done <- false
			return
		} 
		done <- true 
	}(done)

	if <-done == true {
		return article, nil
	}

	return models.Article{}, err
}
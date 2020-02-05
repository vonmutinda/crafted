package crud

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/vonmutinda/crafted/api/channels"
	"github.com/vonmutinda/crafted/api/models"
)

type ArticleCRUD struct {}

func (repo *ArticleCRUD) GetAllArticles() ([]models.Article, error){
	var err error

	// articles slice
	articles := []models.Article{}

	// channels
	done := make(chan bool)

	// go routine to fetch articles
	go func(c chan<- bool){
		if err := repo.db.Debug().Model(&models.Article{}).Find(&articles).Error; err != nil {
			log.Println("Error fetching records :", err)
			c<- false
			return
		}

		// append apropriate author for post before response
		if len(articles) > 0 {
			for i, _ := range articles{
				err = repo.db.Debug().Model(&models.User{}).Where("id = ?", articles[i].AuthorID).Take(&articles[i].Author).Error
				if err != nil {
					c<- false
					return  
				}
			}
		}  
		c<- true 
	}(done) 

	if channels.OK(done){
		close(done)
		return articles,nil 
	}
	return nil, err
}

// save a new article 
func (repo *ArticleCRUD) SaveArticle(article models.Article) (models.Article, error){
	var err error
	// goroutine for saving
	done := make(chan bool)
	go func(c chan<- bool){
		err = repo.db.Debug().Model(&models.Article{}).Create(&article).Error 
		if err != nil {  
			c<- false
			return
		} 
		
		err = repo.db.Debug().Model(&models.User{}).Where("id = ?", article.AuthorID).Take(&article.Author).Error
		if err != nil {
			log.Println("Error associating author", err)
			c<- false
		}
		c<- true 
	}(done)

	if channels.OK(done){
		close(done)
		return articles,nil 
	}
	return models.Article{}, err
}

// delete all articles 
func (repo *ArticleCRUD) DeleteAllArticles() (int64, error){
	// var err error 
	var rep *gorm.DB
	done := make(chan int, 1)

	go func(c chan<- int){  
		rep = repo.db.Debug().Model(&models.Article{}).Delete(&models.Article{})
		c<- 1
	}(done)
	
	<-done

	return rep.RowsAffected, rep.Error
}

// find article by id 
func (repo *ArticleCRUD) FindByID(id uint64) (models.Article, error){ 

	var err error 
	article := models.Article{}  
	done := make(chan bool)

	go func(c chan<- bool){
		if err = repo.db.Debug().Model(&models.Article{}).Where("ID = ?", id).Take(&article).Error; err != nil {
			log.Println(err)
			c<- false
			return
		} 
		c<- true
	}(done)

	if channels.OK(done){ 
		return article, nil
	}

	if gorm.IsRecordNotFoundError(err){
		return models.Article{}, errors.New("Article not found")
	}

	return models.Article{}, err	
}

// delete article by ID
func (repo *ArticleCRUD) DeleteByID(id uint64) (int64, error){ 
	var rep *gorm.DB

	done := make(chan bool) 
	go func(c chan<- bool){ 
		rep = repo.db.Debug().Model(&models.Article{}).Where("id = ?", id).Delete(&models.Article{})
		c<- true
	}(done)

	if channels.OK(done){ 
		return rep.RowsAffected, rep.Error
	}
	return 0, rep.Error
}
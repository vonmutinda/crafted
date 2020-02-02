package controllers

import (
	"net/http"
	"encoding/json" 
	"log"
	"github.com/vonmutinda/crafted/api/repo"
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/repo/crud" 
	"github.com/vonmutinda/crafted/api/responses"

)

// Create new article
func CreateArticle(w http.ResponseWriter, r *http.Request){
	// 1. db connection 
	db, err := database.Connect() 
	if err != nil{
		log.Println(err)
	}

	// 2. article instance + decode json to struct
	article := models.Article{} 
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil{
		responses.ERROR(w,http.StatusBadRequest, err)
	}
	// 3. instance of article repo 
	rep := crud.NewArticleCrud(db) 

	// 4. save article
	func (re repo.ArticlesRepo){
		a, e := re.SaveArticle(article)

		if err != nil{
			responses.ERROR(w, http.StatusInternalServerError, e)
		}
		responses.JSON(w, http.StatusOK, a)

	}(rep)
}

// Fetch all articles
func GetArticles(w http.ResponseWriter, r *http.Request){
	// 1. create db connection  
	db, err := database.Connect()

	if err != nil {
		log.Println(err)
	}

	// 2. article repo instance
	rep := crud.NewArticleCrud(db) 

	// 3. call GetAllArticlesMethod 
	func (re repo.ArticlesRepo){
		a, e := rep.GetAllArticles()
		if e != nil{
			log.Println(e)
			responses.ERROR(w, http.StatusInternalServerError, e)
		}  
		responses.JSON(w, http.StatusOK, a)
		
	}(rep)
}

// Delete all articles
func DeleteAll(w http.ResponseWriter, r *http.Request){
	// 1. db connect 
		db, err := database.Connect()
		if err != nil {
			log.Println("Error connecting to db", err)
		}
	// 2. instantiate repo
		rep := crud.NewArticleCrud(db)

	// 3. call delete all
		func (repo repo.ArticlesRepo){
			if err := rep.DeleteAllArticles(); err != nil {
				log.Println(err)
				responses.ERROR(w, http.StatusInternalServerError, err)
			}
			responses.JSON(w, http.StatusOK, struct{Status string `json:"status"`}{Status: "OK! Deleted!"})

		}(rep)
}
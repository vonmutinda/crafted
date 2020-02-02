package controllers

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/vonmutinda/crafted/api/repo"
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/repo/crud"

)

func CreateArticle(w http.ResponseWriter, r *http.Request){
	// 1. db connection 
	db, err := database.Connect() 
	if err != nil{
		log.Println(err)
	}

	// 2. decode response
	article := models.Article{}
	json.NewDecoder(r.Body).Decode(&article) 
	log.Println("article:",article)

	// 3. save data 
	rep := crud.NewArticleCrud(db) 

	func (re repo.ArticlesRepo){
		re.SaveArticle(article)

		if err := json.NewEncoder(w).Encode(article); err != nil {
			log.Println(err)
		}

	}(rep)
}

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
		} 

		if err := json.NewEncoder(w).Encode(a); err != nil{
			log.Println(err)
		}
		
	}(rep)
}

func DeleteAll(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("All articles"))
}
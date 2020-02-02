package controllers

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json" 
	"github.com/gorilla/mux"
	"github.com/vonmutinda/crafted/api/repo"
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/responses"
	"github.com/vonmutinda/crafted/api/repo/crud" 
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
		responses.ERROR(w,http.StatusUnprocessableEntity, err)
		return
	}
	// 3. instance of article repo 
	rep := crud.NewArticleCrud(db) 

	// 4. validate and save article
	article.Prepare()
	if err = article.Validate(); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	func (re repo.ArticlesRepo){
		a, err := re.SaveArticle(article) 
		if err != nil{
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, a)
		return
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
			responses.ERROR(w, http.StatusUnprocessableEntity, e)
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
			// ra means ===> rows affected
			ra, err := rep.DeleteAllArticles()
			if err != nil {
				log.Println(err)
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
			} 
			responses.JSON(
				w, 
				http.StatusOK, 
				struct{
					Status string `json:"status"`
				}{
					Status: fmt.Sprintf("OK %d Records Deleted!", ra),
				},
			)
		}(rep)
}

// find by id 
func FetchArticleByID(w http.ResponseWriter, r *http.Request){
	//1. connect to db 
	db, err := database.Connect()

	if err != nil {
		log.Println("Error connecting to db",err)
	}

	// 2. Fetch id from url 
	vars := mux.Vars(r) 
	log.Println(vars)
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// 3. instantiate repo
	rep := crud.NewArticleCrud(db)

	// 4. find the record and respond
	func (repo repo.ArticlesRepo){
		article, err := repo.FindByID(id) 
		if err != nil {
			log.Println(err)
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}

		responses.JSON(w, http.StatusOK, article)
	}(rep)
}

// delete article by id 
func DeleteArticleByID(w http.ResponseWriter, r *http.Request){
	//1. connect to db 
	db, err := database.Connect()

	if err != nil {
		log.Println("Error connecting to db",err)
	}

	// 2. Fetch id from url 
	vars := mux.Vars(r) 
	log.Println(vars)
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// 3. instantiate repo
	rep := crud.NewArticleCrud(db)

	// 4. find the record and respond
	func (repo repo.ArticlesRepo){
		ra, err := repo.DeleteByID(id) 
		if err != nil {
			log.Println(err)
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}

		responses.JSON(w, http.StatusOK, 
			struct{
				Status string `json:"status"`
			}{
				Status: fmt.Sprintf("OK %d Records Deleted!", ra),
			},
		)
	}(rep)
}
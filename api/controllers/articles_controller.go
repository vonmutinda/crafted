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
 
	article := models.Article{}

	if err := json.NewDecoder(r.Body).Decode(&article); err != nil{
		responses.ERROR(w,http.StatusUnprocessableEntity, err)
		return
	} 
 
	article.Prepare()
	article.Validate()

	if err := article.Validate(); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	func (re models.ArticlesRepo){

		a, err := re.SaveArticle(article) 
		if err != nil{
			responses.ERROR(w, http.StatusUnprocessableEntity, err) 
		}
		responses.JSON(w, http.StatusCreated, a)
 	}(rep)
}

// Fetch all articles
func GetArticles(w http.ResponseWriter, r *http.Request){  

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

		func (repo models.ArticlesRepo){ 
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
	db := database.GetDB()

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
	db := database.GetDB()

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
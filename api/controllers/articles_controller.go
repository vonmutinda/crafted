package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/log"
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/responses"
	"github.com/vonmutinda/crafted/api/services"
)

// CreateArticle new article
func CreateArticle(w http.ResponseWriter, r *http.Request){ 
 
	article := new(models.Article)

	if err := json.NewDecoder(r.Body).Decode(article); err != nil{
		responses.ERROR(w,http.StatusUnprocessableEntity, err)
		return
	} 
 
	article.Prepare()
	article.Validate()

	if err := article.Validate(); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	service := &services.ArticleService{ 
		Logger: log.GetLogger(),
		DB: database.GetDB(),
	}

	func (servc models.ArticleInterface){ 
		a, err := servc.SaveArticle(article) 
		if err != nil{
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return 
		}
		responses.JSON(w, http.StatusCreated, a)
 	}(service)
}

// GetArticles --FetchAll articles
func GetArticles(w http.ResponseWriter, r *http.Request){  

	service := &services.ArticleService{
		Logger: log.GetLogger(),
		DB:  database.GetDB(),
	}

	func (servc models.ArticleInterface){
		a, e := servc.GetAllArticles()
		if e != nil{
			log.GetLogger().Info("GET/articles", e)
			responses.ERROR(w, http.StatusUnprocessableEntity, e)
			return
		}  
		responses.JSON(w, http.StatusOK, a)
		
	}(service)
}

// DeleteAll - articles
func DeleteAll(w http.ResponseWriter, r *http.Request){  

	service := &services.ArticleService{
		Logger: log.GetLogger(),
		DB: database.GetDB(),
	}

	func (servc models.ArticleInterface){ 
		ra, err := servc.DeleteAllArticles()

		if err != nil { 
			log.GetLogger().Info("DELETE/articles",err)
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
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
	}(service)
}

// FetchArticleByID - 
func FetchArticleByID(w http.ResponseWriter, r *http.Request){ 

	vars := mux.Vars(r)  
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	} 
  
	repo := &services.ArticleService{
		Logger: log.GetLogger(),
		DB: database.GetDB(),
	} 

	func (rep models.ArticleInterface){
		article, err := rep.FetchArticleByID(id)

		if err != nil { 
			responses.ERROR(w, http.StatusNotFound, err)
			return
		} 
		responses.JSON(w, http.StatusOK, article)
	}(repo)
}

// DeleteArticleByID delete article by id 
func DeleteArticleByID(w http.ResponseWriter, r *http.Request){ 

	vars := mux.Vars(r)  
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err) 
		return
	} 

	rep := &services.ArticleService{
		Logger : log.GetLogger(),
		DB : database.GetDB(),
	} 

	func (repo models.ArticleInterface){
		err := repo.DeleteByID(id) 
		if err != nil { 

			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			// return
		}

		w.Header().Set("Entity", fmt.Sprintf("%d", id))
		
		responses.JSON(w, http.StatusOK, 
			struct{
				Status string `json:"status"`
			}{
				Status: fmt.Sprintf("ok Record Deleted!"),
			},
		)
	}(rep)
}

// UpdateArticle pass id and data
func UpdateArticle(w http.ResponseWriter, r *http.Request){

	// url params 
	vars := mux.Vars(r) 
	aid, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	// decode payload 
	newData := new(models.Article)

	if err := json.NewDecoder(r.Body).Decode(newData); err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// update 
	repo := &services.ArticleService{
		Logger : log.GetLogger(),
		DB : database.GetDB(),
	} 

	func(re models.ArticleInterface){

		a, err := re.UpdateArticle(newData, aid)

		if err != nil {
			log.GetLogger().Info("ERROR updating")
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		} 

		responses.JSON(w, http.StatusOK, a)

	}(repo)
}
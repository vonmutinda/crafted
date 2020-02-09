package controllers

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json" 
	"github.com/gorilla/mux" 
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/services" 
	"github.com/vonmutinda/crafted/api/responses" 
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

	rep := &services.ArticleCRUD{}

	func (re models.ArticlesRepo){ 
		a, err := re.SaveArticle(article) 
		if err != nil{
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return 
		}
		responses.JSON(w, http.StatusCreated, a)
 	}(rep)
}

// Fetch all articles
func GetArticles(w http.ResponseWriter, r *http.Request){  

	repo := &services.ArticleCRUD{}
	func (re models.ArticlesRepo){
		a, e := re.GetAllArticles()
		if e != nil{
			log.Println(e)
			responses.ERROR(w, http.StatusUnprocessableEntity, e)
			return
		}  
		responses.JSON(w, http.StatusOK, a)
		
	}(repo)
}

// Delete all articles
type Ids struct {
	Id []int
}

func DeleteAll(w http.ResponseWriter, r *http.Request){  

	repo := &services.ArticleCRUD{}

	func (rep models.ArticlesRepo){ 
		ra, err := rep.DeleteAllArticles()

		if err != nil {
			log.Println(err)
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
	}(repo)
}

// find by id 
func FetchArticleByID(w http.ResponseWriter, r *http.Request){ 

	vars := mux.Vars(r)  
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	} 

	repo := &services.ArticleCRUD{} 

	func (rep models.ArticlesRepo){
		article, err := rep.FetchArticleByID(id)

		if err != nil { 
			responses.ERROR(w, http.StatusNotFound, err)
			return
		} 
		responses.JSON(w, http.StatusOK, article)
	}(repo)
}

// delete article by id 
func DeleteArticleByID(w http.ResponseWriter, r *http.Request){ 

	vars := mux.Vars(r)  
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err) 
	} 

	rep := &services.ArticleCRUD{} 

	func (repo models.ArticlesRepo){
		ra, err := repo.DeleteByID(id) 
		if err != nil { 
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

func UpdateArticle(w http.ResponseWriter, r *http.Request){

	// url params 
	vars := mux.Vars(r) 
	aid, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
	}

	// decode payload 
	newData := models.Article{}

	if err := json.NewDecoder(r.Body).Decode(&newData); err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	// update 
	repo := &services.ArticleCRUD{}

	func(re models.ArticlesRepo){

		a, err := re.UpdateArticle(newData, aid)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		} 

		responses.JSON(w, http.StatusOK, fmt.Sprintf("%d record(s) affected",a))

	}(repo)
}
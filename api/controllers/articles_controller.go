package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vonmutinda/crafted/api/auth"
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

	// is logged in
	uid, err := auth.TokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	// author_id of article should be logedin user id
	if uid != article.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	func (servc models.ArticleInterface){ 
		a, err := servc.SaveArticle(article) 
		if err != nil{
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return 
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s",r.Host, r.RequestURI)) 
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
		w.Header().Set("Location", fmt.Sprintf("%s%s",r.Host, r.RequestURI))  
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
		w.Header().Set("Location", fmt.Sprintf("%s%s",r.Host, r.RequestURI)) 
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

	service := &services.ArticleService{
		Logger : log.GetLogger(),
		DB : database.GetDB(),
	} 

	// look up user id token string
	uid, err := auth.TokenID(r) 
	if err != nil {
		service.Logger.Errorf("cannot find user_id in token string :%v", err)
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	} 

	func (servc models.ArticleInterface){
		ra, err := servc.DeleteByID(id, uid) 
		if err != nil {  
			responses.ERROR(w, http.StatusNotFound, err)
			return
		} 

		w.Header().Set("Entity", fmt.Sprintf("%d", uid)) 
		responses.JSON(w, http.StatusOK, 
			struct{
				Status string `json:"status"`
			}{
				Status: fmt.Sprintf("ok, %d Record(s) Deleted!", ra),
			},
		)
	}(service)
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
	article := new(models.Article) 
	if err := json.NewDecoder(r.Body).Decode(article); err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// validate article
	article.Prepare() 
	if err = article.Validate(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid fields"))
		return
	}

	// instantiate article service 
	service := &services.ArticleService{
		Logger : log.GetLogger(),
		DB : database.GetDB(),
	} 
	
	// logged in ?
	uid, err := auth.TokenID(r)
	if err != nil {
		service.Logger.Errorf("user_id not in token string: %v", err)
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	// authorised ? 
	if uid != article.AuthorID {
		service.Logger.Error("cannot delete another author's article")
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	func(servc models.ArticleInterface){ 

		ra, err := servc.UpdateArticle(article, aid, uid) 

		if err != nil {
			log.GetLogger().Info("ERROR updating")
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		} 

		w.Header().Set("Location", fmt.Sprintf("%s%s",r.Host, r.RequestURI)) 
		responses.JSON(w, http.StatusOK, fmt.Sprintf("Ok! %d record(s) updated!", ra))
	}(service)
}
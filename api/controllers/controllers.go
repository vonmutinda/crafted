package controllers

import (
	"net/http"

)

func GetArticles(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("{'ok':'200'}"))
}

func DeleteAll(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("All articles"))
}
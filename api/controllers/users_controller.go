package controllers

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json" 
	"github.com/gorilla/mux"
	"github.com/vonmutinda/crafted/api/repo" 
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/database" 
	"github.com/vonmutinda/crafted/api/repo/crud"
	"github.com/vonmutinda/crafted/api/responses"
)

// handleFunc methods

// fetch all useers
func GetUsers(w http.ResponseWriter, r *http.Request){
	
	// 1. db connection
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// 2. instantiate user repo
	rep := crud.NewUserCrud(db)

	// 3. fetch all users - implements UserRepository interface
	func (repo repo.UsersRepo){
		users, err := repo.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, users)

	}(rep)
}

// Create a new user
func CreateUser(w http.ResponseWriter, r *http.Request){
	// user instance 
	user := models.User{}  
	// json to struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// db connection
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// initialize repo
	rep := crud.NewUserCrud(db)

	// save new user - implements UsersRepo interface
	user.Prepare()
	if err := user.Validate(""); err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	func (re repo.UsersRepo){
		user, err = re.Save(user)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("location", fmt.Sprintf("%s%s/%d", r.Host, r.URL, user.ID)) 
		responses.JSON(w, http.StatusCreated, user) 
	}(rep)

}

// fetch user by ID
func GetUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) 

	uid, err := strconv.ParseUint( vars["id"], 10, 32 )
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// db connection
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	rep := crud.NewUserCrud(db) 
	// fetch all users - implements UserRepository interface
	func (rep repo.UsersRepo){
		user, err := rep.FindById(uid)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, user)

	}(rep)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Write( []byte("Update Users") )
	// user := models.User{} 
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Write( []byte("Delete User") )
}


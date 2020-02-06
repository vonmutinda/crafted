package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" 
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/responses"
	"github.com/vonmutinda/crafted/api/services"
)

// handleFunc methods

// fetch all useers
func GetUsers(w http.ResponseWriter, r *http.Request){ 

	rep := &services.UserCRUD{}
 
	func (repo models.UsersRepo){
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

	user := models.User{}  
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
  
	
	user.Prepare()
	if err := user.Validate(""); err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err) 
	}
	
	rep := &services.UserCRUD{} 

	func (re models.UsersRepo){
		user, err := re.Save(user)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
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
	} 

	rep := &services.UserCRUD{}  

	func (rep models.UsersRepo){
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


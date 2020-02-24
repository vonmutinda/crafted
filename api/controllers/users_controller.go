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
 

// GetUsers return all useers
func GetUsers(w http.ResponseWriter, r *http.Request){ 

	service := &services.UserService{
		Logger : log.GetLogger(),
		DB : database.GetDB(),
	}
 
	func (servc models.UsersInterface){
		users, err := servc.FindAll()
		if err != nil { 
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, users)

	}(service)
}
 
// CreateUser  new user
func CreateUser(w http.ResponseWriter, r *http.Request){ 

	user := new(models.User)  

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
  
	user.Prepare() 
	if err := user.Validate(); err != nil{
		responses.ERROR(w, http.StatusBadRequest, err) 
		return
	}
	
	service := &services.UserService{
		Logger : log.GetLogger(),
		DB : database.GetDB(),
	} 

	func (servc models.UsersInterface){
		resp, err := servc.Save(user)

		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		w.Header().Set("location", fmt.Sprintf("%s%s/%d", r.Host, r.URL, user.ID)) 
		responses.JSON(w, http.StatusCreated, resp) 
	}(service)

}

// GetUser by ID
func GetUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) 

	uid, err := strconv.ParseUint( vars["id"], 10, 32 )
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	} 

	service := &services.UserService{
		Logger :log.GetLogger(),
		DB: database.GetDB(),
	}  

	func (servc models.UsersInterface){
		user, err := servc.FindUserByID(uid)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		responses.JSON(w, http.StatusOK, user)

	}(service)
}


// UpdateUser - new info
func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Write( []byte("Update Users") )
	// user := models.User{} 
}

// DeleteUser - pass id 
func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Write( []byte("Delete User") )
}


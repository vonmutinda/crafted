package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vonmutinda/crafted/api/auth"
	"github.com/vonmutinda/crafted/api/database"
	"github.com/vonmutinda/crafted/api/log"
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/responses"
)


// Login - 
func Login(w http.ResponseWriter, r *http.Request){

	var err error
	user := new(models.User)

	service := &auth.Auth{
		Logger: log.GetLogger(),
		DB: database.GetDB(),
	}

	if err = json.NewDecoder(r.Body).Decode(user); err != nil {
		service.Logger.Errorf("cannot decode login user data : %v", err)
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	
	user.Prepare()
	// if err = user.Validate(); err != nil {
	// 	service.Logger.Error("cannot validate login data format : %v", err)
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	func (servc models.AuthInterface){ 
		token, err := servc.SignIn(user.Email, user.Password)
		if err != nil {
			service.Logger.Errorf("cannot create token : %v", err)
			responses.ERROR(w, http.StatusNotFound, err)
			return
		} 
		responses.JSON(w, http.StatusOK, struct{Token string `json:"token"`}{Token: token}) 
	}(service)

}
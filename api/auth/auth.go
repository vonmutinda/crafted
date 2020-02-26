package auth

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/vonmutinda/crafted/api/models"
	"github.com/vonmutinda/crafted/api/security"
)

// Auth struct
type Auth struct {
	Logger *logrus.Logger
	DB *gorm.DB
}

// SignIn - authorize
func (a *Auth)SignIn(email string, password string)(string, error) {

	var err error

	user := new(models.User)

	gor := 	a.DB.Raw(`
		SELECT * from users WHERE email = ?
	`, email).Scan(user)

	if err = gor.Error; err != nil {
		a.Logger.Errorf("cannot fetch user for auth. email : %s. error : %v\n", email, err)
		return "", err
	}
 
	err = security.VerifyPassword(user.Password, password)

	if err != nil {
		a.Logger.Errorf("cannot verify password : %v", err)
		return "", errors.New("Password does not match")
	}

	return CreateToken(user.ID)
}
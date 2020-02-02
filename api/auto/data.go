package auto 

import (
	"github.com/vonmutinda/crafted/api/models"

)

// some dummy data to get started with
var users = []models.User{
	models.User{
		Nickname	: "gopher_1",
		Email		: "email.golang.org",
		Password	: "password",
	},
}

var articles = []models.Article{
	models.Article{
		Title: "Golang Dummy Title",
		Body:  "This is the body of go code",
		AuthorID: 1,
	},
}
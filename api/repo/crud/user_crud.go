package crud

import (
	"github.com/jinzhu/gorm"

)

type repoUserCrud struct {
	db *gorm.DB
}

func NewRepoUserCrud (db *gorm.DB) *repoUserCrud{
	return &repoUserCrud{db}
}
package repo

import "github.com/vonmutinda/crafted/api/models"

type UsersRepo interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(uint64)(models.User, error)
	// Update(uint32, models.User) (uint64, error)
	// Delete(uint32) (uint64, error)
	// DeleteAll() (error)
}
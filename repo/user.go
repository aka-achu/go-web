package repo

import (
	"fmt"
	"github.com/aka-achu/go-web/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}


func (UserRepo) Create (user *models.User) error {
	return nil
}

func (UserRepo) Fetch (userID string) (*models.User, error) {
	fmt.Println("I am called", userID)
	return &models.User{}, nil
}
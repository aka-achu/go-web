package repo

import (
	"errors"
	"fmt"
	"github.com/aka-achu/go-web/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo returns an UserRepo object
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// Create, inserts a user record in the database
func (r *UserRepo) Create(user *models.User) error {
	return r.db.Table(user.TableName()).Create(user).Error
}

// Fetch, return a user record from the database for a given user_id
func (r *UserRepo) Fetch(userName string) (*models.User, error) {
	var queryResult models.User
	return &queryResult, r.db.Table(queryResult.TableName()).Where("user_name = ?", userName).First(&queryResult).Error
}

// Exists, returns existence status of the requested user
func (r *UserRepo) Exists(userName string) bool {
	u, err := r.Fetch(userName)
	fmt.Println(u,err)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
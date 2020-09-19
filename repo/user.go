package repo

import (
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
func (r *UserRepo) Fetch(userID string) (*models.User, error) {
	var queryResult models.User
	return &queryResult, r.db.Table(queryResult.TableName()).Where("id = ?", userID).Find(&queryResult).Error
}

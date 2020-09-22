package repo

import (
	"context"
	"errors"
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
func (r *UserRepo) Create(user *models.User, ctx context.Context) error {
	return r.db.
		WithContext(ctx).
		Table(user.TableName()).
		Create(user).
		Error
}

// Fetch, return a user record from the database for a given user_id
func (r *UserRepo) Fetch(userName string, ctx context.Context) (*models.User, error) {
	var queryResult models.User
	return &queryResult, r.db.
		WithContext(ctx).
		Table(queryResult.TableName()).
		Where("user_name = ?", userName).
		First(&queryResult).
		Error
}

// Exists, returns existence status of the requested user
func (r *UserRepo) Exists(userName string, ctx context.Context) bool {
	_, err := r.Fetch(userName, ctx)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
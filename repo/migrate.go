package repo

import (
	"github.com/aka-achu/go-web/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	user :=  &models.User{}
	if !db.Migrator().HasTable(user.TableName()) {
		return db.Table(user.TableName()).AutoMigrate(user)
	}
	return nil
}
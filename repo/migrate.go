package repo

import (
	"github.com/aka-achu/go-web/logging"
	"github.com/aka-achu/go-web/models"
	"gorm.io/gorm"
)

// migrate, syncs the object schema with repo tables
func migrate(db *gorm.DB) error {
	user :=  &models.User{}
	if !db.Migrator().HasTable(user.TableName()) {
		if err:= db.Table(user.TableName()).AutoMigrate(user); err != nil {
			logging.RepoLogger.Errorf("Failed to migrate the model.User schema. Error-%v", err)
			return err
		} else {
			logging.RepoLogger.Errorf("Successfully migrated the model.User schema")
		}
	}
	return nil
}
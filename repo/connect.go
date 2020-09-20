package repo

import (
	"github.com/aka-achu/go-web/logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

// GetConnection return an db connection
func GetConnection() (*gorm.DB, error) {

	// Validating the existence of the path for db
	if _, err := os.Stat(filepath.Join(os.Getenv("DATA_DIR"))); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Join(os.Getenv("DATA_DIR")), os.ModePerm); err != nil {
			logging.RepoLogger.Errorf("Error while establishing connection to the database. Error-%s",
				err.Error())
			return nil, err
		}
	}

	// Establishing db connection
	db, err := gorm.Open(
		sqlite.Open(filepath.Join(os.Getenv("DATA_DIR"), os.Getenv("DSN"))),
		&gorm.Config{},
	)
	if err != nil {
		logging.RepoLogger.Errorf("Error while establishing connection to the database. Error-%s",
			err.Error())
		return nil, err
	}

	// Configuring the connection for pooling
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Validating the db connection
	if err := sqlDB.Ping(); err != nil {
		logging.RepoLogger.Errorf("Error while pinging the database. Error-%s", err.Error())
		return nil, err
	}

	// Migrating schemas
	if err := migrate(db); err != nil {
		logging.RepoLogger.Errorf("Error while migrating schema. Error-%s", err.Error())
		return nil, err
	}
	return db, nil
}

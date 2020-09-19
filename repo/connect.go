package repo

import (
	"github.com/aka-achu/go-web/logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

func GetConnection() (*gorm.DB, error) {
	if _, err := os.Stat(filepath.Join(os.Getenv("DATA_DIR"))); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Join(os.Getenv("DATA_DIR")), os.ModePerm); err != nil {
			logging.RepoLogger.Errorf("Error while establishing connection to the database. Error-%s", err.Error())
			return nil, err
		}
	}
	db, err := gorm.Open(sqlite.Open(filepath.Join(os.Getenv("DATA_DIR"), os.Getenv("DSN"))), &gorm.Config{})
	if err != nil {
		logging.RepoLogger.Errorf("Error while establishing connection to the database. Error-%s", err.Error())
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		logging.RepoLogger.Errorf("Error while pinging the database. Error-%s", err.Error())
		return nil, err
	}
	return db, nil
}

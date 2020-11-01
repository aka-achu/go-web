package repo

import (
	"github.com/aka-achu/go-web/logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
	"time"
)

// GetConnection return an db connection
func GetConnection() (*gorm.DB, error) {

	// Validating the existence of the path for db
	if _, err := os.Stat(filepath.Join(os.Getenv("DATA_DIR"))); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Join(os.Getenv("DATA_DIR")), os.ModePerm); err != nil {
			return nil, err
		}
	}

	// Establishing db connection
	db, err := gorm.Open(
		sqlite.Open(filepath.Join(os.Getenv("DATA_DIR"), os.Getenv("DSN"))),
		&gorm.Config{
			Logger: logger.New(logging.SQLLogger, logger.Config{
				LogLevel: logger.Info,
			}),
		},
	)
	if err != nil {
		return nil, err
	}

	// Configuring the connection for pooling
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Validating the db connection
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

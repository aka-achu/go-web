package repo

import (
	"github.com/aka-achu/go-web/logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)
var Connection *gorm.DB

func Initialize() {
	if _, err := os.Stat(filepath.Join("data")); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Join("data"), os.ModePerm); err != nil {
			log.Fatal("Failed to create the folder for storing the logs", err)
		}
	}
	db, err := gorm.Open(sqlite.Open(filepath.Join("data","gorm.db")), &gorm.Config{})
	if err != nil {
		logging.RepoLogger.Fatalf("Error while establishing connection to the database. Error-%s", err.Error())
	}
	Connection = db
}
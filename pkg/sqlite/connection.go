package sqlite

import (
	"log"
	"os"

	"github.com/edaaltuntas/gomatching/pkg/sqlite/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn = NewConnection()

// Creates a new gorm connection, lookupenv is used for
// heroku-like non-persistent storages. Also it creates
// migrations automatically
func NewConnection() *gorm.DB {
	var path string
	_, m := os.LookupEnv("memory")
	if m {
		path = "file::memory:?cache=shared"
	} else {
		path = "gomatching.db"
	}
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Person{})
	return db
}

package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB, _ = New()
)

func New() (*gorm.DB, error) {

	// if getAppEnv() == "test" {

	// } else {
	driver := "sqlite3"
	connect := "sample.db"

	return gorm.Open(driver, connect)
	// }
}

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Open(source string) error {
	var e error
	database, e = gorm.Open(sqlite.Open(source))
	if e != nil {
		return e
	}
	e = database.AutoMigrate(&DataModel{})
	if e != nil {
		return e
	}
	e = database.AutoMigrate(&DataModel{})
	return e
}

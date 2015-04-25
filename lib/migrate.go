// Package main provides ...
package main

import (
	. "github.com/Pitt-CSC/icarus-backend/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open connection
	db, err := gorm.Open("sqlite3", "tmp/gorm.db")
	if err != nil {
		panic(err)
	}
	db.DB()

	// Run migrations
	db.AutoMigrate(&Talk{})
}

// Package models provides ...
package models

import (
	"github.com/jinzhu/gorm"
)

var db gorm.DB

func InitializeDBConnection(dbconnection gorm.DB) {
	db = dbconnection
}

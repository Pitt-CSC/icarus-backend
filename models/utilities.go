// Package models provides ...
package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

var db gorm.DB

func InitializeDBConnection(dbconnection gorm.DB) {
	db = dbconnection
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

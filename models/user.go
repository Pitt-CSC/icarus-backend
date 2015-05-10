// Package models provides ...
package models

import (
	"time"
)

type User struct {
	ID        int `json:"id" sql:"AUTO_INCREMENT"`
	GithubID  int
	FirstName string     `json:"first-name"`
	LastName  string     `json:"last-name"`
	Email     string     `json:"email"`
	AvatarUrl string     `json:"avatar_url"`
	CreatedAt time.Time  `json:"create-at"`
	UpdatedAt time.Time  `json:"updated-at"`
	DeletedAt *time.Time `json:"deleted-at,omitempty"`
}

func (user *User) Save() (err error) {
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return
}

// Package models provides ...
package models

import (
	"time"
)

type Talk struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Upvotes   int        `json:"upvotes"`
	Downvotes int        `json:"downvotes"`
	CreatedAt time.Time  `json:"created-at"`
	UpdatedAt time.Time  `json:"updated-at"`
	DeletedAt *time.Time `json:"deleted-at,omitempty"`
}

func (talk *Talk) BeforeCreate() (err error) {
	// Initialize vote values
	talk.Upvotes = 0
	talk.Downvotes = 0
	return
}

func (talk *Talk) Upvote() (err error) {
	talk.Upvotes += 1
	return talk.Save()
}

func (talk *Talk) Downvote() (err error) {
	talk.Upvotes -= 1
	return talk.Save()
}

func (talk *Talk) Save() (err error) {
	if err := db.Save(&talk).Error; err != nil {
		return err
	}
	return
}

type TalkResource struct {
	Data Talk `json:"talk"`
}

type TalkCollection struct {
	Data []Talk `json:"talks"`
}

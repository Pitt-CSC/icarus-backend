// Package models provides ...
package models

import (
	"time"
)

type Talk struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created-at"`
	UpdatedAt time.Time  `json:"updated-at"`
	DeletedAt *time.Time `json:"deleted-at,omitempty"`
}

type TalkResource struct {
	Data Talk `json:"talk"`
}

type TalkCollection struct {
	Data []Talk `json:"talks"`
}

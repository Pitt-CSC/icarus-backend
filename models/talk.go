// Package models provides ...
package models

type Talk struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type TalkResource struct {
	Data Talk `json:"talk"`
}

type TalkCollection struct {
	Data []Talk `json:"talks"`
}

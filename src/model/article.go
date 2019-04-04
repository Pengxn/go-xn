package model

import "time"

// Article model
type Article struct {
	ID         uint64
	Title      string
	Content    string
	AuthorID   uint64
	Views      uint
	Status     int
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime time.Time
}

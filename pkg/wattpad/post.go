package wattpad

import "time"

//Model of post
type ModelPost struct {
	ID        uint //PK primary key
	Title     string
	Body      string
	View      string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int   //FK Foreign Key
	CommentID []int //FK Foreign Key
}

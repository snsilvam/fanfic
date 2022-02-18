package wattpad

import "time"

//Model of comment
type ModelComment struct {
	ID          uint //PK primary key
	Title       string
	Description string
	View        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PostID      int //FK Foreign Key
}

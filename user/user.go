package user

import "time"

//Model of user
type ModelUser struct {
	ID        uint //PK primary key
	Name      string
	Age       int
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
	//Slice of idPost
	PostID []int //FK Foreign Key
}

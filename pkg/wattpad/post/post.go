package post

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

//
type Storage interface {
	Migrate() error
}

//
type Service struct {
	storage Storage
}

//Constructor of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

//
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

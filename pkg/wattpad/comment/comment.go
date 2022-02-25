package comment

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

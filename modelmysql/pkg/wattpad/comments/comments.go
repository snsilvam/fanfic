package comments

import "time"

//Model of comment
type ModelComments struct {
	ID        uint //PK primary key = indice de busqueda
	Comment   string
	View      int
	CreatedAt time.Time
	UpdatedAt time.Time
	User_id   int //FK Foreign Key
	Post_id   int //FK Foreign Key
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

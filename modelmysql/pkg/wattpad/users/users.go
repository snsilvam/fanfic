package users

import "time"

//Model of Users
type ModelUsers struct {
	ID        uint //PK primary key = indice de busqueda
	Name      string
	Age       int
	Gender    string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
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

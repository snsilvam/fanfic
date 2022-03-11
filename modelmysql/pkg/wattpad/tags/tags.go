package tags

import "time"

type ModelTags struct {
	ID        uint //PK primary key = indice de busqueda
	NameTag   string
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

package posts

import "time"

//Model of post
type ModelPosts struct {
	ID          uint //PK primary key = indice de busqueda
	Title       string
	Body        string
	View        string
	CreatedAt   time.Time //fecha de publicacion
	UpdatedAt   time.Time
	User_id     int //FK
	Category_id int //FK
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

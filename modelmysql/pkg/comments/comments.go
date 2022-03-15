package comments

import "time"

//The struct ModelComments, is the data model of Comments
type ModelComments struct {
	ID        uint //PK primary key = indice de busqueda
	Comment   string
	View      int
	CreatedAt time.Time
	UpdatedAt time.Time
	User_id   int //FK Foreign Key
	Post_id   int //FK Foreign Key
}

//The Storage interface allows connect our data model ModelComments with the data base. In order to use the storage interface, we must use storage interface's methods
type Storage interface {
	Migrate() error
}

//The struct service have a storage feature, this storage feature is the interface Storage
type Service struct {
	storage Storage
}

//Service's constructor
func NewService(s Storage) *Service {
	return &Service{s}
}

//Implement the interface Storage, because return a error and have a migrate function
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

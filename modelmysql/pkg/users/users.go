package users

import "time"

//The struct ModelUsers, is the data model of users
type ModelUsers struct {
	ID        uint //PK primary key = indice de busqueda
	Name      string
	Age       int
	Gender    string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//The Storage interface allows connect our data model ModelUsers with the data base. In order to use the storage interface, we must use storage interface's methods
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

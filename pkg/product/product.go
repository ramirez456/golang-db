package product

import (
	"errors"
	"time"
)
var (ErrIDNotFound = errors.New("El producto no contiene id"))
type Model struct {
	ID uint
	Name string
	Observations string
	Price float64
	CreatedAt time.Time
	UpdatedAt time.Time

}
type Models []*Model

type Storage interface {
	Migrate() error
	Create(model *Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint)(*Model, error)
	Delete(uint) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return  &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

func (s *Service) Create(m *Model) error{
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}


func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

func (s *Service) GetByID(id uint) (*Model, error){
	return s.storage.GetByID(id)
}

func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}
	m.UpdatedAt = time.Now()

	return s.storage.Update(m)
}

func (s *Service) Delete(id uint) error {
	return  s.storage.Delete(id)
}
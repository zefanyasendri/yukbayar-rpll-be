package pengguna

import (
	"yukbayar-rpll-be/helpers/password"

	"github.com/google/uuid"
)

type Service interface {
	Create(req *Pengguna) (string, bool, error)
	GetAll() ([]Pengguna, error)
	GetByID(ID string) (Pengguna, error)
	UpdateByID(ID string, req *UpdateRequest) error
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(req *Pengguna) (string, bool, error) {
	exists, err := s.repository.GetByEmail(req.Email)
	if exists {
		return "", exists, err
	}

	req.ID = uuid.New().String()
	if req.Password, err = password.HashPassword(req.Password); err != nil {
		return "", false, err
	}

	err = s.repository.Create(req)
	return req.ID, false, err
}

func (s *service) GetAll() ([]Pengguna, error) {
	penggunas, err := s.repository.GetAll()
	return penggunas, err
}

func (s *service) GetByID(ID string) (Pengguna, error) {
	pengguna, err := s.repository.GetByID(ID)
	return pengguna, err
}

func (s *service) UpdateByID(ID string, req *UpdateRequest) error {
	err := s.repository.UpdateByID(ID, req)
	return err
}

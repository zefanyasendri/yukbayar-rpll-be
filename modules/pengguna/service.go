package pengguna

import (
	"yukbayar-rpll-be/helpers/password"

	"github.com/google/uuid"
)

type Service interface {
	Create(req *Pengguna) (string, bool, error)
	GetAll() ([]Pengguna, error)
	GetByID(ID string) (Pengguna, error)
	GetAccount(email string, password string) (string, error)
	UpdateByID(ID string, req *UpdateRequest) error
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(req *Pengguna) (string, bool, error) {
	pengguna, err := s.repository.GetByEmail(req.Email)
	if pengguna.Email == req.Email {
		return "", true, err
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

func (s *service) GetAccount(email string, pass string) (string, error) {
	pengguna, err := s.repository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	match := password.CheckPasswordHash(pengguna.Password, pass)
	if pengguna.Email != email && !match {
		return "", err
	}

	token := "token hahahah"

	return token, err
}

func (s *service) UpdateByID(ID string, req *UpdateRequest) error {
	err := s.repository.UpdateByID(ID, req)
	return err
}

package services

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"

	"github.com/google/uuid"
)

type PenggunaService interface {
	Create(req *models.Pengguna) (string, bool, error)
	GetAll() ([]models.Pengguna, error)
	GetByID(ID string) (models.Pengguna, error)
	GetAccount(email string, password string) (string, error)
	UpdateByID(ID string, req *models.PenggunaUpdateRequest) error
}
type penggunaService struct {
	penggunaRepository repositories.PenggunaRepository
}

func NewPenggunaService(penggunaRepository repositories.PenggunaRepository) *penggunaService {
	return &penggunaService{penggunaRepository}
}

func (s *penggunaService) Create(req *models.Pengguna) (string, bool, error) {
	pengguna, err := s.penggunaRepository.GetByEmail(req.Email)
	if pengguna.Email == req.Email {
		return "", true, err
	}

	req.ID = uuid.New().String()
	if req.Password, err = helpers.HashPassword(req.Password); err != nil {
		return "", false, err
	}

	err = s.penggunaRepository.Create(req)
	return req.ID, false, err
}

func (s *penggunaService) GetAll() ([]models.Pengguna, error) {
	penggunas, err := s.penggunaRepository.GetAll()
	return penggunas, err
}

func (s *penggunaService) GetByID(ID string) (models.Pengguna, error) {
	pengguna, err := s.penggunaRepository.GetByID(ID)
	return pengguna, err
}

func (s *penggunaService) GetAccount(email string, pass string) (string, error) {
	pengguna, err := s.penggunaRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	match := helpers.CheckPasswordHash(pengguna.Password, pass)
	if pengguna.Email != email && !match {
		return "", err
	}

	token := "token"

	return token, err
}

func (s *penggunaService) UpdateByID(ID string, req *models.PenggunaUpdateRequest) error {
	pengguna, err := s.penggunaRepository.GetByID(ID)

	if pengguna.Password != req.Password || req.Password != "" {
		if req.Password, err = helpers.HashPassword(req.Password); err != nil {
			return err
		}
	}

	err = s.penggunaRepository.UpdateByID(ID, req)
	return err
}

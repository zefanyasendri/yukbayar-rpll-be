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
	GetAccount(email string, password string) (models.Pengguna, bool, error)
	UpdateByID(ID string, req *models.PenggunaUpdateRequest) (models.Pengguna, bool, error)
	UpdateSaldoByID(ID string, saldo int) (models.Pengguna, error)
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

func (s *penggunaService) GetAccount(email string, pass string) (models.Pengguna, bool, error) {
	pengguna, err := s.penggunaRepository.GetByEmail(email)
	match := helpers.CheckPasswordHash(pengguna.Password, pass)
	return pengguna, match, err
}

func (s *penggunaService) UpdateByID(ID string, req *models.PenggunaUpdateRequest) (models.Pengguna, bool, error) {
	var match bool
	pengguna, err := s.penggunaRepository.GetByID(ID)

	//update jika pengguna tidak mengubah password
	if (req.OldPassword != "" && req.Password != "") || (req.OldPassword != "" && req.Password == "") {
		match = helpers.CheckPasswordHash(pengguna.Password, req.OldPassword)
		if match {
			if req.Password != "" {
				req.Password, _ = helpers.HashPassword(req.Password)
			}
			req.OldPassword = ""
		}
	}
	if (req.OldPassword == "" && req.Password == "") || match {
		err = s.penggunaRepository.UpdateByID(ID, req)
	}
	pengguna, err = s.penggunaRepository.GetByID(ID)
	return pengguna, match, err
}

func (s *penggunaService) UpdateSaldoByID(ID string, saldo int) (models.Pengguna, error) {
	pengguna, err := s.penggunaRepository.UpdateSaldoByID(ID, saldo)
	return pengguna, err
}

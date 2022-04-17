package pengguna

type Service interface {
	Create(pengguna Pengguna) (Pengguna, error)
	GetAll() ([]Pengguna, error)
	GetByID(ID string) (Pengguna, error)
	UpdateByID(ID string, req *UpdateRequest) error
}
type service struct {
	repository Repository
}

func (s *service) Create(pengguna Pengguna) (Pengguna, error) {
	return pengguna, nil //belum ada logic apa2
}

func NewService(repository Repository) *service {
	return &service{repository}
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

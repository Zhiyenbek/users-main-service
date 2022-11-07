package service

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
)

type doctorService struct {
	doctorRepo repository.DoctorRepository
	cfg        *config.Configs
}

func NewDoctorService(repo *repository.Repository, cfg *config.Configs) DoctorService {
	return &doctorService{
		doctorRepo: repo.DoctorRepository,
		cfg:        cfg,
	}
}

func (s *doctorService) UpdateDoctor(doctorReq *models.UpdateDoctorRequest) (*models.GetDoctorResponse, error) {
	return s.doctorRepo.UpdateDoctor(doctorReq, 0)
}
func (s *doctorService) CreateDoctor(doctorReq *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error) {
	return s.doctorRepo.CreateDoctor(doctorReq)
}
func (s *doctorService) DeleteDoctor(ID int64) error {
	return s.doctorRepo.DeleteDoctor(ID, 0)
}
func (s *doctorService) GetDoctor(ID int64) error {
	return s.doctorRepo.DeleteDoctor(ID, 0)
}

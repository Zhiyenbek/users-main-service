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

func (s *doctorService) UpdateDoctor(models.UpdateDoctorRequest) (models.GetDoctorResponse, error) {

}
func (s *doctorService) CreateDoctor(models.CreateDoctorRequest) (models.CreateDoctorRequest, error) {

}
func (s *doctorService) DeleteDoctor(ID int64) error {

}
func (s *doctorService) GetDoctor(ID int64) error {

}

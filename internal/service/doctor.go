package service

import (
	"log"

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

func (s *doctorService) UpdateDoctor(doctorReq *models.UpdateDoctorRequest) error {
	err := s.doctorRepo.UpdateDoctor(doctorReq, 0)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (s *doctorService) CreateDoctor(doctorReq *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error) {
	res, err := s.doctorRepo.CreateDoctor(doctorReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
func (s *doctorService) DeleteDoctor(ID int64) error {
	err := s.doctorRepo.DeleteDoctor(ID, 0)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (s *doctorService) GetDoctor(ID int64) (*models.GetDoctorResponse, error) {
	res, err := s.doctorRepo.GetDoctor(ID, 0)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

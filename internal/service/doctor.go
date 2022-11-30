package service

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
	"go.uber.org/zap"
)

type doctorService struct {
	doctorRepo repository.DoctorRepository
	logger     *zap.SugaredLogger
	cfg        *config.Configs
}

func NewDoctorService(repo *repository.Repository, logger *zap.SugaredLogger, cfg *config.Configs) DoctorService {
	return &doctorService{
		doctorRepo: repo.DoctorRepository,
		logger:     logger,
		cfg:        cfg,
	}
}

func (s *doctorService) UpdateDoctor(doctorReq *models.UpdateDoctorRequest) error {
	err := s.doctorRepo.UpdateDoctor(doctorReq)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
func (s *doctorService) CreateDoctor(doctorReq *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error) {

	res, err := s.doctorRepo.CreateDoctor(doctorReq)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}
func (s *doctorService) DeleteDoctor(ID int64) error {
	err := s.doctorRepo.DeleteDoctor(ID)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
func (s *doctorService) GetDoctor(ID int64) (*models.GetDoctorResponse, error) {
	res, err := s.doctorRepo.GetDoctor(ID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}

func (s *doctorService) SearchDoctors(searchArgs *models.Search) (*models.SearchDoctorsResponse, error) {
	res, err := s.doctorRepo.SearchDoctors(searchArgs)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}

func (s *doctorService) GetDoctorByDepartment(ID int64, search *models.Search) (*models.SearchDoctorsResponse, error) {
	res, err := s.doctorRepo.SearchDoctorsByDepartment(search, ID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}

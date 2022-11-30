package service

import (
	"log"

	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
	"go.uber.org/zap"
)

type patientService struct {
	patientRepo repository.PatientRepository
	logger      *zap.SugaredLogger
	cfg         *config.Configs
}

func NewPatientService(repo *repository.Repository, logger *zap.SugaredLogger, cfg *config.Configs) PatientService {
	return &patientService{
		patientRepo: repo.PatientRepository,
		logger:      logger,
		cfg:         cfg,
	}
}
func (s *patientService) UpdatePatient(patientReq *models.UpdatePatientRequest) error {
	err := s.patientRepo.UpdatePatient(patientReq)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
func (s *patientService) CreatePatient(patientReq *models.CreatePatientRequest) (*models.CreatePatientResponse, error) {
	res, err := s.patientRepo.CreatePatient(patientReq)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}
func (s *patientService) DeletePatient(ID int64) error {
	err := s.patientRepo.DeletePatient(ID)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
func (s *patientService) GetPatient(ID int64) (*models.GetPatientResponse, error) {
	res, err := s.patientRepo.GetPatient(ID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}
func (s *patientService) GetAllPatients() ([]*models.GetAllPatientsResponse, error) {
	res, err := s.patientRepo.GetAllPatients()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

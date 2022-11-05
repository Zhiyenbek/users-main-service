package service

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
)

type patientService struct {
	PatientRepo repository.PatientRepository
	cfg         *config.Configs
}

func NewPatientService(repo *repository.Repository, cfg *config.Configs) PatientService {
	return &patientService{
		PatientRepo: repo.PatientRepository,
		cfg:         cfg,
	}
}
func (s *patientService) UpdatePatient(patientReq *models.UpdatePatientRequest) (*models.GetPatientResponse, error) {
	return s.UpdatePatient(patientReq)
}
func (s *patientService) CreatePatient(patientReq *models.CreatePatientRequest) (*models.CreatePatientRequest, error) {
	return s.CreatePatient(patientReq)
}
func (s *patientService) DeletePatient(ID int64) error {
	return s.DeletePatient(ID)
}
func (s *patientService) GetPatient(ID int64) error {
	return s.GetPatient(ID)
}

package service

import (
	"log"

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
func (s *patientService) UpdatePatient(patientReq *models.UpdatePatientRequest) error {
	userID, err := s.PatientRepo.GetUserIDbyID(patientReq.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.PatientRepo.UpdatePatient(patientReq, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (s *patientService) CreatePatient(patientReq *models.CreatePatientRequest) (*models.CreatePatientResponse, error) {
	res, err := s.PatientRepo.CreatePatient(patientReq)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
func (s *patientService) DeletePatient(ID int64) error {
	userID, err := s.PatientRepo.GetUserIDbyID(ID)
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.PatientRepo.DeletePatient(ID, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (s *patientService) GetPatient(ID int64) (*models.GetPatientResponse, error) {
	userID, err := s.PatientRepo.GetUserIDbyID(ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := s.PatientRepo.GetPatient(ID, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

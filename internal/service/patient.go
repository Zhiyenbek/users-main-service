package service

import (
	"github.com/Zhiyenbek/users-auth-service/config"
	"github.com/Zhiyenbek/users-auth-service/internal/repository"
)

type PatientService interface {
}

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

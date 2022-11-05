package service

import (
	"github.com/Zhiyenbek/users-auth-service/config"
	"github.com/Zhiyenbek/users-auth-service/internal/repository"
)

type Service struct {
	PatientService
	DoctorService
}

func New(repos *repository.Repository, cfg *config.Configs) *Service {
	return &Service{
		PatientService: NewPatientService(repos, cfg),
		DoctorService:  NewDoctorService(repos, cfg),
	}
}

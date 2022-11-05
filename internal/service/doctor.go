package service

import (
	"github.com/Zhiyenbek/users-auth-service/config"
	"github.com/Zhiyenbek/users-auth-service/internal/repository"
)

type DoctorService interface {
}

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

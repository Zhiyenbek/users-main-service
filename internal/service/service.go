package service

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
)

type DoctorService interface {
	UpdateDoctor(doctorReq *models.UpdateDoctorRequest) (*models.GetDoctorResponse, error)
	CreateDoctor(doctorReq *models.CreateDoctorRequest) (*models.CreateDoctorRequest, error)
	DeleteDoctor(ID int64) error
	GetDoctor(ID int64) error
}
type PatientService interface {
	UpdatePatient(pateintReq *models.UpdatePatientRequest) (*models.GetPatientResponse, error)
	CreatePatient(pateintReq *models.CreatePatientRequest) (*models.CreatePatientRequest, error)
	DeletePatient(ID int64) error
	GetPatient(ID int64) error
}

type AdminService interface {
	CheckAuth(id int64) error
}
type Service struct {
	PatientService
	DoctorService
	AdminService
}

func New(repos *repository.Repository, cfg *config.Configs) *Service {
	return &Service{
		PatientService: NewPatientService(repos, cfg),
		DoctorService:  NewDoctorService(repos, cfg),
	}
}

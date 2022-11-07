package service

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
)

type DoctorService interface {
	UpdateDoctor(doctorReq *models.UpdateDoctorRequest) error
	CreateDoctor(doctorReq *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error)
	DeleteDoctor(ID int64) error
	GetDoctor(ID int64) (*models.GetDoctorResponse, error)
	GetAllDoctors() ([]*models.GetAllDoctorsResponse, error)
}
type PatientService interface {
	UpdatePatient(patientReq *models.UpdatePatientRequest) error
	CreatePatient(pateintReq *models.CreatePatientRequest) (*models.CreatePatientResponse, error)
	DeletePatient(ID int64) error
	GetPatient(ID int64) (*models.GetPatientResponse, error)
	GetAllPatients() ([]*models.GetAllPatientsResponse, error)
}
type AuthService interface {
	Login(creds *models.UserSignInRequest) (*models.Tokens, error)
	RefreshToken(tokenString string) (*models.Tokens, error)
}

type Service struct {
	PatientService
	DoctorService
	AuthService
}

func New(repos *repository.Repository, cfg *config.Configs) *Service {
	return &Service{
		PatientService: NewPatientService(repos, cfg),
		DoctorService:  NewDoctorService(repos, cfg),
		AuthService:    NewAuthService(repos, cfg),
	}
}

package service

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
	"go.uber.org/zap"
)

type DoctorService interface {
	UpdateDoctor(doctorReq *models.UpdateDoctorRequest) error
	CreateDoctor(doctorReq *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error)
	DeleteDoctor(ID int64) error
	GetDoctor(ID int64) (*models.GetDoctorResponse, error)
	SearchDoctors(*models.Search) (*models.SearchDoctorsResponse, error)
	GetDoctorByDepartment(int64, *models.Search) (*models.SearchDoctorsResponse, error)
}
type PatientService interface {
	UpdatePatient(patientReq *models.UpdatePatientRequest) error
	CreatePatient(pateintReq *models.CreatePatientRequest) (*models.CreatePatientResponse, error)
	DeletePatient(ID int64) error
	GetPatient(ID int64) (*models.GetPatientResponse, error)
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

func New(repos *repository.Repository, log *zap.SugaredLogger, cfg *config.Configs) *Service {
	return &Service{
		PatientService: NewPatientService(repos, log, cfg),
		DoctorService:  NewDoctorService(repos, log, cfg),
		AuthService:    NewAuthService(repos, cfg, log),
	}
}

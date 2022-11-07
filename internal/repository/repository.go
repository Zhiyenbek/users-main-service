package repository

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/go-redis/redis/v7"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	DoctorRepository
	PatientRepository
	AuthRepository
	TokenRepository
}
type DoctorRepository interface {
	CreateDoctor(doctor *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error)
	DeleteDoctor(ID int64, userID int64) error
	UpdateDoctor(doctor *models.UpdateDoctorRequest, userID int64) error
	GetDoctor(ID int64, UserID int64) (*models.GetDoctorResponse, error)
	GetUserIDbyID(ID int64) (int64, error)
	GetAllDoctors() ([]*models.GetAllDoctorsResponse, error)
}

type PatientRepository interface {
	CreatePatient(patient *models.CreatePatientRequest) (*models.CreatePatientResponse, error)
	DeletePatient(ID int64, userID int64) error
	UpdatePatient(patient *models.UpdatePatientRequest, userID int64) error
	GetPatient(ID int64, UserID int64) (*models.GetPatientResponse, error)
	GetUserIDbyID(ID int64) (int64, error)
	GetAllPatients() ([]*models.GetAllPatientsResponse, error)
}

type AuthRepository interface {
	GetUserInfoByLogin(login string) (string, int64, error)
}

type TokenRepository interface {
	SetRTToken(token *models.Token) error
	UnsetRTToken(userID int64) error
	GetToken(userID int64) (string, error)
}

func New(db *pgxpool.Pool, cfg *config.Configs, redis *redis.Client) *Repository {
	return &Repository{
		DoctorRepository:  NewDoctorRepository(db, cfg.DB),
		PatientRepository: NewPatientRepository(db, cfg.DB),
		AuthRepository:    NewAuthRepository(db, cfg.DB),
		TokenRepository:   NewTokenRepository(redis),
	}
}

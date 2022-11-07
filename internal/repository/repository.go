package repository

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	DoctorRepository
	PatientRepository
	AdminRepository
}
type DoctorRepository interface {
	CreateDoctor(doctor *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error)
	DeleteDoctor(ID int64, userID int64) error
	UpdateDoctor(doctor *models.UpdateDoctorRequest, userID int64) error
	GetDoctor(ID int64, UserID int64) (*models.GetDoctorResponse, error)
}
type PatientRepository interface {
	CreatePatient(patient *models.CreatePatientRequest) (*models.CreatePatientResponse, error)
	DeletePatient(ID int64, userID int64) error
	UpdatePatient(patient *models.UpdatePatientRequest, userID int64) error
	GetPatient(ID int64, UserID int64) (*models.GetPatientResponse, error)
	GetUserIDbyID(ID int64) (int64, error)
}
type AdminRepository interface {
	CheckAuth(ID int64) error
}

func New(db *pgxpool.Pool, cfg *config.Configs) *Repository {
	return &Repository{
		DoctorRepository:  NewDoctorRepository(db, cfg.DB),
		PatientRepository: NewPatientRepository(db, cfg.DB),
	}
}

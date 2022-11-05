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
	CreateDoctor(*models.CreateDoctorRequest) (*models.CreateDoctorResponse, error)
	DeleteDoctor(ID int64) error
	UpdateDoctor(*models.UpdateDoctorRequest) (*models.GetDoctorResponse, error)
	GetDoctor(ID int64) (*models.GetDoctorResponse, error)
}
type PatientRepository interface {
	CreatePatient(*models.CreatePatientRequest) (*models.CreatePatientResponse, error)
	DeletePatient(ID int64) error
	UpdatePatient(*models.UpdatePatientRequest) (*models.GetPatientResponse, error)
	GetPatient(ID int64) (*models.GetPatientResponse, error)
}
type AdminRepository interface {
	CheckAuth(ID int64) error
}

func New(db *pgxpool.Pool, cfg *config.Configs) *Repository {
	return &Repository{
		DoctorRepository:  NewDoctorRepository(db, cfg),
		PatientRepository: NewPatientRepository(db, cfg),
	}
}

package repository

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	DoctorRepository
	PatientRepository
	AdminRepository
}
type DoctorRepository interface {
}
type PatientRepository interface {
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

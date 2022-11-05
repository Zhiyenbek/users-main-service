package repository

import (
	"github.com/Zhiyenbek/users-auth-service/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	DoctorRepository
	PatientRepository
}
type DoctorRepository interface {
}
type PatientRepository interface {
}

func New(db *pgxpool.Pool, cfg *config.Configs) *Repository {
	return &Repository{
		DoctorRepository:  NewDoctorRepository(db, cfg),
		PatientRepository: NewPatientRepository(db, cfg),
	}
}

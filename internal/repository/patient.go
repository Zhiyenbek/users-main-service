package repository

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type patientRepository struct {
	db  *pgxpool.Pool
	cfg *config.Configs
}

func NewPatientRepository(db *pgxpool.Pool, cfg *config.Configs) PatientRepository {
	return &patientRepository{
		db:  db,
		cfg: cfg,
	}
}

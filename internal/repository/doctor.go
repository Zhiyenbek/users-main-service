package repository

import (
	"github.com/Zhiyenbek/users-auth-service/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type doctorRepository struct {
	db  *pgxpool.Pool
	cfg *config.Configs
}

func NewDoctorRepository(db *pgxpool.Pool, cfg *config.Configs) DoctorRepository {
	return &doctorRepository{
		db:  db,
		cfg: cfg,
	}
}

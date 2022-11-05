package repository

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
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
func (r *patientRepository) CreatePatient(*models.CreatePatientRequest) (*models.CreatePatientResponse, error) {

}
func (r *patientRepository) DeletePatient(ID int64) error {

}
func (r *patientRepository) UpdatePatient(*models.UpdatePatientRequest) (*models.GetPatientResponse, error) {

}
func (r *patientRepository) GetPatient(ID int64) (*models.GetPatientResponse, error) {

}

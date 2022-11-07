package repository

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type doctorRepository struct {
	db  *pgxpool.Pool
	cfg *config.DBConf
}

func NewDoctorRepository(db *pgxpool.Pool, cfg *config.DBConf) DoctorRepository {
	return &doctorRepository{
		db:  db,
		cfg: cfg,
	}
}
func (r *doctorRepository) CreateDoctor(doctor *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error) {
	return nil, nil
}
func (r *doctorRepository) DeleteDoctor(ID int64, userID int64) error {
	return nil
}
func (r *doctorRepository) UpdateDoctor(doctor *models.UpdateDoctorRequest, userID int64) (*models.GetDoctorResponse, error) {
	return nil, nil
}
func (r *doctorRepository) GetDoctor(ID int64, UserID int64) (*models.GetDoctorResponse, error) {
	return nil, nil
}

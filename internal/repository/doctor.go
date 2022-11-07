package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/jackc/pgx/v4"
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
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	var userID, ID int64
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating doctor in users: %v", err)
	}
	query := `INSERT INTO users 
				(first_name, last_name, middle_name, birthdata, iin, phone, address, email)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8) RETURNING ID;`
	dRow, err := tx.Query(ctx, query, doctor.FirstName, doctor.LastName, doctor.MiddleName, doctor.BirthDate, doctor.IIN, doctor.Phone, doctor.Address, doctor.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating deleting patient in users: %v", err)
	}
	err = dRow.Scan(&userID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating deleting patient in users: %v", err)
	}

	query = `INSERT INTO doctors 
		(department_id, spec_id, experience, photo, category, price, schedule, degree, rating, website_url, user_id)
			VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $10);`
	dRow, err = tx.Query(ctx, query, doctor.DepartmentId, doctor.SpecId, doctor.Experience, doctor.Photo, doctor.Category, doctor.Price, doctor.Schedule, doctor.Degree, doctor.Rating, doctor.WebsiteUrl, userID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating deleting patient in patients: %v", err)
	}
	err = dRow.Scan(&ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating deleting patient in patients: %v", err)
	}
	return &models.CreateDoctorResponse{
		ID: ID,
	}, nil
}

func (r *doctorRepository) DeleteDoctor(ID int64, userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	query := `DELETE FROM patients WHERE id=$1`
	_, err = tx.Exec(ctx, query, ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting doctor from users: %v", err)
	}

	query = `DELETE FROM doctors WHERE id=$1`
	_, err = tx.Exec(ctx, query, userID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting doctor from doctors: %v", err)
	}
	return nil
}

func (r *doctorRepository) UpdateDoctor(doctor *models.UpdateDoctorRequest, userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	query := `UPDATE users 
				SET first_name = $1, last_name = $2, middle_name = $3, birthdata = $4, iin = $5, phone = $6, address = $7, email = $8)
			  WHERE id = $9`
	_, err = tx.Exec(ctx, query, doctor.FirstName, doctor.LastName, doctor.MiddleName, doctor.BirthDate, doctor.IIN, doctor.Phone, doctor.Address, doctor.Email, userID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating doctor info in users table: %v", err)
	}

	query = `UPDATE doctors 
				SET deparment_id = $1, spec_id = $2, experience = $3, photo = $4, category = $5, price = $6, schedule = $7, degree = $8, rating = $9, website_url = $10, user_id = $11)
			  WHERE id = $12`
	_, err = tx.Query(ctx, query, doctor.DepartmentId, doctor.SpecId, doctor.Experience, doctor.Photo, doctor.Category, doctor.Price, doctor.Schedule, doctor.Degree, doctor.Rating, doctor.WebsiteUrl, doctor.ID, userID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating doctor INFO in doctors: %v", err)
	}
	return nil
}

func (r *doctorRepository) GetDoctor(ID int64, UserID int64) (*models.GetDoctorResponse, error) {
	res := &models.GetDoctorResponse{}
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	query := `SELECT (first_name, last_name, middle_name, birthdata, iin, phone, address, email) FROM users WHERE id=$1`
	dRow, err := tx.Query(ctx, query, ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, err
	}
	err = dRow.Scan(res.FirstName, res.LastName, res.MiddleName, res.BirthDate, res.IIN, res.Phone, res.Address, res.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while getting patient INFO from users: %v", err)
	}

	query = `SELECT (deparment_id, spec_id, experience, photo, category, price, schedule, degree, rating, website_url, user_id) FROM doctors WHERE id=$1`
	dRow, err = tx.Query(ctx, query, ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, err
	}
	err = dRow.Scan(&res.DepartmentId, &res.SpecId, &res.Experience, &res.Photo, &res.Category, &res.Price, &res.Schedule, &res.Degree, &res.Rating, &res.WebsiteUrl, &res.ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while getting doctor INFO from doctors: %v", err)
	}
	return res, nil
}
func (r *doctorRepository) GetAllDoctors() ([]*models.GetAllDoctorsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	var userID int64
	var first_name, last_name string
	query := `SELECT id, first_name, last_name FROM users`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w error occurred while getting rows from users: %v", models.ErrPatientNotFound, err)
		}
		return nil, fmt.Errorf("error occurred while getting rows from users: %v", err)
	}
	defer rows.Close()
	result := make([]*models.GetAllDoctorsResponse, 0, 100)
	for rows.Next() {
		err := rows.Scan(&userID, &first_name, &last_name)
		if err != nil {
			return nil, fmt.Errorf("%w error occurred while scanning row from users: %v", models.ErrPatientNotFound, err)
		}
		result = append(result, &models.GetAllDoctorsResponse{
			ID:        userID,
			FirstName: first_name,
			LastName:  last_name,
		})
	}
	return result, nil
}

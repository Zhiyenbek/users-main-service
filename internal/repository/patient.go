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

type patientRepository struct {
	db  *pgxpool.Pool
	cfg *config.DBConf
}

func NewPatientRepository(db *pgxpool.Pool, cfg *config.DBConf) PatientRepository {
	return &patientRepository{
		db:  db,
		cfg: cfg,
	}
}
func (r *patientRepository) CreatePatient(patient *models.CreatePatientRequest) (*models.CreatePatientResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	var userID, ID int64
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating deleting patient in users: %v", err)
	}
	query := `INSERT INTO users 
				(first_name, last_name, middle_name, birthdate, iin, phone, address, email)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`
	dRow, err := tx.Query(ctx, query, patient.FirstName, patient.LastName, patient.MiddleName, patient.BirthDate, patient.IIN, patient.Phone, patient.Address, patient.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating deleting patient in users: %v", err)
	}
	if dRow.Next() {
		err = dRow.Scan(&userID)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				log.Printf("ERROR: transaction: %s", errTX)
			}
			return nil, fmt.Errorf("error occurred while creating deleting patient in users: %v", err)
		}
	}
	dRow.Close()
	query = `INSERT INTO patients 
		(blood_type, emergency_contact, marital_status, id)
			VALUES
		($1, $2, $3, $4) RETURNING id;`
	dRow, err = tx.Query(ctx, query, patient.BloodType, patient.EmergencyContact, patient.MaritalStatus, userID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating deleting patient in patients: %v", err)
	}
	if dRow.Next() {
		err = dRow.Scan(&ID)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				log.Printf("ERROR: transaction: %s", errTX)
			}
			return nil, fmt.Errorf("error occurred while creating deleting patient in patients: %v", err)
		}
	}
	dRow.Close()
	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while deleting patient from users: %v", err)
	}
	return &models.CreatePatientResponse{
		ID: ID,
	}, nil

}
func (r *patientRepository) DeletePatient(ID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()

	query := `DELETE FROM users WHERE id=$1`
	rows, err := r.db.Exec(ctx, query, ID)
	if err != nil {
		return fmt.Errorf("error occurred while deleting patient: %v", err)
	}
	if rows.RowsAffected() < 1 {
		return fmt.Errorf("error: no patient in db with such id %d", ID)
	}
	return nil
}
func (r *patientRepository) UpdatePatient(patient *models.UpdatePatientRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	query := `UPDATE users 
				SET first_name = $1, last_name = $2, middle_name = $3, birthdate = $4, iin = $5, phone = $6, address = $7, email = $8
			  WHERE id = $9`
	_, err = tx.Exec(ctx, query, patient.FirstName, patient.LastName, patient.MiddleName, patient.BirthDate, patient.IIN, patient.Phone, patient.Address, patient.Email, patient.ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating patient info: %v", err)
	}

	query = `UPDATE patients 
				SET blood_type = $1, emergency_contact = $2, marital_status = $3
			  WHERE id = $4`
	_, err = tx.Exec(ctx, query, patient.BloodType, patient.EmergencyContact, patient.MaritalStatus, patient.ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating patient INFO in users: %v", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting patient from users: %v", err)
	}
	return nil
}
func (r *patientRepository) GetPatient(ID int64) (*models.GetPatientResponse, error) {
	res := &models.GetPatientResponse{}
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	query := `SELECT first_name, last_name, middle_name, birthdate, iin, phone, address, email FROM users WHERE id=$1`
	dRow, err := tx.Query(ctx, query, ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, err
	}

	if dRow.Next() {
		err = dRow.Scan(&res.FirstName, &res.LastName, &res.MiddleName, &res.BirthDate, &res.IIN, &res.Phone, &res.Address, &res.Email)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				log.Printf("ERROR: transaction: %s", errTX)
			}
			return nil, fmt.Errorf("error occurred while getting patient INFO from users: %v", err)
		}
	}
	dRow.Close()
	query = `SELECT blood_type, emergency_contact, marital_status FROM patients WHERE id=$1`
	dRow, err = tx.Query(ctx, query, ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction: %s", errTX)
		}
		return nil, err
	}
	if dRow.Next() {
		err = dRow.Scan(&res.BloodType, &res.EmergencyContact, &res.MaritalStatus)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				log.Printf("ERROR: transaction: %s", errTX)
			}
			return nil, fmt.Errorf("error occurred while getting patient INFO from patients: %v", err)
		}
	}
	res.ID = ID
	dRow.Close()
	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			log.Printf("ERROR: transaction error: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while deleting patient from users: %v", err)
	}
	return res, nil
}

func (r *patientRepository) GetAllPatients() ([]*models.GetAllPatientsResponse, error) {
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
	result := make([]*models.GetAllPatientsResponse, 0, 100)
	for rows.Next() {
		err := rows.Scan(&userID, &first_name, &last_name)
		if err != nil {
			return nil, fmt.Errorf("%w error occurred while scanning row from users: %v", models.ErrPatientNotFound, err)
		}
		result = append(result, &models.GetAllPatientsResponse{
			ID:        userID,
			FirstName: first_name,
			LastName:  last_name,
		})
	}
	return result, nil
}

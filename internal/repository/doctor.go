package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

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
				(first_name, last_name, middle_name, birthdate, iin, phone, address, email)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`
	dRow, err := tx.Query(ctx, query, doctor.FirstName, doctor.LastName, doctor.MiddleName, doctor.BirthDate, doctor.IIN, doctor.Phone, doctor.Address, doctor.Email)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return nil, fmt.Errorf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating doctor in users: %v", err)
	}
	if dRow.Next() {
		err = dRow.Scan(&userID)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				return nil, fmt.Errorf("ERROR: transaction: %s", errTX)
			}
			return nil, fmt.Errorf("error occurred while scanning doctor in users: %v", err)
		}
	}
	dRow.Close()
	query = `INSERT INTO doctors 
		(department_id, spec_id, experience, photo, category, price, schedule, degree, rating, website_url, id)
			VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id;`
	dRow, err = tx.Query(ctx, query, doctor.DepartmentId, doctor.SpecId, doctor.Experience, doctor.Photo, doctor.Category, doctor.Price, doctor.Schedule, doctor.Degree, doctor.Rating, doctor.WebsiteUrl, userID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return nil, fmt.Errorf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creaing doctors in doctors: %v", err)
	}
	if dRow.Next() {
		err = dRow.Scan(&ID)
		if err != nil {
			errTX := tx.Rollback(ctx)
			if errTX != nil {
				return nil, fmt.Errorf("ERROR: transaction: %s", errTX)
			}
			return nil, fmt.Errorf("error occurred while creating doctors in doctors: %v", err)
		}
	}
	dRow.Close()
	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return nil, fmt.Errorf("ERROR: transaction error: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating doctor: %v", err)
	}
	return &models.CreateDoctorResponse{
		ID: ID,
	}, nil
}

func (r *doctorRepository) DeleteDoctor(ID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()

	query := `DELETE FROM users WHERE id=$1`
	rows, err := r.db.Exec(ctx, query, ID)
	if err != nil {
		return fmt.Errorf("error occurred while deleting doctor: %v", err)
	}
	if rows.RowsAffected() < 1 {
		return fmt.Errorf("error: no doctor in db with such id %d", ID)
	}
	return nil
}

func (r *doctorRepository) UpdateDoctor(doctor *models.UpdateDoctorRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	query := `UPDATE users 
				SET first_name = $1, last_name = $2, middle_name = $3, birthdate = $4, iin = $5, phone = $6, address = $7, email = $8
			  WHERE id = $9`
	_, err = tx.Exec(ctx, query, doctor.FirstName, doctor.LastName, doctor.MiddleName, doctor.BirthDate, doctor.IIN, doctor.Phone, doctor.Address, doctor.Email, doctor.ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return fmt.Errorf("%w ERROR: transaction: %s", models.ErrInternalServer, errTX)
		}
		return fmt.Errorf("%w error occurred while updating doctor info in users table: %v", models.ErrInternalServer, err)
	}

	query = `UPDATE doctors 
				SET department_id = $1, spec_id = $2, experience = $3, photo = $4, category = $5, price = $6, schedule = $7, degree = $8, rating = $9, website_url = $10
			  WHERE id = $11`
	_, err = tx.Exec(ctx, query, doctor.DepartmentId, doctor.SpecId, doctor.Experience, doctor.Photo, doctor.Category, doctor.Price, doctor.Schedule, doctor.Degree, doctor.Rating, doctor.WebsiteUrl, doctor.ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return fmt.Errorf("ERROR: transaction: %s", errTX)
		}
		return fmt.Errorf("error occurred while updating doctor INFO in doctors: %v", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return fmt.Errorf("ERROR: transaction error: %s", errTX)
		}
		return fmt.Errorf("error occurred while deleting doctor from users: %v", err)
	}
	return nil
}

func (r *doctorRepository) GetDoctor(ID int64) (*models.GetDoctorResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()

	query := `SELECT
				users.id,
				users.first_name,
				users.last_name,
				users.middle_name,
				users.birthdate,
				users.phone,
				users.address,
				departments.name,
				specs.name,
				doctors.experience,
				doctors.photo,
				doctors.category,
				doctors.price,
				doctors.degree,
				doctors.rating
			FROM
				users
				INNER JOIN doctors ON doctors.id = users.id
				INNER JOIN specs ON doctors.spec_id = specs.id
				INNER JOIN departments ON doctors.department_id = departments.id
			WHERE
				users.ID = $1
			ORDER BY users.id desc
			`
	response := &models.GetDoctorResponse{}
	err := r.db.QueryRow(ctx, query, ID).Scan(
		&response.ID,
		&response.FirstName,
		&response.LastName,
		&response.MiddleName,
		&response.BirthDate,
		&response.Phone,
		&response.Address,
		&response.Department,
		&response.Spec,
		&response.Experience,
		&response.Phone,
		&response.Category,
		&response.Price,
		&response.Degree,
		&response.Rating,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w no doctors with such id in db", models.ErrDoctorNotFound)
		}
		return nil, fmt.Errorf("%w error occured while getting doctors from db %v", models.ErrInternalServer, err)
	}
	return response, nil
}

func (r *doctorRepository) GetAllDoctors() ([]*models.GetAllDoctorsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	var userID int64
	var first_name, last_name string
	query := `SELECT users.id, first_name, last_name FROM users JOIN doctors ON id=users.id`

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

func (r *doctorRepository) SearchDoctors(searchArgs *models.Search) (*models.SearchDoctorsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	offset := (searchArgs.PageNum - 1) * searchArgs.PageSize
	if offset < 0 {
		offset = 0
	}

	query := `SELECT
				users.id,
				users.first_name,
				users.last_name,
				users.middle_name,
				users.birthdate,
				users.phone,
				users.address,
				departments.name,
				specs.name,
				doctors.experience,
				doctors.photo,
				doctors.category,
				doctors.price,
				doctors.degree,
				doctors.rating
			FROM
				users
				INNER JOIN doctors ON doctors.id = users.id
				INNER JOIN specs ON doctors.spec_id = specs.id
				INNER JOIN departments ON doctors.department_id = departments.id
			WHERE
				users.first_name ILIKE '%' || $1 || '%' OR users.last_name ILIKE '%' || $1 || '%' OR departments.name ILIKE '%' || $1 || '%' OR specs.name ILIKE '%' || $1 || '%'
			LIMIT $2 OFFSET $3 
			`

	rows, err := r.db.Query(ctx, query, searchArgs.Search, searchArgs.PageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("error while searching doctors: %v", err)
	}
	defer rows.Close()
	var res []*models.GetDoctorResponse
	for rows.Next() {
		response := &models.GetDoctorResponse{}
		if err = rows.Scan(
			&response.ID,
			&response.FirstName,
			&response.LastName,
			&response.MiddleName,
			&response.BirthDate,
			&response.Phone,
			&response.Address,
			&response.Department,
			&response.Spec,
			&response.Experience,
			&response.Phone,
			&response.Category,
			&response.Price,
			&response.Degree,
			&response.Rating,
		); err != nil {
			return nil, fmt.Errorf("error while searching doctors in scan: %v", err)
		}
		res = append(res, response)
	}
	var count int
	query = `SELECT 
				COUNT(*)
			FROM 
				users 
					INNER JOIN doctors ON doctors.id = users.id 
					INNER JOIN specs ON doctors.spec_id = specs.id 
					INNER JOIN departments ON doctors.department_id = departments.id
			WHERE
				users.first_name ILIKE '%' || $1 || '%' OR users.last_name ILIKE '%' || $1 || '%' OR departments.name ILIKE '%' || $1 || '%' OR specs.name ILIKE '%' || $1 || '%'`
	err = r.db.QueryRow(ctx, query, searchArgs.Search).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error occured while scanning count in search %v", err)
	}
	return &models.SearchDoctorsResponse{
		Doctors: res,
		Count:   count,
	}, nil
}

func (r *doctorRepository) SearchDoctorsByDepartment(searchArgs *models.Search, ID int64) (*models.SearchDoctorsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	offset := (searchArgs.PageNum - 1) * searchArgs.PageSize
	if offset < 0 {
		offset = 0
	}

	query := `SELECT
				users.id,
				users.first_name,
				users.last_name,
				users.middle_name,
				users.birthdate,
				users.phone,
				users.address,
				departments.name,
				specs.name,
				doctors.experience,
				doctors.photo,
				doctors.category,
				doctors.price,
				doctors.degree,
				doctors.rating
			FROM
				users
				INNER JOIN doctors ON doctors.id = users.id
				INNER JOIN specs ON doctors.spec_id = specs.id
				INNER JOIN departments ON doctors.department_id = departments.id
			WHERE
				doctors.department_id = $1
			ORDER BY users.id desc
			LIMIT $2 OFFSET $3 
			`

	rows, err := r.db.Query(ctx, query, ID, searchArgs.PageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("error while searching doctors: %v", err)
	}
	defer rows.Close()
	var res []*models.GetDoctorResponse
	for rows.Next() {
		response := &models.GetDoctorResponse{}
		if err = rows.Scan(
			&response.ID,
			&response.FirstName,
			&response.LastName,
			&response.MiddleName,
			&response.BirthDate,
			&response.Phone,
			&response.Address,
			&response.Department,
			&response.Spec,
			&response.Experience,
			&response.Phone,
			&response.Category,
			&response.Price,
			&response.Degree,
			&response.Rating,
		); err != nil {
			return nil, fmt.Errorf("error while searching doctors by department in scan: %v", err)
		}
		res = append(res, response)
	}
	var count int
	query = `SELECT 
				COUNT(*)
			FROM 
				users 
					INNER JOIN doctors ON doctors.id = users.id 
					INNER JOIN specs ON doctors.spec_id = specs.id 
					INNER JOIN departments ON doctors.department_id = departments.id
			WHERE
				doctors.department_id = $1`
	err = r.db.QueryRow(ctx, query, ID).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error occured while scanning count in search %v", err)
	}
	return &models.SearchDoctorsResponse{
		Doctors: res,
		Count:   count,
	}, nil
}

func (r *doctorRepository) GetDepartments() (*models.GetDepartments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()
	query := `SELECT id, name from departments`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error occured while quering in getting departments %v", err)
	}
	defer rows.Close()
	var deps []*models.Department
	for rows.Next() {
		dep := &models.Department{}
		err := rows.Scan(&dep.ID, &dep.Name)
		if err != nil {
			return nil, fmt.Errorf("error occured while scanning in getting departments %v", err)
		}
		deps = append(deps, dep)
	}
	return &models.GetDepartments{
		Departments: deps,
	}, nil
}

func (r *doctorRepository) CreateAppointment(doctor *models.CreateAppointmentRequest) (*models.CreateAppointmentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()

	var id int
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("error occurred while creating appointment: %v", err)
	}

	query := `INSERT INTO appointments 
				(doc_id, email, phone, iin, reg_date, reg_time)
			VALUES
				($1, $2, $3, $4, $5, $6) RETURNING doc_id;`
	err = tx.QueryRow(ctx, query, doctor.Doctor_ID, doctor.Email, doctor.Phone, doctor.IIN, doctor.Reg_date, doctor.Reg_time).Scan(&id)
	log.Println(id)
	if err != nil {
		log.Println(err)
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return nil, fmt.Errorf("ERROR: transaction: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating appointment: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			return nil, fmt.Errorf("ERROR: transaction error: %s", errTX)
		}
		return nil, fmt.Errorf("error occurred while creating appointment: %v", err)
	}

	return &models.CreateAppointmentResponse{}, nil
}

func (r *doctorRepository) GetBookedAppointmentsByDate(bookArgs *models.Appointment) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.cfg.TimeOut)
	defer cancel()

	query := `SELECT
				appointments.time
			FROM
				appointments
			WHERE
				appointments.doc_id = $1 AND appointments.date = $2
			ORDER BY appointments.time
			`

	rows, err := r.db.Query(ctx, query, bookArgs.DoctorID, bookArgs.Date)
	if err != nil {
		return nil, fmt.Errorf("error while getting booked appointments: %v", err)
	}
	defer rows.Close()
	var res []string
	for rows.Next() {
		var time time.Time
		if err = rows.Scan(
			&time,
		); err != nil {
			return nil, fmt.Errorf("error while getting booked appointments in scan: %v", err)
		}

		res = append(res, time.Format("15:04"))
	}
	return res, nil
}

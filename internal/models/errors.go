package models

import "errors"

var (
	ErrPatientNotFound = errors.New("PATIENT_NOT_FOUND")
	ErrDoctorNotFound  = errors.New("DOCTOR_NOT_FOUND")
	ErrInvalidInput    = errors.New("INVALID_INPUT")
	ErrInternalServer  = errors.New("INTERNAL_SERVER_ERROR")
)

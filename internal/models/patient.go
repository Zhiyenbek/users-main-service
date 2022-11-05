package models

// import "encoding/json"

type CreatePatientRequest struct {
	BloodType        int32
	EmergencyContact string
	MaritalStatus    string
	FirstName        string
	LastName         string
	MiddleName       string
	BirthDate        string
	IIN              int64
	Phone            string
	Address          string
	Email            string
}

type CreatePatientResponse struct {
	ID int64
}

type UpdatePatientRequest struct {
	ID               int64
	BloodType        int32
	EmergencyContact string
	MaritalStatus    string
	FirstName        string
	LastName         string
	MiddleName       string
	BirthDate        string
	IIN              int64
	Phone            string
	Address          string
	Email            string
}

type GetPatientResponse struct {
	ID               int64
	BloodType        int32
	EmergencyContact string
	MaritalStatus    string
	FirstName        string
	LastName         string
	MiddleName       string
	BirthDate        string
	IIN              int64
	Phone            string
	Address          string
	Email            string
}
package models

// import "encoding/json"

type CreatePatientRequest struct {
	BloodType        int32  `json:"blood_type" binding:"required"`
	EmergencyContact string `json:"emer_contact" binding:"required"`
	MaritalStatus    string `json:"marit_st" binding:"required"`
	FirstName        string `json:"first_name" binding:"required"`
	LastName         string `json:"last_name" binding:"required"`
	MiddleName       string `json:"middle_name" `
	BirthDate        string `json:"birth_date" binding:"required"`
	IIN              string `json:"iin" binding:"required"`
	Phone            string `json:"phone" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Email            string `json:"email" binding:"required"`
}

type CreatePatientResponse struct {
	ID int64
}

type UpdatePatientRequest struct {
	ID               int64
	BloodType        int32  `json:"blood_type" binding:"required"`
	EmergencyContact string `json:"emer_contact" binding:"required"`
	MaritalStatus    string `json:"marit_st" binding:"required"`
	FirstName        string `json:"first_name" binding:"required"`
	LastName         string `json:"last_name" binding:"required"`
	MiddleName       string `json:"middle_name" `
	BirthDate        string `json:"birth_date" binding:"required"`
	IIN              string `json:"iin" binding:"required"`
	Phone            string `json:"phone" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Email            string `json:"email" binding:"required"`
}

type GetPatientResponse struct {
	ID               int64
	BloodType        int32  `json:"blood_type" binding:"required"`
	EmergencyContact string `json:"emer_contact" binding:"required"`
	MaritalStatus    string `json:"marit_st" binding:"required"`
	FirstName        string `json:"first_name" binding:"required"`
	LastName         string `json:"last_name" binding:"required"`
	MiddleName       string `json:"middle_name" `
	BirthDate        string `json:"birth_date" binding:"required"`
	IIN              string `json:"iin" binding:"required"`
	Phone            string `json:"phone" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Email            string `json:"email" binding:"required"`
}

type GetAllPatientsResponse struct {
	ID        int64
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	IIN       string `json:"iin" binding:"required"`
}

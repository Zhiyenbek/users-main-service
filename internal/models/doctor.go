package models

type CreateDoctorRequest struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	MiddleName   string `json:"middle_name" `
	BirthDate    string `json:"birth_date" binding:"required"`
	IIN          string `json:"iin" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Email        string `json:"email" binding:"required"`
	DepartmentId int32  `json:"department_id" binding:"required"`
	SpecId       int32  `json:"spec_id" binding:"required"`
	Experience   int32  `json:"experience" binding:"required"`
	Photo        string `json:"photo" binding:"required"`
	Category     string `json:"category" binding:"required"`
	Price        int32  `json:"price" binding:"required"`
	Schedule     string `json:"schedule" binding:"required"`
	Degree       string `json:"degree" binding:"required"`
	Rating       int32  `json:"rating" binding:"required"`
	WebsiteUrl   string `json:"webstite_url" `
}

type CreateDoctorResponse struct {
	ID int64
}

type UpdateDoctorRequest struct {
	ID           int64
	FirstName    string
	LastName     string
	MiddleName   string
	BirthDate    string
	IIN          string
	Phone        string
	Address      string
	Email        string
	DepartmentId int32
	SpecId       int32
	Experience   int32
	Photo        string
	Category     string
	Price        int32
	Schedule     string
	Degree       string
	Rating       int32
	WebsiteUrl   string
}

type GetDoctorResponse struct {
	ID           int64
	FirstName    string
	LastName     string
	MiddleName   string
	BirthDate    string
	IIN          string
	Phone        string
	Address      string
	Email        string
	DepartmentId int32
	SpecId       int32
	Experience   int32
	Photo        string
	Category     string
	Price        int32
	Schedule     string
	Degree       string
	Rating       int32
	WebsiteUrl   string
}

type GetAllDoctorsResponse struct {
	ID        int64
	FirstName string
	LastName  string
	IIN       string
}

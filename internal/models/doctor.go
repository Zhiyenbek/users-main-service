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

type GetDoctorResponse struct {
	ID         int64  `json:"doctor_id" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	MiddleName string `json:"middle_name" `
	BirthDate  string `json:"birth_date" binding:"required"`
	IIN        string `json:"iin" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Address    string `json:"address" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Department string `json:"department" binding:"required"`
	Spec       string `json:"spec" binding:"required"`
	Experience int32  `json:"experience" binding:"required"`
	Photo      string `json:"photo" binding:"required"`
	Category   string `json:"category" binding:"required"`
	Price      int32  `json:"price" binding:"required"`
	Schedule   string `json:"schedule" binding:"required"`
	Degree     string `json:"degree" binding:"required"`
	Rating     int32  `json:"rating" binding:"required"`
	WebsiteUrl string `json:"webstite_url" `
}

type GetAllDoctorsResponse struct {
	ID        int64
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	IIN       string `json:"iin" binding:"required"`
}

type SearchDoctorsResponse struct {
	Doctors []*GetDoctorResponse `json:"doctors"`
	Count   int                  `json:"count"`
}

package models

type CreateDoctorRequest struct {
	FirstName    string
	LastName     string
	MiddleName   string
	BirthDate    string
	IIN          int64
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

type CreateDoctorResponse struct {
	ID int64
}

type UpdateDoctorRequest struct {
	ID           int64
	FirstName    string
	LastName     string
	MiddleName   string
	BirthDate    string
	IIN          int64
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
	IIN          int64
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

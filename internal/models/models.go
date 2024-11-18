package models

import "time"

type Country struct {
	CName      string `gorm:"primaryKey;column:cname"`
	Population int64  `gorm:"column:population"`
}

type Users struct {
	Email   string `gorm:"primaryKey;column:email"`
	Name    string `gorm:"column:name"`
	Surname string `gorm:"column:surname"`
	Salary  int    `gorm:"column:salary"`
	Phone   string `gorm:"column:phone"`
	CName   string `gorm:"column:cname"`
}

type Patients struct {
	Email string `gorm:"primaryKey;column:email"`
}

type DiseaseType struct {
	ID          int    `gorm:"primaryKey;column:id"`
	Description string `gorm:"column:description"`
}

type Disease struct {
	DiseaseCode string `gorm:"primaryKey;column:disease_code"`
	Pathogen    string `gorm:"column:pathogen"`
	Description string `gorm:"column:description"`
	ID          int    `gorm:"column:id"`
}

type Discover struct {
	CName        string    `gorm:"primaryKey;column:cname"`
	DiseaseCode  string    `gorm:"primaryKey;column:disease_code"`
	FirstEncDate time.Time `gorm:"column:first_enc_date"`
}

type PatientDisease struct {
	Email       string `gorm:"primaryKey;column:email"`
	DiseaseCode string `gorm:"primaryKey;column:disease_code"`
}

type PublicServant struct {
	Email      string `gorm:"primaryKey;column:email"`
	Department string `gorm:"column:department"`
}

type Doctor struct {
	Email  string `gorm:"primaryKey;column:email"`
	Degree string `gorm:"column:degree"`
}

type Specialize struct {
	ID    int    `gorm:"primaryKey;column:id"`
	Email string `gorm:"primaryKey;column:email"`
}

type Record struct {
	Email        string `gorm:"primaryKey;column:email"`
	CName        string `gorm:"primaryKey;column:cname"`
	DiseaseCode  string `gorm:"primaryKey;column:disease_code"`
	TotalDeaths  int    `gorm:"column:total_deaths"`
	TotalPatients int   `gorm:"column:total_patients"`
}


// Requiered for responses
type DoctorInfo struct {
	Name   string `json:"name"`
	Degree string `json:"degree"`
}

type DoctorInfoFullName struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Degree  string `json:"degree"`
}

type CountrySalary struct {
	Country   string  `json:"country"`
	AvgSalary float64 `json:"avg_salary"`
}

type DepartmentCount struct {
	Department    string `json:"department"`
	EmployeeCount int    `json:"employee_count"`
}

type CountryPatientCount struct {
	Country       string `json:"country"`
	TotalPatients int    `json:"total_patients"`
}

type PatientDiseaseView struct {
	FullName          string `json:"full_name"`
	DiseaseDescription string `json:"disease_description"`
}

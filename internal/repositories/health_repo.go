package repositories

import (
	"log"

	"github.com/srmbackisdeveloper/db-assignment3/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HealthRepository struct {
	DB *gorm.DB
}

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL")

	// // Auto-migrate all models
	// if err := db.AutoMigrate(); err != nil {
	// 	return nil, err
	// }
	// log.Println("Database migrated successfully")

	return db, nil
}

func NewHealthRepository(db *gorm.DB) *HealthRepository {
	return &HealthRepository{DB: db}
}

// 1. Basic Retrieval Queries
// 1.1 List all diseases caused by "bacteria" discovered before 2020
func (repo *HealthRepository) GetBacterialDiseasesBefore2020() ([]models.Disease, error) {
	var diseases []models.Disease
	query := `
		SELECT Disease.*
		FROM Disease
		JOIN Discover ON Disease.disease_code = Discover.disease_code
		WHERE pathogen = 'bacteria' AND first_enc_date < '2020-01-01'
	`
	if err := repo.DB.Raw(query).Scan(&diseases).Error; err != nil {
		return nil, err
	}
	return diseases, nil
}

// 1.2 Retrieve names and degrees of doctors not specialized in “infectious diseases”
func (repo *HealthRepository) GetDoctorsNotSpecializedIn(diseaseType string) ([]models.DoctorInfo, error) {
	var doctors []models.DoctorInfo
	query := `
		SELECT Users.name, Doctor.degree
		FROM Users
		JOIN Doctor ON Users.email = Doctor.email
		WHERE Doctor.email NOT IN (
			SELECT email
			FROM Specialize
			JOIN DiseaseType ON Specialize.id = DiseaseType.id
			WHERE DiseaseType.description = ?
		)
	`
	if err := repo.DB.Raw(query, diseaseType).Scan(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

// 1.3 List doctors specialized in more than 2 disease types
func (repo *HealthRepository) GetDoctorsSpecializedInMoreThanTwoDiseases() ([]models.DoctorInfoFullName, error) {
    var doctors []models.DoctorInfoFullName
	query := `
		SELECT Users.name, Users.surname, Doctor.degree
		FROM Users
		JOIN Doctor ON Users.email = Doctor.email
		WHERE Users.email IN (
			SELECT email
			FROM Specialize
			GROUP BY email
			HAVING COUNT(id) > 2
		)
	`
	if err := repo.DB.Raw(query).Scan(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

// 2. Complex Queries with Aggregation
// 2.1 List countries and average salary of doctors specialized in "virology"
func (repo *HealthRepository) GetAverageSalaryByCountryForSpecialization(specialization string) ([]models.CountrySalary, error) {
	var results []models.CountrySalary
	query := `
		SELECT Users.cname AS country, AVG(Users.salary) AS avg_salary
		FROM Users
		JOIN Doctor ON Users.email = Doctor.email
		JOIN Specialize ON Doctor.email = Specialize.email
		JOIN DiseaseType ON Specialize.id = DiseaseType.id
		WHERE DiseaseType.description = ?
		GROUP BY Users.cname
	`
	if err := repo.DB.Raw(query, specialization).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// 2.2 Find departments with public servants reporting "covid-19" cases across multiple
// countries, counting such employees per department. (i.e “Dept1 8” means that in the “Dept1” department there are 8 such employees.)
func (repo *HealthRepository) GetDepartmentsReportingCovid() ([]models.DepartmentCount, error) {
	var results []models.DepartmentCount
	query := `
		SELECT PublicServant.department, COUNT(DISTINCT PublicServant.email) AS employee_count
		FROM PublicServant
		JOIN Record ON PublicServant.email = Record.email
		WHERE Record.disease_code = 'COVID-19'
		GROUP BY PublicServant.department
	`
	if err := repo.DB.Raw(query).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// 3. Update and Maintenance Queries
// 3.1 Double salary of public servants with more than 3 "covid-19" patients
func (repo *HealthRepository) DoubleSalaryForCovidPublicServants() error {
	query := `
		UPDATE Users
		SET salary = salary * 2
		WHERE email IN (
			SELECT email
			FROM Record
			WHERE disease_code = 'COVID-19' AND total_patients > 3
		)
	`
	return repo.DB.Exec(query).Error
}

// 3.2 Delete users whose names contain "bek" or "gul"
func (repo *HealthRepository) DeleteUsersWithNameSubstring() error {
	query := `
		DELETE FROM Users
		WHERE name ILIKE '%bek%' OR name ILIKE '%gul%'
	`
	return repo.DB.Exec(query).Error
}

// 4. Indexing
// 4.1 Create a primary indexing on the "email" field in Users table
func (repo *HealthRepository) CreatePrimaryIndexOnEmail() error {
	query := `
		ALTER TABLE Users
		ADD CONSTRAINT pk_users_email PRIMARY KEY (email)
	`
	return repo.DB.Exec(query).Error
}

// 4.2 Create a secondary indexing on the "disease_code" field in Disease table
func (repo *HealthRepository) CreateSecondaryIndexOnDiseaseCode() error {
	query := `
		CREATE INDEX IF NOT EXISTS idx_disease_code ON Disease(disease_code)
	`
	return repo.DB.Exec(query).Error
}


// 5. Additional Analysis
// 5.1 List top 2 countries with highest total patients recorded
func (repo *HealthRepository) GetTopCountriesByTotalPatients(limit int) ([]models.CountryPatientCount, error) {
	var results []models.CountryPatientCount
	query := `
		SELECT cname AS country, SUM(total_patients) AS total_patients
		FROM Record
		GROUP BY cname
		ORDER BY total_patients DESC
		LIMIT ?
	`
	if err := repo.DB.Raw(query, limit).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// 6. Query with a Derived Attribute
// 6.1 Calculate total number of patients who have covid-19 disease
func (repo *HealthRepository) GetTotalCovidPatients() (int, error) {
	var totalPatients int
	query := `
		SELECT COALESCE(SUM(total_patients), 0) AS total_patients
		FROM Record
		WHERE disease_code = 'COVID-19'
	`
	if err := repo.DB.Raw(query).Scan(&totalPatients).Error; err != nil {
		return 0, err
	}
	return totalPatients, nil
}

// 7. View Operation
// 7.1 Create a view with all patients’ names and diseases
func (repo *HealthRepository) CreatePatientDiseaseView() error {
	query := `
		CREATE OR REPLACE VIEW PatientDiseaseView AS
		SELECT 
			Users.name AS patient_name,
			Users.surname AS patient_surname,
			Disease.description AS disease_description
		FROM Users
		JOIN Patients ON Users.email = Patients.email
		JOIN PatientDisease ON Patients.email = PatientDisease.email
		JOIN Disease ON PatientDisease.disease_code = Disease.disease_code
	`
	return repo.DB.Exec(query).Error
}

// 8. Querying the View
// 8.1 Retrieve list of all patients’ full names and diseases from the view
func (repo *HealthRepository) GetPatientsFromView() ([]models.PatientDiseaseView, error) {
	var results []models.PatientDiseaseView
	query := `
		SELECT 
			TRIM(patient_name || ' ' || patient_surname) AS full_name,
			disease_description
		FROM PatientDiseaseView
	`
	if err := repo.DB.Raw(query).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

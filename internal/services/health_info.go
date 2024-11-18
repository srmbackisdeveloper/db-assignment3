package services

import (
	"github.com/srmbackisdeveloper/db-assignment3/internal/models"
	"github.com/srmbackisdeveloper/db-assignment3/internal/repositories"
)

type HealthService struct {
	HealthRepo *repositories.HealthRepository
}

func NewHealthService(healthRepo *repositories.HealthRepository) *HealthService {
	return &HealthService{
		HealthRepo: healthRepo,
	}
}

// 1. Basic Retrieval Queries
// 1.1 List all diseases caused by "bacteria" discovered before 2020
func (s *HealthService) GetBacterialDiseasesBefore2020() ([]models.Disease, error) {
	return s.HealthRepo.GetBacterialDiseasesBefore2020()
}

// 1.2 Retrieve names and degrees of doctors not specialized in “infectious diseases”
func (s *HealthService) GetDoctorsNotSpecializedIn(diseaseType string) ([]models.DoctorInfo, error) {
	return s.HealthRepo.GetDoctorsNotSpecializedIn(diseaseType)
}

// 1.3 List doctors specialized in more than 2 disease types
func (s *HealthService) GetDoctorsSpecializedInMoreThanTwoDiseases() ([]models.DoctorInfoFullName, error) {
	return s.HealthRepo.GetDoctorsSpecializedInMoreThanTwoDiseases()
}

// 2. Complex Queries with Aggregation
// 2.1 List countries and average salary of doctors specialized in "virology"
func (s *HealthService) GetAverageSalaryByCountryForSpecialization(specialization string) ([]models.CountrySalary, error) {
	return s.HealthRepo.GetAverageSalaryByCountryForSpecialization(specialization)
}

// 2.2 Find departments with public servants reporting "covid-19" cases across multiple countries
func (s *HealthService) GetDepartmentsReportingCovid() ([]models.DepartmentCount, error) {
	return s.HealthRepo.GetDepartmentsReportingCovid()
}

// 3. Update and Maintenance Queries
// 3.1 Double salary of public servants with more than 3 "covid-19" patients
func (s *HealthService) DoubleSalaryForCovidPublicServants() error {
	return s.HealthRepo.DoubleSalaryForCovidPublicServants()
}

// 3.2 Delete users whose names contain "bek" or "gul"
func (s *HealthService) DeleteUsersWithNameSubstring() error {
	return s.HealthRepo.DeleteUsersWithNameSubstring()
}

// 4. Indexing
// 4.1 Create a primary indexing on the "email" field in Users table
func (s *HealthService) CreatePrimaryIndexOnEmail() error {
	return s.HealthRepo.CreatePrimaryIndexOnEmail()
}

// 4.2 Create a secondary indexing on the "disease_code" field in Disease table
func (s *HealthService) CreateSecondaryIndexOnDiseaseCode() error {
	return s.HealthRepo.CreateSecondaryIndexOnDiseaseCode()
}

// 5. Additional Analysis
// 5.1 List top 2 countries with highest total patients recorded
func (s *HealthService) GetTopCountriesByTotalPatients(limit int) ([]models.CountryPatientCount, error) {
	return s.HealthRepo.GetTopCountriesByTotalPatients(limit)
}

// 6. Query with a Derived Attribute
// 6.1 Calculate total number of patients who have covid-19 disease
func (s *HealthService) GetTotalCovidPatients() (int, error) {
	return s.HealthRepo.GetTotalCovidPatients()
}

// 7. View Operation
// 7.1 Create a view with all patients’ names and diseases
func (s *HealthService) CreatePatientDiseaseView() error {
	return s.HealthRepo.CreatePatientDiseaseView()
}

// 8. Querying the View
// 8.1 Retrieve list of all patients’ full names and diseases from the view
func (s *HealthService) GetPatientsFromView() ([]models.PatientDiseaseView, error) {
	return s.HealthRepo.GetPatientsFromView()
}

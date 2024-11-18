package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srmbackisdeveloper/db-assignment3/internal/services"
)

type HealthHandler struct {
	HealthService *services.HealthService
}

func NewHealthHandler(service *services.HealthService) *HealthHandler {
	return &HealthHandler{HealthService: service}
}

// 1. Basic Retrieval Queries
// 1.1 List all diseases caused by "bacteria" discovered before 2020
func (h *HealthHandler) GetBacterialDiseasesBefore2020Handler(c *gin.Context) {
	diseases, err := h.HealthService.GetBacterialDiseasesBefore2020()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, diseases)
}

// 1.2 Retrieve names and degrees of doctors not specialized in "infectious diseases"
func (h *HealthHandler) GetDoctorsNotSpecializedInHandler(c *gin.Context) {
	diseaseType := c.Query("disease_type")
	if diseaseType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'disease_type' is required"})
		return
	}

	doctors, err := h.HealthService.GetDoctorsNotSpecializedIn(diseaseType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
}

// 1.3 List doctors specialized in more than 2 disease types
func (h *HealthHandler) GetDoctorsSpecializedInMoreThanTwoDiseasesHandler(c *gin.Context) {
	doctors, err := h.HealthService.GetDoctorsSpecializedInMoreThanTwoDiseases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
}

// 2. Complex Queries with Aggregation
// 2.1 List countries and average salary of doctors specialized in "virology"
func (h *HealthHandler) GetAverageSalaryByCountryForSpecializationHandler(c *gin.Context) {
	specialization := c.Query("specialization")
	if specialization == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'specialization' is required"})
		return
	}

	results, err := h.HealthService.GetAverageSalaryByCountryForSpecialization(specialization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

// 2.2 Find departments with public servants reporting "covid-19" cases across multiple countries
func (h *HealthHandler) GetDepartmentsReportingCovidHandler(c *gin.Context) {
	results, err := h.HealthService.GetDepartmentsReportingCovid()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

// 3. Update and Maintenance Queries
// 3.1 Double salary of public servants with more than 3 "covid-19" patients
func (h *HealthHandler) DoubleSalaryForCovidPublicServantsHandler(c *gin.Context) {
	err := h.HealthService.DoubleSalaryForCovidPublicServants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Salaries updated successfully"})
}

// 3.2 Delete users whose names contain "bek" or "gul"
func (h *HealthHandler) DeleteUsersWithNameSubstringHandler(c *gin.Context) {
	err := h.HealthService.DeleteUsersWithNameSubstring()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Users deleted successfully"})
}

// 4. Indexing
// 4.1 Create a primary indexing on the "email" field in Users table
func (h *HealthHandler) CreatePrimaryIndexOnEmailHandler(c *gin.Context) {
	err := h.HealthService.CreatePrimaryIndexOnEmail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Primary index on email created successfully"})
}

// 4.2 Create a secondary indexing on the "disease_code" field in Disease table
func (h *HealthHandler) CreateSecondaryIndexOnDiseaseCodeHandler(c *gin.Context) {
	err := h.HealthService.CreateSecondaryIndexOnDiseaseCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Secondary index on disease_code created successfully"})
}

// 5. Additional Analysis
// 5.1 List top 2 countries with highest total patients recorded
func (h *HealthHandler) GetTopCountriesByTotalPatientsHandler(c *gin.Context) {
	limit := 2

	results, err := h.HealthService.GetTopCountriesByTotalPatients(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

// 6. Query with a Derived Attribute
// 6.1 Calculate total number of patients who have covid-19 disease
func (h *HealthHandler) GetTotalCovidPatientsHandler(c *gin.Context) {
	total, err := h.HealthService.GetTotalCovidPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_covid_patients": total})
}

// 7. View Operation
// 7.1 Create a view with all patients’ names and diseases
func (h *HealthHandler) CreatePatientDiseaseViewHandler(c *gin.Context) {
	err := h.HealthService.CreatePatientDiseaseView()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "PatientDiseaseView created successfully"})
}

// 8. Querying the View
// 8.1 Retrieve list of all patients’ full names and diseases from the view
func (h *HealthHandler) GetPatientsFromViewHandler(c *gin.Context) {
	results, err := h.HealthService.GetPatientsFromView()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

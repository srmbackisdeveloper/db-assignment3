package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/srmbackisdeveloper/db-assignment3/config"
	"github.com/srmbackisdeveloper/db-assignment3/internal/handlers"
	"github.com/srmbackisdeveloper/db-assignment3/internal/repositories"
	"github.com/srmbackisdeveloper/db-assignment3/internal/services"
)

func main() {
	// configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// repos
	db, err := repositories.NewPostgresDB(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("failed to connect to PostgreSQL: %v", err)
    }
    defer func() {
        sqlDB, _ := db.DB()
        sqlDB.Close()
    }()

	healthRepo := repositories.NewHealthRepository(db)
	
	// services
	healthService := services.NewHealthService(healthRepo)

	// handlers
    healthHandler := handlers.NewHealthHandler(healthService)

    // server
    router := gin.Default()
    port := cfg.Port


	// Basic Retrieval Queries
	router.GET("/diseases/bacteria", healthHandler.GetBacterialDiseasesBefore2020Handler)
	router.GET("/doctors/not-specialized", healthHandler.GetDoctorsNotSpecializedInHandler)
	router.GET("/doctors/specialized/more-than-two", healthHandler.GetDoctorsSpecializedInMoreThanTwoDiseasesHandler)

	router.GET("/countries/average-salary", healthHandler.GetAverageSalaryByCountryForSpecializationHandler)
	router.GET("/departments/reporting-covid", healthHandler.GetDepartmentsReportingCovidHandler)

	router.PUT("/public-servants/double-salary", healthHandler.DoubleSalaryForCovidPublicServantsHandler)
	router.DELETE("/users/delete-by-name", healthHandler.DeleteUsersWithNameSubstringHandler)

	router.POST("/indexes/email", healthHandler.CreatePrimaryIndexOnEmailHandler)
	router.POST("/indexes/disease-code", healthHandler.CreateSecondaryIndexOnDiseaseCodeHandler)

	router.GET("/countries/top-patients", healthHandler.GetTopCountriesByTotalPatientsHandler)

	router.GET("/covid/total-patients", healthHandler.GetTotalCovidPatientsHandler)

	router.POST("/views/patient-disease", healthHandler.CreatePatientDiseaseViewHandler)

	router.GET("/patients/diseases", healthHandler.GetPatientsFromViewHandler)

    log.Printf("Starting server on port %s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}
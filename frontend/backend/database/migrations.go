package database

import (
	"log"
	"medworld-backend/models"
)

// RunMigrations creates all database tables
func RunMigrations() error {
	log.Println("Running database migrations...")

	// Auto-migrate all models
	err := DB.AutoMigrate(
		&models.User{},
		&models.Cabinet{},
		&models.Doctor{},
		&models.Patient{},
		&models.Assistant{},
		&models.Appointment{},
		&models.Consultation{},
		&models.Message{},
		&models.Calendar{},
	)

	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}

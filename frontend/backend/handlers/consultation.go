package handlers

import (
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ConsultationHandler handles consultation-related endpoints
type ConsultationHandler struct{}

// ListConsultations returns all consultations (filtered by user role)
func (h *ConsultationHandler) ListConsultations(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	query := database.DB.Model(&models.Consultation{}).
		Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Appointment")

	// Filter based on user type
	switch currentUser.Type {
	case models.UserTypeDoctor, models.UserTypeAdmin:
		var doctor models.Doctor
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&doctor).Error; err == nil {
			query = query.Where("doctor_id = ?", doctor.ID)
		}

	case models.UserTypePatient:
		var patient models.Patient
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&patient).Error; err == nil {
			query = query.Where("patient_id = ?", patient.ID)
		}

	case models.UserTypeSuperAdmin:
		// No filter
	}

	var consultations []models.Consultation
	if err := query.Order("created_at DESC").Find(&consultations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch consultations",
		})
	}

	return c.JSON(fiber.Map{
		"data": consultations,
	})
}

// GetConsultationByID returns a consultation by ID
func (h *ConsultationHandler) GetConsultationByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid consultation ID",
		})
	}

	var consultation models.Consultation
	if err := database.DB.Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Appointment").
		First(&consultation, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Consultation not found",
		})
	}

	return c.JSON(consultation)
}

// CreateConsultation creates a new consultation
func (h *ConsultationHandler) CreateConsultation(c *fiber.Ctx) error {
	var consultation models.Consultation
	if err := c.BodyParser(&consultation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Validate required fields
	if consultation.DoctorID == 0 || consultation.PatientID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Doctor and patient are required",
		})
	}

	// Create consultation
	if err := database.DB.Create(&consultation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create consultation",
		})
	}

	// Reload with associations
	database.DB.Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Appointment").
		First(&consultation, consultation.ID)

	return c.Status(fiber.StatusCreated).JSON(consultation)
}

// UpdateConsultation updates a consultation
func (h *ConsultationHandler) UpdateConsultation(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid consultation ID",
		})
	}

	var consultation models.Consultation
	if err := database.DB.First(&consultation, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Consultation not found",
		})
	}

	// Parse updates
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Update consultation
	if err := database.DB.Model(&consultation).Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update consultation",
		})
	}

	// Reload with associations
	database.DB.Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Appointment").
		First(&consultation, id)

	return c.JSON(consultation)
}

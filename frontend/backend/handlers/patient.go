package handlers

import (
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// PatientHandler handles patient-related endpoints
type PatientHandler struct{}

// ListPatients returns all patients
func (h *PatientHandler) ListPatients(c *fiber.Ctx) error {
	var patients []models.Patient
	query := database.DB.Preload("User")

	if err := query.Find(&patients).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch patients",
		})
	}

	// Transform to include user data
	result := make([]map[string]interface{}, 0)
	for _, patient := range patients {
		if patient.User != nil {
			userData := getUserWithExtendedInfo(patient.User)
			result = append(result, userData.(map[string]interface{}))
		}
	}

	return c.JSON(fiber.Map{
		"data": result,
	})
}

// GetPatientByID returns a patient by ID
func (h *PatientHandler) GetPatientByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid patient ID",
		})
	}

	var patient models.Patient
	if err := database.DB.Preload("User").First(&patient, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Patient not found",
		})
	}

	userData := getUserWithExtendedInfo(patient.User)
	return c.JSON(userData)
}

// GetPatientAppointments returns appointments for a patient
func (h *PatientHandler) GetPatientAppointments(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid patient ID",
		})
	}

	var appointments []models.Appointment
	query := database.DB.Where("patient_id = ?", id).
		Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Cabinet").
		Order("date DESC")

	if err := query.Find(&appointments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch appointments",
		})
	}

	return c.JSON(fiber.Map{
		"data": transformAppointments(appointments),
	})
}

// GetPatientConsultations returns consultations for a patient
func (h *PatientHandler) GetPatientConsultations(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid patient ID",
		})
	}

	var consultations []models.Consultation
	query := database.DB.Where("patient_id = ?", id).
		Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Appointment").
		Order("created_at DESC")

	if err := query.Find(&consultations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch consultations",
		})
	}

	return c.JSON(fiber.Map{
		"data": consultations,
	})
}

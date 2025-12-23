package handlers

import (
	"medworld-backend/database"
	"medworld-backend/models"

	"github.com/gofiber/fiber/v2"
)

// AllHandler handles cross-entity listing endpoints (admin only)
type AllHandler struct{}

// ListAllAppointments returns all appointments in the system
func (h *AllHandler) ListAllAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	if err := database.DB.
		Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Cabinet").
		Order("date DESC").
		Find(&appointments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch appointments",
		})
	}

	return c.JSON(fiber.Map{
		"data": transformAppointments(appointments),
	})
}

// ListAllDoctors returns all doctors in the system
func (h *AllHandler) ListAllDoctors(c *fiber.Ctx) error {
	var doctors []models.Doctor
	if err := database.DB.Preload("User").Preload("Cabinet").Find(&doctors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch doctors",
		})
	}

	result := make([]map[string]interface{}, 0)
	for _, doctor := range doctors {
		if doctor.User != nil {
			userData := getUserWithExtendedInfo(doctor.User)
			result = append(result, userData.(map[string]interface{}))
		}
	}

	return c.JSON(fiber.Map{
		"data": result,
	})
}

// ListAllPatients returns all patients in the system
func (h *AllHandler) ListAllPatients(c *fiber.Ctx) error {
	var patients []models.Patient
	if err := database.DB.Preload("User").Find(&patients).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch patients",
		})
	}

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

// ListAllAssistants returns all assistants in the system
func (h *AllHandler) ListAllAssistants(c *fiber.Ctx) error {
	var assistants []models.Assistant
	if err := database.DB.Preload("User").Preload("Cabinet").Find(&assistants).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch assistants",
		})
	}

	result := make([]map[string]interface{}, 0)
	for _, assistant := range assistants {
		if assistant.User != nil {
			userData := getUserWithExtendedInfo(assistant.User)
			result = append(result, userData.(map[string]interface{}))
		}
	}

	return c.JSON(fiber.Map{
		"data": result,
	})
}

// ListAllCabinets returns all cabinets in the system
func (h *AllHandler) ListAllCabinets(c *fiber.Ctx) error {
	var cabinets []models.Cabinet
	if err := database.DB.Preload("Admin").Find(&cabinets).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch cabinets",
		})
	}

	return c.JSON(fiber.Map{
		"data": cabinets,
	})
}

// ListAllUsers returns all users in the system
func (h *AllHandler) ListAllUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
		})
	}

	result := make([]map[string]interface{}, 0)
	for i := range users {
		userData := getUserWithExtendedInfo(&users[i])
		result = append(result, userData.(map[string]interface{}))
	}

	return c.JSON(fiber.Map{
		"data": result,
	})
}

// ListAllConsultations returns all consultations in the system
func (h *AllHandler) ListAllConsultations(c *fiber.Ctx) error {
	var consultations []models.Consultation
	if err := database.DB.
		Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Appointment").
		Order("created_at DESC").
		Find(&consultations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch consultations",
		})
	}

	return c.JSON(fiber.Map{
		"data": consultations,
	})
}

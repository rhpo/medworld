package handlers

import (
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AppointmentHandler handles appointment-related endpoints
type AppointmentHandler struct{}

// ListAppointments returns all appointments (filtered by user role)
func (h *AppointmentHandler) ListAppointments(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	query := database.DB.Model(&models.Appointment{}).
		Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Cabinet")

	// Filter based on user type
	switch currentUser.Type {
	case models.UserTypeDoctor, models.UserTypeAdmin:
		// Get doctor record
		var doctor models.Doctor
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&doctor).Error; err == nil {
			query = query.Where("doctor_id = ?", doctor.ID)
		}

	case models.UserTypePatient:
		// Get patient record
		var patient models.Patient
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&patient).Error; err == nil {
			query = query.Where("patient_id = ?", patient.ID)
		}

	case models.UserTypeAssistant:
		// Get assistant's cabinet
		var assistant models.Assistant
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&assistant).Error; err == nil {
			query = query.Where("cabinet_id = ?", assistant.CabinetID)
		}

	case models.UserTypeSuperAdmin:
		// No filter - can see all
	}

	var appointments []models.Appointment
	if err := query.Order("date DESC").Find(&appointments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch appointments",
		})
	}

	return c.JSON(fiber.Map{
		"data": transformAppointments(appointments),
	})
}

// GetAppointmentByID returns an appointment by ID
func (h *AppointmentHandler) GetAppointmentByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid appointment ID",
		})
	}

	var appointment models.Appointment
	if err := database.DB.Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Cabinet").
		First(&appointment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Appointment not found",
		})
	}

	aptData := transformAppointments([]models.Appointment{appointment})[0]
	return c.JSON(aptData)
}

// CreateAppointment creates a new appointment
func (h *AppointmentHandler) CreateAppointment(c *fiber.Ctx) error {
	var appointment models.Appointment
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Set default status if not provided
	if appointment.Status == "" {
		appointment.Status = models.AppointmentStatusScheduled
	}

	// Validate required fields
	if appointment.PatientID == 0 || appointment.DoctorID == 0 || appointment.CabinetID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Patient, doctor, and cabinet are required",
		})
	}

	// Create appointment
	if err := database.DB.Create(&appointment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create appointment",
		})
	}

	// Reload with associations
	database.DB.Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Cabinet").
		First(&appointment, appointment.ID)

	aptData := transformAppointments([]models.Appointment{appointment})[0]
	return c.Status(fiber.StatusCreated).JSON(aptData)
}

// UpdateAppointment updates an appointment
func (h *AppointmentHandler) UpdateAppointment(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid appointment ID",
		})
	}

	var appointment models.Appointment
	if err := database.DB.First(&appointment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Appointment not found",
		})
	}

	// Parse updates
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Update appointment
	if err := database.DB.Model(&appointment).Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update appointment",
		})
	}

	// Reload with associations
	database.DB.Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Cabinet").
		First(&appointment, id)

	aptData := transformAppointments([]models.Appointment{appointment})[0]
	return c.JSON(aptData)
}

// DeleteAppointment deletes an appointment
func (h *AppointmentHandler) DeleteAppointment(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid appointment ID",
		})
	}

	if err := database.DB.Delete(&models.Appointment{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete appointment",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Appointment deleted successfully",
	})
}

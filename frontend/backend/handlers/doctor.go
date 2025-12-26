package handlers

import (
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DoctorHandler handles doctor-related endpoints
type DoctorHandler struct{}

// ListDoctors returns all doctors
func (h *DoctorHandler) ListDoctors(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	var doctors []models.Doctor
	query := database.DB.Preload("User").Preload("Cabinet")

	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		if cabinetID := getCurrentUserCabinetID(currentUser); cabinetID != 0 {
			query = query.Where("cabinet_id = ?", cabinetID)
		}
	}

	if err := query.Find(&doctors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch doctors",
		})
	}

	// Transform to include user data
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

// GetDoctorByID returns a doctor by ID
func (h *DoctorHandler) GetDoctorByID(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid doctor ID",
		})
	}

	var doctor models.Doctor
	if err := database.DB.Preload("User").Preload("Cabinet").First(&doctor, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Doctor not found",
		})
	}

	// Sandbox: non-superadmin can only view doctors in their cabinet
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		cabinetID := getCurrentUserCabinetID(currentUser)
		if cabinetID == 0 || doctor.CabinetID != cabinetID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden (outside your cabinet)",
			})
		}
	}

	userData := getUserWithExtendedInfo(doctor.User)
	return c.JSON(userData)
}

// GetDoctorAppointments returns appointments for a doctor
func (h *DoctorHandler) GetDoctorAppointments(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid doctor ID",
		})
	}

	// Sandbox: non-superadmin can only access doctor resources within their cabinet
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		cabinetID := getCurrentUserCabinetID(currentUser)
		var targetDoctor models.Doctor
		if err := database.DB.First(&targetDoctor, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Doctor not found",
			})
		}
		if cabinetID == 0 || targetDoctor.CabinetID != cabinetID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden (outside your cabinet)",
			})
		}
	}

	var appointments []models.Appointment
	query := database.DB.Where("doctor_id = ?", id).
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

// GetDoctorConsultations returns consultations for a doctor
func (h *DoctorHandler) GetDoctorConsultations(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid doctor ID",
		})
	}

	// Sandbox: non-superadmin can only access doctor resources within their cabinet
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		cabinetID := getCurrentUserCabinetID(currentUser)
		var targetDoctor models.Doctor
		if err := database.DB.First(&targetDoctor, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Doctor not found",
			})
		}
		if cabinetID == 0 || targetDoctor.CabinetID != cabinetID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden (outside your cabinet)",
			})
		}
	}

	var consultations []models.Consultation
	query := database.DB.Where("doctor_id = ?", id).
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

// GetDoctorPatients returns patients for a doctor
func (h *DoctorHandler) GetDoctorPatients(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid doctor ID",
		})
	}

	// Sandbox: non-superadmin can only access doctor resources within their cabinet
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		cabinetID := getCurrentUserCabinetID(currentUser)
		var targetDoctor models.Doctor
		if err := database.DB.First(&targetDoctor, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Doctor not found",
			})
		}
		if cabinetID == 0 || targetDoctor.CabinetID != cabinetID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden (outside your cabinet)",
			})
		}
	}

	// Get unique patients from appointments
	var appointments []models.Appointment
	query := database.DB.Where("doctor_id = ?", id).Preload("Patient.User")

	// Optional cabinet filter: only allowed for SuperAdmin
	if currentUser != nil && currentUser.Type == models.UserTypeSuperAdmin {
		cabinetID := c.Query("cabinet_id")
		if cabinetID != "" {
			query = query.Where("cabinet_id = ?", cabinetID)
		}
	}

	if err := query.Find(&appointments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch patients",
		})
	}

	// Get unique patients
	patientMap := make(map[uint]*models.Patient)
	for _, apt := range appointments {
		if apt.Patient != nil && apt.Patient.User != nil {
			patientMap[apt.Patient.ID] = apt.Patient
		}
	}

	result := make([]interface{}, 0)
	for _, patient := range patientMap {
		userData := getUserWithExtendedInfo(patient.User)
		result = append(result, userData)
	}

	return c.JSON(fiber.Map{
		"data": result,
	})
}

// SearchDoctors searches doctors by filters
func (h *DoctorHandler) SearchDoctors(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	query := database.DB.Model(&models.Doctor{}).
		Preload("User").
		Preload("Cabinet")

	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		if cabinetID := getCurrentUserCabinetID(currentUser); cabinetID != 0 {
			query = query.Where("cabinet_id = ?", cabinetID)
		}
	}

	// Apply filters
	if speciality := c.Query("speciality"); speciality != "" {
		query = query.Where("LOWER(speciality) LIKE ?", "%"+speciality+"%")
	}

	// cabinet_id filter: only allowed for SuperAdmin (others are already scoped)
	if currentUser != nil && currentUser.Type == models.UserTypeSuperAdmin {
		if cabinetID := c.Query("cabinet_id"); cabinetID != "" {
			query = query.Where("cabinet_id = ?", cabinetID)
		}
	}

	if priceMax := c.Query("price_max"); priceMax != "" {
		query = query.Where("consultation_price <= ?", priceMax)
	}

	if priceMin := c.Query("price_min"); priceMin != "" {
		query = query.Where("consultation_price >= ?", priceMin)
	}

	var doctors []models.Doctor
	if err := query.Find(&doctors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to search doctors",
		})
	}

	// Transform results
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

// Helper function to transform appointments with nested data
func transformAppointments(appointments []models.Appointment) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for _, apt := range appointments {
		aptData := map[string]interface{}{
			"id":        apt.ID,
			"date":      apt.Date,
			"status":    apt.Status,
			"createdAt": apt.CreatedAt,
			"updatedAt": apt.UpdatedAt,
		}

		if apt.Patient != nil && apt.Patient.User != nil {
			aptData["patient"] = getUserWithExtendedInfo(apt.Patient.User)
			aptData["patientId"] = apt.PatientID
		}

		if apt.Doctor != nil && apt.Doctor.User != nil {
			aptData["doctor"] = getUserWithExtendedInfo(apt.Doctor.User)
			aptData["doctorId"] = apt.DoctorID
		}

		if apt.Cabinet != nil {
			aptData["cabinet"] = apt.Cabinet
			aptData["cabinetId"] = apt.CabinetID
		}

		result = append(result, aptData)
	}
	return result
}

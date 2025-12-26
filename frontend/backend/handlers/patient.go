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
	currentUser := c.Locals("user").(*models.User)

	var patients []models.Patient
	query := database.DB.Model(&models.Patient{}).Preload("User")

	// Sandbox:
	// - SuperAdmin: can see all
	// - Doctor/Admin/Assistant: only patients with at least one appointment in the user's cabinet
	// - Patient: only themselves
	if currentUser != nil {
		switch currentUser.Type {
		case models.UserTypeSuperAdmin:
			// no filter
		case models.UserTypePatient:
			var patient models.Patient
			if err := database.DB.Where("user_id = ?", currentUser.ID).First(&patient).Error; err == nil {
				query = query.Where("patients.id = ?", patient.ID)
			} else {
				query = query.Where("1 = 0")
			}
		case models.UserTypeDoctor, models.UserTypeAdmin, models.UserTypeAssistant:
			cabinetID := getCurrentUserCabinetID(currentUser)
			if cabinetID != 0 {
				query = query.
					Joins("JOIN appointments ON appointments.patient_id = patients.id").
					Where("appointments.cabinet_id = ?", cabinetID).
					Distinct("patients.id")
			} else {
				query = query.Where("1 = 0")
			}
		default:
			query = query.Where("1 = 0")
		}
	}

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
	currentUser := c.Locals("user").(*models.User)

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

	// Sandbox checks
	if currentUser != nil {
		switch currentUser.Type {
		case models.UserTypeSuperAdmin:
			// allowed
		case models.UserTypePatient:
			// only self
			if patient.UserID != currentUser.ID {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "Forbidden",
				})
			}
		case models.UserTypeDoctor, models.UserTypeAdmin, models.UserTypeAssistant:
			cabinetID := getCurrentUserCabinetID(currentUser)
			if cabinetID == 0 {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "Forbidden (outside your cabinet)",
				})
			}
			// Ensure patient has at least one appointment in this cabinet
			var count int64
			if err := database.DB.Model(&models.Appointment{}).
				Where("patient_id = ? AND cabinet_id = ?", patient.ID, cabinetID).
				Count(&count).Error; err == nil {
				if count == 0 {
					return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"message": "Forbidden (outside your cabinet)",
					})
				}
			} else {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "Forbidden",
				})
			}
		default:
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}
	}

	userData := getUserWithExtendedInfo(patient.User)
	return c.JSON(userData)
}

// GetPatientAppointments returns appointments for a patient
func (h *PatientHandler) GetPatientAppointments(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid patient ID",
		})
	}

	// Sandbox: patient can only see their own appointments; staff can only see appointments in their cabinet
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		if currentUser.Type == models.UserTypePatient {
			var patient models.Patient
			if err := database.DB.Where("user_id = ?", currentUser.ID).First(&patient).Error; err == nil {
				if patient.ID != uint(id) {
					return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"message": "Forbidden",
					})
				}
			} else {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "Forbidden",
				})
			}
		} else {
			cabinetID := getCurrentUserCabinetID(currentUser)
			if cabinetID != 0 {
				c.Locals("_cabinetID", cabinetID)
			}
		}
	}

	var appointments []models.Appointment
	query := database.DB.Where("patient_id = ?", id).
		Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Cabinet").
		Order("date DESC")
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin && currentUser.Type != models.UserTypePatient {
		if cabinetID, ok := c.Locals("_cabinetID").(uint); ok && cabinetID != 0 {
			query = query.Where("cabinet_id = ?", cabinetID)
		}
	}

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
	currentUser := c.Locals("user").(*models.User)

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid patient ID",
		})
	}

	// Sandbox: patient can only see their own consultations; staff can only see consultations in their cabinet
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		if currentUser.Type == models.UserTypePatient {
			var patient models.Patient
			if err := database.DB.Where("user_id = ?", currentUser.ID).First(&patient).Error; err == nil {
				if patient.ID != uint(id) {
					return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
						"message": "Forbidden",
					})
				}
			} else {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "Forbidden",
				})
			}
		} else {
			cabinetID := getCurrentUserCabinetID(currentUser)
			if cabinetID != 0 {
				c.Locals("_cabinetID", cabinetID)
			}
		}
	}

	var consultations []models.Consultation
	query := database.DB.Where("patient_id = ?", id).
		Preload("Patient.User").
		Preload("Doctor.User").
		Preload("Appointment").
		Order("created_at DESC")
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin && currentUser.Type != models.UserTypePatient {
		if cabinetID, ok := c.Locals("_cabinetID").(uint); ok && cabinetID != 0 {
			query = query.Joins("JOIN appointments ON appointments.id = consultations.appointment_id").Where("appointments.cabinet_id = ?", cabinetID)
		}
	}

	if err := query.Find(&consultations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch consultations",
		})
	}

	return c.JSON(fiber.Map{
		"data": consultations,
	})
}

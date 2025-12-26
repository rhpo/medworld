package handlers

import (
	"encoding/json"
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type assignAssistantDoctorRequest struct {
	DoctorID uint `json:"doctorId"`
}

// CabinetHandler handles cabinet-related endpoints
type CabinetHandler struct{}

// ListCabinets returns all cabinets
func (h *CabinetHandler) ListCabinets(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	var cabinets []models.Cabinet
	query := database.DB.Preload("Admin")

	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		cabinetID := getCurrentUserCabinetID(currentUser)
		if cabinetID != 0 {
			query = query.Where("id = ?", cabinetID)
		} else {
			query = query.Where("1 = 0")
		}
	}

	if err := query.Find(&cabinets).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch cabinets",
		})
	}

	return c.JSON(fiber.Map{
		"data": cabinets,
	})
}

// GetCabinetByID returns a cabinet by ID
func (h *CabinetHandler) GetCabinetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	var cabinet models.Cabinet
	if err := database.DB.Preload("Admin").First(&cabinet, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cabinet not found",
		})
	}

	return c.JSON(cabinet)
}

// GetCabinetDoctors returns doctors for a cabinet
func (h *CabinetHandler) GetCabinetDoctors(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	var doctors []models.Doctor
	if err := database.DB.Where("cabinet_id = ?", id).
		Preload("User").
		Find(&doctors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch doctors",
		})
	}

	// Transform results
	result := make([]map[string]any, 0)
	for _, doctor := range doctors {
		if doctor.User != nil {
			userData := getUserWithExtendedInfo(doctor.User)
			result = append(result, userData.(map[string]any))
		}
	}

	return c.JSON(fiber.Map{
		"data": result,
	})
}

// GetCabinetAppointments returns appointments for a cabinet
func (h *CabinetHandler) GetCabinetAppointments(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	var appointments []models.Appointment
	if err := database.DB.Where("cabinet_id = ?", id).
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

// GetCabinetAssistants returns assistants for a cabinet
func (h *CabinetHandler) GetCabinetAssistants(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	var assistants []models.Assistant
	if err := database.DB.Where("cabinet_id = ?", id).
		Preload("User").
		Find(&assistants).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch assistants",
		})
	}

	// Transform results
	result := make([]map[string]any, 0)
	for _, assistant := range assistants {
		if assistant.User != nil {
			userData := getUserWithExtendedInfo(assistant.User)
			result = append(result, userData.(map[string]any))
		}
	}

	return c.JSON(fiber.Map{
		"data": result,
	})
}

// CreateCabinet creates a new cabinet
func (h *CabinetHandler) CreateCabinet(c *fiber.Ctx) error {
	var req struct {
		Name             string `json:"name"`
		Phone            string `json:"phone"`
		Image            string `json:"image"`
		Location         any    `json:"location"`
		OpeningHours     any    `json:"openingHours"`
		AccessHandicap   *bool  `json:"accessHandicap"`
		HasParking       *bool  `json:"hasParking"`
		HasWifi          *bool  `json:"hasWifi"`
		AcceptsUrgent    *bool  `json:"acceptsUrgent"`
		AcceptsInsurance *bool  `json:"acceptsInsurance"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	locationJSON, _ := json.Marshal(req.Location)
	openingHoursJSON, _ := json.Marshal(req.OpeningHours)

	cabinet := models.Cabinet{
		Name:             req.Name,
		Phone:            req.Phone,
		Image:            req.Image,
		Location:         string(locationJSON),
		OpeningHours:     string(openingHoursJSON),
		AccessHandicap:   req.AccessHandicap,
		HasParking:       req.HasParking,
		HasWifi:          req.HasWifi,
		AcceptsUrgent:    req.AcceptsUrgent,
		AcceptsInsurance: req.AcceptsInsurance,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := database.DB.Create(&cabinet).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create cabinet",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(cabinet)
}

// UpdateCabinet updates an existing cabinet
func (h *CabinetHandler) UpdateCabinet(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	var cabinet models.Cabinet
	if err := database.DB.First(&cabinet, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cabinet not found",
		})
	}

	var req struct {
		Name             string `json:"name"`
		Phone            string `json:"phone"`
		Image            string `json:"image"`
		Location         any    `json:"location"`
		OpeningHours     any    `json:"openingHours"`
		AccessHandicap   *bool  `json:"accessHandicap"`
		HasParking       *bool  `json:"hasParking"`
		HasWifi          *bool  `json:"hasWifi"`
		AcceptsUrgent    *bool  `json:"acceptsUrgent"`
		AcceptsInsurance *bool  `json:"acceptsInsurance"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	locationJSON, _ := json.Marshal(req.Location)
	openingHoursJSON, _ := json.Marshal(req.OpeningHours)

	updates := models.Cabinet{
		Name:             req.Name,
		Phone:            req.Phone,
		Image:            req.Image,
		Location:         string(locationJSON),
		OpeningHours:     string(openingHoursJSON),
		AccessHandicap:   req.AccessHandicap,
		HasParking:       req.HasParking,
		HasWifi:          req.HasWifi,
		AcceptsUrgent:    req.AcceptsUrgent,
		AcceptsInsurance: req.AcceptsInsurance,
		UpdatedAt:        time.Now(),
	}

	if err := database.DB.Model(&cabinet).Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update cabinet",
		})
	}

	return c.JSON(cabinet)
}

// DeleteCabinet deletes a cabinet
func (h *CabinetHandler) DeleteCabinet(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	if err := database.DB.Delete(&models.Cabinet{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete cabinet",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cabinet deleted successfully",
	})
}

// AddDoctorToCabinet adds an existing doctor to the cabinet
func (h *CabinetHandler) AddDoctorToCabinet(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	var req struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Find user by email to get doctor profile
	var user models.User
	if err := database.DB.Where("email = ? AND type = ?", req.Email, models.UserTypeDoctor).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Doctor not found with this email",
		})
	}

	// Update doctor's cabinet_id
	if err := database.DB.Model(&models.Doctor{}).Where("user_id = ?", user.ID).Update("cabinet_id", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add doctor to cabinet",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Doctor added successfully",
	})
}

// RemoveDoctorFromCabinet removes a doctor from the cabinet
func (h *CabinetHandler) RemoveDoctorFromCabinet(c *fiber.Ctx) error {
	// cabinetId is in Params("id"), but we primarily need doctorId to nullify their cabinet
	doctorId, err := strconv.ParseUint(c.Params("doctorId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid doctor ID",
		})
	}

	// Set cabinet_id to null (0 in this case as we might not have nullable setup, or 0 indicates no cabinet)
	// Assuming 0 is "no cabinet" or checking model definition.
	// Update: GORM defaults uint to 0. If foreign key is nullable, we should use map or pointer.
	// Looking at models/doctor.go: CabinetID uint `gorm:"index" json:"cabinetId"` - it is NOT a pointer, so 0.

	if err := database.DB.Model(&models.Doctor{}).Where("id = ?", doctorId).Update("cabinet_id", 0).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to remove doctor from cabinet",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Doctor removed successfully",
	})
}

// AddAssistantToCabinet adds an existing assistant to the cabinet
func (h *CabinetHandler) AddAssistantToCabinet(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	var req struct {
		Email    string `json:"email"`
		DoctorID uint   `json:"doctorId"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	var user models.User
	if err := database.DB.Where("email = ? AND type = ?", req.Email, models.UserTypeAssistant).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Assistant not found with this email",
		})
	}

	// Resolve which doctor to assign (assistant is assigned to exactly one doctor)
	assignedDoctorID := req.DoctorID
	if assignedDoctorID != 0 {
		// Validate doctor belongs to this cabinet
		var doctor models.Doctor
		if err := database.DB.Where("id = ? AND cabinet_id = ?", assignedDoctorID, id).First(&doctor).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid doctorId for this cabinet",
			})
		}

		// Enforce: a doctor can have at most one assistant
		var existing models.Assistant
		if err := database.DB.Where("doctor_id = ?", assignedDoctorID).First(&existing).Error; err == nil {
			// Allow idempotent re-link if it's the same assistant
			if existing.UserID != user.ID {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"message": "This doctor already has an assigned assistant",
				})
			}
		}
	} else {
		// Fallback to first doctor in the cabinet
		var doctor models.Doctor
		if err := database.DB.Where("cabinet_id = ?", id).First(&doctor).Error; err == nil {
			assignedDoctorID = doctor.ID
		}
	}

	// Assign assistant to cabinet + doctor
	updates := map[string]interface{}{
		"cabinet_id": id,
	}
	if assignedDoctorID != 0 {
		updates["doctor_id"] = assignedDoctorID
	}
	if err := database.DB.Model(&models.Assistant{}).Where("user_id = ?", user.ID).Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add assistant to cabinet",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Assistant added successfully",
	})
}

// AssignAssistantToDoctor assigns/unassigns an assistant to a doctor within the same cabinet.
// doctorId=0 unassigns from any doctor but keeps the assistant linked to the cabinet.
func (h *CabinetHandler) AssignAssistantToDoctor(c *fiber.Ctx) error {
	cabinetID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	assistantID, err := strconv.ParseUint(c.Params("assistantId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assistant ID",
		})
	}

	var req assignAssistantDoctorRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	var assistant models.Assistant
	if err := database.DB.Where("id = ? AND cabinet_id = ?", assistantID, cabinetID).First(&assistant).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Assistant not found in this cabinet",
		})
	}

	if req.DoctorID != 0 {
		// Validate doctor belongs to cabinet
		var doctor models.Doctor
		if err := database.DB.Where("id = ? AND cabinet_id = ?", req.DoctorID, cabinetID).First(&doctor).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid doctorId for this cabinet",
			})
		}

		// Enforce one assistant per doctor
		var existing models.Assistant
		if err := database.DB.Where("doctor_id = ?", req.DoctorID).First(&existing).Error; err == nil {
			if existing.ID != assistant.ID {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"message": "This doctor already has an assigned assistant",
				})
			}
		}
	}

	if err := database.DB.Model(&models.Assistant{}).Where("id = ?", assistant.ID).Update("doctor_id", req.DoctorID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update assistant assignment",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Assistant assignment updated",
	})
}

// CreateAssistantInCabinet creates a new assistant user and links them to the cabinet + assigned doctor.
func (h *CabinetHandler) CreateAssistantInCabinet(c *fiber.Ctx) error {
	cabinetID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	var req struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		PhoneNumber string `json:"phoneNumber"`
		Address     string `json:"address"`
		Gender      string `json:"gender"`
		DoctorID    uint   `json:"doctorId"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if req.Email == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "firstName, lastName, email and password are required",
		})
	}

	// Validate doctor assignment
	if req.DoctorID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "doctorId is required",
		})
	}
	var doctor models.Doctor
	if err := database.DB.Where("id = ? AND cabinet_id = ?", req.DoctorID, cabinetID).First(&doctor).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid doctorId for this cabinet",
		})
	}

	// Enforce one assistant per doctor
	var existing models.Assistant
	if err := database.DB.Where("doctor_id = ?", req.DoctorID).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "This doctor already has an assigned assistant",
		})
	}

	// Ensure email not taken
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email already registered",
		})
	}

	user := models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Gender:      req.Gender,
		Type:        models.UserTypeAssistant,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	assistant := models.Assistant{
		UserID:    user.ID,
		CabinetID: uint(cabinetID),
		DoctorID:  req.DoctorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := database.DB.Create(&assistant).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create assistant",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Assistant created successfully",
		"data":    getUserWithExtendedInfo(&user),
	})
}

// RemoveAssistantFromCabinet removes an assistant from the cabinet
func (h *CabinetHandler) RemoveAssistantFromCabinet(c *fiber.Ctx) error {
	assistantId, err := strconv.ParseUint(c.Params("assistantId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid assistant ID",
		})
	}

	if err := database.DB.Model(&models.Assistant{}).Where("id = ?", assistantId).Updates(map[string]interface{}{
		"cabinet_id": 0,
		"doctor_id":  0,
	}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to remove assistant from cabinet",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Assistant removed successfully",
	})
}

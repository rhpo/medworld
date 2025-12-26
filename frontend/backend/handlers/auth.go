package handlers

import (
	"fmt"
	"medworld-backend/database"
	"medworld-backend/models"
	"medworld-backend/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthHandler handles authentication endpoints
type AuthHandler struct{}

func getUserWithExtendedInfo(user *models.User) interface{} {
	return getUserWithExtendedInfoDepth(user, 2)
}

func getUserWithExtendedInfoDepth(user *models.User, depth int) interface{} {
	if user == nil {
		return nil
	}

	baseUser := map[string]interface{}{
		"id":          user.ID,
		"firstName":   user.FirstName,
		"lastName":    user.LastName,
		"fullName":    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		"email":       user.Email,
		"phoneNumber": user.PhoneNumber,
		"avatarUrl":   user.AvatarURL,
		"address":     user.Address,
		"gender":      user.Gender,
		"dateOfBirth": user.DateOfBirth,
		"type":        user.Type,
		"createdAt":   user.CreatedAt,
	}

	// Add messages (both sent and received)
	var messages []models.Message
	if err := database.DB.Where("receiver_id = ? OR sender_id = ?", user.ID, user.ID).
		Preload("Sender").
		Preload("Receiver").
		Order("created_at DESC").
		Find(&messages).Error; err == nil {
		baseUser["messages"] = messages
	} else {
		baseUser["messages"] = []models.Message{}
	}

	if depth <= 0 {
		return baseUser
	}

	// Add type-specific data
	switch user.Type {
	case models.UserTypeDoctor, models.UserTypeAdmin:
		var doctor models.Doctor
		if err := database.DB.Where("user_id = ?", user.ID).Preload("Cabinet").Preload("Calendars").First(&doctor).Error; err == nil {
			baseUser["doctorId"] = doctor.ID
			baseUser["speciality"] = doctor.Speciality
			baseUser["licenseNumber"] = doctor.LicenseNumber
			baseUser["careerStart"] = doctor.CareerStart
			baseUser["consultationPrice"] = doctor.ConsultationPrice
			baseUser["consultationDuration"] = doctor.ConsultationDuration
			baseUser["cabinet"] = doctor.Cabinet
			baseUser["cabinetId"] = doctor.CabinetID
			baseUser["calendars"] = doctor.Calendars

			// Include consultations
			var consultations []models.Consultation
			if err := database.DB.
				Where("doctor_id = ?", doctor.ID).
				Preload("Doctor.User").
				Preload("Patient.User").
				Preload("Appointment").
				Order("created_at DESC").
				Find(&consultations).Error; err == nil {
				baseUser["consultations"] = consultations
			} else {
				baseUser["consultations"] = []models.Consultation{}
			}

			// Single assistant (doctor has exactly one assistant)
			var assistant models.Assistant
			if err := database.DB.Where("doctor_id = ?", doctor.ID).
				Preload("User").
				Preload("Cabinet").
				First(&assistant).Error; err == nil {
				baseUser["assistantId"] = assistant.ID
				if assistant.User != nil {
					baseUser["assistant"] = getUserWithExtendedInfoDepth(assistant.User, depth-1)
				} else {
					baseUser["assistant"] = nil
				}
			} else {
				baseUser["assistantId"] = nil
				baseUser["assistant"] = nil
			}
		}

	case models.UserTypePatient:
		var patient models.Patient
		if err := database.DB.Where("user_id = ?", user.ID).First(&patient).Error; err == nil {
			baseUser["patientId"] = patient.ID
			baseUser["emergencyContact"] = patient.EmergencyContact
			baseUser["bloodType"] = patient.BloodType
			baseUser["weight"] = patient.Weight
			baseUser["medicalHistory"] = patient.MedicalHistory
			baseUser["allergies"] = patient.Allergies
		}

	case models.UserTypeAssistant:
		var assistant models.Assistant
		if err := database.DB.Where("user_id = ?", user.ID).
			Preload("Cabinet").
			Preload("Doctor").
			Preload("Doctor.User").
			Preload("Doctor.Cabinet").
			Preload("Doctor.Calendars").
			First(&assistant).Error; err == nil {
			// Backfill older data
			if assistant.DoctorID == 0 && assistant.CabinetID != 0 {
				var doctor models.Doctor
				if err := database.DB.Where("cabinet_id = ?", assistant.CabinetID).First(&doctor).Error; err == nil {
					if err := database.DB.Model(&models.Assistant{}).Where("id = ?", assistant.ID).Update("doctor_id", doctor.ID).Error; err == nil {
						assistant.DoctorID = doctor.ID
						database.DB.Where("id = ?", assistant.ID).
							Preload("Cabinet").
							Preload("Doctor").
							Preload("Doctor.User").
							Preload("Doctor.Cabinet").
							Preload("Doctor.Calendars").
							First(&assistant)
					}
				}
			}

			baseUser["assistantId"] = assistant.ID
			baseUser["cabinet"] = assistant.Cabinet
			baseUser["cabinetId"] = assistant.CabinetID
			baseUser["doctorId"] = assistant.DoctorID
			if assistant.Doctor != nil && assistant.Doctor.User != nil {
				baseUser["doctor"] = getUserWithExtendedInfoDepth(assistant.Doctor.User, depth-1)
			} else {
				baseUser["doctor"] = nil
			}
		}
	}

	return baseUser
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegisterRequest represents registration data
type RegisterRequest struct {
	FirstName   string          `json:"firstName" validate:"required"`
	LastName    string          `json:"lastName" validate:"required"`
	Email       string          `json:"email" validate:"required,email"`
	Password    string          `json:"password" validate:"required,min=6"`
	PhoneNumber string          `json:"phoneNumber"`
	Address     string          `json:"address"`
	Gender      string          `json:"gender"`
	DateOfBirth string          `json:"dateOfBirth"`
	Type        models.UserType `json:"type" validate:"required"`
}

// Login handles user login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	// Find user by email
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, string(user.Type))
	if err != nil {
		fmt.Println("Error of token: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	// Get user data with additional info
	userData := getUserWithExtendedInfo(&user)

	return c.JSON(fiber.Map{
		"user":  userData,
		"token": token,
	})
}

// Register handles user registration
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	// Check if email already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Email already registered",
		})
	}

	// Create user
	user := models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password, // Will be hashed by BeforeSave hook
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Gender:      req.Gender,
		Type:        req.Type,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, string(user.Type))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	// Get user data with additional info
	userData := getUserWithExtendedInfo(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":  userData,
		"token": token,
	})
}

// Me returns current authenticated user
func (h *AuthHandler) Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	userData := getUserWithExtendedInfo(user)

	return c.JSON(fiber.Map{
		"user": userData,
	})
}

// Logout handles user logout (placeholder - JWT is stateless)
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// With JWT, logout is typically handled client-side by removing the token
	// For sessions or token blacklisting, you would implement it here
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

package handlers

import (
	"encoding/json"
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// UserHandler handles user-related endpoints
type UserHandler struct{}

// GetUserByID returns a user by ID
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Get extended user info
	userData := getUserWithExtendedInfo(&user)

	return c.JSON(userData)
}

// UpdateUser updates a user profile
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	currentUser := c.Locals("user").(*models.User)

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Check permission
	isSelf := currentUser.ID == user.ID
	isAdmin := currentUser.Type == models.UserTypeSuperAdmin || currentUser.Type == models.UserTypeAdmin
	isMedicUpdatingPatient := (currentUser.Type == models.UserTypeDoctor || currentUser.Type == models.UserTypeAssistant) && user.Type == models.UserTypePatient

	if !isSelf && !isAdmin && !isMedicUpdatingPatient {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You can only update your own profile",
		})
	}

	// Parse update data
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Remove sensitive fields that shouldn't be updated directly
	delete(updates, "id")
	delete(updates, "password")
	delete(updates, "type")
	delete(updates, "email") // Email changes require verification

	// Identify specialized fields to avoid SQL errors on users table
	// We use snake_case because the frontend transforms camelCase to snake_case
	specializedFields := map[string]bool{
		"speciality": true, "license_number": true, "consultation_price": true,
		"consultation_duration": true, "years_of_experience": true,
		"blood_type": true, "emergency_contact": true, "weight": true,
		"medical_history": true, "allergies": true,
	}

	userUpdates := make(map[string]interface{})
	for k, v := range updates {
		if !specializedFields[k] {
			userUpdates[k] = v
		}
	}

	// Update user base info
	if len(userUpdates) > 0 {
		if err := database.DB.Model(&user).Updates(userUpdates).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to update user base info",
			})
		}
	}

	// Update specialized tables based on type
	switch user.Type {
	case models.UserTypeDoctor, models.UserTypeAdmin:
		docUpdates := make(map[string]interface{})
		if val, ok := updates["speciality"]; ok {
			docUpdates["speciality"] = val
		}
		if val, ok := updates["license_number"]; ok {
			docUpdates["license_number"] = val
		}
		if val, ok := updates["consultation_price"]; ok {
			docUpdates["consultation_price"] = val
		}
		if val, ok := updates["consultation_duration"]; ok {
			docUpdates["consultation_duration"] = val
		}
		if val, ok := updates["years_of_experience"]; ok {
			years := 0
			switch v := val.(type) {
			case float64:
				years = int(v)
			case int:
				years = v
			}
			if years > 0 {
				careerStart := time.Now().AddDate(-years, 0, 0)
				docUpdates["career_start"] = &careerStart
			}
		}
		if len(docUpdates) > 0 {
			database.DB.Model(&models.Doctor{}).Where("user_id = ?", user.ID).Updates(docUpdates)
		}

	case models.UserTypePatient:
		patientUpdates := make(map[string]interface{})
		if val, ok := updates["blood_type"]; ok {
			patientUpdates["blood_type"] = val
		}
		if val, ok := updates["emergency_contact"]; ok {
			patientUpdates["emergency_contact"] = val
		}
		if val, ok := updates["weight"]; ok {
			patientUpdates["weight"] = val
		}
		if val, ok := updates["medical_history"]; ok {
			// If it's already a string (from frontend stringify), keep it as is
			// If it's an object/array, marshal it
			if str, isStr := val.(string); isStr {
				patientUpdates["medical_history"] = str
			} else {
				jsonVal, _ := json.Marshal(val)
				patientUpdates["medical_history"] = string(jsonVal)
			}
		}
		if val, ok := updates["allergies"]; ok {
			if str, isStr := val.(string); isStr {
				patientUpdates["allergies"] = str
			} else {
				jsonVal, _ := json.Marshal(val)
				patientUpdates["allergies"] = string(jsonVal)
			}
		}
		if len(patientUpdates) > 0 {
			database.DB.Model(&models.Patient{}).Where("user_id = ?", user.ID).Updates(patientUpdates)
		}
	}

	// Reload user
	database.DB.First(&user, id)
	userData := getUserWithExtendedInfo(&user)

	return c.JSON(userData)
}

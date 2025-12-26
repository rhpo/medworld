package middleware

import (
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// RequireCabinetAccess ensures the authenticated user can only access resources within their own cabinet.
// SuperAdmin is allowed to access any cabinet.
// Admin/Doctor are scoped by their Doctor.CabinetID.
// Assistant is scoped by their Assistant.CabinetID.
func RequireCabinetAccess(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	cabinetIDParam, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid cabinet ID",
		})
	}

	// SuperAdmin can access any cabinet.
	if currentUser.Type == models.UserTypeSuperAdmin {
		return c.Next()
	}

	switch currentUser.Type {
	case models.UserTypeAdmin, models.UserTypeDoctor:
		var doctor models.Doctor
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&doctor).Error; err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Doctor profile not found",
			})
		}
		if doctor.CabinetID == 0 || uint(cabinetIDParam) != doctor.CabinetID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden (outside your cabinet)",
			})
		}

	case models.UserTypeAssistant:
		var assistant models.Assistant
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&assistant).Error; err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Assistant profile not found",
			})
		}
		if assistant.CabinetID == 0 || uint(cabinetIDParam) != assistant.CabinetID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden (outside your cabinet)",
			})
		}

	default:
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	return c.Next()
}

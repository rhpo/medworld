package handlers

import (
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"

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

	// Check permission - users can only update their own profile unless admin/superadmin
	if currentUser.ID != uint(id) && currentUser.Type != models.UserTypeSuperAdmin && currentUser.Type != models.UserTypeAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You can only update your own profile",
		})
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
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

	// Update user
	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	// Reload user
	database.DB.First(&user, id)
	userData := getUserWithExtendedInfo(&user)

	return c.JSON(userData)
}

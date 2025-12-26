package handlers

import (
	"encoding/json"
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CalendarHandler struct{}

// GetCalendarByID retrieves a calendar by ID with doctor and cabinet information
func (h *CalendarHandler) GetCalendarByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid calendar ID",
		})
	}

	var calendar models.Calendar
	if err := database.DB.
		Preload("Doctor.User").
		Preload("Cabinet").
		First(&calendar, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Calendar not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch calendar",
		})
	}

	// Parse availability JSON
	var availability []fiber.Map
	if calendar.Availability != "" {
		if err := json.Unmarshal([]byte(calendar.Availability), &availability); err != nil {
			availability = []fiber.Map{}
		}
	}

	return c.JSON(fiber.Map{
		"id":           calendar.ID,
		"doctorId":     calendar.DoctorID,
		"cabinetId":    calendar.CabinetID,
		"doctor":       calendar.Doctor,
		"cabinet":      calendar.Cabinet,
		"availability": availability,
		"createdAt":    calendar.CreatedAt,
		"updatedAt":    calendar.UpdatedAt,
	})
}

// UpdateCalendar updates the calendar's availability
func (h *CalendarHandler) UpdateCalendar(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid calendar ID",
		})
	}

	var req struct {
		Availability []fiber.Map `json:"availability"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Convert availability to JSON string
	availabilityJSON, err := json.Marshal(req.Availability)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid availability format",
		})
	}

	var calendar models.Calendar
	if err := database.DB.First(&calendar, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Calendar not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch calendar",
		})
	}

	// Update calendar
	calendar.Availability = string(availabilityJSON)
	if err := database.DB.Save(&calendar).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update calendar",
		})
	}

	// Reload with relationships
	if err := database.DB.
		Preload("Doctor.User").
		Preload("Cabinet").
		First(&calendar, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch updated calendar",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":           calendar.ID,
		"doctorId":     calendar.DoctorID,
		"cabinetId":    calendar.CabinetID,
		"doctor":       calendar.Doctor,
		"cabinet":      calendar.Cabinet,
		"availability": req.Availability,
		"createdAt":    calendar.CreatedAt,
		"updatedAt":    calendar.UpdatedAt,
	})
}

// ListCalendars retrieves all calendars for a doctor or cabinet
func (h *CalendarHandler) ListCalendars(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)

	doctorID := c.Query("doctorId")
	cabinetID := c.Query("cabinetId")

	var calendars []models.Calendar
	query := database.DB.Preload("Doctor.User").Preload("Cabinet")

	// Sandbox: non-superadmin can only access calendars in their cabinet.
	if currentUser != nil && currentUser.Type != models.UserTypeSuperAdmin {
		cabinetID = ""
		if cID := getCurrentUserCabinetID(currentUser); cID != 0 {
			query = query.Where("cabinet_id = ?", cID)
		} else {
			query = query.Where("1 = 0")
		}
	}

	if doctorID != "" {
		query = query.Where("doctor_id = ?", doctorID)
	}

	if cabinetID != "" {
		query = query.Where("cabinet_id = ?", cabinetID)
	}

	if err := query.Find(&calendars).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch calendars",
		})
	}

	result := make([]fiber.Map, len(calendars))
	for i, cal := range calendars {
		var availability []fiber.Map
		if cal.Availability != "" {
			if err := json.Unmarshal([]byte(cal.Availability), &availability); err != nil {
				availability = []fiber.Map{}
			}
		}

		result[i] = fiber.Map{
			"id":           cal.ID,
			"doctorId":     cal.DoctorID,
			"cabinetId":    cal.CabinetID,
			"doctor":       cal.Doctor,
			"cabinet":      cal.Cabinet,
			"availability": availability,
			"createdAt":    cal.CreatedAt,
			"updatedAt":    cal.UpdatedAt,
		}
	}

	return c.JSON(result)
}

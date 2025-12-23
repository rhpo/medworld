package handlers

import (
	"encoding/json"
	"medworld-backend/database"
	"medworld-backend/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CabinetHandler handles cabinet-related endpoints
type CabinetHandler struct{}

// ListCabinets returns all cabinets
func (h *CabinetHandler) ListCabinets(c *fiber.Ctx) error {
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

// CreateCabinet creates a new cabinet
func (h *CabinetHandler) CreateCabinet(c *fiber.Ctx) error {
	var req struct {
		Name             string      `json:"name"`
		Phone            string      `json:"phone"`
		Location         interface{} `json:"location"`
		OpeningHours     interface{} `json:"openingHours"`
		AccessHandicap   bool        `json:"accessHandicap"`
		HasParking       *bool       `json:"hasParking"`
		HasWifi          *bool       `json:"hasWifi"`
		AcceptsUrgent    *bool       `json:"acceptsUrgent"`
		AcceptsInsurance *bool       `json:"acceptsInsurance"`
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

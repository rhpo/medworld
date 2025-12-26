package handlers

import (
	"strings"
	"time"

	"medworld-backend/database"
	"medworld-backend/models"

	"github.com/gofiber/fiber/v2"
)

// MessageHandler handles message-related endpoints
type MessageHandler struct{}

type createMessageRequest struct {
	SenderID   uint   `json:"senderId"`
	ReceiverID uint   `json:"receiverId"`
	Content    string `json:"content"`
}

// ListMyMessages returns all messages where the current user is sender or receiver
func (h *MessageHandler) ListMyMessages(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var messages []models.Message
	if err := database.DB.
		Preload("Sender").
		Preload("Receiver").
		Where("sender_id = ? OR receiver_id = ?", currentUser.ID, currentUser.ID).
		Order("created_at DESC").
		Find(&messages).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch messages",
		})
	}

	return c.JSON(fiber.Map{
		"data": messages,
	})
}

// CreateMessage creates a new message from the authenticated user to a receiver
func (h *MessageHandler) CreateMessage(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var req createMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	content := strings.TrimSpace(req.Content)
	if req.ReceiverID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "receiverId is required",
		})
	}
	if content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "content is required",
		})
	}

	// ensure receiver exists
	var receiver models.User
	if err := database.DB.First(&receiver, req.ReceiverID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Receiver not found",
		})
	}

	message := models.Message{
		SenderID:   currentUser.ID,
		ReceiverID: req.ReceiverID,
		Content:    content,
		IsRead:     false,
		CreatedAt:  time.Now(),
	}

	if err := database.DB.Create(&message).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to send message",
		})
	}

	// Reload with associations for UI convenience
	database.DB.Preload("Sender").Preload("Receiver").First(&message, message.ID)

	return c.Status(fiber.StatusCreated).JSON(message)
}

func (h *MessageHandler) ListRecipients(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*models.User)
	if currentUser == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if currentUser.Type != models.UserTypeDoctor && currentUser.Type != models.UserTypeAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Only doctors can list recipients",
		})
	}

	var doctor models.Doctor
	if err := database.DB.Where("user_id = ?", currentUser.ID).First(&doctor).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Doctor profile not found",
		})
	}

	staff := make([]interface{}, 0)
	patients := make([]interface{}, 0)

	if doctor.CabinetID != 0 {
		var doctors []models.Doctor
		if err := database.DB.Where("cabinet_id = ?", doctor.CabinetID).
			Preload("User").
			Find(&doctors).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to fetch staff",
			})
		}

		for _, d := range doctors {
			if d.User != nil && d.User.ID != currentUser.ID {
				staff = append(staff, getUserWithExtendedInfo(d.User))
			}
		}

		var assistants []models.Assistant
		if err := database.DB.Where("cabinet_id = ?", doctor.CabinetID).
			Preload("User").
			Find(&assistants).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to fetch staff",
			})
		}

		for _, a := range assistants {
			if a.User != nil && a.User.ID != currentUser.ID {
				staff = append(staff, getUserWithExtendedInfo(a.User))
			}
		}
	}

	var appointments []models.Appointment
	if err := database.DB.Where("doctor_id = ?", doctor.ID).
		Preload("Patient.User").
		Find(&appointments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch patients",
		})
	}

	patientUserIDs := make(map[uint]struct{})
	for _, apt := range appointments {
		if apt.Patient != nil && apt.Patient.User != nil {
			if apt.Patient.User.ID == currentUser.ID {
				continue
			}
			patientUserIDs[apt.Patient.User.ID] = struct{}{}
		}
	}
	for userID := range patientUserIDs {
		var u models.User
		if err := database.DB.First(&u, userID).Error; err == nil {
			patients = append(patients, getUserWithExtendedInfo(&u))
		}
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"staff":    staff,
			"patients": patients,
		},
	})
}

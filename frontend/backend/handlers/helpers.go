package handlers

import (
	"medworld-backend/database"
	"medworld-backend/models"
)

func getCurrentUserCabinetID(currentUser *models.User) uint {
	if currentUser == nil {
		return 0
	}

	switch currentUser.Type {
	case models.UserTypeAdmin, models.UserTypeDoctor:
		var doctor models.Doctor
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&doctor).Error; err == nil {
			return doctor.CabinetID
		}
	case models.UserTypeAssistant:
		var assistant models.Assistant
		if err := database.DB.Where("user_id = ?", currentUser.ID).First(&assistant).Error; err == nil {
			return assistant.CabinetID
		}
	}

	return 0
}

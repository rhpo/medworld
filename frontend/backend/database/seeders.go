package database

import (
	"encoding/json"
	"log"
	"medworld-backend/models"
	"time"
)

// SeedDatabase populates the database with initial data
func SeedDatabase() error {
	log.Println("Seeding database...")

	// Check if data already exists
	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)
	if userCount > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	// Create SuperAdmin
	dob := time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	superAdmin := models.User{
		FirstName:   "Super",
		LastName:    "Admin",
		Email:       "admin@medworld.com",
		Password:    "admin123",
		PhoneNumber: "+1234567890",
		Gender:      "male",
		DateOfBirth: &dob,
		Type:        models.UserTypeSuperAdmin,
		Address:     "Admin Office, Medical Plaza",
	}
	if err := DB.Create(&superAdmin).Error; err != nil {
		return err
	}

	// Create Cabinet
	location, _ := json.Marshal(map[string]interface{}{
		"address":   "123 Medical Street, Healthcare City",
		"latitude":  40.7128,
		"longitude": -74.0060,
	})

	openingHours, _ := json.Marshal(map[string]interface{}{
		"monday":    map[string]string{"open": "09:00", "close": "17:00"},
		"tuesday":   map[string]string{"open": "09:00", "close": "17:00"},
		"wednesday": map[string]string{"open": "09:00", "close": "17:00"},
		"thursday":  map[string]string{"open": "09:00", "close": "17:00"},
		"friday":    map[string]string{"open": "09:00", "close": "17:00"},
		"saturday":  map[string]string{"open": "10:00", "close": "14:00"},
		"sunday":    map[string]string{"open": "closed", "close": "closed"},
	})

	trueVal := true
	cabinet := models.Cabinet{
		Name:             "Central Medical Clinic",
		Phone:            "+1234567891",
		AdminID:          superAdmin.ID,
		Location:         string(location),
		OpeningHours:     string(openingHours),
		AccessHandicap:   true,
		HasParking:       &trueVal,
		HasWifi:          &trueVal,
		AcceptsUrgent:    &trueVal,
		AcceptsInsurance: &trueVal,
	}
	if err := DB.Create(&cabinet).Error; err != nil {
		return err
	}

	// Create Doctor User
	doctorDob := time.Date(1985, 5, 15, 0, 0, 0, 0, time.UTC)
	doctorUser := models.User{
		FirstName:   "John",
		LastName:    "Smith",
		Email:       "doctor@medworld.com",
		Password:    "doctor123",
		PhoneNumber: "+1234567892",
		Gender:      "male",
		DateOfBirth: &doctorDob,
		Type:        models.UserTypeDoctor,
		Address:     "456 Doctor Ave",
	}
	if err := DB.Create(&doctorUser).Error; err != nil {
		return err
	}

	// Create Doctor Profile
	careerStart := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)
	doctor := models.Doctor{
		UserID:               doctorUser.ID,
		Speciality:           "Cardiology",
		CareerStart:          &careerStart,
		ConsultationPrice:    150.00,
		ConsultationDuration: 30,
		CabinetID:            cabinet.ID,
	}
	if err := DB.Create(&doctor).Error; err != nil {
		return err
	}

	// Create Admin Doctor User
	adminDoctorDob := time.Date(1982, 3, 20, 0, 0, 0, 0, time.UTC)
	adminDoctorUser := models.User{
		FirstName:   "Sarah",
		LastName:    "Johnson",
		Email:       "admin.doctor@medworld.com",
		Password:    "admin123",
		PhoneNumber: "+1234567893",
		Gender:      "female",
		DateOfBirth: &adminDoctorDob,
		Type:        models.UserTypeAdmin,
		Address:     "789 Admin Blvd",
	}
	if err := DB.Create(&adminDoctorUser).Error; err != nil {
		return err
	}

	// Create Admin Doctor Profile
	adminCareerStart := time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC)
	adminDoctor := models.Doctor{
		UserID:               adminDoctorUser.ID,
		Speciality:           "General Practice",
		CareerStart:          &adminCareerStart,
		ConsultationPrice:    120.00,
		ConsultationDuration: 30,
		CabinetID:            cabinet.ID,
	}
	if err := DB.Create(&adminDoctor).Error; err != nil {
		return err
	}

	// Create Patient User
	patientDob := time.Date(1990, 8, 10, 0, 0, 0, 0, time.UTC)
	patientUser := models.User{
		FirstName:   "Alice",
		LastName:    "Williams",
		Email:       "patient@medworld.com",
		Password:    "patient123",
		PhoneNumber: "+1234567894",
		Gender:      "female",
		DateOfBirth: &patientDob,
		Type:        models.UserTypePatient,
		Address:     "321 Patient Lane",
	}
	if err := DB.Create(&patientUser).Error; err != nil {
		return err
	}

	// Create Patient Profile
	medicalHistory, _ := json.Marshal([]string{"Hypertension", "Diabetes"})
	allergies, _ := json.Marshal([]string{"Penicillin", "Peanuts"})
	patient := models.Patient{
		UserID:           patientUser.ID,
		EmergencyContact: "+1234567895",
		BloodType:        "A+",
		Weight:           65.5,
		MedicalHistory:   string(medicalHistory),
		Allergies:        string(allergies),
	}
	if err := DB.Create(&patient).Error; err != nil {
		return err
	}

	// Create Assistant User
	assistantDob := time.Date(1992, 11, 25, 0, 0, 0, 0, time.UTC)
	assistantUser := models.User{
		FirstName:   "Emma",
		LastName:    "Brown",
		Email:       "assistant@medworld.com",
		Password:    "assistant123",
		PhoneNumber: "+1234567896",
		Gender:      "female",
		DateOfBirth: &assistantDob,
		Type:        models.UserTypeAssistant,
		Address:     "654 Assistant St",
	}
	if err := DB.Create(&assistantUser).Error; err != nil {
		return err
	}

	// Create Assistant Profile
	assistant := models.Assistant{
		UserID:    assistantUser.ID,
		CabinetID: cabinet.ID,
	}
	if err := DB.Create(&assistant).Error; err != nil {
		return err
	}

	// Create Sample Appointment
	appointmentDate := time.Now().Add(24 * time.Hour)
	appointment := models.Appointment{
		Date:      appointmentDate,
		Status:    models.AppointmentStatusScheduled,
		PatientID: patient.ID,
		DoctorID:  doctor.ID,
		CabinetID: cabinet.ID,
	}
	if err := DB.Create(&appointment).Error; err != nil {
		return err
	}

	// Create Sample Consultation
	prescriptions, _ := json.Marshal([]string{"Aspirin 100mg - 1 tablet daily", "Lisinopril 10mg - 1 tablet daily"})
	consultation := models.Consultation{
		DoctorID:      doctor.ID,
		PatientID:     patient.ID,
		AppointmentID: appointment.ID,
		Notes:         "Patient doing well. Blood pressure under control. Continue current medication.",
		Prescriptions: string(prescriptions),
	}
	if err := DB.Create(&consultation).Error; err != nil {
		return err
	}

	log.Println("Database seeded successfully")
	return nil
}

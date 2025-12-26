package database

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"medworld-backend/models"
	"time"
)

// SeedDatabase populates the database with initial data
func SeedDatabase() error {
	log.Println("Seeding database...")

	// Check if data already exists - if so, we generally skip or could allow force reseeding
	var userCount int64
	DB.Model(&models.User{}).Count(&userCount)
	if userCount > 50 { // Increased threshold or we can check for specific critical data
		log.Println("Database appears relatively populated, skipping basic seed...")
		return nil
	}

	// 1. Create Super Admin
	if err := seedSuperAdmin(); err != nil {
		return err
	}

	// 2. Create Cabinets
	cabinets, err := seedCabinets()
	if err != nil {
		return err
	}

	// 3. Create Doctors & Admin Doctors
	doctors, err := seedDoctors(cabinets)
	if err != nil {
		return err
	}

	// 4. Create Assistants
	if err := seedAssistants(cabinets); err != nil {
		return err
	}

	// 5. Create Patients
	patients, err := seedPatients()
	if err != nil {
		return err
	}

	// 6. Create Appointments & Consultations
	if err := seedAppointmentsAndConsultations(doctors, patients, cabinets); err != nil {
		return err
	}

	// 7. Create Calendars
	if err := seedCalendars(doctors); err != nil {
		return err
	}

	// 8. Create Messages
	if err := seedMessages(); err != nil {
		return err
	}

	log.Println("Database seeded successfully with rich data!")
	return nil
}

func seedSuperAdmin() error {
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
		Address:     "MedWorld HQ",
	}
	// Check if exists
	var count int64
	DB.Model(&models.User{}).Where("email = ?", superAdmin.Email).Count(&count)
	if count == 0 {
		return DB.Create(&superAdmin).Error
	}
	return nil
}

func seedCabinets() ([]models.Cabinet, error) {
	cabinetData := []struct {
		Name     string
		Address  string
		Location map[string]interface{}
	}{
		{
			Name:     "Central Medical Clinic",
			Address:  "123 Healthcare Blvd",
			Location: map[string]interface{}{"address": "123 Healthcare Blvd", "latitude": 40.7128, "longitude": -74.0060},
		},
		{
			Name:     "Northside Family Practice",
			Address:  "45 North Way",
			Location: map[string]interface{}{"address": "45 North Way", "latitude": 40.7580, "longitude": -73.9855},
		},
		{
			Name:     "West End Specialists",
			Address:  "88 West End Ave",
			Location: map[string]interface{}{"address": "88 West End Ave", "latitude": 40.7899, "longitude": -73.9799},
		},
	}

	openingHours, _ := json.Marshal(map[string]interface{}{
		"monday":    map[string]string{"open": "09:00", "close": "17:00"},
		"tuesday":   map[string]string{"open": "09:00", "close": "17:00"},
		"wednesday": map[string]string{"open": "09:00", "close": "17:00"},
		"thursday":  map[string]string{"open": "09:00", "close": "17:00"},
		"friday":    map[string]string{"open": "09:00", "close": "16:00"},
	})

	var createdCabinets []models.Cabinet
	trueVal := true

	for _, data := range cabinetData {
		locJSON, _ := json.Marshal(data.Location)
		cab := models.Cabinet{
			Name:             data.Name,
			Phone:            fmt.Sprintf("+1555%06d", rand.Intn(1000000)),
			Location:         string(locJSON),
			OpeningHours:     string(openingHours),
			AccessHandicap:   &trueVal,
			HasParking:       &trueVal,
			HasWifi:          &trueVal,
			AcceptsUrgent:    &trueVal,
			AcceptsInsurance: &trueVal,
		}
		// AdminID is required, linking to SuperAdmin for now or first user
		var sysAdmin models.User
		DB.Where("type = ?", models.UserTypeSuperAdmin).First(&sysAdmin)
		cab.AdminID = sysAdmin.ID

		if err := DB.FirstOrCreate(&cab, models.Cabinet{Name: cab.Name}).Error; err != nil {
			return nil, err
		}
		createdCabinets = append(createdCabinets, cab)
	}
	return createdCabinets, nil
}

func seedDoctors(cabinets []models.Cabinet) ([]models.Doctor, error) {
	specialties := []string{"Cardiology", "General Practice", "Dermatology", "Pediatrics", "Neurology"}
	firstNames := []string{"John", "Sarah", "Emily", "Michael", "David", "Jessica", "Robert", "Jennifer", "William", "Elizabeth", "James", "Linda"}
	lastNames := []string{"Smith", "Johnson", "Brown", "Williams", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez"}

	var createdDoctors []models.Doctor

	// Ensure we have our main Admin Doctor for testing
	adminEmail := "admin.doctor@medworld.com"
	if err := createSpecificDoctor("Sarah", "Johnson", adminEmail, "admin123", models.UserTypeAdmin, cabinets[0].ID, "General Practice"); err != nil {
		return nil, err
	}

	// Ensure we have our main regular Doctor for testing
	docEmail := "doctor@medworld.com"
	if err := createSpecificDoctor("John", "Smith", docEmail, "doctor123", models.UserTypeDoctor, cabinets[0].ID, "Cardiology"); err != nil {
		return nil, err
	}

	// Retrieve them to add to list
	var specificDocs []models.Doctor
	DB.Preload("User").Find(&specificDocs) // This gets all, which is fine
	createdDoctors = append(createdDoctors, specificDocs...)

	// Create random doctors
	for i := 0; i < 10; i++ {
		fName := firstNames[rand.Intn(len(firstNames))]
		lName := lastNames[rand.Intn(len(lastNames))]
		email := fmt.Sprintf("doctor%d@medworld.com", i+1)
		cabID := cabinets[rand.Intn(len(cabinets))].ID
		spec := specialties[rand.Intn(len(specialties))]
		uType := models.UserTypeDoctor
		if i%3 == 0 {
			uType = models.UserTypeAdmin // Make some admins
		}

		// Check if email taken by specific docs
		if email == adminEmail || email == docEmail {
			continue
		}

		if err := createSpecificDoctor(fName, lName, email, "password", uType, cabID, spec); err != nil {
			log.Printf("Skipping duplicate doctor %s: %v", email, err)
		}
	}

	// Refetch all doctors
	DB.Preload("User").Find(&createdDoctors)
	return createdDoctors, nil
}

func createSpecificDoctor(fName, lName, email, password string, uType models.UserType, cabID uint, spec string) error {
	var existing models.User
	if err := DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil // Already exists
	}

	dob := time.Date(1980+rand.Intn(15), time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC)
	user := models.User{
		FirstName:   fName,
		LastName:    lName,
		Email:       email,
		Password:    password,
		PhoneNumber: fmt.Sprintf("+1555%06d", rand.Intn(1000000)),
		Gender:      []string{"male", "female"}[rand.Intn(2)],
		DateOfBirth: &dob,
		Type:        uType,
		Address:     fmt.Sprintf("%d Doctor Lane", rand.Intn(999)),
	}
	if err := DB.Create(&user).Error; err != nil {
		return err
	}

	start := time.Date(2005+rand.Intn(15), 1, 1, 0, 0, 0, 0, time.UTC)
	doc := models.Doctor{
		UserID:               user.ID,
		Speciality:           spec,
		CareerStart:          &start,
		ConsultationPrice:    float64(50 + rand.Intn(200)),
		ConsultationDuration: 30,
		CabinetID:            cabID,
	}
	return DB.Create(&doc).Error
}

func seedAssistants(cabinets []models.Cabinet) error {
	firstNames := []string{"Emma", "Olivia", "Ava", "Isabella", "Sophia", "Mia"}
	lastNames := []string{"Charlotte", "Amelia", "Harper", "Evelyn", "Abigail", "Emily"}

	// Specific assistant
	admAssist := "assistant@medworld.com"
	if err := createSpecificAssistant("Emma", "Brown", admAssist, "assistant123", cabinets[0].ID); err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		fName := firstNames[rand.Intn(len(firstNames))]
		lName := lastNames[rand.Intn(len(lastNames))]
		email := fmt.Sprintf("assistant%d@medworld.com", i+1)
		cabID := cabinets[rand.Intn(len(cabinets))].ID

		if email == admAssist {
			continue
		}
		if err := createSpecificAssistant(fName, lName, email, "password", cabID); err != nil {
			continue
		}
	}
	return nil
}

func createSpecificAssistant(fName, lName, email, password string, cabID uint) error {
	var existing models.User
	if err := DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil
	}

	dob := time.Date(1990+rand.Intn(10), time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC)
	user := models.User{
		FirstName:   fName,
		LastName:    lName,
		Email:       email,
		Password:    password,
		PhoneNumber: fmt.Sprintf("+1555%06d", rand.Intn(1000000)),
		Gender:      []string{"male", "female"}[rand.Intn(2)],
		DateOfBirth: &dob,
		Type:        models.UserTypeAssistant,
		Address:     fmt.Sprintf("%d Assistant Way", rand.Intn(999)),
	}
	if err := DB.Create(&user).Error; err != nil {
		return err
	}

	assist := models.Assistant{
		UserID:    user.ID,
		CabinetID: cabID,
	}

	// Ensure assistant has exactly one doctor (first doctor in the cabinet)
	var doctor models.Doctor
	if err := DB.Where("cabinet_id = ?", cabID).First(&doctor).Error; err == nil {
		assist.DoctorID = doctor.ID
	}

	return DB.Create(&assist).Error
}

func seedPatients() ([]models.Patient, error) {
	firstNames := []string{"James", "Mary", "Robert", "Patricia", "John", "Jennifer", "Michael", "Linda", "David", "Elizabeth", "William", "Barbara", "Richard", "Susan", "Joseph", "Jessica", "Thomas", "Sarah", "Charles", "Karen"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin"}

	var createdPatients []models.Patient

	// Specific patient
	patEmail := "patient@medworld.com"
	if err := createSpecificPatient("Alice", "Williams", patEmail, "patient123"); err != nil {
		return nil, err
	}

	for i := 0; i < 20; i++ {
		fName := firstNames[rand.Intn(len(firstNames))]
		lName := lastNames[rand.Intn(len(lastNames))]
		email := fmt.Sprintf("patient%d@medworld.com", i+1)

		if email == patEmail {
			continue
		}
		if err := createSpecificPatient(fName, lName, email, "password"); err != nil {
			log.Printf("Error creating patient %s: %v", email, err)
			continue
		}
	}

	DB.Preload("User").Find(&createdPatients)
	return createdPatients, nil
}

func createSpecificPatient(fName, lName, email, password string) error {
	var existing models.User
	if err := DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil
	}

	dob := time.Date(1950+rand.Intn(50), time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC)
	user := models.User{
		FirstName:   fName,
		LastName:    lName,
		Email:       email,
		Password:    password,
		PhoneNumber: fmt.Sprintf("+1555%06d", rand.Intn(1000000)),
		Gender:      []string{"male", "female"}[rand.Intn(2)],
		DateOfBirth: &dob,
		Type:        models.UserTypePatient,
		Address:     fmt.Sprintf("%d Patient Rd", rand.Intn(999)),
	}
	if err := DB.Create(&user).Error; err != nil {
		return err
	}

	medicalHistory, _ := json.Marshal([]string{"Hypertension", "Diabetes"})
	allergies, _ := json.Marshal([]string{"Penicillin"})
	if rand.Intn(2) == 0 {
		medicalHistory, _ = json.Marshal([]string{})
		allergies, _ = json.Marshal([]string{})
	}

	patient := models.Patient{
		UserID:           user.ID,
		EmergencyContact: fmt.Sprintf("+1555%06d", rand.Intn(1000000)),
		BloodType:        []string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"}[rand.Intn(8)],
		Weight:           float64(50 + rand.Intn(50)),
		MedicalHistory:   string(medicalHistory),
		Allergies:        string(allergies),
	}
	return DB.Create(&patient).Error
}

func seedAppointmentsAndConsultations(doctors []models.Doctor, patients []models.Patient, cabinets []models.Cabinet) error {
	statuses := []models.AppointmentStatus{models.AppointmentStatusScheduled, models.AppointmentStatusCompleted, models.AppointmentStatusCancelled, models.AppointmentStatusNoShow}

	// 1. Force appointments for John Smith
	var johnSmith models.Doctor
	foundJohn := false
	for _, d := range doctors {
		if d.User.Email == "doctor@medworld.com" {
			johnSmith = d
			foundJohn = true
			break
		}
	}

	if foundJohn {
		// Assign first 5 patients to John Smith
		for i := 0; i < 5 && i < len(patients); i++ {
			appt := models.Appointment{
				Date:      time.Now().AddDate(0, 0, rand.Intn(10)+1).Truncate(time.Hour),
				Status:    models.AppointmentStatusScheduled,
				PatientID: patients[i].ID,
				DoctorID:  johnSmith.ID,
				CabinetID: johnSmith.CabinetID,
			}
			DB.Create(&appt)
		}
	}

	for _, patient := range patients {
		// Each patient has 1-3 appointments
		count := 1 + rand.Intn(3)
		for i := 0; i < count; i++ {
			doctor := doctors[rand.Intn(len(doctors))]
			cabinetID := doctor.CabinetID
			if cabinetID == 0 {
				cabinetID = cabinets[0].ID
			}

			// Random date in the last year or next month
			daysOffset := rand.Intn(395) - 365
			apptDate := time.Now().AddDate(0, 0, daysOffset).Truncate(time.Hour)
			status := statuses[rand.Intn(len(statuses))]
			if daysOffset > 0 {
				status = models.AppointmentStatusScheduled
			} else if status == models.AppointmentStatusScheduled {
				status = models.AppointmentStatusCompleted // Past appointments usually completed
			}

			appt := models.Appointment{
				Date:      apptDate,
				Status:    status,
				PatientID: patient.ID,
				DoctorID:  doctor.ID,
				CabinetID: cabinetID,
			}
			if err := DB.Create(&appt).Error; err != nil {
				return err
			}

			// If completed, maybe create consultation
			if status == models.AppointmentStatusCompleted {
				prescriptions, _ := json.Marshal([]string{"Generic Medicine 500mg"})
				consult := models.Consultation{
					DoctorID:      doctor.ID,
					PatientID:     patient.ID,
					AppointmentID: appt.ID,
					Notes:         "Routine checkup. Patient stable.",
					Prescriptions: string(prescriptions),
				}
				DB.Create(&consult)
			}
		}
	}
	return nil
}

func seedCalendars(doctors []models.Doctor) error {
	// Create calendars with YYYY-MM-DD date format
	baseDate := time.Now()

	for _, doctor := range doctors {
		availabilityData := []map[string]interface{}{
			{
				"date":  baseDate.AddDate(0, 0, 1).Format("2006-01-02"),
				"slots": []string{"09:00", "09:30", "10:00", "11:00"},
			},
			{
				"date":  baseDate.AddDate(0, 0, 2).Format("2006-01-02"),
				"slots": []string{"14:00", "14:30", "15:00", "16:00"},
			},
			{
				"date":  baseDate.AddDate(0, 0, 3).Format("2006-01-02"),
				"slots": []string{"09:00", "09:30", "10:00"},
			},
			{
				"date":  baseDate.AddDate(0, 0, 4).Format("2006-01-02"),
				"slots": []string{"10:00", "11:00", "14:00"},
			},
			{
				"date":  baseDate.AddDate(0, 0, 5).Format("2006-01-02"),
				"slots": []string{"08:30", "09:00", "09:30"},
			},
		}

		availability, _ := json.Marshal(availabilityData)

		cal := models.Calendar{
			DoctorID:     doctor.ID,
			CabinetID:    doctor.CabinetID,
			Availability: string(availability),
		}
		if err := DB.Create(&cal).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedMessages() error {
	// Find our main doctor
	var doctorUser models.User
	if err := DB.Where("email = ?", "doctor@medworld.com").First(&doctorUser).Error; err != nil {
		return nil
	}

	// Find admin
	var adminUser models.User
	if err := DB.Where("email = ?", "admin.doctor@medworld.com").First(&adminUser).Error; err != nil {
		return nil
	}

	// Create conversation
	msgs := []models.Message{
		{
			SenderID:   adminUser.ID,
			ReceiverID: doctorUser.ID,
			Content:    "Welcome to the team, Dr. Smith!",
			CreatedAt:  time.Now().Add(-24 * time.Hour),
		},
		{
			SenderID:   doctorUser.ID,
			ReceiverID: adminUser.ID,
			Content:    "Thank you! Glad to be here.",
			CreatedAt:  time.Now().Add(-23 * time.Hour),
		},
		{
			SenderID:   adminUser.ID,
			ReceiverID: doctorUser.ID,
			Content:    "Please check your calendar for the upcoming week.",
			CreatedAt:  time.Now().Add(-2 * time.Hour),
		},
	}

	for _, m := range msgs {
		DB.Create(&m)
	}
	return nil
}

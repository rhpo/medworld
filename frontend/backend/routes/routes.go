package routes

import (
	"medworld-backend/handlers"
	"medworld-backend/middleware"
	"medworld-backend/models"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// Initialize handlers
	authHandler := &handlers.AuthHandler{}
	userHandler := &handlers.UserHandler{}
	doctorHandler := &handlers.DoctorHandler{}
	patientHandler := &handlers.PatientHandler{}
	appointmentHandler := &handlers.AppointmentHandler{}
	consultationHandler := &handlers.ConsultationHandler{}
	cabinetHandler := &handlers.CabinetHandler{}
	allHandler := &handlers.AllHandler{}

	// API v1 group
	api := app.Group("/api/v1")

	// Public routes - Authentication
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
	auth.Post("/logout", authHandler.Logout)

	// Protected routes - require authentication
	auth.Get("/me", middleware.AuthMiddleware, authHandler.Me)

	// User routes
	users := api.Group("/users", middleware.AuthMiddleware)
	users.Get("/:id", userHandler.GetUserByID)
	users.Put("/:id", userHandler.UpdateUser)

	// Doctor routes
	doctors := api.Group("/doctors", middleware.AuthMiddleware)
	doctors.Get("/", doctorHandler.ListDoctors)
	doctors.Get("/:id", doctorHandler.GetDoctorByID)
	doctors.Get("/:id/appointments", doctorHandler.GetDoctorAppointments)
	doctors.Get("/:id/consultations", doctorHandler.GetDoctorConsultations)
	doctors.Get("/:id/patients", doctorHandler.GetDoctorPatients)
	doctors.Get("/search/filter", doctorHandler.SearchDoctors)

	// Patient routes
	patients := api.Group("/patients", middleware.AuthMiddleware)
	patients.Get("/", patientHandler.ListPatients)
	patients.Get("/:id", patientHandler.GetPatientByID)
	patients.Get("/:id/appointments", patientHandler.GetPatientAppointments)
	patients.Get("/:id/consultations", patientHandler.GetPatientConsultations)

	// Appointment routes
	appointments := api.Group("/appointments", middleware.AuthMiddleware)
	appointments.Get("/", appointmentHandler.ListAppointments)
	appointments.Get("/:id", appointmentHandler.GetAppointmentByID)
	appointments.Post("/",
		middleware.RequireAnyPermission(middleware.PermBookAppointment, middleware.PermEditAppointment),
		appointmentHandler.CreateAppointment)
	appointments.Put("/:id",
		middleware.RequirePermission(middleware.PermEditAppointment),
		appointmentHandler.UpdateAppointment)
	appointments.Delete("/:id",
		middleware.RequirePermission(middleware.PermCancelAppointment),
		appointmentHandler.DeleteAppointment)

	// Consultation routes
	consultations := api.Group("/consultations", middleware.AuthMiddleware)
	consultations.Get("/", consultationHandler.ListConsultations)
	consultations.Get("/:id", consultationHandler.GetConsultationByID)
	consultations.Post("/",
		middleware.RequirePermission(middleware.PermAddConsultation),
		consultationHandler.CreateConsultation)
	consultations.Put("/:id",
		middleware.RequirePermission(middleware.PermEditConsultation),
		consultationHandler.UpdateConsultation)

	// Cabinet routes
	cabinets := api.Group("/cabinets", middleware.AuthMiddleware)
	cabinets.Get("/", cabinetHandler.ListCabinets)
	cabinets.Get("/:id", cabinetHandler.GetCabinetByID)
	cabinets.Get("/:id/doctors", cabinetHandler.GetCabinetDoctors)
	cabinets.Get("/:id/appointments", cabinetHandler.GetCabinetAppointments)
	cabinets.Get("/:id/assistants", cabinetHandler.GetCabinetAssistants)
	cabinets.Post("/",
		middleware.RequirePermission(middleware.PermAddCabinet),
		cabinetHandler.CreateCabinet)

	// All routes (admin/superadmin only)
	all := api.Group("/all",
		middleware.AuthMiddleware,
		middleware.RequireUserType(models.UserTypeSuperAdmin, models.UserTypeAdmin))
	all.Get("/appointments", allHandler.ListAllAppointments)
	all.Get("/doctors", allHandler.ListAllDoctors)
	all.Get("/patients", allHandler.ListAllPatients)
	all.Get("/assistants", allHandler.ListAllAssistants)
	all.Get("/cabinets", allHandler.ListAllCabinets)
	all.Get("/users", allHandler.ListAllUsers)
	all.Get("/consultations", allHandler.ListAllConsultations)

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "MedWorld API is running",
		})
	})
}

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
	calendarHandler := &handlers.CalendarHandler{}
	messageHandler := &handlers.MessageHandler{}

	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
	auth.Post("/logout", authHandler.Logout)

	// Protected routes thatt require authentication
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
	cabinets.Get("/:id", middleware.RequireCabinetAccess, cabinetHandler.GetCabinetByID)
	cabinets.Get("/:id/doctors", middleware.RequireCabinetAccess, cabinetHandler.GetCabinetDoctors)
	cabinets.Get("/:id/appointments", middleware.RequireCabinetAccess, cabinetHandler.GetCabinetAppointments)
	cabinets.Get("/:id/assistants", middleware.RequireCabinetAccess, cabinetHandler.GetCabinetAssistants)
	cabinets.Post("/",
		middleware.RequirePermission(middleware.PermAddCabinet),
		cabinetHandler.CreateCabinet)
	cabinets.Put("/:id",
		middleware.RequireCabinetAccess,
		middleware.RequirePermission(middleware.PermEditCabinet),
		cabinetHandler.UpdateCabinet)
	cabinets.Delete("/:id",
		middleware.RequireCabinetAccess,
		middleware.RequirePermission(middleware.PermRemoveCabinet),
		cabinetHandler.DeleteCabinet)
	cabinets.Post("/:id/doctors",
		middleware.RequireCabinetAccess,
		middleware.RequirePermission(middleware.PermEditCabinet),
		cabinetHandler.AddDoctorToCabinet)
	cabinets.Delete("/:id/doctors/:doctorId",
		middleware.RequireCabinetAccess,
		middleware.RequirePermission(middleware.PermEditCabinet),
		cabinetHandler.RemoveDoctorFromCabinet)
	cabinets.Post("/:id/assistants",
		middleware.RequireCabinetAccess,
		middleware.RequirePermission(middleware.PermEditCabinet),
		cabinetHandler.AddAssistantToCabinet)
	cabinets.Post("/:id/assistants/create",
		middleware.RequireCabinetAccess,
		middleware.RequirePermission(middleware.PermAssignAssistant),
		cabinetHandler.CreateAssistantInCabinet)
	cabinets.Put("/:id/assistants/:assistantId/assign",
		middleware.RequireCabinetAccess,
		middleware.RequirePermission(middleware.PermAssignAssistant),
		cabinetHandler.AssignAssistantToDoctor)
	cabinets.Delete("/:id/assistants/:assistantId",
		middleware.RequireCabinetAccess,
		middleware.RequirePermission(middleware.PermEditCabinet),
		cabinetHandler.RemoveAssistantFromCabinet)

	// All routes (admin/superadmin only)
	all := api.Group("/all",
		middleware.AuthMiddleware,
		middleware.RequireUserType(models.UserTypeSuperAdmin))
	all.Get("/appointments", allHandler.ListAllAppointments)
	all.Get("/doctors", allHandler.ListAllDoctors)
	all.Get("/patients", allHandler.ListAllPatients)
	all.Get("/assistants", allHandler.ListAllAssistants)
	all.Get("/cabinets", allHandler.ListAllCabinets)
	all.Get("/users", allHandler.ListAllUsers)
	all.Get("/consultations", allHandler.ListAllConsultations)

	// Calendar routes
	calendars := api.Group("/calendars", middleware.AuthMiddleware)
	calendars.Get("/", calendarHandler.ListCalendars)
	calendars.Get("/:id", calendarHandler.GetCalendarByID)
	calendars.Put("/:id",
		middleware.RequireAnyPermission(middleware.PermEditCalendar),
		calendarHandler.UpdateCalendar)

	// Message routes
	messages := api.Group("/messages", middleware.AuthMiddleware)
	messages.Get("/",
		middleware.RequireAnyPermission(middleware.PermViewMessage),
		messageHandler.ListMyMessages)
	messages.Get("/recipients",
		middleware.RequireAnyPermission(middleware.PermViewMessage),
		messageHandler.ListRecipients)
	messages.Post("/",
		middleware.RequireAnyPermission(middleware.PermSendMessage),
		messageHandler.CreateMessage)

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "MedWorld API is running",
		})
	})
}

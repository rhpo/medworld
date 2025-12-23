package middleware

import (
	"medworld-backend/models"

	"github.com/gofiber/fiber/v2"
)

// Permission represents a specific action a user can perform
type Permission string

const (
	// Cabinet permissions
	PermViewCabinet   Permission = "view_cabinet"
	PermAddCabinet    Permission = "add_cabinet"
	PermEditCabinet   Permission = "edit_cabinet"
	PermManageCabinet Permission = "manage_cabinet"
	PermRemoveCabinet Permission = "remove_cabinet"

	// Admin permissions
	PermSetAdminDoctor Permission = "set_admin_doctor"

	// Doctor permissions
	PermViewDoctor   Permission = "view_doctor"
	PermAddDoctor    Permission = "add_doctor"
	PermEditDoctor   Permission = "edit_doctor"
	PermRemoveDoctor Permission = "remove_doctor"

	// Calendar permissions
	PermViewCalendar Permission = "view_calendar"
	PermEditCalendar Permission = "edit_calendar"

	// Assistant permissions
	PermViewAssistant   Permission = "view_assistant"
	PermAddAssistant    Permission = "add_assistant"
	PermEditAssistant   Permission = "edit_assistant"
	PermRemoveAssistant Permission = "remove_assistant"
	PermAssignAssistant Permission = "assign_assistant"

	// Appointment permissions
	PermCancelAppointment Permission = "cancel_appointment"
	PermEditAppointment   Permission = "edit_appointment"
	PermViewAppointment   Permission = "view_appointment"
	PermBookAppointment   Permission = "book_appointment"

	// Patient permissions
	PermViewPatient   Permission = "view_patient"
	PermAddPatient    Permission = "add_patient"
	PermEditPatient   Permission = "edit_patient"
	PermRemovePatient Permission = "remove_patient"

	// Consultation permissions
	PermViewConsultation   Permission = "view_consultation"
	PermAddConsultation    Permission = "add_consultation"
	PermEditConsultation   Permission = "edit_consultation"
	PermRemoveConsultation Permission = "remove_consultation"

	// Prescription permissions
	PermAddPrescription  Permission = "add_prescription"
	PermEditPrescription Permission = "edit_prescription"

	// Medical folder permissions
	PermViewMedicalFolder   Permission = "view_medical_folder"
	PermAddMedicalFolder    Permission = "add_medical_folder"
	PermEditMedicalFolder   Permission = "edit_medical_folder"
	PermRemoveMedicalFolder Permission = "remove_medical_folder"

	// Message permissions
	PermViewMessage Permission = "view_message"
	PermSendMessage Permission = "send_message"

	// Patient-specific permissions
	PermRateDoctor Permission = "rate_doctor"

	// All permissions
	PermAll Permission = "all"
)

// getPermissionsForUserType returns all permissions for a given user type
func getPermissionsForUserType(userType models.UserType) []Permission {
	switch userType {
	case models.UserTypeSuperAdmin:
		return []Permission{
			// Doctor management
			PermViewDoctor, PermAddDoctor, PermEditDoctor, PermRemoveDoctor,
			// Cabinet management
			PermViewCabinet, PermAddCabinet, PermEditCabinet, PermRemoveCabinet,
			// Assistant management
			PermViewAssistant, PermAddAssistant, PermEditAssistant, PermRemoveAssistant,
			// Messages
			PermViewMessage, PermSendMessage,
		}

	case models.UserTypeAdmin:
		return []Permission{
			// Doctor permissions (Admin is also a doctor)
			PermViewMedicalFolder, PermAddMedicalFolder, PermEditMedicalFolder, PermRemoveMedicalFolder,
			PermViewCalendar, PermEditCalendar,
			PermViewPatient, PermAddPatient, PermEditPatient, PermRemovePatient,
			PermViewConsultation, PermAddConsultation, PermEditConsultation,
			PermViewAppointment, PermEditAppointment, PermCancelAppointment,
			PermViewMessage, PermSendMessage,
			PermAssignAssistant,
			// Admin-specific permissions
			PermEditCabinet, PermAddDoctor, PermEditDoctor, PermSetAdminDoctor,
		}

	case models.UserTypeDoctor:
		return []Permission{
			// Medical folder management
			PermViewMedicalFolder, PermAddMedicalFolder, PermEditMedicalFolder, PermRemoveMedicalFolder,
			// Calendar management
			PermViewCalendar, PermEditCalendar,
			// Patient management
			PermViewPatient, PermAddPatient, PermEditPatient, PermRemovePatient,
			// Consultation management
			PermViewConsultation, PermAddConsultation, PermEditConsultation,
			// Appointment management
			PermViewAppointment, PermEditAppointment, PermCancelAppointment,
			// Messages
			PermViewMessage, PermSendMessage,
			// Assistant assignment
			PermAssignAssistant,
		}

	case models.UserTypeAssistant:
		return []Permission{
			// Appointment management
			PermViewAppointment, PermEditAppointment, PermCancelAppointment,
			// Consultation management
			PermViewConsultation, PermAddConsultation, PermEditConsultation,
			// Prescription management
			PermAddPrescription, PermEditPrescription,
			// Calendar
			PermViewCalendar, PermEditCalendar,
			// Patient
			PermAddPatient, PermEditPatient,
		}

	case models.UserTypePatient:
		return []Permission{
			// Appointment
			PermBookAppointment, PermViewAppointment, PermCancelAppointment,
			PermRateDoctor,
		}

	default:
		return []Permission{}
	}
}

// hasPermission checks if a user has a specific permission
func hasPermission(userType models.UserType, permission Permission) bool {
	permissions := getPermissionsForUserType(userType)
	for _, p := range permissions {
		if p == permission || p == PermAll {
			return true
		}
	}
	return false
}

// RequirePermission middleware checks if user has required permission
func RequirePermission(permission Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		if !hasPermission(user.Type, permission) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "You don't have permission to perform this action",
			})
		}

		return c.Next()
	}
}

// RequireAnyPermission middleware checks if user has any of the required permissions
func RequireAnyPermission(permissions ...Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		for _, permission := range permissions {
			if hasPermission(user.Type, permission) {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You don't have permission to perform this action",
		})
	}
}

// RequireUserType middleware checks if user has specific user type
func RequireUserType(types ...models.UserType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		for _, t := range types {
			if user.Type == t {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You don't have permission to perform this action",
		})
	}
}

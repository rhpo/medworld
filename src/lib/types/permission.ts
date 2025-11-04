import { Users, type User, type UserType } from "./users/";

export type Permission =
    // Cabinets
    'view_cabinet' |
    'add_cabinet' |
    'edit_cabinet' |
    'remove_cabinet' |

    // Admin
    'set_admin_doctor' |

    // Doctor
    'view_doctor' |
    'add_doctor' |
    'edit_doctor' |
    'remove_doctor' |

    // Doctor's Calendar
    'view_calendar' |
    'edit_calendar' |

    // Assistant
    'view_assistant' |
    'add_assistant' |
    'edit_assistant' |
    'remove_assistant' |
    'assign_assistant' |

    // Appointment
    'cancel_appointment' |
    'edit_appointment' |

    // Patient
    'view_patient' |
    'add_patient' |
    'edit_patient' |
    'remove_patient' |

    // Consultation
    'view_consultation' |
    'add_consultation' |
    'edit_consultation' |
    'remove_consultation' |

    // MedicalFolder
    'view_medical_folder' |
    'add_medical_folder' |
    'edit_medical_folder' |
    'remove_medical_folder' |

    // Appointment
    'view_appointment' |
    'book_appointment' |

    // Patient-level permissions
    'rate_doctor' |
    'view_messages'

    'all';

export const ConsultationManagement: Permission[] = [
    'view_consultation',
    'add_consultation',
    'edit_consultation',
]

export const CabinetManagement: Permission[] = [
    'view_cabinet',
    'add_cabinet',
    'edit_cabinet',
    'remove_cabinet',
]

export const PatientManagement: Permission[] = [
    'view_patient',
    'add_patient',
    'edit_patient',
    'remove_patient',
]

export const AssistantManagement: Permission[] = [
    'view_assistant',
    'add_assistant',
    'edit_assistant',
    'remove_assistant',
]

export const DoctorManagement: Permission[] = [
    'view_doctor',
    'add_doctor',
    'edit_doctor',
    'remove_doctor',
]

export const CalendarManagement: Permission[] = [
    'view_calendar',
    'edit_calendar',
]

export const MedicalFolderManagement: Permission[] = [
    'view_medical_folder',
    'add_medical_folder',
    'edit_medical_folder',
    'remove_medical_folder',
]

export const DoctorPermissions: Permission[] = [
    // Gestion Dossier Medical
    ...MedicalFolderManagement,

    // Gestion de calendar
    ...CalendarManagement,

    //Gestion des patients
    ...PatientManagement,

    // Gestion de consultations
    ...ConsultationManagement,

    'assign_assistant'

];

export const SuperAdminPermissions: Permission[] = [
    ...DoctorManagement,

    // Gestion de cabinets
    ...CabinetManagement,

    // Gestion d'assistants
    ...AssistantManagement,

];

export const AdminPermissions: Permission[] = [

    // Il a les mêmes permissions de docteur.
    ...DoctorPermissions,

    // Modifier Cabinet
    'edit_cabinet',
    'add_doctor',
    'edit_doctor',
    'set_admin_doctor'

];

export const AssistantPermissions: Permission[] = [
    // Appointments
    'view_appointment',
    'cancel_appointment',
    'edit_appointment',


    // Calendar
    'view_calendar',
    'edit_calendar',

    // Patients
    'add_patient',
    'edit_patient',

];

export const PatientPermissions: Permission[] = [
    // View/Book/cancel Appointments
    'book_appointment',
    'view_appointment',
    'cancel_appointment'
];

// Usqble Functions
export function getPermissionsFromUserType(type: UserType): Permission[] {
    switch (type) {
        case Users.SuperAdmin:
            return SuperAdminPermissions;
        case Users.Admin:
            return AdminPermissions;
        case Users.Doctor:
            return DoctorPermissions;
        case Users.Assistant:
            return AssistantPermissions;
        case Users.Patient:
            return PatientPermissions;
        default:
            return [];
    }
}

export function hasPermission(user: User<any>, permission: Permission): boolean {
    return getPermissionsFromUserType(user.type).includes(permission);
}

export type Group = 'patients' | 'assistants' | 'doctors' | 'appointments' | 'calendar' | 'messages';

export function permissionGroupMap(permissions: Permission[]): Record<Group, boolean> {
        return {
            patients:
                permissions.filter((e) => e.endsWith("_patient")).length > 0,
            assistants:
                permissions.filter((e) => e.endsWith("_assistant")).length > 0,
            doctors:
                permissions.filter((e) => e.endsWith("_doctor")).length > 0,
            appointments:
                permissions.filter((e) => e.endsWith("_appointment")).length >
                0,
            calendar: permissions.includes("view_calendar"),
            messages: permissions.includes("view_messages"),
        };
    }

export function groupPermissions(permissions: Permission[]): Group[] {
    const map = permissionGroupMap(permissions);
    return Object.keys(map).filter((key) => map[key as Group]) as Group[];
}

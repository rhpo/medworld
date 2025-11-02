import type { User } from './users';
import { Users } from './users';
import type { Appointment } from './appointment';
import type { Consultation } from './consultation';
import type { Cabinet } from './cabinet';
import type { Doctor } from './users/doctor';
import type { Assistant } from './users/assistant';
import { Speciality } from './speciality';
import type { Admin } from './users/admin';
import { PlanID, plans } from './plan';

// Generate SuperAdmin users first (they oversee the entire platform)
export const fakeSuperAdmin: User<Users.SuperAdmin> =
    {
        id: 0,
        firstName: 'Houria',
        lastName: 'Aichi',
        email: 'houria.aichi@medworld.dz',
        phoneNumber: '+213 555 222 000',
        type: Users.SuperAdmin,
        getFullName: () => 'Houria Aichi',
        createdAt: new Date('2023-01-01')
    }

// Generate Admin users (they manage individual cabinets)
export const fakeAdmins: Array<Admin> = [
    {
        id: 3,
        firstName: 'Kamel',
        lastName: 'Daoud',
        email: 'kamel.daoud@medworld.dz',
        phoneNumber: '+213 555 123 456',
        type: Users.Admin,
        getFullName: () => 'Kamel Daoud',
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.basic]
    } as any,
    {
        id: 4,
        firstName: 'Yasmina',
        lastName: 'Khadra',
        email: 'yasmina.khadra@medworld.dz',
        phoneNumber: '+213 555 567 890',
        type: Users.Admin,
        getFullName: () => 'Yasmina Khadra',
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.standard]

    },
    {
        id: 5,
        firstName: 'Mohammed',
        lastName: 'Dib',
        email: 'mohammed.dib@medworld.dz',
        phoneNumber: '+213 555 901 234',
        type: Users.Admin,
        getFullName: () => 'Mohammed Dib',
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.premium]

    },
    {
        id: 6,
        firstName: 'Rachid',
        lastName: 'Boudjedra',
        email: 'rachid.boudjedra@medworld.dz',
        phoneNumber: '+213 555 345 678',
        type: Users.Admin,
        getFullName: () => 'Rachid Boudjedra',
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.basic]

    },
    {
        id: 7,
        firstName: 'Kateb',
        lastName: 'Yacine',
        email: 'kateb.yacine@medworld.dz',
        phoneNumber: '+213 555 789 012',
        type: Users.Admin,
        getFullName: () => 'Kateb Yacine',
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.standard]

    }
];

// Create helper function to make an assistant
const makeAssistant = (
    id: number,
    firstName: string,
    lastName: string,
    email: string,
    phoneNumber: string
): Assistant => ({
    id,
    firstName,
    lastName,
    email,
    phoneNumber,
    type: Users.Assistant,
    getFullName: () => `${firstName} ${lastName}`,
    createdAt: new Date('2023-03-01'),
    doctors: [], // Will be populated later
    cabinet: {} as Cabinet, // Will be populated later
    appointments: [],
    planAppointment: () => {},
    cancelAppointment: () => {},
    recordPayment: () => {}
});

// Generate assistants (each linked to a cabinet and doctor)
export const fakeAssistants: Assistant[] = [
    makeAssistant(8, 'Amira', 'Bouraoui', 'amira.bouraoui@medworld.dz', '+213 555 234 567'),
    makeAssistant(9, 'Assia', 'Djebar', 'assia.djebar@medworld.dz', '+213 555 678 901'),
    makeAssistant(10, 'Malika', 'Mokeddem', 'malika.mokeddem@medworld.dz', '+213 555 012 345'),
    makeAssistant(11, 'Leila', 'Aissaoui', 'leila.aissaoui@medworld.dz', '+213 555 456 789'),
    makeAssistant(12, 'Warda', 'Al-Jazairia', 'warda.aljazairia@medworld.dz', '+213 555 890 123')
];

// Generate base user arrays
const makeDoctor = (
    id: number,
    firstName: string,
    lastName: string,
    email: string,
    phoneNumber: string,
    speciality: Speciality,
    avatarUrl?: string
): Doctor => ({
    id,
    firstName,
    lastName,
    email,
    phoneNumber,
    avatarUrl: avatarUrl || 'https://i.pravatar.cc/50',
    type: Users.Doctor,
    speciality,
    dateOfBirth: new Date('1980-01-01'),
    careerStart: new Date('2005-01-01'),
    messages: [],
    reviews: [],
    calendars: [],
    consultations: [],
    cabinets: [],
    consultationDuration: 20,
    consultationPrice: 2000,
    assistants: [], // Will be populated after creating relationships
    getFullName: () => `${firstName} ${lastName}`,
    getYearsOfExperience: () => {
        const diff = new Date().getTime() - new Date('2005-01-01').getTime();
        return Math.floor(diff / (1000 * 60 * 60 * 24 * 365));
    }
});

// Generate doctors first so we can reference them in cabinets
export const fakeDoctors: Doctor[] = [
    makeDoctor(21, 'Rayan', 'El-miloudi', 'rayan.miloudi@medworld.dz', '+213 555 123 456', Speciality.CARDIOLOGY, 'https://cdn.discordapp.com/attachments/1269361766622564465/1434614521502372041/image.png?ex=6908f863&is=6907a6e3&hm=d3b284bec0be0c2eec0ce1811cba6a4861e854e4a26ee548b6530ee93f3c7f7a&'),
    makeDoctor(22, 'Manel', 'Belaiouer', 'manel.manel@medworld.dz', '+213 555 234 567', Speciality.FAMILY_MEDICINE, 'https://i.ibb.co/MDfvPPPS/image.png'),
    makeDoctor(23, 'Micipsa', 'Nekkmouche', 'micipsa.nekkmouche@medworld.dz', '+213 555 345 678', Speciality.PEDIATRICS, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS1aqGoiYxd-zaVibJFGg9xTrZdMuknj4LuGmiQ5Xizhoq4Tj9OiQIzXss6D5b0YRtHmKE7ofSA8uF5IXjuYAqNqlvHV1VpWhRxJ8IikA&s=10'),
    makeDoctor(24, 'Sarah', 'Zerrouki', 'sarah.zerrouki@medworld.dz', '+213 555 456 789', Speciality.DERMATOLOGY)
];

// Set up doctor-assistant relationships
// Dr. Benlamri has two assistants
fakeDoctors[0].assistants = [fakeAssistants[0], fakeAssistants[1]];
fakeAssistants[0].doctors = [fakeDoctors[0]];
fakeAssistants[1].doctors = [fakeDoctors[0]];

// Dr. Belkacem shares an assistant with Dr. Benchaabane
fakeDoctors[1].assistants = [fakeAssistants[2]];
fakeDoctors[2].assistants = [fakeAssistants[2]];
fakeAssistants[2].doctors = [fakeDoctors[1], fakeDoctors[2]];

// Dr. Zerrouki has one dedicated assistant
fakeDoctors[3].assistants = [fakeAssistants[3]];
fakeAssistants[3].doctors = [fakeDoctors[3]];

// Assistant Al-Jazairia is not currently assigned to any doctor
fakeAssistants[4].doctors = [];

// Generate fake cabinets
export const fakeCabinets: Cabinet[] = [
    {
        id: 1,
        name: 'Centre Médical Ibn Sina',
        address: '123 Rue Didouche Mourad, Alger',
        phone: '+213 21 23 45 67',
        createdAt: new Date('2024-01-01'),
        admin: fakeAdmins[0],
        doctors: [fakeDoctors[0], fakeDoctors[1]],
        assistants: [fakeAssistants[0]],
        openingHours: {
            'monday': { open: '08:00', close: '17:00' },
            'tuesday': { open: '08:00', close: '17:00' },
            'wednesday': { open: '08:00', close: '17:00' },
            'thursday': { open: '08:00', close: '17:00' },
            'friday': { open: '08:00', close: '12:00' }
        },accessHandicap: true,
        ratings: [],
        isClosed: false
    },
    {
        id: 2,
        name: 'Clinique El Djazair',
        address: '45 Boulevard Zighout Youcef, Alger',
        phone: '+213 21 34 56 78',
        createdAt: new Date('2024-02-01'),
        admin: fakeAdmins[1],
        doctors: [fakeDoctors[2]],
        assistants: [fakeAssistants[1]],
        openingHours: {
            'monday': { open: '09:00', close: '18:00' },
            'tuesday': { open: '09:00', close: '18:00' },
            'wednesday': { open: '09:00', close: '18:00' },
            'thursday': { open: '09:00', close: '18:00' },
            'friday': { open: '09:00', close: '13:00' }
        },
        accessHandicap: true,
        ratings: [],
        isClosed: false
    },
    {
        id: 3,
        name: 'Cabinet Médical Al Razi',
        address: '78 Avenue Hassiba Ben Bouali, Alger',
        phone: '+213 21 45 67 89',
        createdAt: new Date('2024-03-01'),
        admin: fakeAdmins[2],
        doctors: [fakeDoctors[3]],
        assistants: [fakeAssistants[2]],
        openingHours: {
            'monday': { open: '08:30', close: '17:30' },
            'tuesday': { open: '08:30', close: '17:30' },
            'wednesday': { open: '08:30', close: '17:30' },
            'thursday': { open: '08:30', close: '17:30' },
            'friday': { open: '08:30', close: '12:30' }
        },accessHandicap: true,
        ratings: [],
        isClosed: false
    },
    {
        id: 4,
        name: 'Centre de Santé Al Shifa',
        address: '123 Rue Larbi Ben M\'Hidi, Alger',
        phone: '+213 21 56 78 90',
        createdAt: new Date('2024-04-01'),
        admin: fakeAdmins[3],
        doctors: [fakeDoctors[0]],
        assistants: [fakeAssistants[3]],
        openingHours: {
            'monday': { open: '09:00', close: '18:00' },
            'tuesday': { open: '09:00', close: '18:00' },
            'wednesday': { open: '09:00', close: '18:00' },
            'thursday': { open: '09:00', close: '18:00' },
            'friday': { open: '09:00', close: '13:00' }
        },
        accessHandicap: true,
        ratings: [],
        isClosed: false
    },
    {
        id: 5,
        name: 'Clinique Al Amal',
        address: '56 Boulevard Colonel Amirouche, Alger',
        phone: '+213 21 67 89 01',
        createdAt: new Date('2024-05-01'),
        admin: fakeAdmins[4],
        doctors: [fakeDoctors[1]],
        assistants: [fakeAssistants[4]],
        openingHours: {
            'monday': { open: '08:00', close: '17:00' },
            'tuesday': { open: '08:00', close: '17:00' },
            'wednesday': { open: '08:00', close: '17:00' },
            'thursday': { open: '08:00', close: '17:00' },
            'friday': { open: '08:00', close: '12:00' }
        },
        accessHandicap: true,
        ratings: [],
        isClosed: false
    }
];


// Generate fake patients with Algerian names
export const fakePatients: Array<User<Users.Patient>> = [
    {
        id: 25,
        firstName: 'Riad',
        lastName: 'Mahrez',
        email: 'riad.mahrez@email.dz',
        phoneNumber: '+213 555 567 890',
        type: Users.Patient,
        avatarUrl: 'https://i.pravatar.cc/50',
        getFullName: () => 'Riad Mahrez'
    },
    {
        id: 26,
        firstName: 'Amira',
        lastName: 'Brahimi',
        email: 'amira.brahimi@email.dz',
        phoneNumber: '+213 555 678 901',
        type: Users.Patient,
        avatarUrl: 'https://i.pravatar.cc/50',

        getFullName: () => 'Amira Brahimi'
    },
    {
        id: 27,
        firstName: 'Karim',
        lastName: 'Ziani',
        email: 'karim.ziani@email.dz',
        phoneNumber: '+213 555 789 012',
        type: Users.Patient,
        avatarUrl: 'https://i.pravatar.cc/50',

        getFullName: () => 'Karim Ziani'
    },
    {
        id: 28,
        firstName: 'Louisa',
        lastName: 'Hanoune',
        email: 'louisa.hanoune@email.dz',
        phoneNumber: '+213 555 890 123',
        type: Users.Patient,
        avatarUrl: 'https://i.pravatar.cc/50',

        getFullName: () => 'Louisa Hanoune'
    }
];

// Generate fake appointments
export const fakeAppointments: Appointment[] = [
    {
        id: 1,
        date: new Date('2025-11-01T09:00:00'),
        status: 'confirmed',
        doctor: fakeDoctors[0],
        patient: fakePatients[0],
        cabinet: fakeCabinets[0],
        createdAt: new Date('2024-01-01'),
        updatedAt: new Date('2024-01-02')
    },
    {
        id: 2,
        date: new Date('2025-11-01T10:00:00'),
        status: 'pending',
        doctor: fakeDoctors[0],
        patient: fakePatients[1],
        cabinet: fakeCabinets[0],
        createdAt: new Date('2024-01-01')
    },
    {
        id: 3,
        date: new Date('2025-11-01T11:00:00'),
        status: 'confirmed',
        doctor: fakeDoctors[1],
        patient: fakePatients[2],
        cabinet: fakeCabinets[1],
        createdAt: new Date('2024-01-01')
    },
    {
        id: 4,
        date: new Date('2025-11-02T09:00:00'),
        status: 'pending',
        doctor: fakeDoctors[2],
        patient: fakePatients[3],
        cabinet: fakeCabinets[2],
        createdAt: new Date('2024-01-01')
    }
];

// Generate fake consultations
export const fakeConsultations: Consultation[] = [
    {
        id: 1,
        doctor: fakeDoctors[0],
        patient: fakePatients[0],
        appointment: fakeAppointments[0],
        notes: 'Patient presents with mild hypertension. Prescribed lifestyle changes and monitoring.',
        prescriptions: ['Lisinopril 10mg - 1x daily', 'Regular blood pressure monitoring'],
        attachments: ['bp_chart_20251101.pdf'],
        createdAt: new Date('2025-11-01T10:00:00'),
    },
    {
        id: 2,
        doctor: fakeDoctors[1],
        patient: fakePatients[1],
        appointment: fakeAppointments[1],
        notes: 'Routine checkup, all vitals normal.',
        prescriptions: ['Vitamin D3 supplement'],
        createdAt: new Date('2025-11-01T11:00:00')
    },
    {
        id: 3,
        doctor: fakeDoctors[2],
        patient: fakePatients[2],
        appointment: fakeAppointments[2],
        notes: 'Follow-up visit for previous condition. Showing improvement.',
        prescriptions: ['Continue current medication'],
        createdAt: new Date('2025-11-01T12:00:00')
    }
];

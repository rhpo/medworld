import { Users } from './users';
import type { Appointment } from './appointment';
import type { Consultation } from './consultation';
import type { Cabinet } from './cabinet';
import type { Doctor } from './users/doctor';
import type { Assistant } from './users/assistant';
import { Speciality } from './speciality';
import type { Calendar } from './calendar';
import type { Review } from './reviews';
import type { Rating } from './rating';
import type { Admin } from './users/admin';
import { PlanID, plans } from './plan';
import type { SuperAdmin } from './users/superadmin';
import type { Patient } from './users/patient';


export const fakeSuperAdmin: SuperAdmin =
{
    id: 0,
    firstName: 'Houria',
    lastName: 'Aichi',
    email: 'houria.aichi@medworld.dz',
    phoneNumber: '+213 555 222 000',
    type: Users.SuperAdmin,
    address: 'Algiers, Algeria',
    gender: 'female',
    dateOfBirth: new Date('1985-06-15'),
    getFullName: () => 'Houria Aichi',
    createdAt: new Date('2023-01-01'),
    createdCabinets: [],
    createCabinet: () => { },
    deleteCabinet: () => { },
    sendGlobalNotification: () => { }
}


export const fakeAdmins: Admin[] = [
    {
        id: 3,
        firstName: 'Kamel',
        lastName: 'Daoud',
        email: 'kamel.daoud@medworld.dz',
        phoneNumber: '+213 555 123 456',
        address: 'Algiers',
        gender: 'male',
        dateOfBirth: new Date('1978-04-12'),
        avatarUrl: 'https://i.pravatar.cc/50?img=3',
        type: Users.Admin,
        speciality: Speciality.FAMILY_MEDICINE,
        careerStart: new Date('2002-01-01'),
        messages: [],
        reviews: [],
        calendars: [],
        consultations: [],
        cabinets: [],
        consultationDuration: 30,
        consultationPrice: 1500,
        assistants: [],
        getFullName: () => 'Kamel Daoud',
        getYearsOfExperience: () => 20,
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.basic]
    },
    {
        id: 4,
        firstName: 'Yasmina',
        lastName: 'Khadra',
        email: 'yasmina.khadra@medworld.dz',
        phoneNumber: '+213 555 567 890',
        address: 'Algiers',
        gender: 'female',
        dateOfBirth: new Date('1984-07-08'),
        avatarUrl: 'https://i.pravatar.cc/50?img=4',
        type: Users.Admin,
        speciality: Speciality.DERMATOLOGY,
        careerStart: new Date('2008-01-01'),
        messages: [],
        reviews: [],
        calendars: [],
        consultations: [],
        cabinets: [],
        consultationDuration: 25,
        consultationPrice: 1800,
        assistants: [],
        getFullName: () => 'Yasmina Khadra',
        getYearsOfExperience: () => 17,
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.standard]
    },
    {
        id: 5,
        firstName: 'Mohammed',
        lastName: 'Dib',
        email: 'mohammed.dib@medworld.dz',
        phoneNumber: '+213 555 901 234',
        address: 'Algiers',
        gender: 'male',
        dateOfBirth: new Date('1975-03-20'),
        avatarUrl: 'https://i.pravatar.cc/50?img=5',
        type: Users.Admin,
        speciality: Speciality.CARDIOLOGY,
        careerStart: new Date('1999-01-01'),
        messages: [],
        reviews: [],
        calendars: [],
        consultations: [],
        cabinets: [],
        consultationDuration: 30,
        consultationPrice: 2500,
        assistants: [],
        getFullName: () => 'Mohammed Dib',
        getYearsOfExperience: () => 26,
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.premium]
    },
    {
        id: 6,
        firstName: 'Rachid',
        lastName: 'Boudjedra',
        email: 'rachid.boudjedra@medworld.dz',
        phoneNumber: '+213 555 345 678',
        address: 'Algiers',
        gender: 'male',
        dateOfBirth: new Date('1982-11-02'),
        avatarUrl: 'https://i.pravatar.cc/50?img=6',
        type: Users.Admin,
        speciality: Speciality.FAMILY_MEDICINE,
        careerStart: new Date('2006-01-01'),
        messages: [],
        reviews: [],
        calendars: [],
        consultations: [],
        cabinets: [],
        consultationDuration: 20,
        consultationPrice: 1400,
        assistants: [],
        getFullName: () => 'Rachid Boudjedra',
        getYearsOfExperience: () => 19,
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.basic]
    },
    {
        id: 7,
        firstName: 'Kateb',
        lastName: 'Yacine',
        email: 'kateb.yacine@medworld.dz',
        phoneNumber: '+213 555 789 012',
        address: 'Algiers',
        gender: 'male',
        dateOfBirth: new Date('1987-09-10'),
        avatarUrl: 'https://i.pravatar.cc/50?img=7',
        type: Users.Admin,
        speciality: Speciality.DERMATOLOGY,
        careerStart: new Date('2010-01-01'),
        messages: [],
        reviews: [],
        calendars: [],
        consultations: [],
        cabinets: [],
        consultationDuration: 30,
        consultationPrice: 1700,
        assistants: [],
        getFullName: () => 'Kateb Yacine',
        getYearsOfExperience: () => 15,
        createdAt: new Date('2023-02-01'),
        plan: plans[PlanID.standard]
    }
];


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
    address: 'Algiers',
    gender: 'female',
    dateOfBirth: new Date('1992-01-01'),
    doctors: [],
    cabinet: {} as Cabinet,
    appointments: [],
    planAppointment: () => { },
    cancelAppointment: () => { },
    recordPayment: () => { }
});


export const fakeAssistants: Assistant[] = [
    makeAssistant(8, 'Amira', 'Bouraoui', 'amira.bouraoui@medworld.dz', '+213 555 234 567'),
    makeAssistant(9, 'Assia', 'Djebar', 'assia.djebar@medworld.dz', '+213 555 678 901'),
    makeAssistant(10, 'Malika', 'Mokeddem', 'malika.mokeddem@medworld.dz', '+213 555 012 345'),
    makeAssistant(11, 'Leila', 'Aissaoui', 'leila.aissaoui@medworld.dz', '+213 555 456 789'),
    makeAssistant(12, 'Warda', 'Al-Jazairia', 'warda.aljazairia@medworld.dz', '+213 555 890 123')
];



export const fakeCabinets: Cabinet[] = [
    {
        id: 1,
        name: 'Centre Médical Ibn Sina',
        phone: '+213 21 23 45 67',
        createdAt: new Date('2024-01-01'),
        admin: fakeAdmins[0],
        doctors: [],
        image: 'https://picsum.photos/1200/700?random=' + Math.random(),
        assistants: [fakeAssistants[0]],
        openingHours: {
            'monday': { open: '08:00', close: '17:00' },
            'tuesday': { open: '08:00', close: '17:00' },
            'wednesday': { open: '08:00', close: '17:00' },
            'thursday': { open: '08:00', close: '17:00' },
            'friday': { open: '08:00', close: '12:00' }
        },
        accessHandicap: true,
        hasParking: true,
        hasWifi: true,
        acceptsInsurance: true,
        acceptsUrgent: true,
        location: {
            address: '123 Rue Didouche Mourad, Alger',
            latitude: 36.7525,
            longitude: 3.042
        },
        ratings: [{
            date: new Date('2024-04-15'),
            equippement: {
                disponibility: true,
                rating: 5,
            },
            cabinet: null as any,
            userExperience: {
                reception: 5,
                hygiene: 5,
                waitTime: 5,
                communication: 5,
                emplacement: 5,
                professionalism: 5,
            },
            review: "BEST CLINIC EVER! The staff were super friendly and the equipment was top-notch.",
            id: 1,
            patient: fakeAdmins[1] as any,
        }],
    },
    {
        id: 2,
        name: 'Clinique El Djazair',
        phone: '+213 21 34 56 78',
        image: 'https://picsum.photos/1200/700?random=' + Math.random(),

        createdAt: new Date('2024-02-01'),
        admin: fakeAdmins[1],
        doctors: [],
        assistants: [fakeAssistants[1]],
        openingHours: {
            'monday': { open: '09:00', close: '18:00' },
            'tuesday': { open: '09:00', close: '18:00' },
            'wednesday': { open: '09:00', close: '18:00' },
            'thursday': { open: '09:00', close: '18:00' },
            'friday': { open: '09:00', close: '13:00' }
        },
        accessHandicap: true,
        hasParking: true,
        hasWifi: true,
        acceptsInsurance: true,
        acceptsUrgent: false,
        ratings: [{
            date: new Date('2024-04-15'),
            equippement: {
                disponibility: true,
                rating: 5,
            },
            cabinet: null as any,
            userExperience: {
                reception: 1,
                hygiene: 3,
                waitTime: 5,
                communication: 4,
                emplacement: 5,
                professionalism: 4,
            },
            review: "Great experience overall, the clinic was a masterpiecee",
            id: 1,
            patient: fakeAdmins[1] as any,
        }],
        location: {
            address: '45 Boulevard Zighout Youcef, Alger',
            latitude: 36.7530,
            longitude: 3.0425
        },
    },
    {
        id: 3,
        name: 'Cabinet Médical Al Razi',
        phone: '+213 21 45 67 89',
        image: 'https://picsum.photos/1200/700?random=' + Math.random(),

        createdAt: new Date('2024-03-01'),
        admin: fakeAdmins[2],
        doctors: [],
        assistants: [fakeAssistants[2]],
        openingHours: {
            'monday': { open: '08:30', close: '17:30' },
            'tuesday': { open: '08:30', close: '17:30' },
            'wednesday': { open: '08:30', close: '17:30' },
            'thursday': { open: '08:30', close: '17:30' },
            'friday': { open: '08:30', close: '12:30' }
        },
        accessHandicap: false,
        hasParking: false,
        hasWifi: true,
        acceptsInsurance: true,
        acceptsUrgent: true,
        ratings: [
            {
                date: new Date('2024-04-15'),
                equippement: {
                    disponibility: true,
                    rating: 4,
                },
                cabinet: null as any,
                userExperience: {
                    reception: 1,
                    hygiene: 2,
                    waitTime: 3,
                    communication: 4,
                    emplacement: 1,
                    professionalism: 0,
                },
                review: "Bruh, the wait time was unbearable, but at least the equipment was decent.",
                id: 1,
                patient: fakeAdmins[1] as any,
            }
        ],
        location: {
            address: '78 Avenue Pasteur, Alger',
            latitude: 35.7540,
            longitude: 3.0430
        }
    },
    {
        id: 4,
        name: 'Centre de Santé Al Shifa',
        phone: '+213 21 56 78 90',
        image: 'https://picsum.photos/1200/700?random=' + Math.random(),

        createdAt: new Date('2024-04-01'),
        admin: fakeAdmins[3],
        doctors: [],
        assistants: [fakeAssistants[3]],
        openingHours: {
            'monday': { open: '09:00', close: '18:00' },
            'tuesday': { open: '09:00', close: '18:00' },
            'wednesday': { open: '09:00', close: '18:00' },
            'thursday': { open: '09:00', close: '18:00' },
            'friday': { open: '09:00', close: '13:00' }
        },
        accessHandicap: false,
        location: {
            address: '123 Rue Larbi Ben M\'Hidi, Alger',
            latitude: 33.7550,
            longitude: 3.0440
        },
        ratings: [{
            date: new Date('2024-04-15'),
            equippement: {
                disponibility: true,
                rating: 2,
            },
            cabinet: null as any,
            userExperience: {
                reception: 4,
                hygiene: 1,
                waitTime: 1,
                communication: 2,
                emplacement: 1,
                professionalism: 2,
            },
            id: 1,
            review: "HORRIBLE experience! The equipment was outdated and the staff were unprofessional, never goingf here again!",
            patient: fakeAdmins[1] as any,
        }],
    },
    {
        id: 5,
        name: 'Clinique Al Amal',
        phone: '+213 21 67 89 01',
        image: 'https://picsum.photos/1200/700?random=' + Math.random(),

        createdAt: new Date('2024-05-01'),
        admin: fakeAdmins[4],
        doctors: [],
        assistants: [fakeAssistants[4]],
        openingHours: {
            'monday': { open: '08:00', close: '17:00' },
            'tuesday': { open: '08:00', close: '17:00' },
            'wednesday': { open: '08:00', close: '17:00' },
            'thursday': { open: '08:00', close: '17:00' },
            'friday': { open: '08:00', close: '12:00' }
        },
        accessHandicap: false,
        location: {
            address: '56 Boulevard Colonel Amirouche, Alger',
            latitude: 36.7560,
            longitude: 3.0450
        },
        ratings: [{
            date: new Date('2024-04-15'),
            equippement: {
                disponibility: true,
                rating: 4,
            },
            cabinet: null as any,
            userExperience: {
                reception: 1,
                hygiene: 2,
                waitTime: 0,
                communication: 2,
                emplacement: 0,
                professionalism: 1,
            },
            id: 1,
            review: "The equipment was good, but the professionalism could be improved.",
            patient: fakeAdmins[1] as any,
        }],
    }
];


const makeDoctor = (
    id: number,
    firstName: string,
    lastName: string,
    email: string,
    phoneNumber: string,
    speciality: Speciality,
    avatarUrl?: string,
): Doctor => ({
    id,
    firstName,
    lastName,
    email,
    phoneNumber,
    avatarUrl: avatarUrl || 'https://i.pravatar.cc/50',
    type: Users.Doctor,
    speciality,
    address: 'Algiers',
    gender: 'male',
    dateOfBirth: new Date('1980-01-01'),
    careerStart: new Date('2005-01-01'),
    messages: [],
    reviews: [],
    calendars: [],
    consultations: [],
    cabinets: fakeCabinets,
    consultationDuration: 20,
    consultationPrice: 2000,
    assistants: [],
    getFullName: () => `${firstName} ${lastName}`,
    getYearsOfExperience: () => {
        const diff = new Date().getTime() - new Date('2005-01-01').getTime();
        return Math.floor(diff / (1000 * 60 * 60 * 24 * 365));
    }
});


export const fakeDoctors: Doctor[] = [
    makeDoctor(21, 'Rayan', 'El-miloudi', 'rayan.miloudi@medworld.dz', '+213 555 123 456', Speciality.CARDIOLOGY, 'https://i.pravatar.cc/500'),
    makeDoctor(22, 'Manel', 'Manel', 'manel.manel@medworld.dz', '+213 555 234 567', Speciality.FAMILY_MEDICINE, 'https://i.pravatar.cc/500'),
    makeDoctor(23, 'Ramy', 'Hadid', 'ramy.ramyyy@medworld.dz', '+213 555 345 678', Speciality.PEDIATRICS, 'https://i.pravatar.cc/500'),
    makeDoctor(24, 'Sarah', 'Zerrouki', 'sarah.zerrouki@medworld.dz', '+213 555 456 789', Speciality.DERMATOLOGY)
];

for (const cab of fakeCabinets) {

    (cab.doctors) = fakeDoctors;
}

fakeDoctors[0].assistants = [fakeAssistants[0], fakeAssistants[1]];
fakeAssistants[0].doctors = [fakeDoctors[0]];
fakeAssistants[1].doctors = [fakeDoctors[0]];


fakeDoctors[1].assistants = [fakeAssistants[2]];
fakeDoctors[2].assistants = [fakeAssistants[2]];
fakeAssistants[2].doctors = [fakeDoctors[1], fakeDoctors[2]];


fakeDoctors[3].assistants = [fakeAssistants[3]];
fakeAssistants[3].doctors = [fakeDoctors[3]];


fakeAssistants[4].doctors = [];




for (const cab of fakeCabinets) {
    for (const asst of cab.assistants) {

        (asst as Assistant).cabinet = cab;
    }
}


let _calId = 1;
const fakeCalendars: Calendar[] = [];
for (const doc of fakeDoctors) {

    const workCabinets = fakeCabinets.filter((c) => c.doctors.some((d) => d.id === doc.id));
    const assigned = workCabinets.length ? workCabinets : [fakeCabinets[0]];
    for (const cab of assigned) {
        const cal: Calendar = {
            id: _calId++,
            cabinet: cab,
            doctor: doc,

            availability: doc.id === fakeDoctors[0].id ? [
                {
                    date: '2023-10-01',
                    slots: ['09:00', '10:00', '11:00']
                },
                {
                    date: '2023-10-02',
                    slots: ['13:00', '14:00']
                }
            ] : doc.id === fakeDoctors[1].id ? [
                {
                    date: '2023-10-01',
                    slots: ['09:00', '10:00', '11:00']
                },
                {
                    date: '2023-10-02',
                    slots: ['13:00', '14:00']
                }
            ] : doc.id === fakeDoctors[2].id ? [
                {
                    date: '2023-10-01',
                    slots: ['09:00', '10:00', '11:00']
                },
                {
                    date: '2023-10-02',
                    slots: ['13:00', '14:00']
                }
            ] : [
                {
                    date: '2023-10-01',
                    slots: ['09:00', '10:00', '11:00']
                },
                {
                    date: '2023-10-02',
                    slots: ['13:00', '14:00']
                }
            ],
        };
        fakeCalendars.push(cal);

        if (!doc.calendars) doc.calendars = [];
        doc.calendars.push(cal);
    }
}



export const fakePatients: Patient[] = [
    {
        id: 25,
        firstName: 'Riad',
        lastName: 'Mahrez',
        email: 'riad.mahrez@email.dz',
        phoneNumber: '+213 555 567 890',
        type: Users.Patient,
        avatarUrl: 'https://i.pravatar.cc/50',
        address: 'Algiers',
        gender: 'male',
        dateOfBirth: new Date('1990-05-05'),
        emergencyContact: '+213 555 000 111',
        bloodType: 'A+',
        appointments: [],
        consultations: [],
        reviews: [],
        medicalHistory: [
            'Hypertension - Diagnosed 2020',
            'Type 2 Diabetes - Controlled with medication',
            'Right knee arthroscopy (2018)',
            'Family history of cardiovascular disease'
        ],
        allergies: ['Penicillin', 'Sulfa drugs'],
        getFullName: () => 'Riad Mahrez',
        addReview: (review: Review) => { }
    },
    {
        id: 26,
        firstName: 'Amira',
        lastName: 'Brahimi',
        email: 'amira.brahimi@email.dz',
        phoneNumber: '+213 555 678 901',
        type: Users.Patient,
        avatarUrl: 'https://i.pravatar.cc/50',
        address: 'Algiers',
        gender: 'female',
        dateOfBirth: new Date('1992-07-12'),
        emergencyContact: '+213 555 111 222',
        bloodType: 'O+',
        appointments: [],
        consultations: [],
        reviews: [],
        medicalHistory: [
            'Asthma - Mild, controlled with inhalers',
            'Migraines - Chronic, managed with medication',
            'Appendectomy (2015)',
            'Regular wellness checkups'
        ],
        allergies: ['Dust mites', 'Pollen', 'Aspirin'],
        getFullName: () => 'Amira Brahimi',
        addReview: (review: Review) => { }
    },
    {
        id: 27,
        firstName: 'Karim',
        lastName: 'Ziani',
        email: 'karim.ziani@email.dz',
        phoneNumber: '+213 555 789 012',
        type: Users.Patient,
        avatarUrl: 'https://i.pravatar.cc/50',
        address: 'Algiers',
        gender: 'male',
        dateOfBirth: new Date('1988-01-20'),
        emergencyContact: '+213 555 222 333',
        bloodType: 'B+',
        appointments: [],
        consultations: [],
        reviews: [],
        medicalHistory: [
            'GERD - Managed with diet and medication',
            'Mild anxiety - Under treatment',
            'Sports injury - Left shoulder (2021)',
            'Annual physical examinations'
        ],
        allergies: ['Shellfish', 'Ibuprofen'],
        getFullName: () => 'Karim Ziani',
        addReview: (review: Review) => { }
    },
    {
        id: 28,
        firstName: 'Louisa',
        lastName: 'Hanoune',
        email: 'louisa.hanoune@email.dz',
        phoneNumber: '+213 555 890 123',
        type: Users.Patient,
        avatarUrl: 'https://i.pravatar.cc/50',
        address: 'Algiers',
        gender: 'female',
        dateOfBirth: new Date('1995-09-02'),
        emergencyContact: '+213 555 333 444',
        bloodType: 'AB+',
        appointments: [],
        consultations: [],
        reviews: [],
        medicalHistory: [
            'Hypothyroidism - Managed with Levothyroxine',
            'Iron deficiency anemia - Under treatment',
            'Cesarean section (2019)',
            'Regular gynecological check-ups'
        ],
        allergies: ['Latex', 'Codeine'],
        getFullName: () => 'Louisa Hanoune',
        addReview: (review: Review) => { }
    }
];


export const fakeAppointments: Appointment[] = [

    {
        id: 1,
        date: new Date('2024-01-15T09:00:00'),
        status: 'completed',
        doctor: fakeDoctors[0],
        patient: fakePatients[0],
        cabinet: fakeCabinets[0],
        createdAt: new Date('2024-01-01'),
        updatedAt: new Date('2024-01-02')
    },
    {
        id: 2,
        date: new Date('2024-04-15T09:30:00'),
        status: 'completed',
        doctor: fakeDoctors[0],
        patient: fakePatients[0],
        cabinet: fakeCabinets[0],
        createdAt: new Date('2024-03-01')
    },
    {
        id: 3,
        date: new Date('2024-07-15T10:00:00'),
        status: 'confirmed',
        doctor: fakeDoctors[0],
        patient: fakePatients[0],
        cabinet: fakeCabinets[0],
        createdAt: new Date('2024-06-01')
    },

    {
        id: 4,
        date: new Date('2024-02-10T14:00:00'),
        status: 'completed',
        doctor: fakeDoctors[1],
        patient: fakePatients[1],
        cabinet: fakeCabinets[0],
        createdAt: new Date('2024-01-15')
    },
    {
        id: 5,
        date: new Date('2024-05-20T15:30:00'),
        status: 'confirmed',
        doctor: fakeDoctors[1],
        patient: fakePatients[1],
        cabinet: fakeCabinets[1],
        createdAt: new Date('2024-04-01')
    },

    {
        id: 6,
        date: new Date('2024-03-05T11:00:00'),
        status: 'completed',
        doctor: fakeDoctors[1],
        patient: fakePatients[2],
        cabinet: fakeCabinets[1],
        createdAt: new Date('2024-02-01')
    },
    {
        id: 7,
        date: new Date('2024-06-15T10:30:00'),
        status: 'confirmed',
        doctor: fakeDoctors[1],
        patient: fakePatients[2],
        cabinet: fakeCabinets[1],
        createdAt: new Date('2024-05-01')
    },

    {
        id: 8,
        date: new Date('2024-01-20T09:00:00'),
        status: 'completed',
        doctor: fakeDoctors[1],
        patient: fakePatients[3],
        cabinet: fakeCabinets[2],
        createdAt: new Date('2024-01-01')
    },
    {
        id: 9,
        date: new Date('2024-04-25T14:30:00'),
        status: 'confirmed',
        doctor: fakeDoctors[1],
        patient: fakePatients[3],
        cabinet: fakeCabinets[2],
        createdAt: new Date('2024-03-15')
    },
    {
        id: 10,
        date: new Date('2024-07-20T10:00:00'),
        status: 'pending',
        doctor: fakeDoctors[1],
        patient: fakePatients[3],
        cabinet: fakeCabinets[2],
        createdAt: new Date('2024-06-01')
    }
];


export const fakeConsultations: Consultation[] = [

    {
        id: 1,
        doctor: fakeDoctors[0],
        patient: fakePatients[0],
        appointment: fakeAppointments[0],
        notes: 'Initial consultation for hypertension and diabetes management. BP: 145/90. Blood sugar: 7.8 mmol/L (fasting). Patient reports occasional dizziness and fatigue. Physical examination normal. ECG shows mild LVH.',
        prescriptions: [
            'Lisinopril 10mg - 1x daily',
            'Metformin 500mg - 2x daily with meals',
            'Daily blood pressure monitoring',
            'Blood sugar monitoring 3x daily'
        ],
        attachments: ['bp_chart_20240115.pdf', 'ecg_20240115.pdf', 'lab_results_20240115.pdf'],
        createdAt: new Date('2024-01-15T10:00:00'),
    },
    {
        id: 2,
        doctor: fakeDoctors[0],
        patient: fakePatients[0],
        appointment: fakeAppointments[1],
        notes: 'Follow-up visit. BP improved to 135/85. Blood sugar levels better controlled (avg 6.5 mmol/L). Patient adhering to medication schedule. Reports improved energy levels. Weight loss of 3kg since last visit.',
        prescriptions: [
            'Continue Lisinopril 10mg daily',
            'Continue Metformin 500mg twice daily',
            'Added daily 30-minute walking routine'
        ],
        attachments: ['progress_report_20240415.pdf'],
        createdAt: new Date('2024-04-15T10:30:00')
    },

    {
        id: 3,
        doctor: fakeDoctors[1],
        patient: fakePatients[1],
        appointment: fakeAppointments[3],
        notes: 'Patient reports increased frequency of migraines (2-3x/week) and seasonal asthma symptoms. Lung examination clear. Peak flow 400 L/min. Migraine diary shows correlation with stress and poor sleep.',
        prescriptions: [
            'Salbutamol inhaler - as needed',
            'Fluticasone/Salmeterol inhaler - 2 puffs twice daily',
            'Sumatriptan 50mg - for acute migraines',
            'Propranolol 40mg - daily for migraine prevention'
        ],
        attachments: ['migraine_diary_20240210.pdf', 'pef_chart_20240210.pdf'],
        createdAt: new Date('2024-02-10T15:00:00')
    },

    {
        id: 4,
        doctor: fakeDoctors[1],
        patient: fakePatients[2],
        appointment: fakeAppointments[5],
        notes: 'Follow-up for GERD and anxiety. Reports improvement in reflux symptoms with medication and dietary changes. Anxiety symptoms stable but still present during work hours. Sleep improved with relaxation techniques.',
        prescriptions: [
            'Omeprazole 20mg - daily before breakfast',
            'Escitalopram 10mg - daily in the morning',
            'Continue stress management techniques',
            'Dietary restrictions maintained'
        ],
        attachments: ['anxiety_assessment_20240305.pdf', 'diet_log_20240305.pdf'],
        createdAt: new Date('2024-03-05T12:00:00')
    },

    {
        id: 5,
        doctor: fakeDoctors[1],
        patient: fakePatients[3],
        appointment: fakeAppointments[7],
        notes: 'Regular follow-up for hypothyroidism and anemia. TSH levels normalized (2.8 mU/L). Hemoglobin improved to 11.8 g/dL. Patient reports better energy levels and reduced cold sensitivity. Weight stable.',
        prescriptions: [
            'Levothyroxine 75mcg - daily on empty stomach',
            'Iron sulfate 325mg - daily with vitamin C',
            'Continue regular exercise routine',
            'Follow-up blood work in 3 months'
        ],
        attachments: ['thyroid_panel_20240120.pdf', 'cbc_results_20240120.pdf'],
        createdAt: new Date('2024-01-20T10:00:00')
    }
];


let _ratingId = 1;
for (const cab of fakeCabinets) {
    const rating: Rating = {
        id: _ratingId++,
        patient: fakePatients[Math.floor(Math.random() * fakePatients.length)],
        cabinet: cab,
        date: new Date(),
        equippement: { disponibility: true, rating: 4 },
        userExperience: {
            reception: 4,
            hygiene: 4,
            waitTime: 3,
            communication: 4,
            professionalism: 5,
            emplacement: 4,
        },
    };
    cab.ratings = cab.ratings || [];
    cab.ratings.push(rating);
}



for (const cons of fakeConsultations) {

    const patient = fakePatients.find((x) => x.id === cons.patient.id);
    if (patient && patient.consultations) {
        patient.consultations.push(cons);
    }


    const doctor = fakeDoctors.find((x) => x.id === cons.doctor.id);
    if (doctor && doctor.consultations) {
        doctor.consultations.push(cons);
    }
}


for (const doc of fakeDoctors) {
    if (doc.reviews) {
        for (const review of doc.reviews) {

            const patient = fakePatients.find((x) => x.id === review.from.id);
            if (patient && patient.reviews) {
                patient.reviews.push(review);
            }
        }
    }
}


for (const appt of fakeAppointments) {

    const patient = fakePatients.find((x) => x.id === appt.patient.id);
    if (patient && patient.appointments) {
        patient.appointments.push(appt);
    }


    for (const assistant of fakeAssistants) {
        if (assistant.appointments && assistant.doctors.some(doc => doc.id === appt.doctor.id)) {
            assistant.appointments.push(appt);
        }
    }
}


for (const doc of fakeDoctors) {
    doc.reviews = [];


    if (doc.id === fakeDoctors[0].id) {
        doc.reviews.push({
            from: fakePatients[0],
            to: doc,
            rating: 5,
            message: 'Dr. El-miloudi is an exceptional cardiologist. He took the time to explain my hypertension condition thoroughly and created a comprehensive treatment plan. His follow-up care has been excellent. Very pleased with the improvement in my blood pressure control.'
        });
        doc.reviews.push({
            from: fakePatients[2],
            to: doc,
            rating: 4,
            message: 'Very knowledgeable and professional doctor. The wait times can be a bit long, but the quality of care is worth it. Explains everything clearly and answers all questions patiently.'
        });
    } else if (doc.id === fakeDoctors[1].id) {
        doc.reviews.push({
            from: fakePatients[1],
            to: doc,
            rating: 5,
            message: 'Dr. Belaiouer has been managing my asthma and migraines wonderfully. She is very attentive and her holistic approach to treatment has made a significant difference. Her staff is also very helpful and friendly.'
        });
        doc.reviews.push({
            from: fakePatients[3],
            to: doc,
            rating: 4,
            message: 'Great family doctor who really listens to her patients. Has been very helpful with managing my thyroid condition. Always follows up and ensures prescriptions are working well.'
        });
    } else if (doc.id === fakeDoctors[2].id) {
        doc.reviews.push({
            from: fakePatients[1],
            to: doc,
            rating: 5,
            message: 'Excellent pediatrician! Dr. Nekkmouche is amazing with children and very thorough in his examinations. Makes both parents and kids feel comfortable. Highly recommended!'
        });
        doc.reviews.push({
            from: fakePatients[0],
            to: doc,
            rating: 4,
            message: 'Very good with kids, explains everything clearly to parents. The only downside is sometimes the wait can be long, but the care is excellent.'
        });
    } else if (doc.id === fakeDoctors[3].id) {
        doc.reviews.push({
            from: fakePatients[2],
            to: doc,
            rating: 5,
            message: 'Dr. Zerrouki is a skilled dermatologist. Very thorough in her examinations and explains treatment options clearly. Modern approach to skincare and treatment. The clinic is well-equipped.'
        });
        doc.reviews.push({
            from: fakePatients[3],
            to: doc,
            rating: 3,
            message: 'Competent doctor but the wait times can be very long. Would appreciate better scheduling system. Treatment was effective though.'
        });
    }


    doc.messages = [];


    const doctorConsultations = fakeConsultations.filter(cons => cons.doctor.id === doc.id);
    const latestConsultation = doctorConsultations.length > 0 ? doctorConsultations[0] : undefined;

    if (doc.id === fakeDoctors[0].id) {

        if (latestConsultation) {

            doc.messages.push({
                sender: doc as Doctor,
                cabinet: fakeCabinets[0],
                receiver: fakeDoctors[1] as Doctor,
                date: new Date('2024-01-20T09:30:00'),
                content: {
                    message: 'Regarding patient Riad Mahrez: Blood pressure now stabilized with Lisinopril. Please monitor during regular check-ups. Will review again in 3 months.',
                    attachment: latestConsultation
                },
                status: 'seen'
            });


            doc.messages.push({
                sender: doc as Doctor,
                cabinet: fakeCabinets[0],
                receiver: fakeDoctors[1] as Doctor,
                date: new Date('2024-01-21T10:15:00'),
                content: {
                    message: 'Thank you for the update on Mr. Mahrez. Will monitor BP during his next visit. Also noting good diabetes control with current regimen.',
                    attachment: latestConsultation
                },
                status: 'seen'
            });


            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[0],
                receiver: fakeDoctors[2],
                date: new Date('2024-02-05T11:00:00'),
                content: {
                    message: 'Organizing a case conference next week to discuss recent cardiovascular findings in young patients. Would value your pediatric perspective. Available Tuesday or Thursday afternoon.',
                    attachment: latestConsultation
                },
                status: 'seen'
            });


            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[0],
                receiver: fakeDoctors[3],
                date: new Date('2024-02-10T15:45:00'),
                content: {
                    message: 'Patient with unusual skin manifestations possibly related to new BP medication. Could you provide a quick consultation? Sending photos and current medication list.',
                    attachment: latestConsultation
                },
                status: 'unseen'
            });


            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[0],
                receiver: fakeDoctors[1],
                date: new Date('2024-02-15T08:30:00'),
                content: {
                    message: 'Initiating a study on hypertension management in diabetic patients. Would you be interested in collaborating? Have 20 potential candidates from our shared patient pool.',
                    attachment: latestConsultation
                },
                status: 'unseen'
            });
        }
    }

    if (doc.id === fakeDoctors[1].id) {
        const familyDocConsultation = fakeConsultations.find(cons => cons.doctor.id === doc.id);
        if (familyDocConsultation) {

            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[1],
                receiver: fakeDoctors[2],
                date: new Date('2024-02-15T14:20:00'),
                content: {
                    message: "Would appreciate your input on recent cases of respiratory infections in children. Can we schedule a brief discussion?",
                    attachment: familyDocConsultation
                },
                status: 'seen'
            });


            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[0],
                receiver: fakeDoctors[0],
                date: new Date('2024-02-16T09:00:00'),
                content: {
                    message: "Monthly staff meeting scheduled for next Tuesday at 8:30 AM. Please review updated clinic protocols before the meeting.",
                    attachment: familyDocConsultation
                },
                status: 'unseen'
            });
        }
    }

    if (doc.id === fakeDoctors[2].id) {
        const pediatricConsultation = fakeConsultations.find(cons => cons.doctor.id === doc.id);
        if (pediatricConsultation) {

            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[1],
                receiver: fakeDoctors[1],
                date: new Date('2024-02-15T16:45:00'),
                content: {
                    message: "Available for discussion tomorrow at 13:00. Have noticed similar trend in my practice. Will share my observations and treatment protocols.",
                    attachment: pediatricConsultation
                },
                status: 'seen'
            });


            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[2],
                receiver: fakeDoctors[1],
                date: new Date('2024-02-20T11:30:00'),
                content: {
                    message: "Need to update our pediatric examination protocols. Will share findings at next staff meeting.",
                    attachment: pediatricConsultation
                },
                status: 'unseen'
            });
        }
    }

    if (doc.id === fakeDoctors[3].id) {
        const dermatologyConsultation = fakeConsultations.find(cons => cons.doctor.id === doc.id);
        if (dermatologyConsultation) {

            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[2],
                receiver: fakeDoctors[1],
                date: new Date('2024-03-01T10:00:00'),
                content: {
                    message: "Interesting case of systemic lupus with skin manifestations. Would appreciate your input on managing the general health aspects.",
                    attachment: dermatologyConsultation
                },
                status: 'seen'
            });


            doc.messages.push({
                sender: doc,
                cabinet: fakeCabinets[2],
                receiver: fakeDoctors[2],
                date: new Date('2024-03-02T08:30:00'),
                content: {
                    message: "Several appointment requests for next week. Should we extend clinic hours on Wednesday?",
                    attachment: dermatologyConsultation
                },
                status: 'unseen'
            });
        }
    }
}

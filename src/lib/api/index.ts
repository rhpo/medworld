import { Users, type User } from '../types/users';
import type { Cabinet } from '../types/cabinet';
import type { Appointment } from '../types/appointment';
import type { Consultation } from '../types/consultation';
import type { Doctor } from '../types/users/doctor';
import {
    fakeCabinets,
    fakeDoctors,
    fakePatients,
    fakeAppointments,
    fakeConsultations,
    fakeAssistants,
    fakeAdmins,
    fakeSuperAdmin
} from '../types/fakedata';
import type { Patient } from '$lib/types/users/patient';
import type { Assistant } from '$lib/types/users/assistant';

// Simulate network delay for realistic API behavior
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

// Random delay between 200ms and 800ms
const randomDelay = () => delay(Math.random() * 600 + 200);

export const AllAPI = {
    async listAllAppointments(): Promise<Appointment[]> {
        await randomDelay();
        return [...fakeAppointments];
    },

    async listAllDoctors(): Promise<Doctor[]> {
        await randomDelay();
        return [...fakeDoctors];
    },

    async listAllPatients(): Promise<Patient[]> {
        await randomDelay();
        return [...fakePatients] as Patient[];
    },

    async listAllAssistants(): Promise<Assistant[]> {
        await randomDelay();
        return [...fakeAssistants] as Assistant[];
    },

    async listAllCabinets(): Promise<Cabinet[]> {
        await randomDelay();
        return [...fakeCabinets];
    },

    async listAllUsers(): Promise<User<any>[]> {
        await randomDelay();
        return [
            fakeSuperAdmin,
            ...fakeAdmins,
            ...fakeDoctors,
            ...fakePatients,
            ...fakeAssistants,
        ];
    },
};

// Cabinet API
export const CabinetAPI = {
    async list(): Promise<Cabinet[]> {
        await randomDelay();
        return [...fakeCabinets];
    },

    async getById(id: number): Promise<Cabinet | null> {
        await randomDelay();
        const cabinet = fakeCabinets.find(c => c.id === id);
        return cabinet || null;
    },

    async getDoctors(cabinetId: number): Promise<Array<User<Users.Doctor>>> {
        await randomDelay();
        const cabinet = fakeCabinets.find(c => c.id === cabinetId);
        return cabinet?.doctors || [];
    },

    async getAppointments(cabinetId: number): Promise<Appointment[]> {
        await randomDelay();
        return fakeAppointments.filter(app => app.cabinet.id === cabinetId);
    },

    async getAssistants(cabinetId: number): Promise<Array<User<Users.Assistant>>> {
        await randomDelay();
        const cabinet = fakeCabinets.find(c => c.id === cabinetId);
        return cabinet?.assistants || [];
    },
};

// Doctor API
export const UserAPI = {
    async UpdateProfile(user: User<any>, newUser: User<any>) {
        if (user.type === Users.Doctor) {
            // fakeDoctors = fakeDoctors.map((doctor: Doctor) => {
            //     if (doctor.id === user.id) {
            //         return user as Doctor;
            //     }
            // })
        }
    },

    async getById(userID: User<any>['id']) {
        return (await AllAPI.listAllUsers()).find(user => user.id === userID) || null;
    }
}

// Doctor API
export const DoctorAPI = {
    async list(): Promise<Doctor[]> {
        await randomDelay();
        return [...fakeDoctors];
    },

    async getById(id: number): Promise<Doctor | null> {
        await randomDelay();
        const doctor = fakeDoctors.find(d => d.id === id);
        return doctor || null;
    },

    async getAppointments(doctorId: number): Promise<Appointment[]> {
        await randomDelay();
        return fakeAppointments.filter(app => app.doctor.id === doctorId);
    },

    async getConsultations(doctorId: number): Promise<Consultation[]> {
        await randomDelay();
        return fakeConsultations.filter(cons => cons.doctor.id === doctorId);
    },


    async getPatients(doctorId: User<any>['id'], cabinetId?: Cabinet['id']) {
        if (typeof cabinetId === 'undefined') {
            return fakePatients.filter(p =>
                fakeAppointments.some(app =>
                    app.doctor.id === doctorId && app.patient.id === p.id
                )
            );
        } else {
            return fakePatients.filter(p =>
                fakeAppointments.some(app =>
                    app.doctor.id === doctorId &&
                    app.cabinet.id === cabinetId &&
                    app.patient.id === p.id
                )
            );
        }
    },
};

// Patient API
export const PatientAPI = {
    async list(): Promise<Array<User<Users.Patient>>> {
        await randomDelay();
        return [...fakePatients];
    },

    async getById(id: number): Promise<User<Users.Patient> | null> {
        await randomDelay();
        const patient = fakePatients.find(p => p.id === id);
        return patient || null;
    },

    async getAppointments(patientId: number): Promise<Appointment[]> {
        await randomDelay();
        return fakeAppointments.filter(app => app.patient.id === patientId);
    },

    async getConsultations(patientId: number): Promise<Consultation[]> {
        await randomDelay();
        return fakeConsultations.filter(cons => cons.patient.id === patientId);
    }
};

// Appointment API
export const AppointmentAPI = {
    async list(): Promise<Appointment[]> {
        await randomDelay();
        return [...fakeAppointments];
    },

    async getById(id: number): Promise<Appointment | null> {
        await randomDelay();
        const appointment = fakeAppointments.find(a => a.id === id);
        return appointment || null;
    },

    async create(appointment: Omit<Appointment, 'id' | 'createdAt' | 'updatedAt'>): Promise<Appointment> {
        await randomDelay();
        const newAppointment: Appointment = {
            ...appointment,
            id: Math.max(...fakeAppointments.map(a => a.id)) + 1,
            createdAt: new Date(),
            status: 'pending'
        };
        fakeAppointments.push(newAppointment);
        return newAppointment;
    },

    async updateStatus(id: number, status: Appointment['status']): Promise<Appointment> {
        await randomDelay();
        const appointment = fakeAppointments.find(a => a.id === id);
        if (!appointment) {
            throw new Error('Appointment not found');
        }
        appointment.status = status;
        appointment.updatedAt = new Date();
        return appointment;
    },

    async delete(id: number): Promise<void> {
        await randomDelay();
        const index = fakeAppointments.findIndex(a => a.id === id);
        if (index === -1) {
            throw new Error('Appointment not found');
        }
        fakeAppointments.splice(index, 1);
    }
};

// Consultation API
export const ConsultationAPI = {
    async list(): Promise<Consultation[]> {
        await randomDelay();
        return [...fakeConsultations];
    },

    async getById(id: number): Promise<Consultation | null> {
        await randomDelay();
        const consultation = fakeConsultations.find(c => c.id === id);
        return consultation || null;
    },

    async create(consultation: Omit<Consultation, 'id' | 'createdAt' | 'updatedAt'>): Promise<Consultation> {
        await randomDelay();
        const newConsultation: Consultation = {
            ...consultation,
            id: Math.max(...fakeConsultations.map(c => c.id)) + 1,
            createdAt: new Date(),
            notes: consultation.notes || '',
            prescriptions: consultation.prescriptions || []
        };
        fakeConsultations.push(newConsultation);
        return newConsultation;
    },

    async update(id: number, updates: Partial<Consultation>): Promise<Consultation> {
        await randomDelay();
        const consultation = fakeConsultations.find(c => c.id === id);
        if (!consultation) {
            throw new Error('Consultation not found');
        }
        Object.assign(consultation, updates);
        consultation.updatedAt = new Date();
        return consultation;
    }
};

// Global search functionality
export const SearchAPI = {
    async searchDoctors(query: string): Promise<Doctor[]> {
        await randomDelay();
        const q = query.toLowerCase();
        return fakeDoctors.filter(d =>
            d.firstName.toLowerCase().includes(q) ||
            d.lastName.toLowerCase().includes(q)
        );
    },

    async searchPatients(query: string): Promise<Array<User<Users.Patient>>> {
        await randomDelay();
        const q = query.toLowerCase();
        return fakePatients.filter(p =>
            p.firstName.toLowerCase().includes(q) ||
            p.lastName.toLowerCase().includes(q)
        );
    },

    async searchCabinets(query: string): Promise<Cabinet[]> {
        await randomDelay();
        const q = query.toLowerCase();
        return fakeCabinets.filter(c =>
            c.name.toLowerCase().includes(q) ||
            c.address.toLowerCase().includes(q)
        );
    }
};

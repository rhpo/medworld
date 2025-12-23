import { apiClient } from './client';
import { Users, type IUser, type User } from '../types/users';
import type { Cabinet } from '../types/cabinet';
import type { Appointment } from '../types/appointment';
import type { Consultation } from '../types/consultation';
import type { Doctor } from '../types/users/doctor';
import type { Patient } from '../types/users/patient';
import type { Assistant } from '../types/users/assistant';

// ==================== AUTH API ====================
export const AuthAPI = {
    async login(email: string, password: string): Promise<{ user: User<any>; token: string }> {
        return apiClient.post('/auth/login', { email, password });
    },

    async register(userData: {
        first_name: string;
        last_name: string;
        email: string;
        password: string;
        phone_number?: string;
        address?: string;
        gender?: 'male' | 'female';
        date_of_birth?: string;
        type: 'patient' | 'doctor' | 'admin' | 'assistant' | 'superadmin';
    }): Promise<{ user: User<any>; token: string }> {
        return apiClient.post('/auth/register', userData);
    },

    async logout(): Promise<void> {
        return apiClient.post('/auth/logout');
    },

    async me(): Promise<{ user: User<any> }> {
        return apiClient.get('/auth/me');
    }
};

// ==================== ALL API ====================
export const AllAPI = {
    async listAllAppointments(): Promise<Appointment[]> {
        const response = await apiClient.get<{ data: Appointment[] }>('/all/appointments');
        return response.data || [];
    },

    async listAllDoctors(): Promise<Doctor[]> {
        const response = await apiClient.get<{ data: Doctor[] }>('/all/doctors');
        return response.data || [];
    },

    async listAllPatients(): Promise<Patient[]> {
        const response = await apiClient.get<{ data: Patient[] }>('/all/patients');
        return response.data || [];
    },

    async listAllAssistants(): Promise<Assistant[]> {
        const response = await apiClient.get<{ data: Assistant[] }>('/all/assistants');
        return response.data || [];
    },

    async listAllCabinets(): Promise<Cabinet[]> {
        const response = await apiClient.get<Cabinet[]>('/all/cabinets');
        return response || [];
    },

    async listAllUsers(): Promise<User<any>[]> {
        const response = await apiClient.get<{ data: User<any>[] }>('/all/users');
        return response.data || [];
    },

    async listAllConsultations(): Promise<Consultation[]> {
        const response = await apiClient.get<{ data: Consultation[] }>('/all/consultations');
        return response.data || [];
    }
};

// ==================== CABINET API ====================
export const CabinetAPI = {
    async list(): Promise<Cabinet[]> {
        const response = await apiClient.get<{ data: Cabinet[] }>('/cabinets');
        return response.data || [];
    },

    async getById(id: number): Promise<Cabinet | null> {
        try {
            return await apiClient.get(`/cabinets/${id}`);
        } catch (error) {
            return null;
        }
    },

    async getDoctors(cabinetId: number): Promise<Doctor[]> {
        const response = await apiClient.get<{ data: Doctor[] }>(`/cabinets/${cabinetId}/doctors`);
        return response.data || [];
    },

    async getAppointments(cabinetId: number): Promise<Appointment[]> {
        const response = await apiClient.get<{ data: Appointment[] }>(`/cabinets/${cabinetId}/appointments`);
        return response.data || [];
    },

    async getAssistants(cabinetId: number): Promise<Assistant[]> {
        const response = await apiClient.get<{ data: Assistant[] }>(`/cabinets/${cabinetId}/assistants`);
        return response.data || [];
    },

    async create(cabinet: Partial<Cabinet>): Promise<Cabinet> {
        return apiClient.post('/cabinets', cabinet);
    }
};

// ==================== USER API ====================
export const UserAPI = {
    async UpdateProfile(user: User<any>, newUser: User<any>) {
        return apiClient.put(`/users/${user.id}`, newUser);
    },

    async getById(userID: User<any>['id']): Promise<IUser | null> {
        try {
            return await apiClient.get(`/users/${userID}`);
        } catch (error) {
            return null;
        }
    }
};

// ==================== DOCTOR API ====================
export const DoctorAPI = {
    async list(): Promise<Doctor[]> {
        const response = await apiClient.get<{ data: Doctor[] }>('/doctors');
        return response.data || [];
    },

    async getById(id: number): Promise<Doctor | null> {
        try {
            return await apiClient.get(`/doctors/${id}`);
        } catch (error) {
            return null;
        }
    },

    async getAppointments(doctorId: number): Promise<Appointment[]> {
        const response = await apiClient.get<{ data: Appointment[] }>(`/doctors/${doctorId}/appointments`);
        return response.data || [];
    },

    async getConsultations(doctorId: number): Promise<Consultation[]> {
        const response = await apiClient.get<{ data: Consultation[] }>(`/doctors/${doctorId}/consultations`);
        return response.data || [];
    },

    async getPatients(doctorId: User<any>['id'], cabinetId?: Cabinet['id']): Promise<Patient[]> {
        let endpoint = `/doctors/${doctorId}/patients`;
        if (cabinetId !== undefined) {
            endpoint += `?cabinet_id=${cabinetId}`;
        }
        const response = await apiClient.get<{ data: Patient[] }>(endpoint);
        return response.data || [];
    },

    async search(filters: {
        speciality?: string;
        cabinetId?: number;
        priceMax?: number;
        priceMin?: number;
        available?: boolean;
    }): Promise<Doctor[]> {
        const params = new URLSearchParams();
        if (filters.speciality) params.append('speciality', filters.speciality);
        if (filters.cabinetId) params.append('cabinet_id', filters.cabinetId.toString());
        if (filters.priceMax !== undefined) params.append('price_max', filters.priceMax.toString());
        if (filters.priceMin !== undefined) params.append('price_min', filters.priceMin.toString());
        if (filters.available !== undefined) params.append('available', filters.available.toString());

        const endpoint = `/doctors/search/filter?${params.toString()}`;
        const response = await apiClient.get<{ data: Doctor[] }>(endpoint);
        return response.data || [];
    }
};

// ==================== PATIENT API ====================
export const PatientAPI = {
    async list(): Promise<Patient[]> {
        const response = await apiClient.get<{ data: Patient[] }>('/patients');
        return response.data || [];
    },

    async getById(id: number): Promise<Patient | null> {
        try {
            return await apiClient.get(`/patients/${id}`);
        } catch (error) {
            return null;
        }
    },

    async getAppointments(patientId: number): Promise<Appointment[]> {
        const response = await apiClient.get<{ data: Appointment[] }>(`/patients/${patientId}/appointments`);
        return response.data || [];
    },

    async getConsultations(patientId: number): Promise<Consultation[]> {
        const response = await apiClient.get<{ data: Consultation[] }>(`/patients/${patientId}/consultations`);
        return response.data || [];
    }
};

// ==================== APPOINTMENT API ====================
export const AppointmentAPI = {
    async list(): Promise<Appointment[]> {
        const response = await apiClient.get<{ data: Appointment[] }>('/appointments');
        return response.data || [];
    },

    async getById(id: number): Promise<Appointment | null> {
        try {
            return await apiClient.get(`/appointments/${id}`);
        } catch (error) {
            return null;
        }
    },

    async create(appointment: Omit<Appointment, 'id' | 'createdAt' | 'updatedAt'>): Promise<Appointment> {
        return apiClient.post('/appointments', appointment);
    },

    async updateStatus(id: number, status: Appointment['status']): Promise<Appointment> {
        return apiClient.put(`/appointments/${id}`, { status });
    },

    async delete(id: number): Promise<void> {
        return apiClient.delete(`/appointments/${id}`);
    }
};

// ==================== CONSULTATION API ====================
export const ConsultationAPI = {
    async list(): Promise<Consultation[]> {
        const response = await apiClient.get<{ data: Consultation[] }>('/consultations');
        return response.data || [];
    },

    async getById(id: number): Promise<Consultation | null> {
        try {
            return await apiClient.get(`/consultations/${id}`);
        } catch (error) {
            return null;
        }
    },

    async create(consultation: Omit<Consultation, 'id' | 'createdAt' | 'updatedAt'>): Promise<Consultation> {
        return apiClient.post('/consultations', consultation);
    },

    async update(id: number, updates: Partial<Consultation>): Promise<Consultation> {
        return apiClient.put(`/consultations/${id}`, updates);
    }
};

// ==================== SEARCH API ====================
export const SearchAPI = {
    async searchDoctors(query: string): Promise<Doctor[]> {
        const allDoctors = await DoctorAPI.list();
        const q = query.toLowerCase();
        return allDoctors.filter(d =>
            d.firstName?.toLowerCase().includes(q) ||
            d.lastName?.toLowerCase().includes(q) ||
            d.speciality?.toLowerCase().includes(q)
        );
    },

    async searchPatients(query: string): Promise<Patient[]> {
        const allPatients = await PatientAPI.list();
        const q = query.toLowerCase();
        return allPatients.filter(p =>
            p.firstName?.toLowerCase().includes(q) ||
            p.lastName?.toLowerCase().includes(q) ||
            p.email?.toLowerCase().includes(q)
        );
    },

    async searchCabinets(query: string): Promise<Cabinet[]> {
        const allCabinets = await CabinetAPI.list();
        const q = query.toLowerCase();
        return allCabinets.filter(c =>
            c.name?.toLowerCase().includes(q) ||
            c.location?.address?.toLowerCase().includes(q)
        );
    }
};

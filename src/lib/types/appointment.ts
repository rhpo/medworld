import type { User, Users } from "./users";
import type { Cabinet } from "./cabinet";
import type { Doctor } from "./users/doctor";
import type { Patient } from "./users/patient";

export interface Appointment {
    id: number;
    date: Date;
    status: 'pending' | 'confirmed' | 'cancelled' | 'completed';

    doctor: Doctor;
    patient: Patient;
    cabinet: Cabinet;

    createdAt: Date;
    updatedAt?: Date;
}

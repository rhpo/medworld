import type { Appointment } from "./appointment";
import type { User, Users } from "./users";
import type { Doctor } from "./users/doctor";

export interface Consultation {
    id: number;
    doctor: Doctor;
    patient: User<Users.Patient>;
    appointment: Appointment;

    notes: string;
    prescriptions: string[];
    attachments?: string[]; // file URLs

    createdAt: Date;
    updatedAt?: Date;
}

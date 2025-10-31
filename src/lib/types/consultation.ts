import type { Appointment } from "./appointment";
import type { User, Users } from "./users";

export interface Consultation {
    id: number;
    doctor: User<Users.Doctor>;
    patient: User<Users.Patient>;
    appointment: Appointment;

    notes: string;
    prescriptions: string[];
    attachments?: string[]; // file URLs

    createdAt: Date;
    updatedAt?: Date;
}

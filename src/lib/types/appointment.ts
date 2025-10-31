import type { User, Users } from "./users";
import type { Cabinet } from "./cabinet";

export interface Appointment {
    id: number;
    date: Date;
    status: 'pending' | 'confirmed' | 'cancelled' | 'completed';

    doctor: User<Users.Doctor>;
    patient: User<Users.Patient>;
    cabinet: Cabinet;

    createdAt: Date;
    updatedAt?: Date;
}

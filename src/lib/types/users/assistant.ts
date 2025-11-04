import type { User, Users } from "./index";
import type { Appointment } from "../appointment";
import type { Cabinet } from "../cabinet";
import type { Doctor } from "./doctor";
import type { Admin } from "./admin";

export interface Assistant extends User<Users.Assistant> {
    cabinet: Cabinet;
    appointments: Appointment[];

    doctors: (Doctor | Admin)[];

    planAppointment(): void;
    cancelAppointment(): void;
    recordPayment(patientId: number, amount: number): void;
}

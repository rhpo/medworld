import type { Appointment } from "./appointment";
import type { User, Users } from "./users/";

export type Consultation = {
    doctor: User<Users.Doctor>;
    patient: User<Users.Patient>;
    attachment: Appointment;
}

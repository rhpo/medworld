import type { User, Users } from "./index";
import type { Appointment } from "../appointment";
import type { Consultation } from "../consultation";
import type { Review } from "../reviews";

export interface Patient extends User<Users.Patient> {
    appointments: Appointment[];
    consultations: Consultation[];
    reviews: Review[];

    medicalHistory: string[];
    allergies?: string[];

    addReview(review: Review): void;
}

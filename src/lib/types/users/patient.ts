import type { User, Users } from "./index";
import type { Appointment } from "../appointment";
import type { Consultation } from "../consultation";
import type { Review } from "../reviews";
import type { Prescription } from "../prescription";

export interface Patient extends User<Users.Patient> {
    emergencyContact: string;
    bloodType: string;
    appointments: Appointment[];
    consultations: Consultation[];
    prescriptions?: Prescription[];

    reviews: Review[];
    weight: number;

    medicalHistory: string[];
    allergies?: string[];

    addReview(review: Review): void;
}

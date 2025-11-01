import type { User, Users } from ".";
import type { Cabinet } from "../cabinet";
import type { Calendar } from "../calendar";
import type { Consultation } from "../consultation";
import type { Message } from "../message";
import type { Review } from "../reviews";
import type { Speciality } from "../speciality";

export interface Doctor extends User<Users.Doctor> {
    // Speciality
    speciality: Speciality;

    // Dates
    dateOfBirth: Date;
    careerStart: Date;

    // Messages
    messages: Message[];

    // Reviews
    reviews: Review[];

    // Calendars at different cabinets
    calendars: Calendar[];

    // Consultationss
    consultations: Consultation[];

    // Cabinets
    cabinets: Cabinet[];

    // Assistants
    assistants: User<Users.Assistant>[];

    // Price of consultations (per doctor)
    consultationPrice: number;
    consultationDuration: number; // in minutes

    // To Calculate Years Of Experience
    getYearsOfExperience(): number;

}

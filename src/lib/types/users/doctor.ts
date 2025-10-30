import type { User, Users } from ".";
import type { Calendar } from "../calendar";
import type { Message } from "../message";
import type { Review } from "../reviews";
import type { Speciality } from "../speciality";

export interface Doctor extends User<Users.Doctor> {
    type: Users.Doctor,

    speciality: Speciality,

    // Dates
    dateOfBirth: Date;
    careerStart: Date;

    // Messages
    messages: Message[];

    // Reviews
    reviews: Review[];

    // Calendars at different cabinets
    calendars: Calendar[];

    // To Calculate Years Of Experience
    getYearsOfExperience(start: Date): number;
}

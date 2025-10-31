import type { Cabinet } from "./cabinet";
import type { User, Users } from "./users";

export interface Calendar {
    id: number;
    cabinet: Cabinet;
    doctor: User<Users.Doctor>;
    availability: number[][]; // [dayIndex, hourIndex]
}

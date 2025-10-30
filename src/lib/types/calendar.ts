import type { Cabinet } from "./cabinet";
import type { User, Users } from "./users";

export type Calendar = {

    cabinet: Cabinet;
    doctor: User<Users.Doctor> | User<Users.Admin>;

    availability: number[][]; // [[1, 3], ... x7]

}

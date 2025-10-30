import type { Cabinet } from "./cabinet";
import type { User, Users } from "./users/";

export type Appointment = {
    time:                      Date;
    doctor:      User<Users.Doctor>;
    patient:     User<Users.Patient>;
    cabinet:                Cabinet;
}

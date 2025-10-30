import type { User, Users } from "./users"

export type Review = {
    from: User<Users.Patient>;
    to:   User<Users.Doctor> | User<Users.Admin>;
    rating: number;
    message: string;
}

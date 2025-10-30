import type { User, Users } from ".";
import type { Plan } from "../plan";

export interface Admin extends User<Users.SuperAdmin> {
    type: Users.SuperAdmin,

    plan: Plan;

}

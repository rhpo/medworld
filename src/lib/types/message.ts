import type { Cabinet } from "./cabinet";
import type { Consultation } from "./consultation";
import type { User, Users } from "./users/";

export type MessageContent = {
    message?: string;
    attachment: Consultation;
}

export type Message = {
    sender:             User<Users.Doctor>;
    cabinet:            Cabinet;
    receiver:           User<Users.Doctor> | User<Users.Admin>;
    date:               Date;
    content:            MessageContent;

    status:             'seen' | 'unseen'
}

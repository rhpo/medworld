import type { Cabinet } from "./cabinet";
import type { Consultation } from "./consultation";
import type { User, Users } from "./users/";
import type { Admin } from "./users/admin";
import type { Doctor } from "./users/doctor";

export type MessageContent = {
    message?: string;
    attachment: Consultation;
}

export type Message = {
    id:                 number;
    sender:             Doctor | Admin;
    cabinet:            Cabinet;
    receiver:           Doctor | Admin;
    date:               Date;
    content:            MessageContent;

    status:             'seen' | 'unseen'
}


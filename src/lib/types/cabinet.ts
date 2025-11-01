import type { Rating } from "./rating";
import type { User, Users } from "./users";

export interface Cabinet {
    id: number;
    name: string;
    address: string;
    phone: string;
    createdAt: Date;

    // Links
    admin: User<Users.Admin>;
    doctors: User<Users.Doctor>[];
    assistants: User<Users.Assistant>[];

    ratings: Rating[];

    // General Information
    image?: string;
    accessHandicap: boolean;

    // General Settings
    openingHours: Record<string, { open: string; close: string }>;

    isClosed?: boolean;
}

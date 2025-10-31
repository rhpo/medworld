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

    // Pack info
    premiumPack?: string;

    // General Settings
    openingHours?: Record<string, { open: string; close: string }>;
    consultationDuration?: number; // in minutes
    isClosed?: boolean;
}

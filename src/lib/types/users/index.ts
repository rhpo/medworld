
export type UserType = 'superadmin' | 'admin' | 'doctor' | 'assistant' | 'patient';

export const enum Users {
    SuperAdmin = 'superadmin',
    Admin = 'admin',
    Doctor = 'doctor',
    Assistant = 'assistant',
    Patient = 'patient'
}

export interface User<T extends UserType> {
    id: number;
    firstName: string;
    lastName: string;
    createdAt?: Date;

    type: T;

    email: string;
    password?: string;
    phoneNumber: string;
    avatarUrl?: string;

    getFullName(): string;
}

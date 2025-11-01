import type { Cabinet } from "./cabinet";
import type { Patient } from "./users/patient";

// Si il a l'equippement + Si il sait l'utiliser.
export type EquippementRating = {
    disponibility: boolean;
    rating: number;
}

export type UserExperience = {
    reception: number;
    hygiene: number;
    waitTime: number;
    communication: number;
    professionalism: number;
    emplacement: number;
}

export type Rating = {
    id: number;

    // Entre:
    patient: Patient;
    cabinet: Cabinet;

    date: Date;

    // Equippement
    equippement: EquippementRating;

    // Acceuil dans le cabinet
    userExperience: UserExperience;

}

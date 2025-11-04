import { writable, type Writable } from "svelte/store";
import type { IUser, User } from "./types/users";
import type { Group } from "./types/permission";
import type { Cabinet, Location } from "./types/cabinet";

export const user: Writable<IUser | null> = writable(null);

export const currentBlock: Writable<Group | null> = writable(null);

export const currentCabinet = writable<Cabinet | null>(null);

export const userLocation: Writable<Location> = writable({
    // Algiers
    address: "Algiers, Algeria",
    latitude: 36.7538,
    longitude: 3.0588,
})

//ik but this code is for messagestore so i put it here
// You should use: let messages = ($user as Doctor).messages // gives Message[]
//  {#each messages as message}
// h1 {message.sender.name}
// You
// {message.content}

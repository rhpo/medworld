import { writable, type Writable } from "svelte/store";
import type { User } from "./types/users";

export const user: Writable<User<any> | null> = writable(null);

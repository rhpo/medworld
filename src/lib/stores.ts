import { writable, type Writable } from "svelte/store";
import type { User } from "./types/users";
import type { Doctor } from "./types/users/doctor";
import type { SuperAdmin } from "./types/users/superadmin";
import type { Admin } from "./types/users/admin";
import type { Assistant } from "./types/users/assistant";
import type { Patient } from "./types/users/patient";

export const user: Writable<SuperAdmin | Admin | Doctor | Assistant | Patient | null> = writable(null);

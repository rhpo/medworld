import type { User, UserType, Users } from "./users/";

export type Cabinet = {
  id: string;
  name: string;
  date: Date;
  doctor: any;
  manager: User<Users.SuperAdmin> | User<Users.Admin>;
  assistant: User<Users.Assistant>;
  patient: User<Users.Patient>;
  owner: User<Users.Admin>;
};

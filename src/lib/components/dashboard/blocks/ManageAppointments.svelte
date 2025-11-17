<script lang="ts">
  import { AllAPI, AppointmentAPI } from "$lib/api";
  import Block from "$lib/components/ui/Block.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import IconButton from "$lib/components/ui/IconButton.svelte";
  import SearchInput from "$lib/components/ui/SearchInput.svelte";
  import { AppointmentStatus, type Appointment } from "$lib/types/appointment";
  import type { Cabinet } from "$lib/types/cabinet";
  import type { Permission } from "$lib/types/permission";
  import { Users, type IDoctor } from "$lib/types/users";
  import type { Assistant } from "$lib/types/users/assistant";
  import { Pen, Star, Timer, Trash } from "@lucide/svelte";
  import { onMount } from "svelte";

  let {
    user,
    permissions,
    cabinet = null,
  }: {
    user: Assistant | IDoctor;
    permissions: Permission[];
    cabinet: Cabinet | null;
  } = $props();

  let appointment: Appointment[] = $state([]);
  let realAppointments = $derived(
    appointment.filter((apt) => {
      if (
        apt.cabinet.name.toLowerCase().includes(query.toLowerCase()) ||
        apt.doctor.getFullName().toLowerCase().includes(query.toLowerCase()) ||
        apt.date.toString().toLowerCase().includes(query.toLowerCase()) ||
        apt.patient.getFullName().toLowerCase().includes(query.toLowerCase()) ||
        apt.status.toLowerCase().includes(query.toLowerCase())
      ) {
        return true;
      } else return false;
    })
  );

  let query: string = $state("");

  function sortAppointments(appointments: Appointment[]): Appointment[] {
    const priority: Record<AppointmentStatus, number> = {
      [AppointmentStatus.CONFIRMED]: 1,
      [AppointmentStatus.COMPLETED]: 0,
      [AppointmentStatus.CANCELLED]: 3,
      [AppointmentStatus.IN_PROGRESS]: 2,
      [AppointmentStatus.SCHEDULED]: 4,
      [AppointmentStatus.NO_SHOW]: 5,
    };

    return appointments.sort((a, b) => {
      return priority[a.status] - priority[b.status];
    });
  }

  onMount(async () => {
    if (user.type === Users.Assistant) {
      appointment = appointment;
    } else if (user.type === Users.Admin || user.type === Users.Doctor) {
      let allAppts = (await AllAPI.listAllAppointments()).filter(
        (appt) => appt.doctor.id === user.id
      );

      if (!cabinet) appointment = allAppts;
      else
        appointment = allAppts.filter((appt) => appt.cabinet.id === cabinet.id);
    } else {
      appointment = await AllAPI.listAllAppointments();
    }

    appointment = sortAppointments(appointment);
  });
</script>

<Block group="appointments" title="Appointments" Icon={Timer}>
  <SearchInput
    bind:value={query}
    placeholder="Search an Appointment/Patient..."
  />

  <br /><br />

  {#if realAppointments.length > 0}
    <table>
      <thead>
        <tr>
          <th>Date</th>
          <th>Patient</th>
          <th>Status</th>
          <th>Doctor</th>
          <th>Cabinet</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each realAppointments as appointment}
          <tr>
            <td>{new Date(appointment.date).toDateString()}</td>
            <td>{appointment.patient.getFullName()}</td>
            <td
              ><p class="status {appointment.status}">
                {appointment.status}
              </p></td
            >
            <td>{appointment.doctor.getFullName()}</td>
            <td>{appointment.cabinet.name}</td>
            <td class="actions">
              {#if permissions.includes("edit_appointment")}
                <IconButton Icon={Pen} label="Edit Appointment"></IconButton>
              {/if}
              {#if permissions.includes("cancel_appointment")}
                <IconButton type="error" Icon={Trash} label="Cancel Appointment"
                ></IconButton>
              {/if}
              {#if permissions.includes("view_appointment")}
                <IconButton
                  Icon={Timer}
                  color="#f74b00"
                  label="View Appointment Details"
                ></IconButton>
              {/if}
              {#if permissions.includes("rate_doctor") && appointment.status === AppointmentStatus.COMPLETED}
                <IconButton
                  Icon={Star}
                  color="darkblue"
                  label="Rate Doctor"
                  href="/doctors/{appointment.doctor.id}/rate"
                  target="_blank"
                ></IconButton>
              {/if}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {:else}
    <p>No Appointments.</p>
  {/if}
</Block>

<style>
  .actions {
    display: flex;
    gap: 0.5rem;
  }

  .status {
    border-radius: 9999px;
    border: 1px solid;
    width: fit-content;
    padding: 0.25rem 1rem;
    text-transform: capitalize;
  }

  /* SCHEDULED = 'SCHEDULED',
    CONFIRMED = 'CONFIRMED',
    IN_PROGRESS = 'IN_PROGRESS',
    COMPLETED = 'COMPLETED',
    CANCELLED = 'CANCELLED',
    NO_SHOW = 'NO_SHOW' */

  .status.COMPLETED {
    background: rgba(176, 176, 176, 0.128);
    border-color: rgba(133, 133, 133, 0.409);
  }

  .status.CONFIRMED {
    background: rgba(47, 255, 158, 0.128);
    border-color: rgba(4, 135, 102, 0.409);
  }

  .status.IN_PROGRESS {
    background: rgba(255, 245, 47, 0.128);
    border-color: rgba(205, 224, 35, 0.586);
  }

  .status.CANCELLED {
    background: rgba(255, 47, 47, 0.128);
    border-color: rgba(218, 40, 40, 0.409);
  }

  .status.NO_SHOW {
    background: rgba(242, 75, 75, 0.128);
    border-color: rgba(101, 12, 12, 0.409);
  }
</style>

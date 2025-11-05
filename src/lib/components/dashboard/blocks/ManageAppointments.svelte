<script lang="ts">
    import { AllAPI, AppointmentAPI } from "$lib/api";
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import IconButton from "$lib/components/ui/IconButton.svelte";
    import SearchInput from "$lib/components/ui/SearchInput.svelte";
    import type { Appointment } from "$lib/types/appointment";
    import type { Cabinet } from "$lib/types/cabinet";
    import type { Permission } from "$lib/types/permission";
    import { Users, type IDoctor } from "$lib/types/users";
    import type { Assistant } from "$lib/types/users/assistant";
    import { Pen, Timer, Trash } from "@lucide/svelte";
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
                apt.doctor
                    .getFullName()
                    .toLowerCase()
                    .includes(query.toLowerCase()) ||
                apt.date
                    .toString()
                    .toLowerCase()
                    .includes(query.toLowerCase()) ||
                apt.patient
                    .getFullName()
                    .toLowerCase()
                    .includes(query.toLowerCase()) ||
                apt.status.toLowerCase().includes(query.toLowerCase())
            ) {
                return true;
            } else return false;
        }),
    );

    let query: string = $state("");

    onMount(async () => {
        if (user.type === Users.Assistant) {
            appointment = appointment;
        } else if (user.type === Users.Admin || user.type === Users.Doctor) {
            let allAppts = (await AllAPI.listAllAppointments()).filter(
                (appt) => appt.doctor.id === user.id,
            );

            if (!cabinet) appointment = allAppts;
            else
                appointment = allAppts.filter(
                    (appt) => appt.cabinet.id === cabinet.id,
                );
        } else {
            appointment = await AllAPI.listAllAppointments();
        }
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
                                <IconButton Icon={Pen} label="Edit Appointment"
                                ></IconButton>
                            {/if}
                            {#if permissions.includes("cancel_appointment")}
                                <IconButton
                                    type="error"
                                    Icon={Trash}
                                    label="Cancel Appointment"
                                ></IconButton>
                            {/if}
                            {#if permissions.includes("view_appointment")}
                                <IconButton
                                    Icon={Timer}
                                    color="orange"
                                    label="View Appointment Details"
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

    .status.completed {
        background: rgba(172, 255, 47, 0.128);
        border-color: rgba(172, 255, 47, 0.409);
    }

    .status.confirmed {
        background: rgba(47, 255, 158, 0.128);
        border-color: rgba(4, 135, 102, 0.409);
    }

    .status.in_progress {
        background: rgba(255, 245, 47, 0.128);
        border-color: rgba(234, 255, 47, 0.409);
    }

    .status.canceled {
        background: rgba(255, 47, 47, 0.128);
        border-color: rgba(218, 40, 40, 0.409);
    }
</style>

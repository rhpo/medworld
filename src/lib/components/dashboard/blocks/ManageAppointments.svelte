<script lang="ts">
    import { AllAPI } from "$lib/api";
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import IconButton from "$lib/components/ui/IconButton.svelte";
    import type { Appointment } from "$lib/types/appointment";
    import type { Permission } from "$lib/types/permission";
    import { Users } from "$lib/types/users";
    import type { Assistant } from "$lib/types/users/assistant";
    import { Pen, Timer, Trash } from "@lucide/svelte";
    import { onMount } from "svelte";

    let {
        user,
        permissions,
    }: {
        user: Assistant;
        permissions: Permission[];
    } = $props();

    let appointment: Appointment[] = $state([]);

    onMount(async () => {
        if (user.type === Users.Assistant) {
            appointment = appointment;
        } else {
            appointment = await AllAPI.listAllAppointments();
        }
    });
</script>

<Block group="appointments" title="Manage Appointments" Icon={Timer}>
    <table>
        <thead>
            <tr>
                <th>Id</th>
                <th>Date</th>
                <th>Status</th>
                <th>Doctor</th>
                <th>Cabinet</th>
                <th>Created at</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            {#each appointment as appointment}
                <tr>
                    <td>{appointment.patient.getFullName()}</td>
                    <td>{appointment.date}</td>
                    <td>{appointment.status}</td>
                    <td>{appointment.doctor.getFullName()}</td>
                    <td>{appointment.cabinet.name}</td>
                    <td>{new Date(appointment.createdAt).toDateString()}</td>
                    <td class="actions">
                        {#if permissions.includes("edit_appointment")}
                            <IconButton
                                Icon={Pen}
                                title="Edit Appointment"
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
                                label="View Appointment Details"
                            ></IconButton>
                        {/if}
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</Block>

<style>
    .actions {
        display: flex;
        gap: 0.5rem;
    }
</style>

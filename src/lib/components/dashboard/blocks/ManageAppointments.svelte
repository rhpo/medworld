<script lang="ts">
    import { AllAPI } from "$lib/api";
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import type { Appointment } from "$lib/types/appointment";
    import type { Permission } from "$lib/types/permission";
    import { Users } from "$lib/types/users";
    import type { Assistant } from "$lib/types/users/assistant";
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

<main>
    <Block title="Manage Appointments">
        <table>
            <thead>
                <tr>
                    <th>Id</th>
                    <th>Date</th>
                    <th>Status</th>
                    <th>Doctor</th>
                    <th>Cabinet</th>
                    <th>Created at</th>
                    <th>Updated at</th>
                </tr>
            </thead>
            <tbody>
                {#each appointment as appointment}
                    <tr>
                        <td>{appointment.patient}</td>
                        <td>{appointment.date}</td>
                        <td>{appointment.status}</td>
                        <td>{appointment.doctor}</td>
                        <td>{appointment.cabinet}</td>
                        <td>{appointment.createdAt}</td>
                        <td>{appointment.updatedAt}</td>
                        <td class="actions">
                            {#if permissions.includes("edit_appointment")}
                                <Button>Edit</Button>
                            {/if}
                            {#if permissions.includes("cancel_appointment")}
                                <Button>Delete</Button>
                            {/if}
                            {#if permissions.includes("view_appointment")}
                                <Button>View</Button>
                            {/if}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </Block>
</main>

<style>
    .actions {
        display: flex;
        gap: 0.5rem;
    }
</style>

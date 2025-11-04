<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import type { Permission } from "$lib/types/permission";
    import type { Assistant } from "$lib/types/users/assistant";
    import { onMount } from "svelte";
    import type { Calendar } from "$lib/types/calendar";
    import IconButton from "$lib/components/ui/IconButton.svelte";
    import { CalendarIcon, Eye } from "@lucide/svelte";
    import type { Doctor } from "$lib/types/users/doctor";
    import type { Admin } from "$lib/types/users/admin";
    import { Users } from "$lib/types/users";
    import type { Cabinet } from "$lib/types/cabinet";

    let {
        user,
        permissions,
        cabinet,
    }: {
        user: Doctor | Admin | Assistant;
        permissions: Permission[];
        cabinet: Cabinet;
    } = $props();

    let calendars: Calendar[] = $state([]);
    let filteredCalendars = $derived(
        calendars.filter((cal) => cal.cabinet.id === cabinet.id),
    );

    onMount(() => {
        switch (user.type) {
            case Users.Assistant:
                const assistant = user as Assistant;
                calendars = assistant.doctors.flatMap(
                    (doctor) => doctor.calendars,
                );
                break;
            case Users.Doctor:
                const doctor = user as Doctor;
                calendars = doctor.calendars;
                break;
            case Users.Admin:
                const admin = user as Admin;
                calendars = admin.calendars;
                break;
        }
    });
</script>

<Block group="calendar" title={`Calendar`} Icon={CalendarIcon}>
    <table>
        <thead>
            <tr>
                <th>Day</th>
                <th>Available Slots</th>
            </tr>
        </thead>
        <tbody>
            {#each filteredCalendars as calendar}
                <tr>
                    <td>
                        {calendar.availability.length} slots available
                    </td>
                    <td class="actions">
                        {#if permissions.includes("view_calendar")}
                            <IconButton
                                Icon={Eye}
                                label="View Calendar"
                                href="/dashboard/calendar/{calendar.id}"
                            />
                        {/if}
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>

    {#if filteredCalendars.length === 0}
        <p class="no-calendars">No calendars found for this cabinet.</p>
    {/if}
</Block>

<style>
    table {
        width: 100%;
        border-collapse: collapse;
        margin-top: 1rem;
    }

    th,
    td {
        padding: 0.75rem;
        text-align: left;
        border-bottom: 1px solid #e2e8f0;
    }

    th {
        font-weight: 600;
        color: #4a5568;
        background-color: #f7fafc;
    }

    .actions {
        display: flex;
        gap: 0.5rem;
    }

    .no-calendars {
        text-align: center;
        padding: 2rem;
        color: #718096;
        font-size: 1.1rem;
    }

    tr:hover {
        background-color: #f8fafc;
    }
</style>

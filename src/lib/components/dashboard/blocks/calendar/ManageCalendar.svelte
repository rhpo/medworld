<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import type { Permission } from "$lib/types/permission";
    import type { Assistant } from "$lib/types/users/assistant";
    import { onMount } from "svelte";
    import type { Calendar } from "$lib/types/calendar";
    import IconButton from "$lib/components/ui/IconButton.svelte";
    import { CalendarIcon, Eye, Pen } from "@lucide/svelte";
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
        // calendars.filter((cal) => cal.cabinet.id === cabinet.id),
        calendars,
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
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            {#each filteredCalendars as calendar}
                {#each calendar.availability as availability}
                    <tr>
                        <td>
                            {new Date().toDateString()}
                        </td>

                        <td>
                            <div class="slots">
                                {#each availability.slots as slot}
                                    <p>
                                        {slot}
                                    </p>
                                {/each}
                            </div>
                        </td>

                        <td>
                            <div class="actions">
                                {#if permissions.includes("view_calendar")}
                                    <!-- <IconButton
                                    Icon={Eye}
                                    label="View Calendar"
                                    href="/dashboard/calendar/{calendar.id}"
                                /> -->

                                    <IconButton
                                        Icon={Eye}
                                        label="View Calendar"
                                    />
                                {/if}

                                {#if permissions.includes("edit_calendar")}
                                    <IconButton
                                        Icon={Pen}
                                        color="orange"
                                        label="Edit Calendar"
                                    />
                                {/if}
                            </div>
                        </td>
                    </tr>
                {/each}
            {/each}
        </tbody>
    </table>

    {#if filteredCalendars.length === 0}
        <p class="no-calendars">No calendars found for this cabinet.</p>
    {/if}
</Block>

<style>
    .no-calendars {
        text-align: center;
        padding: 2rem;
        color: #718096;
        font-size: 1.1rem;
    }

    .slots {
        display: flex;
        flex-wrap: wrap;
        gap: 0.25rem;
    }

    .slots p {
        background-color: var(--color-primary);
        padding: 0.5rem 1rem;
        width: fit-content;
        color: white;
        border-radius: 1rem;
    }

    table .actions {
        display: flex;
        gap: 0.5rem;
    }

    tr:hover {
        background-color: #f8fafc;
    }
</style>

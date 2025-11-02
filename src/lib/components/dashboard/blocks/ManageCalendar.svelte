<script lang="ts">
    import { AllAPI } from "$lib/api";
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import type { Appointment } from "$lib/types/appointment";
    import type { Permission } from "$lib/types/permission";
    import { Users } from "$lib/types/users";
    import type { Assistant } from "$lib/types/users/assistant";
    import { onMount } from "svelte";
    import type { Calendar } from "$lib/types/calendar";

    let {
        user,
        permissions,
    }: {
        user: Assistant;
        permissions: Permission[];
    } = $props();

    let calendar: Calendar[] = $state([]);

    onMount(async () => {
        if (user.type === Users.Assistant) {
            calendar = calendar;
        }
    });
</script>

<main>
    <Block title="Calendar">
        <table>
            <thead>
                <tr>
                    <th>Id</th>
                    <th>Doctor</th>
                    <th>Cabinet</th>
                    <th>Availability</th>
                </tr>
            </thead>
            <tbody>
                {#each calendar as calendar}
                    <tr>
                        <td>{calendar.id}</td>
                        <td>{calendar.doctor}</td>
                        <td>{calendar.cabinet}</td>
                        <td>{calendar.availability}</td>

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

<script lang="ts">
    import ManagePatients from "./blocks/ManagePatients.svelte";
    import { getPermissionsFromUserType } from "$lib/types/permission";

    import type { Assistant } from "$lib/types/users/assistant";
    import ManageAppointments from "./blocks/ManageAppointments.svelte";
    import ManageCalendar from "./blocks/ManageCalendar.svelte";

    interface IProps {
        assistant: Assistant;
    }

    let { assistant }: IProps = $props();
    let permissions = getPermissionsFromUserType(assistant.type);
</script>

<main>
    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={assistant} />
    {/if}

    {#if permissions.find((e) => e.endsWith("_appointment"))}
        <ManageAppointments user={assistant} {permissions} />
    {/if}

    {#if permissions.find((e) => e.endsWith("_calendar"))}
        <ManageCalendar user={assistant} {permissions} />
    {/if}
</main>

<style>
    main {
        padding-top: 2rem;
        display: flex;
        flex-direction: column;
        gap: 4rem;
        margin-bottom: 4rem;
    }
</style>

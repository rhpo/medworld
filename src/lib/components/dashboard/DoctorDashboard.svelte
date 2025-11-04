<script lang="ts">
    import ManagePatients from "./blocks/ManagePatients.svelte";
    import ManageAssistants from "./blocks/ManageAssistants.svelte";
    import {
        getPermissionsFromUserType,
        groupPermissions,
    } from "$lib/types/permission";

    import type { Doctor } from "$lib/types/users/doctor";
    import { currentBlock, currentCabinet } from "$lib/stores";
    import ManageCalendar from "./blocks/calendar/ManageCalendar.svelte";
    import type { Cabinet } from "$lib/types/cabinet";
    import CabinetSelector from "./blocks/cabinets/CabinetSelector.svelte";

    interface IProps {
        doctor: Doctor;
    }

    let { doctor }: IProps = $props();
    let permissions = getPermissionsFromUserType(doctor.type);
    let permissionGroups = groupPermissions(permissions);
</script>

<main class:grid={$currentBlock === null}>
    {#if permissionGroups.includes("patients")}
        <ManagePatients user={doctor} cabinet={$currentCabinet as Cabinet} />
    {/if}

    {#if permissionGroups.includes("assistants")}
        <ManageAssistants user={doctor} {permissions} />
    {/if}

    {#if permissionGroups.includes("calendar")}
        <ManageCalendar
            user={doctor}
            {permissions}
            cabinet={$currentCabinet as Cabinet}
        />
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

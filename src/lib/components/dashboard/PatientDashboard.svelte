<script lang="ts">
    import ManagePatients from "./blocks/ManagePatients.svelte";
    import {
        getPermissionsFromUserType,
        type Permission,
    } from "$lib/types/permission";

    import type { Patient } from "$lib/types/users/patient";
    import ManageAppointments from "./blocks/ManageAppointments.svelte";
    import BookAppointment from "./blocks/BookAppointment.svelte";
    import type { Cabinet } from "$lib/types/cabinet";
    import { currentBlock, currentCabinet } from "$lib/stores";

    interface IProps {
        patient: Patient;
    }

    let { patient }: IProps = $props();
    let permissions = getPermissionsFromUserType(patient.type);
</script>

<main class:grid={$currentBlock === null}>
    {#if permissions.find((e) => e === "book_appointment")}
        <BookAppointment />
    {/if}

    {#if permissions.find( (e) => (["view_appointment", "edit_appointment", "cancel_appointment"] as Permission[]).includes(e), )}
        <ManageAppointments
            cabinet={$currentCabinet as Cabinet}
            user={patient as any}
            {permissions}
        />
    {/if}

    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={patient} />
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

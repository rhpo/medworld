<script lang="ts">
    import ManagePatients from "./blocks/ManagePatients.svelte";
    import { getPermissionsFromUserType } from "$lib/types/permission";

    import type { Patient } from "$lib/types/users/patient";
    import ManageAppointments from "./blocks/ManageAppointments.svelte";

    interface IProps {
        patient: Patient;
    }

    let { patient }: IProps = $props();
    let permissions = getPermissionsFromUserType(patient.type);
</script>

<main>
    <!-- Show Blocks based on the permissions that the patient has -->
    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={patient} />
    {/if}

    <!-- Show Blocks based on the permissions that the patient has -->
    {#if permissions.find((e) => e.endsWith("_appointment"))}
        <ManageAppointments user={patient as any} {permissions} />
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

<script lang="ts">
    import { getPermissionsFromUserType } from "$lib/types/permission";
    import type { Doctor } from "$lib/types/users/doctor";
    import ManageAssistants from "./blocks/ManageAssistants.svelte";
    import ManagePatients from "./blocks/ManagePatients.svelte";

    interface IProps {
        doctor: Doctor;
    }

    let { doctor }: IProps = $props();
    let permissions = getPermissionsFromUserType(doctor.type);
</script>

<h1>
    Welcome, Dr. {doctor.firstName}
</h1>

<main>
    <!-- Show Blocks based on the permissions that the doctor has -->
    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={doctor} />
    {/if}

    {#if permissions.find((permission) => permission.endsWith("_assistant"))}
        <ManageAssistants user={doctor} {permissions} />
    {/if}
</main>

<style>
    main {
        padding-top: 2rem;
        display: flex;
        flex-direction: column;
        gap: 2rem;
        margin-bottom: 4rem;
    }
</style>

<script lang="ts">
    import ManagePatients from "./blocks/ManagePatients.svelte";
    import ManageAssistants from "./blocks/ManageAssistants.svelte";
    import { getPermissionsFromUserType } from "$lib/types/permission";

    import type { Doctor } from "$lib/types/users/doctor";

    interface IProps {
        doctor: Doctor;
    }

    let { doctor }: IProps = $props();
    let permissions = getPermissionsFromUserType(doctor.type);
</script>

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
        gap: 4rem;
        margin-bottom: 4rem;
    }
</style>

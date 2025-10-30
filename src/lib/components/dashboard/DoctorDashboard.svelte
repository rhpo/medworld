<script lang="ts">
    import { getPermissionsFromUserType } from "$lib/types/permission";
    import type { Doctor } from "$lib/types/users/doctor";
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
    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={doctor} />
    {/if}
</main>

<style>
    main {
        padding-top: 2rem;
    }

    h1 {
    }
</style>

<script lang="ts">
    import { currentBlock, currentCabinet, user } from "$lib/stores";
    import type { Cabinet } from "$lib/types/cabinet";
    import { getPermissionsFromUserType } from "$lib/types/permission";

    import type { SuperAdmin } from "$lib/types/users/superadmin";

    import ManageCabinet from "./blocks/ManageCabinet.svelte";
    import ManageDoctors from "./blocks/ManageDoctors.svelte";
    import ManagePatients from "./blocks/ManagePatients.svelte";

    interface IProps {
        superadmin: SuperAdmin;
    }

    let { superadmin }: IProps = $props();
    let permissions = getPermissionsFromUserType(superadmin.type);
</script>

<main class:grid={$currentBlock === null}>
    {#if permissions.find((e) => e.endsWith("_doctor"))}
        <ManageDoctors user={superadmin} {permissions} />
    {/if}

    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={superadmin} />
    {/if}

    {#if permissions.find((p) => p.endsWith("cabinet"))}
        <ManageCabinet
            user={$user as SuperAdmin}
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

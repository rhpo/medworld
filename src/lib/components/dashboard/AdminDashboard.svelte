<script lang="ts">
    import { getPermissionsFromUserType } from "$lib/types/permission";
    import type { Admin } from "$lib/types/users/admin";
    import ManageAssistants from "./blocks/ManageAssistants.svelte";
    import ManagePatients from "./blocks/ManagePatients.svelte";
    import ManageDoctors from "./blocks/ManageDoctors.svelte";
    import Avatar from "../ui/Avatar.svelte";

    interface IProps {
        admin: Admin;
    }

    let { admin }: IProps = $props();
    let permissions = getPermissionsFromUserType(admin.type);
</script>

<main>
    {#if permissions.find((permission) => permission.endsWith("_doctor"))}
        <ManageDoctors user={admin} {permissions} />
    {/if}

    {#if permissions.find((permission) => permission.endsWith("_assistant"))}
        <ManageAssistants user={admin} {permissions} />
    {/if}

    <!-- Show Blocks based on the permissions that the admin has -->
    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={admin} />
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

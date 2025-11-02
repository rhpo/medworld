<script lang="ts">
    import { getPermissionsFromUserType } from "$lib/types/permission";
    import type { Admin } from "$lib/types/users/admin";
    import ManageAssistants from "./blocks/ManageAssistants.svelte";
    import ManagePatients from "./blocks/ManagePatients.svelte";
    import ManageDoctors from "./blocks/ManageDoctors.svelte";
    import Plan from "./blocks/Plan.svelte";

    import type { AnyUser } from "$lib/types/users";

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

    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={admin as AnyUser} />
    {/if}

    <Plan />
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

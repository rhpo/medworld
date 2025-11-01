<script lang="ts">
    import AdminDashboard from "$lib/components/dashboard/AdminDashboard.svelte";
    import DoctorDashboard from "$lib/components/dashboard/DoctorDashboard.svelte";
    import Avatar from "$lib/components/ui/Avatar.svelte";
    import View from "$lib/components/ui/View.svelte";

    import { user } from "$lib/stores";

    import {
        fakeAdmins,
        fakeDoctors,
        fakeSuperAdmin,
    } from "$lib/types/fakedata";

    import { Users, type UserType } from "$lib/types/users";
    import type { Admin } from "$lib/types/users/admin";
    import type { Doctor } from "$lib/types/users/doctor";
    import { onMount } from "svelte";

    let currentDemoAccountType: UserType = "doctor";

    onMount(() => {
        user.set(fakeDoctors[0]);
    });
</script>

{#if $user}
    <View style="padding-top: 2rem;">
        <div class="welcome">
            <Avatar
                size="4rem"
                avatarUrl={$user.avatarUrl}
                alt={$user.firstName + " " + $user.lastName}
                original
            />

            <h1>
                Welcome, {$user.type === Users.Doctor ||
                $user.type === Users.Admin
                    ? "Dr. "
                    : ""}
                {$user.firstName}
                {$user.type === Users.Admin ? " (Admin)" : ""}
            </h1>
        </div>

        {#if $user.type === Users.Doctor}
            <DoctorDashboard doctor={$user as Doctor} />
        {:else if $user.type === Users.Admin}
            <AdminDashboard admin={$user as Admin} />
        {:else}
            <h2 style="color: var(--error)">
                Unsupported Account type: {$user.type}
            </h2>
        {/if}

        <div style="margin-bottom: 2rem;">
            <h4>Demo Helper</h4>

            Account Type:
            <!-- I'm here! -->
            <select bind:value={currentDemoAccountType}>
                <option value={Users.Doctor}>Doctor</option>
                <option value={Users.Admin}>Admin</option>
                <option value={Users.SuperAdmin}>Super Admin</option>
            </select>

            <br />

            Available {currentDemoAccountType}s :
            <select bind:value={$user}>
                {#if currentDemoAccountType === Users.Doctor}
                    {#each fakeDoctors as doctor}
                        <option value={doctor}>{doctor.getFullName()}</option>
                    {/each}
                {:else if currentDemoAccountType === Users.SuperAdmin}
                    <option value={fakeSuperAdmin}
                        >{fakeSuperAdmin.getFullName()}</option
                    >
                {:else if currentDemoAccountType === Users.Admin}
                    {#each fakeAdmins as user}
                        <option value={user}>{user.getFullName()}</option>
                    {/each}
                {:else if currentDemoAccountType === Users.Assistant}
                    {#each fakeDoctors as user}
                        <option value={user}>{user.getFullName()}</option>
                    {/each}
                {:else if currentDemoAccountType === Users.Patient}
                    {#each fakeDoctors as user}
                        <option value={user}>{user.getFullName()}</option>
                    {/each}
                {/if}
            </select>
        </div>
    </View>
{/if}

<style>
    .welcome {
        display: flex;

        align-items: center;
        gap: 1rem;

        padding-bottom: 2rem;
        margin-bottom: 1rem;
        border-bottom: 1px solid var(--border-color-light);
    }

    .welcome h1 {
        font-size: 3rem;
    }
</style>

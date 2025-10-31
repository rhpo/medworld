<script lang="ts">
    import DoctorDashboard from "$lib/components/dashboard/DoctorDashboard.svelte";
    import View from "$lib/components/ui/View.svelte";
    import {
        fakeAdmins,
        fakeDoctors,
        fakeSuperAdmin,
    } from "$lib/types/fakedata";
    import { Users, type User, type UserType } from "$lib/types/users";

    let currentDemoAccountType: UserType = "doctor";
    let user: any = fakeDoctors[0];
</script>

<View style="padding-top: 2rem;">
    <div style="margin-bottom: 1rem;">
        <!-- I'm here! -->
        <select bind:value={currentDemoAccountType}>
            <option value={Users.Doctor}>Doctor</option>
            <option value={Users.Admin}>Admin</option>
            <option value={Users.SuperAdmin}>Super Admin</option>
        </select>
    </div>

    <div style="margin-bottom: 1rem;">
        <select bind:value={user}>
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

    {#if user.type === Users.Doctor}
        <DoctorDashboard doctor={user} />
    {:else}
        <h2 style="color: var(--error)">
            Unsupported Account type: {user.type}
        </h2>
    {/if}
</View>

<script lang="ts">
    import DoctorDashboard from "$lib/components/dashboard/DoctorDashboard.svelte";
    import View from "$lib/components/ui/View.svelte";
    import { Speciality } from "$lib/types/speciality";
    import { Users } from "$lib/types/users";
    import type { Doctor } from "$lib/types/users/doctor";

    // @ts-ignore
    let user: Doctor = {
        type: Users.Doctor,
        firstName: "Manel",
        lastName: "BELAIOUER",
        id: 0,
        email: "manel@gmail.com",
        phoneNumber: "+213 552 428 434",
        careerStart: new Date("22/12/1986"),
        speciality: Speciality.CARDIOLOGY,
    };
</script>

<View style="padding-top: 2rem;">
    {#if user.type === Users.Doctor}
        <DoctorDashboard doctor={user} />
    {:else}
        <h2 style="color: var(--error)">
            Unsupported Account type: {user.type}
        </h2>
    {/if}
</View>

<script lang="ts">
    import { getPermissionsFromUserType } from "$lib/types/permission";

    import Dashboard from "./Dashboard.svelte";
    import ManagePatients from "./blocks/ManagePatients.svelte";
    import ManageCalendar from "./blocks/calendar/ManageCalendar.svelte";
    import ManageAppointments from "./blocks/ManageAppointments.svelte";
    import ManageConsultations from "./blocks/ManageConsultations.svelte";

    import type { Cabinet } from "$lib/types/cabinet";
    import type { Assistant } from "$lib/types/users/assistant";

    interface IProps {
        assistant: Assistant;
    }

    let { assistant }: IProps = $props();
    let permissions = getPermissionsFromUserType(assistant.type);
</script>

<!-- I removed cabvinet selector logic cuz assistant has one cabinet (romy told me) -->

<Dashboard>
    {#if permissions.find((e) => e.endsWith("_patient"))}
        <ManagePatients user={assistant} />
    {/if}

    {#if permissions.find((e) => e.endsWith("_appointment"))}
        <ManageAppointments
            user={assistant}
            {permissions}
            cabinet={assistant.cabinet}
        />
    {/if}

    {#if permissions.find((e) => e.endsWith("_consultation"))}
        <ManageConsultations
            user={assistant}
            patients={assistant.doctors
                .map((d) => d.consultations.map((e) => e.patient))
                .flat()}
        />
    {/if}

    {#if permissions.find((e) => e.endsWith("_calendar"))}
        <ManageCalendar
            user={assistant}
            {permissions}
            cabinet={assistant.cabinet as Cabinet}
        />
    {/if}
</Dashboard>

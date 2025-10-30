<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import type { User, Users } from "$lib/types/users";

    interface IProps {
        user: User<any>;
    }

    let patients: User<Users.Patient>[] = [];

    let { user }: IProps = $props();
</script>

<Block title="Manage Patients">
    {#if patients.length === 0}
        <h2>No patients.</h2>
    {:else}
        <!-- Display the patients list. -->
    {/if}
</Block>

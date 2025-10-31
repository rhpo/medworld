<script lang="ts">
    import { DoctorAPI } from "$lib/api";
    import Block from "$lib/components/ui/Block.svelte";
    import SearchInput from "$lib/components/ui/SearchInput.svelte";
    import type { User, Users } from "$lib/types/users";
    import { onMount } from "svelte";

    interface IProps {
        user: User<any>;
    }

    let patients: User<Users.Patient>[] = $state([]);
    let searchQuery: string = $state("");
    let { user }: IProps = $props();

    onMount(async () => {
        patients = await DoctorAPI.getPatients(user.id);
    });

    let filteredPatients = $derived(
        patients.filter((patient) => {
            if (!searchQuery) return true;
            const searchTerm = searchQuery.toLowerCase();
            return (
                patient.getFullName().toLowerCase().includes(searchTerm) ||
                patient.email.toLowerCase().includes(searchTerm) ||
                (patient.phoneNumber &&
                    patient.phoneNumber.includes(searchTerm))
            );
        }),
    );
</script>

<Block title="Manage Patients">
    <div class="search-container">
        <SearchInput
            bind:value={searchQuery}
            placeholder="Search patients..."
        />
    </div>

    {#if filteredPatients.length === 0}
        <h2>No patients found.</h2>
    {:else}
        <table>
            <thead>
                <tr>
                    <th>Avatar</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Phone</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each filteredPatients as patient}
                    <tr>
                        <td>
                            <img
                                src={patient.avatarUrl || "/default-avatar.png"}
                                alt="{patient.getFullName()}'s avatar"
                            />
                        </td>
                        <td>{patient.getFullName()}</td>
                        <td>{patient.email}</td>
                        <td>{patient.phoneNumber || "N/A"}</td>
                        <td>
                            <button>View</button>
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}
</Block>

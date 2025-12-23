<script lang="ts">
    import SearchInput from "$lib/components/ui/SearchInput.svelte";
    import type { Cabinet } from "$lib/types/cabinet";
    import {
        cabinets as cabinetsStore,
        loadAllCabinets,
        isLoadingCabinets,
    } from "$lib/stores/data";
    import type { IUser } from "$lib/types/users";
    import type { Admin } from "$lib/types/users/admin";
    import type { Doctor } from "$lib/types/users/doctor";
    import { onMount } from "svelte";
    import CabinetCard from "./CabinetCard.svelte";
    import { scale } from "svelte/transition";

    let {
        selectedCabinet = $bindable(null),
        user,
    }: {
        selectedCabinet: Cabinet | null;
        user: IUser;
    } = $props();

    let cabinets: Cabinet[] = $derived($cabinetsStore);
    let search: string = $state("");

    onMount(async () => {
        // Load cabinets from API if not already loaded
        if ($cabinetsStore.length === 0) {
            await loadAllCabinets();
        }

        // Filter cabinets based on user type
        switch (user.type) {
            case "superadmin":
                cabinets = $cabinetsStore;
                break;
            case "doctor":
                cabinets = (user as Doctor).cabinet
                    ? [(user as Doctor).cabinet]
                    : $cabinetsStore;
                break;
            case "admin":
                cabinets = (user as Admin).cabinet
                    ? [(user as Admin).cabinet]
                    : $cabinetsStore;
                break;
            default:
                cabinets = [];
        }
    });

    function searchCabinet(cabinets: Cabinet[], query: string): Cabinet[] {
        let result = cabinets;
        if (result.length > 0) {
            query = query.toLowerCase().trim();

            result = result.filter((cabinet) => {
                if (
                    cabinet.name.toLowerCase().includes(query) ||
                    cabinet.doctors
                        .map((d) => d.getFullName().toLowerCase())
                        .includes(query) ||
                    cabinet.doctors.some((d) =>
                        query.includes(d.speciality.toLowerCase()),
                    ) ||
                    cabinet.location.address.toLowerCase().includes(query)
                ) {
                    return true;
                } else return false;
            });
        }

        return result;
    }
</script>

<div class="cabinet-selector">
    <div class="search-container">
        <SearchInput
            placeholder="Search cabinets..."
            bind:value={search}
            class="search-input"
        />
    </div>

    {#if cabinets.length === 0}
        <pre
            style="text-align: center; font-size: 2rem;">No cabinets found.</pre>
    {:else}
        <div class="cabinets-grid">
            {#each searchCabinet(cabinets, search) as cabinet}
                <button
                    onclick={() => (selectedCabinet = cabinet)}
                    transition:scale
                >
                    <CabinetCard {cabinet} />
                </button>
            {/each}
        </div>
    {/if}
</div>

<style>
    button {
        background-color: transparent;
        /*é */
        border: none;
        outline: none;
    }

    .cabinet-selector {
        width: 100%;
        margin-bottom: 1rem;
    }

    .search-container {
        width: 100%;
        margin-bottom: 1rem;
    }

    .cabinets-grid {
        display: grid;
        grid-template-columns: 1fr;
        gap: 1rem;
    }

    @media (min-width: 768px) {
        .cabinets-grid {
            grid-template-columns: repeat(2, 1fr);
        }
    }

    @media (min-width: 1024px) {
        .cabinets-grid {
            grid-template-columns: repeat(3, 1fr);
        }
    }
</style>

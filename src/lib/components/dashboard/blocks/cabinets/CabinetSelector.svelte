<script lang="ts">
    import SearchInput from "$lib/components/ui/SearchInput.svelte";
    import type { Cabinet } from "$lib/types/cabinet";
    import { fakeCabinets } from "$lib/types/fakedata";
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

    let cabinets: Cabinet[] = $state([]);
    let search: string = $state("");

    function getDayOfWeek(): string {
        const days = [
            "Sunday",
            "Monday",
            "Tuesday",
            "Wednesday",
            "Thursday",
            "Friday",
            "Saturday",
        ];
        return days[new Date().getDay()];
    }

    onMount(() => {
        switch (user.type) {
            case "doctor":
                cabinets = (user as Doctor).cabinets;
                break;
            case "admin":
                cabinets = (user as Admin).cabinets;
                break;
            default:
                cabinets = [];
        }
    });

    function searchCabinet(cabinets: Cabinet[], query: string): Cabinet[] {
        let result = cabinets;
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

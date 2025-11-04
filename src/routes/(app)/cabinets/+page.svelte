<script lang="ts">
    import {
        calculateDistance,
        type Cabinet,
        type Location,
    } from "$lib/types/cabinet";
    import Block from "$lib/components/ui/Block.svelte";
    import View from "$lib/components/ui/View.svelte";
    import Face from "$lib/components/ui/Face.svelte";
    import SearchInput from "$lib/components/ui/SearchInput.svelte";
    import {
        MapPin,
        Phone,
        Users as UsersIcon,
        Plus,
        AlarmSmoke,
        CreditCard,
        ParkingCircle,
        Wifi,
        Accessibility,
        Clock,
    } from "@lucide/svelte";
    import { fakeCabinets } from "$lib/types/fakedata";
    import CabinetCard from "$lib/components/dashboard/blocks/cabinets/CabinetCard.svelte";
    import { fade, scale } from "svelte/transition";
    import { calculateRatingScore } from "$lib/types/rating";
    import { userLocation } from "$lib/stores";
    import { Specialities, type Speciality } from "$lib/types/speciality";

    let allCabinets: Cabinet[] = fakeCabinets;

    type SearchSettings = {
        term: string;
        criteresElemination: {
            acceptsUrgent: boolean;
            acceptsInsurance: boolean;
            hasParking: boolean;
            hasWifi: boolean;
            isOpen: boolean;
            hasAccessHandicap: boolean;
        };
        speciality: Speciality | null;
        sortBy: "name" | "location" | "rating" | "reviews";
        orderBy: "asc" | "desc";
    };

    let search: SearchSettings = $state({
        term: "",
        criteresElemination: {
            acceptsUrgent: false,
            acceptsInsurance: false,
            hasParking: false,
            hasWifi: false,
            hasAccessHandicap: false,
            isOpen: false,
        },
        sortBy: "rating",
        speciality: null,
        orderBy: "desc",
    });

    function performSearch(
        cabinets: Cabinet[],
        search: SearchSettings,
    ): Cabinet[] {
        let result = cabinets;

        result = result.filter((cabinet) => {
            if (
                search.speciality !== null &&
                cabinet.doctors.every(
                    (doc) => doc.speciality !== search.speciality,
                )
            ) {
                return false;
            }

            if (
                search.criteresElemination.acceptsUrgent &&
                !cabinet.acceptsUrgent
            )
                return false;
            if (
                search.criteresElemination.acceptsInsurance &&
                !cabinet.acceptsInsurance
            )
                return false;
            if (search.criteresElemination.hasParking && !cabinet.hasParking)
                return false;
            if (search.criteresElemination.hasWifi && !cabinet.hasWifi)
                return false;
            if (
                search.criteresElemination.hasAccessHandicap &&
                !cabinet.accessHandicap
            )
                return false;
            return true;
        });

        if (search.term.trim() !== "") {
            const term = search.term.toLowerCase();
            result = result.filter((cabinet) => {
                return (
                    cabinet.name.toLowerCase().includes(term) ||
                    cabinet.location.address.toLowerCase().includes(term) ||
                    cabinet.admin.getFullName().toLowerCase().includes(term)
                );
            });
        }

        result = result.sort((a, b) => {
            let compareA: string | number = "";
            let compareB: string | number = "";

            switch (search.sortBy) {
                case "name":
                    compareA = a.name.toLowerCase();
                    compareB = b.name.toLowerCase();
                    break;
                case "location":
                    compareA = calculateDistance(a, $userLocation);
                    compareB = calculateDistance(b, $userLocation);
                    break;
                case "rating":
                    compareA = calculateRatingScore(a);
                    compareB = calculateRatingScore(b);
                    break;
                case "reviews":
                    compareA = a.ratings.length;
                    compareB = b.ratings.length;
                    break;
            }

            if (compareA < compareB) return search.orderBy === "asc" ? -1 : 1;
            if (compareA > compareB) return search.orderBy === "asc" ? 1 : -1;
            return 0;
        });

        return result;
    }

    let filteredCabinets: Cabinet[] = $derived(
        performSearch(allCabinets, search),
    );
</script>

<Face title="Cabinets" description="Browse and manage cabinets">
    <View>
        <div class="page-inner">
            <div class="list-header">
                <div class="search">
                    <SearchInput
                        placeholder="Search cabinets by name, address or admin..."
                        bind:value={search.term}
                    />

                    <div class="filters-container" transition:fade>
                        <div class="filters-section sorting">
                            <h3>Sort Options</h3>
                            <div class="sort-controls">
                                <select
                                    bind:value={search.sortBy}
                                    class="select-field"
                                >
                                    <option value="name">Sort by Name</option>
                                    <option value="location"
                                        >Sort by Location</option
                                    >
                                    <option value="rating"
                                        >Sort by Rating</option
                                    >
                                    <option value="reviews"
                                        >Sort by Reviews</option
                                    >
                                </select>
                                <select
                                    bind:value={search.orderBy}
                                    class="select-field order-select"
                                >
                                    <option value="asc">Ascending</option>
                                    <option value="desc">Descending</option>
                                </select>
                            </div>
                        </div>

                        <div class="filters-section specialities-section">
                            <h3>Specialities</h3>
                            <div class="specialities">
                                <select
                                    class="select-field"
                                    bind:value={search.speciality}
                                >
                                    <option value={null}
                                        >All Specialities</option
                                    >
                                    {#each Specialities as speciality}
                                        <option value={speciality}
                                            >{speciality}</option
                                        >
                                    {/each}
                                </select>
                            </div>
                        </div>

                        <div class="filters-section features">
                            <h3>Cabinet Features</h3>
                            <div class="feature-grid">
                                <label class="feature-label">
                                    <input
                                        type="checkbox"
                                        bind:checked={
                                            search.criteresElemination
                                                .acceptsUrgent
                                        }
                                    />
                                    <span class="feature-icon"
                                        ><AlarmSmoke /></span
                                    >
                                    <span class="feature-text"
                                        >Accepts Urgent</span
                                    >
                                </label>
                                <label class="feature-label">
                                    <input
                                        type="checkbox"
                                        bind:checked={
                                            search.criteresElemination
                                                .acceptsInsurance
                                        }
                                    />
                                    <span class="feature-icon"
                                        ><CreditCard /></span
                                    >
                                    <span class="feature-text"
                                        >Accepts Insurance</span
                                    >
                                </label>
                                <label class="feature-label">
                                    <input
                                        type="checkbox"
                                        bind:checked={
                                            search.criteresElemination
                                                .hasParking
                                        }
                                    />
                                    <span class="feature-icon"
                                        ><ParkingCircle /></span
                                    >
                                    <span class="feature-text">Has Parking</span
                                    >
                                </label>
                                <label class="feature-label">
                                    <input
                                        type="checkbox"
                                        bind:checked={
                                            search.criteresElemination.hasWifi
                                        }
                                    />
                                    <span class="feature-icon"><Wifi /></span>
                                    <span class="feature-text">Has Wifi</span>
                                </label>
                                <label class="feature-label">
                                    <input
                                        type="checkbox"
                                        bind:checked={
                                            search.criteresElemination
                                                .hasAccessHandicap
                                        }
                                    />
                                    <span class="feature-icon"
                                        ><Accessibility /></span
                                    >
                                    <span class="feature-text">Accessible</span>
                                </label>

                                <label class="feature-label">
                                    <input
                                        type="checkbox"
                                        bind:checked={
                                            search.criteresElemination.isOpen
                                        }
                                    />
                                    <span class="feature-icon"><Clock /></span>
                                    <span class="feature-text">Open</span>
                                </label>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="cabinets-grid">
                {#each filteredCabinets as cabinet (cabinet.id)}
                    <a
                        href={`/cabinets/${cabinet.id}`}
                        target="_blank"
                        class="cabinet-wrapper"
                        transition:scale
                    >
                        <CabinetCard {cabinet} />
                    </a>
                {/each}
            </div>

            {#if filteredCabinets.length === 0}
                <div class="no-results">
                    No cabinets found with the current search and filters.
                </div>
            {/if}
        </div>
    </View>
</Face>

<style>
    .cabinet-wrapper {
        text-decoration: none;
        width: 100%;
    }

    .specialities-section {
        display: block;
    }

    .specialities > label {
        width: fit-content;
    }

    .filters-container {
        background-color: var(--background-secondary);
        padding: 1.5rem;
        margin-top: 1rem;
        border-radius: 1.25rem;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
    }

    .filters-section {
        margin-bottom: 1.5rem;
    }

    .filters-section:last-child {
        margin-bottom: 0;
    }

    .filters-section h3 {
        font-size: 0.9rem;
        font-weight: 600;
        color: var(--text-color);
        margin: 0 0 1rem 0;
        opacity: 0.8;
    }

    .sort-controls {
        display: flex;
        gap: 0.75rem;
        flex-wrap: wrap;
    }

    .select-field {
        flex: 1;
        min-width: 150px;
        padding: 0.75rem 1rem;
        border: 1px solid var(--color-primary-alpha);
        border-radius: 0.75rem;
        background-color: var(--background-primary);
        color: var(--text-color);
        font-size: 0.875rem;
        cursor: pointer;
        transition: all 0.2s ease;
        appearance: none;
        background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke='currentColor'%3E%3Cpath stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M19 9l-7 7-7-7'%3E%3C/path%3E%3C/svg%3E");
        background-repeat: no-repeat;
        background-position: right 0.75rem center;
        background-size: 1rem;
        padding-right: 2.5rem;
    }

    .select-field:hover {
        border-color: var(--color-primary);
    }

    .select-field:focus {
        outline: none;
        border-color: var(--color-primary);
        box-shadow: 0 0 0 2px var(--color-primary-alpha);
    }

    .order-select {
        max-width: 120px;
    }

    .feature-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
        gap: 1rem;
    }

    .feature-label {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        padding: 0.75rem 1rem;
        background-color: var(--background-primary);
        border: 1px solid var(--color-primary-alpha);
        border-radius: 0.75rem;
        cursor: pointer;
        transition: all 0.2s ease;
    }

    .feature-label:hover {
        border-color: var(--color-primary);
        transform: translateY(-1px);
    }

    .feature-icon {
        font-size: 1.25rem;
        opacity: 0.8;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .feature-text {
        font-size: 0.875rem;
        color: var(--text-color);
    }

    input[type="checkbox"] {
        width: 1.125rem;
        height: 1.125rem;
        border-radius: 0.25rem;
        border: 2px solid var(--color-primary-alpha);
        appearance: none;
        cursor: pointer;
        position: relative;
        transition: all 0.2s ease;
    }

    input[type="checkbox"]:checked {
        background-color: var(--color-primary);
        border-color: var(--color-primary);
    }

    input[type="checkbox"]:checked::after {
        content: "✓";
        color: white;
        font-size: 0.75rem;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }

    input[type="checkbox"]:focus {
        outline: none;
        box-shadow: 0 0 0 2px var(--color-primary-alpha);
    }

    .list-header {
        display: flex;
        gap: 1rem;
        align-items: center;
        margin-bottom: 1.5rem;
    }

    .search {
        flex: 1;
    }

    .cabinets-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
        gap: 1rem;
        margin-top: 1rem;
    }

    .no-results {
        text-align: center;
        color: var(--text-muted);
        padding: 2rem 0;
    }

    .page-inner {
        padding-bottom: 6rem;
    }

    @media (max-width: 768px) {
        .list-header {
            flex-direction: column;
            align-items: stretch;
        }

        .filters-container {
            padding: 1rem;
        }

        .filters-section {
            margin-bottom: 1.25rem;
        }

        .sort-controls {
            flex-direction: column;
        }

        .select-field {
            width: 100%;
        }

        .order-select {
            max-width: 100%;
        }

        .feature-grid {
            grid-template-columns: 1fr;
            gap: 0.5rem;
        }

        .feature-label {
            padding: 0.625rem 0.875rem;
        }

        .cabinets-grid {
            grid-template-columns: 1fr;
            gap: 0.75rem;
        }
    }

    @media (min-width: 769px) and (max-width: 1024px) {
        .feature-grid {
            grid-template-columns: repeat(2, 1fr);
        }
    }

    @media (min-width: 1025px) {
        .feature-grid {
            grid-template-columns: repeat(3, 1fr);
        }
    }

    @media (prefers-reduced-motion: reduce) {
        .feature-label {
            transform: none;
        }

        .feature-label:hover {
            transform: none;
        }
    }
</style>

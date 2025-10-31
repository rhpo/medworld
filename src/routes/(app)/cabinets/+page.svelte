<script lang="ts">
    import type { Cabinet } from "$lib/types/cabinet";
    import Block from "$lib/components/ui/Block.svelte";
    import View from "$lib/components/ui/View.svelte";
    import Face from "$lib/components/ui/Face.svelte";
    import SearchInput from "$lib/components/ui/SearchInput.svelte";
    import { MapPin, Phone, Users as UsersIcon, Plus } from "@lucide/svelte";
    import { fakeCabinets } from "$lib/types/fakedata";

    // cabinets list
    let allCabinets: Cabinet[] = fakeCabinets;
    let searchTerm = "";
    let filteredCabinets: Cabinet[] = allCabinets;

    $: filteredCabinets = searchTerm
        ? allCabinets.filter((c) => {
              const term = searchTerm.toLowerCase();
              return (
                  c.name.toLowerCase().includes(term) ||
                  c.address.toLowerCase().includes(term) ||
                  c.admin.getFullName().toLowerCase().includes(term)
              );
          })
        : allCabinets;

    function handleSearch(term: string) {
        searchTerm = term;
    }
</script>

<Face title="Cabinets" description="Browse and manage cabinets">
    <View>
        <div class="page-inner">
            <div class="list-header">
                <div class="search">
                    <SearchInput
                        placeholder="Search cabinets by name, address or admin..."
                        onSearch={handleSearch}
                        bind:value={searchTerm}
                    />
                </div>
            </div>

            <div class="cabinets-grid">
                {#each filteredCabinets as cabinet (cabinet.id)}
                    <Block>
                        <a
                            href={`/cabinets/${cabinet.id}`}
                            class="cabinet-card"
                        >
                            <h3 class="cabinet-name">{cabinet.name}</h3>
                            <div class="card-body">
                                <div class="line">
                                    <MapPin size={16} />
                                    <span>{cabinet.address}</span>
                                </div>
                                <div class="line">
                                    <Phone size={16} />
                                    <span>{cabinet.phone}</span>
                                </div>
                                <div class="line">
                                    <UsersIcon size={16} />
                                    <span
                                        >Admin: {cabinet.admin.getFullName()}</span
                                    >
                                </div>
                            </div>
                        </a>
                    </Block>
                {/each}
            </div>

            {#if filteredCabinets.length === 0}
                <div class="no-results">
                    No cabinets found for "{searchTerm}"
                </div>
            {/if}
        </div>
    </View>
</Face>

<style>
    .list-header {
        display: flex;
        gap: 1rem;
        align-items: center;
        margin-bottom: 1.5rem;
    }

    .search {
        flex: 1;
        max-width: 600px;
        margin: 0 auto;
    }

    .cabinets-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
        gap: 1rem;
        margin-top: 1rem;
    }

    .cabinet-card {
        display: flex;
        flex-direction: column;
        text-decoration: none;
        color: inherit;
        padding: 1rem;
        border-radius: 8px;
        transition: background-color 0.2s;
    }

    .cabinet-card:hover {
        background-color: var(--background-hover);
    }

    .cabinet-name {
        margin: 0 0 1rem 0;
        font-size: 1.2rem;
        color: var(--text-primary);
    }

    .card-body .line {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        color: var(--text-muted);
        margin-bottom: 0.75rem;
        font-size: 0.9rem;
    }

    .card-body .line:last-child {
        margin-bottom: 0;
    }

    .no-results {
        text-align: center;
        color: var(--text-muted);
        padding: 2rem 0;
    }

    .page-inner {
        padding-bottom: 6rem;
    }

    .card-body .line :global(svg) {
        color: var(--primary-color);
        flex-shrink: 0;
    }

    @media (max-width: 768px) {
        .list-header {
            flex-direction: column;
            align-items: stretch;
        }

        .cabinets-grid {
            grid-template-columns: 1fr;
            gap: 0.75rem;
        }

        .cabinet-card {
            padding: 0.75rem;
        }

        .cabinet-name {
            font-size: 1.1rem;
            margin-bottom: 0.75rem;
        }
    }
</style>

<script lang="ts">
    import { Search } from "@lucide/svelte";

    let {
        placeholder = "Rechercher un livre...",
        onSearch = () => {},
        value = $bindable(""),
        searchInput = $bindable(),
    } = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<main onclick={() => searchInput.focus()}>
    <input
        type="text"
        {placeholder}
        bind:this={searchInput}
        bind:value
        onkeydown={(e) => {
            if (e.key === "Enter") {
                onSearch(value);
            }
        }}
    />

    <button onclick={() => onSearch(value)}>
        <Search />
    </button>
</main>

<style>
    :root {
        --padding-search: 2rem;
    }

    :global(html[data-rtl="true"]) main input {
        padding-right: var(--padding-search);
        padding-left: 0;
    }

    main {
        width: 100%;
        max-width: 56rem;
        display: inline-flex;
        align-items: center;

        background-color: var(--background-secondary);

        border-radius: 50px;

        cursor: text;

        height: 60px;
        padding: 0.3rem 0;
    }

    main input {
        flex: 1;
        height: 100%;
        width: 100%;
        padding-left: 1.5rem;
        font-size: 1.2rem;
        background-color: transparent;
        border: none;
    }

    main input:focus {
        outline: none;
    }

    main input::placeholder {
        font-family: var(--font-secondary);
    }

    button {
        display: grid;
        place-items: center;
        cursor: pointer;

        height: 100%;
        aspect-ratio: 1 / 1;
        border-radius: 50%;
        margin: calc(var(--padding-search) / 4);

        border: none;
        background-color: var(--background-primary);
        color: var(--color-primary-dark);

        transition: all var(--transition-duration) var(--transition-easing);
    }

    button:hover {
        transform: scale(1.05) rotate(5deg);
    }

    @media screen and (max-width: 600px) {
        :root {
            --padding-search: 1rem;
        }

        main input {
            font-size: 0.875rem;
        }
    }
</style>

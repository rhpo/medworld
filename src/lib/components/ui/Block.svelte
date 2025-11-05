<script lang="ts">
    import { currentBlock } from "$lib/stores";
    import type { Group } from "$lib/types/permission";
    import { ArrowLeft } from "@lucide/svelte";
    import { blur, fly, scale } from "svelte/transition";

    interface IProps {
        title?: string;
        group: Group;
        Icon?: any;
        children?: any;
    }

    let { title = "", children, group, Icon }: IProps = $props();
</script>

{#if $currentBlock === group}
    <main>
        <button class="back" onclick={() => currentBlock.set(null)}>
            <div class="back-icon">
                <ArrowLeft />
            </div>
            <h2>{title}</h2>
        </button>

        <div class="container">
            {#if children}
                {@render children()}
            {/if}
        </div>
    </main>
{:else if $currentBlock === null}
    <button
        class="card"
        onclick={() => currentBlock.set(group)}
        transition:scale
    >
        {#if Icon}
            <Icon />
        {/if}

        <h3>
            {title}
        </h3>
    </button>
{/if}

<style>
    .back {
        display: flex;
        align-items: center;
        gap: 1rem;
        background: none;
        border: none;
        cursor: pointer;
        margin-bottom: 1.5rem;
        padding: 0;
    }

    .back h2 {
        font-family: var(--font-secondary);
        font-size: 2rem;
        font-weight: 400;
        margin: 0;
    }

    .back-icon {
        display: flex;
        align-items: center;
        justify-content: center;
        transition: transform 0.2s ease;
    }

    .back:hover .back-icon {
        transform: translateX(-4px);
    }

    .back:hover h2 {
        text-decoration: underline;
    }

    .card {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        padding: 1.5rem;
        border-radius: 1rem;
        background-color: var(--color-primary-dark);
        box-shadow:
            0 4px 6px -1px rgba(0, 0, 0, 0.1),
            0 2px 4px -1px rgba(0, 0, 0, 0.06);
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.18);
        cursor: pointer;
        transition: all var(--transition-duration) var(--transition-easing);
        width: 100%;
        height: 100%;
    }

    .card :global(*) {
        color: white;
    }

    .card h3 {
        margin: 1rem 0 0 0;
        font-size: 1.5rem;
        font-weight: 400;
    }

    .card:hover {
        transform: translateY(-2px);
        box-shadow:
            0 10px 15px -3px rgba(0, 0, 0, 0.1),
            0 4px 6px -2px rgba(0, 0, 0, 0.05);
        background-color: var(--color-primary);
    }

    main {
        display: flex;
        flex-direction: column;
        gap: 0;
    }

    h2 {
        font-family: var(--font-secondary);
        margin: 0;
        font-size: 2.5rem;
        font-weight: 300;
    }

    .container {
        background-color: rgba(255, 255, 255, 0.8);
        border-radius: 1rem;
        padding: 2rem;
        box-shadow:
            0 4px 6px -1px rgba(0, 0, 0, 0.1),
            0 2px 4px -1px rgba(0, 0, 0, 0.06);
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.18);
        transition:
            transform 0.2s ease,
            box-shadow 0.2s ease;
    }

    .container:hover {
        transform: translateY(-2px);
        box-shadow:
            0 10px 15px -3px rgba(0, 0, 0, 0.1),
            0 4px 6px -2px rgba(0, 0, 0, 0.05);
    }

    @media (max-width: 768px) {
        .back h2 {
            font-size: 1.1rem;
            font-weight: 500;
        }

        .container {
            padding: 0;
            background-color: transparent;
            box-shadow: none;
            overflow-x: auto;
        }

        .container:hover {
            padding: 0;
            background-color: transparent;
            box-shadow: none;
            overflow-x: auto;
        }
    }
</style>

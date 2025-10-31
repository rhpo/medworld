<script lang="ts">
    interface IProps {
        title?: string;
        children?: any;
    }

    let { title = "", children }: IProps = $props();
</script>

<main>
    {#if title}
        <h2>{title}</h2>
    {/if}

    <div class="container">
        {#if children}
            {@render children()}
        {/if}
    </div>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        gap: 0;
    }

    h2 + .container {
        margin-top: 2rem;
    }

    h2 {
        font-family: var(--font-secondary);
        margin: 0;
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
</style>

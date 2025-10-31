<script lang="ts">
    import type { Snippet } from "svelte";

    interface Props {
        type?: "primary" | "secondary" | "error";
        href?: string;
        label?: string;
        autoWidth?: boolean;
        Icon?: any;
        onClick?: () => void;
        children?: Snippet;
        selected?: boolean;
        [key: string]: any;
    }

    let {
        type = "primary",
        href = "",
        label = "",
        autoWidth = false,
        Icon,
        onClick = () => {},
        selected = false,
        children,
        ...rest
    }: Props = $props();
</script>

{#if href}
    <a
        class="button"
        aria-label="Panier"
        class:auto-width={autoWidth}
        {href}
        class:selected
        {...rest}
    >
        {@render children?.()}

        {#if Icon}
            <Icon />
        {/if}
    </a>
{:else}
    <button
        class="button {type}"
        aria-label="Panier"
        class:auto-width={autoWidth}
        onclick={onClick}
        class:selected
        {...rest}
    >
        {@render children?.()}

        {#if Icon}
            <Icon />
        {/if}
    </button>
{/if}

<style>
    :root {
        --border-icon: 2px;
        --padding-icon: calc(1rem - var(--border-icon));
    }

    .button {
        text-decoration: none;
        padding: var(--padding-icon);
        border-radius: var(--border-radius-md);
        transition: all var(--transition-normal);
        position: relative;
        display: flex;
        align-items: center;
        justify-content: center;
        background: transparent;
        color: var(--white);
        padding: var(--padding-icon);
        border-radius: var(--border-radius-md);
        cursor: pointer;
        transition: all var(--transition-normal);
        display: flex;
        align-items: center;
        justify-content: center;
        border: none;
    }

    .button:hover {
        background: var(--color-secondary);
        color: var(--text-primary);
        transform: translateY(-2px);
    }

    .button.primary {
        border: var(--border-icon) solid var(--border);
    }

    .button:hover {
        border-color: var(--primary);
        color: var(--primary);
        transform: translateY(-2px);
    }

    .button.selected {
        background: var(--primary) !important;
        color: var(--primary) !important;
        transform: scale(1.05) rotate(5deg) translateY(-2px);
        border: var(--border-icon) solid var(--primary) !important;
    }

    .button.selected :global(svg) {
        fill: var(--white) !important;
        transform: scale(1.2);
        transition: all var(--transition-normal);
    }

    .button.selected:hover {
        background: var(--primary);
        color: var(--white) !important;
        transform: scale(1.05) rotate(5deg) translateY(-2px);
    }

    .button.secondary {
        background: var(--color-secondary);
        color: var(--white);
    }

    .button.secondary:hover {
        background: var(--primary-dark);
        color: var(--white);
    }

    .button.error {
        background: var(--error);
        color: var(--white);
    }

    .button.error:hover {
        background: var(--error-dark);
        color: var(--white);
    }

    /*
	@media screen and (max-width: 768px) {
		.button.selected {
		}
	} */
</style>

<script lang="ts">
    interface Props {
        category?: "textarea" | "input";
        theme?: "primary" | "secondary";
        type?: string;
        label?: string;
        value?: string;
        placeholder?: string;
        error?: string;
        required?: boolean;
        name?: string;
    }

    let {
        category = "input",
        type = "text",
        label = "",
        theme = "primary",
        value = $bindable(""),
        placeholder = "",
        error = "",
        required = false,
        name = "",
    }: Props = $props();
</script>

<main>
    {#if label}
        <label for={label.toLowerCase().replace(/\s+/g, "-")}>
            {label}
        </label>
    {/if}

    {#if category === "textarea"}
        <textarea
            class="input"
            id={label.toLowerCase().replace(/\s+/g, "-")}
            bind:value
            class:primary={theme === "primary"}
            class:secondary={theme === "secondary"}
            {placeholder}
            class:error
            {required}
            {name}
        ></textarea>
    {:else}
        <input
            class="input"
            id={label.toLowerCase().replace(/\s+/g, "-")}
            {type}
            class:primary={theme === "primary"}
            class:secondary={theme === "secondary"}
            bind:value
            {placeholder}
            class:error
            {required}
            {name}
        />
    {/if}
    {#if error}
        <span class="error-message">{error}</span>
    {/if}
</main>

<style>
    main {
        border-bottom: 1px solid var(--border-color);
        margin-bottom: 1.5rem;
    }

    label {
        display: block;
        margin-bottom: 0.5rem;
        font-weight: 500;
        color: var(--text-color);
        font-size: 13px;
        text-transform: uppercase;
        letter-spacing: 0.05em;
    }

    .input.primary {
        width: 100%;
        padding: 16px 0;
        border: none;

        font-size: 16px;
        transition: var(--transition);
        background: transparent;
        font-family: var(--font-secondary);
        color: var(--text-color);
    }

    .input.primary:focus {
        outline: none;
        border-bottom-color: var(--text-color);
    }

    .input.primary::placeholder {
        color: var(--text-secondary);
    }

    .input.secondary {
        width: 100%;
        padding: 12px 16px;
        border: 1px solid var(--border-color);
        border-radius: 8px;

        font-size: 16px;
        transition: var(--transition);
        background: var(--background-secondary);
        font-family: var(--font-secondary);
        color: var(--text-color);
    }

    .input.secondary:focus {
        outline: none;
        border-color: var(--text-color);
    }

    .input.secondary::placeholder {
        color: var(--text-secondary);
    }

    .input.error {
        border-bottom-color: #e74c3c;
    }

    .error-message {
        color: #e74c3c;
        font-size: 12px;
        margin-top: 0.5rem;
        display: block;
    }
</style>

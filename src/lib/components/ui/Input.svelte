<script lang="ts">
    import tippy from "sveltejs-tippy";

    interface Props {
        category?: "textarea" | "input" | "display";
        theme?: "primary" | "secondary";
        type?: string;
        label?: string;
        value?: string;
        placeholder?: string;
        error?: string;
        required?: boolean;
        name?: string;
        showLabel?: boolean;
        showTooltip?: boolean;
        onInput?: any;
    }

    let {
        category = "input",
        type = "text",
        label = "",
        showLabel = false,
        showTooltip = false,
        theme = "primary",
        value = $bindable(""),
        placeholder = "",
        error = "",
        required = false,
        name = "",
        onInput = () => {},
    }: Props = $props();

    label = label || placeholder || "Enter text here...";
</script>

<main>
    {#if label && showLabel}
        <label for={label.toLowerCase().replace(/\s+/g, "-")}>
            {label}
        </label>
    {/if}

    {#if category === "textarea"}
        <textarea
            use:tippy={{ content: label, placement: "auto" }}
            class="input"
            id={label.toLowerCase().replace(/\s+/g, "-")}
            bind:value
            class:primary={theme === "primary"}
            class:secondary={theme === "secondary"}
            {placeholder}
            oninput={onInput}
            class:error
            {required}
            {name}
        ></textarea>
    {:else if category === "input"}
        <input
            class="input"
            use:tippy={{ content: label, placement: "auto" }}
            id={label.toLowerCase().replace(/\s+/g, "-")}
            {type}
            class:primary={theme === "primary"}
            class:secondary={theme === "secondary"}
            bind:value
            {placeholder}
            oninput={onInput}
            class:error
            {required}
            {name}
        />
    {:else}
        <p>{value}</p>
    {/if}
    {#if error}
        <span class="error-message">{error}</span>
    {/if}
</main>

<style>
    main {
        margin-bottom: 1.5rem;
        text-align: left !important;
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

        border-bottom: 1px solid var(--border-color);
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
        border: 2px solid var(--border-color);
        border-radius: 8px;

        font-size: 16px;
        transition: var(--transition);
        background: var(--background-secondary);
        font-family: var(--font-secondary);
        color: var(--text-color);
    }

    .input.secondary:focus {
        outline: none;
        border-color: 3px solid var(--border-color);
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

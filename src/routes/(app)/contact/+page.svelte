<script lang="ts">
    import { SITE_NAME } from "$lib";

    import Button from "$lib/components/ui/Button.svelte";
    import Input from "$lib/components/ui/Input.svelte";
    import View from "$lib/components/ui/View.svelte";

    import { Phone, Plane, Send } from "@lucide/svelte";

    let data = $state({
        name: "",
        email: "",
        phone: "",
        message: "",
    });

    let submitStatus = $state({
        loading: false,
        success: false,
        error: null,
    });

    function handleSubmit(event: Event) {
        event.preventDefault();
        const form = event.target as HTMLFormElement;
        const formData = new FormData(form);

        submitStatus.loading = true;

        fetch("/api/submit", {
            method: "POST",
            body: formData,
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                return response.json();
            })
            .then((data) => {
                submitStatus.loading = false;
                submitStatus.success = true;
                console.log("Form submitted successfully:", data);
            })
            .catch((error) => {
                submitStatus.loading = false;
                submitStatus.error = error.message;
                console.error("Error submitting form:", error);
            });
    }
</script>

<svelte:head>
    <title>Contact Us | {SITE_NAME}</title>
    <meta
        name="description"
        content="Contact us for any questions or requests for quotes. We are here to help."
    />
</svelte:head>

<main>
    <View align="column" center>
        <div class="hero-section">
            <Phone size="5rem" style="margin-bottom: 2rem;" />
            <h1 data-aos="fade-up" data-aos-duration="800">
                Contact {SITE_NAME}
            </h1>
            <p data-aos="fade-up" data-aos-duration="800">
                Feel free to contact us for any questions or requests for
                quotes, or via Whatsapp or phone.
            </p>
        </div>

        <form
            onsubmit={handleSubmit}
            class="contact-form"
            data-aos="fade-up"
            data-aos-duration="800"
        >
            <Input
                type="text"
                name="name"
                required
                bind:value={data.name}
                label="Name"
                placeholder="Your Name"
            />

            <Input
                type="email"
                name="email"
                required
                bind:value={data.email}
                label="Email"
                placeholder="your@email.com"
            />

            <Input
                type="tel"
                name="phone"
                required
                bind:value={data.phone}
                label="Phone"
                placeholder="+213 678 123 456"
            />

            <Input
                category="textarea"
                bind:value={data.message}
                name="message"
                required
                label="Message"
                placeholder="Your message here..."
            />

            <Button
                variant="primary"
                style="width: 100%;"
                large
                iconPosition="right"
                Icon={Send}
                type="submit"
                fullWidth>Send</Button
            >

            <p class="status">
                {#if submitStatus.loading}
                    <span class="loading">Sending...</span>
                {:else if submitStatus.success}
                    <span class="success">Message sent successfully!</span>
                {:else if submitStatus.error}
                    <span class="error">Error: {submitStatus.error}</span>
                {/if}
            </p>
        </form></View
    >
</main>

<style>
    .status {
        text-align: center;
        margin-top: 1rem;
        font-size: 0.875rem;
    }

    main {
        min-height: 100vh;
        padding-top: 6rem;
        padding-bottom: 4rem;
    }

    form {
        max-width: 1000px;
        width: 100%;

        background-color: var(--background-light);
        padding: 2rem;
    }

    .hero-section {
        text-align: center;
        max-width: 800px;

        color: var(--color-primary-dark);
    }

    .hero-section h1 {
        font-size: clamp(2.5rem, 8vw, 4rem);
        font-weight: 600;
        margin-bottom: 1.5rem;
        color: var(--text-color);
        font-family: var(--font-primary);
    }

    .hero-section p {
        font-size: 1.125rem;
        color: var(--text-light);
        line-height: 1.7;
        max-width: 600px;
        margin: 0 auto;
    }

    @media (max-width: 768px) {
        main {
            padding-top: calc(var(--header-height) + 1rem);
            padding-bottom: 2rem;
        }

        .hero-section {
            margin-bottom: 2.5rem;
            padding: 0 0.5rem;
        }

        .hero-section h1 {
            font-size: 2.5rem;
        }

        .hero-section p {
            font-size: 1rem;
        }
    }

    @media (max-width: 480px) {
        .hero-section h1 {
            font-size: 2rem;
        }
    }
</style>

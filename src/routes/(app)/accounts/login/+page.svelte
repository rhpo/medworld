<script lang="ts">
    import Button from "$lib/components/ui/Button.svelte";
    import Input from "$lib/components/ui/Input.svelte";

    let data = {
        email: "",
        password: "",
    };

    function logIn() {
        window.location.href = "/dashboard";
    }
</script>

<main class="auth-container">
    <!-- Left Side -->
    <section class="left">
        <div class="overlay"></div>
    </section>

    <!-- Right Side -->
    <section class="right">
        <div class="form-container">
            <h2>Log in to your account</h2>

            <Input
                bind:value={data.email}
                placeholder="Email"
                type="email"
                theme="secondary"
                required
            />

            <Input
                bind:value={data.password}
                placeholder="Password"
                type="password"
                required
            />

            <Button type="submit" category="primary">Log In</Button>

            <p class="switch">
                Don't have an account?
                <a href="/accounts/create">Create one</a>.
            </p>
        </div>
    </section>
</main>

<style>
    .auth-container {
        display: flex;
        min-height: calc(
            100vh - var(--nav-height) - var(--notification-height)
        );
    }

    /* Left side */
    .left {
        flex: 1;
        background: url("/cabinet.webp") center/cover no-repeat;
        position: relative;
        color: white;
    }

    .overlay {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        text-align: center;
    }

    /* Right side */
    .right {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .form-container {
        width: 100%;
        max-width: 380px;
        text-align: center;
    }

    form {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        margin-top: 2rem;
    }

    .switch {
        margin-top: 1.5rem;
        font-size: 0.9rem;
    }

    .switch a {
        background: none;
        border: none;
        color: var(--color-primary);
        cursor: pointer;
    }

    @media (max-width: 768px) {
        .auth-container {
            flex-direction: column;
        }
        .left,
        .right {
            flex: none;
            width: 100%;
            height: auto;
        }
    }
</style>

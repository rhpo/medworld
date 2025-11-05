<script lang="ts">
    import Button from "$lib/components/ui/Button.svelte";
    import Input from "$lib/components/ui/Input.svelte";
    import { User } from "@lucide/svelte";

    let data = {
        email: "",
        password: "",
    };

    function logIn() {
        window.location.href = "/dashboard";
    }
</script>

<main class="auth-container" data-aos="fade-up" data-aos-duration="1200">
    <section class="left">
        <div class="overlay"></div>
    </section>

    <section class="right" data-aos="flip-left" data-aos-duration="1200">
        <div class="form-container">
            <User size="72" style="margin-bottom: 1rem;" />

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
                theme="secondary"
                required
            />

            <Button
                type="submit"
                category="primary"
                Icon={User}
                href="/dashboard"
                style="width: 100%">Log In</Button
            >

            <p class="switch">
                Don't have an account?
                <a href="/accounts/create">Create one</a>.
                <br />
                <a href="/dashboard">Goto Dashboard</a>
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
        width: 100%;
        height: 100%;

        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;

        background: linear-gradient(
            to bottom,
            rgba(0, 0, 0, 0),
            rgba(0, 0, 0, 0.3)
        );
    }

    /* Right side */
    .right {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .form-container h2 {
        margin-bottom: 2rem;
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

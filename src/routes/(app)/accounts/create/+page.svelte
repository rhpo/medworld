<script lang="ts">
    import Button from "$lib/components/ui/Button.svelte";
    import Input from "$lib/components/ui/Input.svelte";
    import { Users, type User } from "$lib/types/users";
    import { User as UserIcon, UserPlus } from "@lucide/svelte";

    let data: User<Users.Patient> = $state({
        id: 0,
        firstName: "",
        lastName: "",
        email: "",
        type: Users.Patient,
        phoneNumber: "",

        getFullName() {
            return `${this.firstName} ${this.lastName}`;
        },
    });

    function signUp() {
        window.location.href = "/dashboard";
    }
</script>

<main class="auth-container" data-aos="fade-up" data-aos-duration="1200">
    <!-- Left Side -->
    <section class="left">
        <div class="overlay"></div>
    </section>

    <!-- Right Side -->
    <section class="right" data-aos="flip-right" data-aos-duration="1200">
        <div class="form-container">
            <UserPlus size="72" style="margin-bottom: 1rem;" />

            <h2>Create your account</h2>

            <Input
                bind:value={data.firstName}
                placeholder="First Name"
                type="text"
                theme="secondary"
                required
            />

            <Input
                bind:value={data.lastName}
                placeholder="Last Name"
                type="text"
                theme="secondary"
                required
            />

            <Input
                bind:value={data.email}
                placeholder="Email"
                type="email"
                theme="secondary"
                required
            />

            <Input
                bind:value={data.phoneNumber}
                placeholder="Phone Number"
                type="tel"
                theme="secondary"
                required
            />

            <Button
                onClick={() => {
                    signUp();
                }}
                Icon={UserIcon}
                style="width: 100%;"
                category="primary">Sign Up</Button
            >

            <p class="switch">
                Already have an account?
                <a href="/accounts/login">Log in</a>
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

    .right {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: center;

        padding: 2rem 0;
    }

    .form-container {
        width: 100%;
        max-width: 380px;
        text-align: center;
    }

    .form-container h2 {
        margin-bottom: 2rem;
    }

    .switch {
        margin-top: 1.5rem;
        font-size: 0.9rem;
    }

    .switch a {
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

<script lang="ts">
    import { page } from "$app/state";
    import { AllAPI, UserAPI } from "$lib/api";
    import Avatar from "$lib/components/ui/Avatar.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import View from "$lib/components/ui/View.svelte";
    import { user as me } from "$lib/stores";
    import { Users, type IDoctor } from "$lib/types/users";
    import type { Doctor } from "$lib/types/users/doctor";
    import type { Patient } from "$lib/types/users/patient";
    import { Plus } from "@lucide/svelte";

    const userId = Number(page.params.id);
    let result = UserAPI.getById(userId);
</script>

<main>
    <View>
        {#await result}
            <h2>Loading user...</h2>
        {:then user}
            {#if !user}
                <h2>User not found.</h2>
            {:else}
                <main>
                    <div class="avatar">
                        <div class="avatar-wrapper">
                            <div class="profile-picture">
                                <img
                                    src={user.avatarUrl}
                                    alt="Profile"
                                    class="avatar-image"
                                />
                            </div>
                        </div>
                    </div>

                    <section>
                        <div class="form">
                            <!-- Basic Information for All Users -->
                            <div class="form-section">
                                <h3>Basic Information</h3>
                                <div class="form-row">
                                    <div class="form-group">
                                        <label for="firstName"
                                            >First Name *</label
                                        >
                                        {user.firstName}
                                    </div>

                                    <div class="form-group">
                                        <label for="lastName">Last Name *</label
                                        >
                                        {user.lastName}
                                    </div>
                                </div>

                                <div class="form-row">
                                    <div class="form-group">
                                        <label for="email">Email *</label>
                                        {user.email}
                                    </div>

                                    <div class="form-group">
                                        <label for="phoneNumber"
                                            >Phone Number</label
                                        >
                                        {user.phoneNumber}
                                    </div>
                                </div>

                                <div class="form-row">
                                    <div class="form-group">
                                        <label for="dateOfBirth"
                                            >Date of Birth</label
                                        >

                                        {new Date(
                                            user.dateOfBirth,
                                        ).toDateString()}
                                    </div>

                                    <div class="form-group">
                                        <label for="gender">Gender</label>
                                        <p>{(user as Patient).gender}</p>
                                    </div>

                                    <div class="form-group">
                                        <label for="gender">Weiight</label>
                                        <p>
                                            {(user as Patient).weight || "90kg"}
                                        </p>
                                    </div>

                                    <div class="form-group">
                                        <label for="gender"
                                            >Medical Histories</label
                                        >
                                        <ul>
                                            {#if (user as Patient).medicalHistory.length > 0}
                                                {#each (user as Patient).medicalHistory as history}
                                                    <li>
                                                        {history}
                                                    </li>
                                                {/each}
                                            {:else}
                                                <p>None.</p>
                                            {/if}
                                        </ul>
                                    </div>
                                </div>

                                <div class="form-group">
                                    <label for="address">Address</label>
                                    {user.address}
                                </div>
                            </div>

                            {#if user.type === Users.Doctor}
                                <div class="form-section">
                                    <h3>Professional Information</h3>
                                    <div class="form-row">
                                        <div class="form-group">
                                            <label for="specialization"
                                                >Specialization</label
                                            >
                                            <p>
                                                {(user as IDoctor).speciality}
                                            </p>
                                        </div>
                                    </div>

                                    <div class="form-group">
                                        <label for="yearsOfExperience"
                                            >Years of Experience</label
                                        >

                                        {(
                                            user as Doctor
                                        ).getYearsOfExperience()}
                                    </div>
                                </div>
                            {/if}

                            <!-- Patient Specific Fields -->
                            {#if user.type === Users.Patient}
                                <div class="form-section">
                                    <h3>Medical Information</h3>
                                    <div class="form-row">
                                        <div class="form-group">
                                            <label for="bloodType"
                                                >Blood Type</label
                                            >
                                            <p>{(user as Patient).bloodType}</p>
                                        </div>

                                        <div class="form-group">
                                            <label for="emergencyContact"
                                                >Emergency Contact</label
                                            >

                                            {(user as Patient).emergencyContact}
                                        </div>
                                    </div>

                                    <div class="form-group">
                                        <label>Allergies</label>
                                        <div class="chip-list">
                                            {#each (user as Patient).allergies as allergy, index}
                                                <div class="chip">
                                                    {allergy}
                                                </div>
                                            {/each}
                                        </div>
                                    </div>
                                </div>
                            {/if}
                        </div>

                        {#if $me && $me.type === Users.Doctor}
                            <div class="actions">
                                <Button
                                    href="/dashboard/users/{user.id}/modify"
                                    target="_blank"
                                    Icon={Plus}>Add Medical History</Button
                                >
                            </div>
                        {/if}
                    </section>
                </main>
            {/if}
        {/await}
    </View>
</main>

<style>
    .avatar {
        position: sticky;
        top: calc(var(--nav-height) + 5rem);

        aspect-ratio: 1 / 1;
        width: 100%;
        height: 100%;
        max-width: 400px;
        max-height: 400px;
    }

    .avatar-wrapper {
        position: relative;
        width: 100%;
        height: 100%;

        flex: 1;
        border-radius: 50%;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .profile-picture {
        width: 100%;
        height: 100%;
        border-radius: 50%;
        position: relative;
    }

    .avatar-image {
        width: 100%;
        height: 100%;
        border-radius: 50%;
        object-fit: cover;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: #222;
        font-weight: 600;
        font-size: 1.5rem;
        color: #fff;
    }

    main {
        display: flex;
        gap: 4rem;
        padding-bottom: 4rem;
        padding-top: 2rem;

        position: relative;
    }

    .avatar {
        position: sticky;
        top: calc(var(--nav-height) + 5rem);

        aspect-ratio: 1 / 1;
        width: 100%;
        height: 100%;
        max-width: 400px;
    }

    .avatar-wrapper {
        position: relative;
        width: 100%;
        height: 100%;

        flex: 1;
        border-radius: 50%;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .profile-picture {
        width: 100%;
        height: 100%;
        border-radius: 50%;
        position: relative;
    }

    section {
        flex: 1;
    }

    .form-section {
        margin-bottom: 2rem;
        padding: 1.5rem;
        background: var(--background, #f8f9fa);
        border-radius: 8px;
        border: 1px solid var(--border, #e1e8ed);
    }

    .form-section h3 {
        margin: 0 0 1rem 0;
        color: var(--text-primary, #2c3e50);
        font-size: 1.25rem;
        font-weight: 600;
    }

    .form-row {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 1rem;
    }

    .form-group {
        margin-bottom: 1rem;
    }

    label {
        display: block;
        margin-bottom: 0.5rem;
        font-weight: 600;
        color: var(--text-primary, #2c3e50);
    }

    .chip-list {
        display: flex;
        flex-wrap: wrap;
        gap: 0.5rem;
        margin-top: 0.5rem;
    }

    .chip {
        background: var(--primary-color, #3498db);
        color: white;
        padding: 0.25rem 0.75rem;
        border-radius: 20px;
        font-size: 0.8rem;
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .chip-remove {
        background: none;
        border: none;
        color: white;
        cursor: pointer;
        padding: 0;
        font-size: 1rem;
    }

    .actions {
        display: flex;
        gap: 1rem;
        margin-top: 2rem;
    }

    @media (max-width: 768px) {
        main {
            flex-direction: column;
        }

        .profile-picture {
            width: 120px;
            height: 120px;
            margin: 0 auto;
        }

        .form-row {
            grid-template-columns: 1fr;
        }

        .actions {
            flex-direction: column;
        }

        .avatar {
            position: static;
            padding: 0;
        }

        .avatar-wrapper,
        .avatar {
            width: auto;
            height: auto;
            aspect-ratio: unset;
        }
    }
</style>

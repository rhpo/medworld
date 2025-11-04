<script lang="ts">
    import { DoctorAPI, UserAPI, PatientAPI, AllAPI } from "$lib/api";
    import { Users, type User } from "$lib/types/users";
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import { page } from "$app/state";
    import type { Doctor } from "$lib/types/users/doctor";
    import type { Patient } from "$lib/types/users/patient";
    import Input from "$lib/components/ui/Input.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import { Specialities } from "$lib/types/speciality";
    import View from "$lib/components/ui/View.svelte";
    import { Image } from "@lucide/svelte";

    const userId = Number(page.params.id);

    let user: User<any> | null;
    let isLoading = true;
    let errorMessage = "";

    let formData = {
        firstName: "",
        lastName: "",
        email: "",
        phoneNumber: "",
        avatarUrl: "",
        dateOfBirth: new Date(),
        gender: "",

        medicalHistory: [],
        address: "",

        specialization: "",
        licenseNumber: "",
        yearsOfExperience: "",

        bloodType: "",
        allergies: [] as string[],
        emergencyContact: "",
    };

    const bloodTypes = ["A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"];

    onMount(async () => {
        try {
            if (!userId) throw new Error("User ID is required");

            user = await UserAPI.getById(userId);

            if (user === null) return;

            formData.firstName = user.firstName;
            formData.lastName = user.lastName;
            formData.email = user.email;
            formData.phoneNumber = user.phoneNumber || "";
            formData.avatarUrl = user.avatarUrl || "";
            formData.dateOfBirth = user.dateOfBirth || new Date();
            formData.gender = user.gender || "";
            formData.address = user.address || "";
            (formData.medicalHistory as string[]) =
                (user as Patient).medicalHistory || [];

            if (user.type === Users.Doctor) {
                const doctor = user as Doctor;
                formData.specialization = doctor.speciality || "";
                formData.yearsOfExperience =
                    doctor.getYearsOfExperience()?.toString() || "";
            } else if (user.type === Users.Patient) {
                const patient = user as Patient;
                formData.bloodType = patient.bloodType || "";
                formData.allergies = patient.allergies || [];
                formData.emergencyContact = patient.emergencyContact || "";
            }

            isLoading = false;
        } catch (error: any) {
            errorMessage = error.message;
            isLoading = false;
        }
    });

    function addAllergy() {
        const allergy = prompt("Enter allergy:");
        if (allergy && !formData.allergies.includes(allergy)) {
            formData.allergies = [...formData.allergies, allergy];
        }
    }

    function removeAllergy(index: number) {
        formData.allergies.splice(index, 1);
        formData.allergies = formData.allergies;
    }

    function handleFileChange(e: Event) {
        const target = e.target as HTMLInputElement;
        const file = target.files?.[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = () => {
                formData.avatarUrl = reader.result as string;
            };
            reader.readAsDataURL(file);
        }
    }

    async function handleSubmit() {
        try {
            isLoading = true;

            const updateData: any = {
                firstName: formData.firstName,
                lastName: formData.lastName,
                email: formData.email,
                phoneNumber: formData.phoneNumber,
                avatarUrl: formData.avatarUrl,
                dateOfBirth: formData.dateOfBirth,
                gender: formData.gender,
                address: formData.address,
            };

            if (user?.type === Users.Doctor) {
                updateData.speciality = formData.specialization;
                updateData.yearsOfExperience =
                    parseInt(formData.yearsOfExperience) || 0;
            } else if (user?.type === Users.Patient) {
                updateData.bloodType = formData.bloodType;
                updateData.allergies = formData.allergies;
                updateData.emergencyContact = formData.emergencyContact;
            }

            await UserAPI.UpdateProfile(user!, updateData);

            alert("Information Modified Successfully");
            goto("/dashboard/users");
        } catch (error: any) {
            errorMessage = error.message;
        } finally {
            isLoading = false;
        }
    }
</script>

<View>
    <main>
        <div class="avatar">
            <div class="avatar-wrapper">
                <input
                    type="file"
                    accept="image/*"
                    class="avatar-input"
                    on:change={handleFileChange}
                />
                <div class="profile-picture">
                    {#if formData.avatarUrl}
                        <img
                            src={formData.avatarUrl}
                            alt="Profile"
                            class="avatar-image"
                        />
                    {:else}
                        <div class="avatar-placeholder">
                            {formData.firstName?.charAt(
                                0,
                            )}{formData.lastName?.charAt(0)}
                        </div>
                    {/if}
                </div>
                <div class="avatar-overlay">
                    <Image size="24" />
                    <span>Upload New Image</span>
                </div>
            </div>
        </div>

        <section>
            {#if isLoading}
                <p>Loading...</p>
            {:else if errorMessage}
                <p class="error">{errorMessage}</p>
            {:else if user}
                <form on:submit|preventDefault={handleSubmit}>
                    <!-- Basic Information for All Users -->
                    <div class="form-section">
                        <h3>Basic Information</h3>
                        <div class="form-row">
                            <div class="form-group">
                                <label for="firstName">First Name *</label>
                                <Input
                                    type="text"
                                    bind:value={formData.firstName}
                                    required
                                />
                            </div>

                            <div class="form-group">
                                <label for="lastName">Last Name *</label>
                                <Input
                                    type="text"
                                    bind:value={formData.lastName}
                                    required
                                />
                            </div>
                        </div>

                        <div class="form-row">
                            <div class="form-group">
                                <label for="email">Email *</label>
                                <Input
                                    type="email"
                                    bind:value={formData.email}
                                    required
                                />
                            </div>

                            <div class="form-group">
                                <label for="phoneNumber">Phone Number</label>
                                <Input
                                    type="tel"
                                    bind:value={formData.phoneNumber}
                                />
                            </div>
                        </div>

                        <div class="form-row">
                            <div class="form-group">
                                <label for="dateOfBirth">Date of Birth</label>
                                <Input
                                    type="date"
                                    bind:value={formData.dateOfBirth as any}
                                />
                            </div>

                            <div class="form-group">
                                <label for="gender">Gender</label>
                                <select
                                    class="form-select"
                                    bind:value={formData.gender}
                                >
                                    <option value="">Select Gender</option>
                                    <option value="male">Male</option>
                                    <option value="female">Female</option>
                                </select>
                            </div>
                        </div>

                        <div class="form-group">
                            <label for="address">Address</label>
                            <Input type="text" bind:value={formData.address} />
                        </div>
                    </div>

                    <!-- Doctor Specific Fields -->
                    {#if user.type === Users.Doctor}
                        <div class="form-section">
                            <h3>Professional Information</h3>
                            <div class="form-row">
                                <div class="form-group">
                                    <label for="specialization"
                                        >Specialization</label
                                    >
                                    <select
                                        class="form-select"
                                        bind:value={formData.specialization}
                                    >
                                        <option value=""
                                            >Select Specialization</option
                                        >
                                        {#each Specialities as spec}
                                            <option value={spec}>{spec}</option>
                                        {/each}
                                    </select>
                                </div>

                                <div class="form-group">
                                    <label for="licenseNumber"
                                        >License Number</label
                                    >
                                    <Input
                                        type="text"
                                        bind:value={formData.licenseNumber}
                                    />
                                </div>
                            </div>

                            <div class="form-group">
                                <label for="yearsOfExperience"
                                    >Years of Experience</label
                                >
                                <Input
                                    type="number"
                                    bind:value={formData.yearsOfExperience}
                                />
                            </div>
                        </div>
                    {/if}

                    <!-- Patient Specific Fields -->
                    {#if user.type === Users.Patient}
                        <div class="form-section">
                            <h3>Medical Information</h3>
                            <div class="form-row">
                                <div class="form-group">
                                    <label for="bloodType">Blood Type</label>
                                    <select
                                        class="form-select"
                                        bind:value={formData.bloodType}
                                    >
                                        <option value=""
                                            >Select Blood Type</option
                                        >
                                        {#each bloodTypes as bloodType}
                                            <option value={bloodType}
                                                >{bloodType}</option
                                            >
                                        {/each}
                                    </select>
                                </div>

                                <div class="form-group">
                                    <label for="bloodType">Blood Type</label>
                                    <select
                                        class="form-select"
                                        bind:value={formData.bloodType}
                                    >
                                        <option value=""
                                            >Select Blood Type</option
                                        >
                                        {#each bloodTypes as bloodType}
                                            <option value={bloodType}
                                                >{bloodType}</option
                                            >
                                        {/each}
                                    </select>
                                </div>

                                <div class="form-group">
                                    <label for="emergencyContact"
                                        >Emergency Contact</label
                                    >
                                    <Input
                                        type="text"
                                        bind:value={formData.emergencyContact}
                                    />
                                </div>
                            </div>

                            <div class="form-group">
                                <label for="chip-list">Allergies</label>
                                <div class="chip-list">
                                    {#each formData.allergies as allergy, index}
                                        <div class="chip">
                                            {allergy}
                                            <button
                                                type="button"
                                                class="chip-remove"
                                                on:click={() =>
                                                    removeAllergy(index)}
                                                >x</button
                                            >
                                        </div>
                                    {/each}
                                    <button
                                        type="button"
                                        class="add-chip"
                                        on:click={addAllergy}
                                        >+ Add Allergy</button
                                    >
                                </div>
                            </div>
                        </div>
                    {/if}

                    <!-- Form Actions -->
                    <div class="actions">
                        <Button
                            type="submit"
                            disabled={isLoading}
                            variant="primary"
                        >
                            {isLoading ? "Saving..." : "Save Changes"}
                        </Button>
                        <Button type="button" href="#" category="secondary">
                            Cancel
                        </Button>
                    </div>
                </form>
            {:else}
                <p class="error">User not found</p>
            {/if}
        </section>
    </main>
</View>

<style>
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

    .avatar-input {
        position: absolute;
        inset: 0;
        opacity: 0;
        cursor: pointer;
        z-index: 3;
    }

    .profile-picture {
        width: 100%;
        height: 100%;
        border-radius: 50%;
        position: relative;
    }

    .avatar-image,
    .avatar-placeholder {
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

    .avatar-overlay {
        position: absolute;
        inset: 0;
        background-color: rgba(0, 0, 0, 0.6);
        color: #fff;
        margin: 0.5rem;
        border: 2px dotted rgba(255, 255, 255, 0.5);
        border-radius: 50%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        gap: 6px;
        font-size: 0.85rem;
        font-weight: 500;
        text-align: center;
        opacity: 0.4;
        transition: opacity 0.25s ease;
        z-index: 2;
    }

    .avatar-wrapper:hover .avatar-overlay {
        opacity: 1;
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

    .form-select {
        width: 100%;
        padding: 0.75rem;
        border: 2px solid var(--border, #e1e8ed);
        border-radius: 6px;
        font-size: 1rem;
        background: white;
        color: var(--text-primary, #2c3e50);
    }

    .form-select:focus {
        outline: none;
        border-color: var(--primary-color, #3498db);
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

    .add-chip {
        background: var(--success, #27ae60);
        color: white;
        border: none;
        padding: 0.25rem 0.75rem;
        border-radius: 20px;
        cursor: pointer;
        font-size: 0.8rem;
    }
    .error {
        color: var(--error, #e74c3c);
        background: var(--error-bg, #fee);
        padding: 1rem;
        border-radius: 6px;
        border-left: 4px solid var(--error, #e74c3c);
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

        .avatar-placeholder {
            font-size: 2.5rem;
        }

        .form-row {
            grid-template-columns: 1fr;
        }

        .actions {
            flex-direction: column;
        }
    }
</style>

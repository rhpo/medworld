<script lang="ts">
    import { UserAPI } from "$lib/api";
    import { Users, type IDoctor, type User } from "$lib/types/users";
    import { onMount } from "svelte";
    import { page } from "$app/state";
    import type { Doctor } from "$lib/types/users/doctor";
    import type { Patient } from "$lib/types/users/patient";
    import Input from "$lib/components/ui/Input.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import { Specialities } from "$lib/types/speciality";
    import View from "$lib/components/ui/View.svelte";
    import { Image } from "@lucide/svelte";
    import InputArray from "$lib/components/ui/InputArray.svelte";

    const userId = Number(page.params.id);

    let {
        edit = false,
    }: {
        edit: boolean;
    } = $props();

    let user: User<any> | null = $state(null);
    let isLoading = $state(true);
    let errorMessage = $state("");

    let formData = {
        firstName: "",
        lastName: "",
        email: "",
        phoneNumber: "",
        avatarUrl: "",
        dateOfBirth: new Date(),
        gender: "",

        medicalHistory: [] as string[],
        address: "",

        speciality: "",
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
                formData.speciality = doctor.speciality || "";
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
        if (edit) {
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
                    updateData.speciality = formData.speciality;
                    updateData.yearsOfExperience =
                        parseInt(formData.yearsOfExperience) || 0;
                } else if (user?.type === Users.Patient) {
                    updateData.bloodType = formData.bloodType;
                    updateData.allergies = formData.allergies;
                    updateData.emergencyContact = formData.emergencyContact;
                }

                await UserAPI.UpdateProfile(user!, updateData);

                alert("Information Modified Successfully");
            } catch (error: any) {
                errorMessage = error.message;
            } finally {
                isLoading = false;
            }
        }
    }
</script>

<View>
    <main>
        <div class="avatar">
            <div class="avatar-wrapper">
                {#if edit}
                    <input
                        type="file"
                        accept="image/*"
                        class="avatar-input"
                        onchange={handleFileChange}
                    />
                {/if}
                <div class="profile-picture">
                    {#if user && user.avatarUrl}
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
                {#if edit}
                    <div class="avatar-overlay">
                        <Image size="24" />
                        <span>Upload New Image</span>
                    </div>
                {/if}
            </div>
        </div>

        <section>
            {#if isLoading}
                <p>Loading...</p>
            {:else if errorMessage}
                <p class="error">{errorMessage}</p>
            {:else if user}
                <form
                    onsubmit={(e) => {
                        e.preventDefault();
                        handleSubmit();
                    }}
                >
                    <div class="form-section">
                        <h3>Basic Information</h3>
                        <div class="form-row">
                            <div class="form-group">
                                <Input
                                    label="First Name"
                                    showLabel
                                    category={edit ? "input" : "display"}
                                    type="text"
                                    bind:value={formData.firstName}
                                    required
                                />
                            </div>

                            <div class="form-group">
                                <Input
                                    label="Last Name"
                                    showLabel
                                    category={edit ? "input" : "display"}
                                    type="text"
                                    bind:value={formData.lastName}
                                    required
                                />
                            </div>
                        </div>

                        <div class="form-row">
                            <div class="form-group">
                                <Input
                                    label="Email"
                                    showLabel
                                    category={edit ? "input" : "display"}
                                    type="email"
                                    bind:value={formData.email}
                                    required
                                />
                            </div>

                            <div class="form-group">
                                <Input
                                    label="Phone Number"
                                    showLabel
                                    category={edit ? "input" : "display"}
                                    type="tel"
                                    bind:value={formData.phoneNumber}
                                />
                            </div>
                        </div>

                        <div class="form-row">
                            <div class="form-group">
                                <Input
                                    label="Date of Birth"
                                    showLabel
                                    category={edit ? "input" : "display"}
                                    type="date"
                                    bind:value={formData.dateOfBirth as any}
                                />
                            </div>

                            <div class="form-group">
                                {#if edit}
                                    <select
                                        class="form-select"
                                        bind:value={formData.gender}
                                    >
                                        <option value="">Select Gender</option>
                                        <option value="male">Male</option>
                                        <option value="female">Female</option>
                                    </select>
                                {:else}
                                    <Input
                                        label="Gender"
                                        showLabel
                                        category="display"
                                        value={user.gender}
                                    />
                                {/if}
                            </div>
                        </div>

                        <div class="form-group">
                            <Input
                                label="Address"
                                showLabel
                                category={edit ? "input" : "display"}
                                type="text"
                                bind:value={formData.address}
                            />
                        </div>
                    </div>

                    {#if user.type === Users.Doctor}
                        <div class="form-section">
                            <h3>Professional Information</h3>
                            <div class="form-row">
                                <div class="form-group">
                                    {#if edit}
                                        <select
                                            class="form-select"
                                            bind:value={formData.speciality}
                                        >
                                            <option value=""
                                                >Select Specialization</option
                                            >
                                            {#each Specialities as spec}
                                                <option value={spec}
                                                    >{spec}</option
                                                >
                                            {/each}
                                        </select>
                                    {:else}
                                        <Input
                                            label="Speciality"
                                            showLabel
                                            category="display"
                                            value={(user as IDoctor).speciality}
                                        />
                                    {/if}
                                </div>

                                <div class="form-group">
                                    <Input
                                        label="License Number"
                                        showLabel
                                        category={edit ? "input" : "display"}
                                        type="text"
                                        bind:value={formData.licenseNumber}
                                    />
                                </div>
                            </div>

                            <div class="form-group">
                                <Input
                                    label="Years of Experience"
                                    showLabel
                                    category={edit ? "input" : "display"}
                                    type="number"
                                    bind:value={formData.yearsOfExperience}
                                />
                            </div>
                        </div>
                    {/if}

                    {#if user.type === Users.Patient}
                        <div class="form-section">
                            <h3>Medical Information</h3>
                            <div class="form-row">
                                <div class="form-group">
                                    {#if edit}
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
                                    {:else}
                                        {(user as Patient).bloodType}
                                    {/if}
                                </div>

                                <div class="form-group">
                                    <Input
                                        label="Emergency Contact"
                                        showLabel
                                        category={edit ? "input" : "display"}
                                        type="text"
                                        bind:value={formData.emergencyContact}
                                    />
                                </div>

                                <div class="form-group">
                                    <InputArray
                                        show={!edit}
                                        bind:value={formData.allergies}
                                    />
                                </div>

                                <div class="form-group">
                                    <InputArray
                                        show={!edit}
                                        bind:value={formData.medicalHistory}
                                    />
                                </div>
                            </div>
                        </div>
                    {/if}

                    <div class="actions">
                        <Button
                            type="submit"
                            disabled={isLoading}
                            variant="primary"
                        >
                            {isLoading ? "Saving..." : "Save Changes"}
                        </Button>
                        <Button
                            type="button"
                            onClick={() => window.history.back()}
                            category="secondary"
                        >
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

        .avatar {
            position: static;
            padding: 0;
        }

        .avatar-input {
            display: flex;
            justify-content: center;
            align-items: center;
        }

        .avatar-wrapper,
        .avatar {
            width: auto;
            height: auto;
            aspect-ratio: unset;
        }

        .avatar-overlay {
            display: none;
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

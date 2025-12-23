<script lang="ts">
    import { page } from "$app/state";
    import { Users } from "$lib/types/users";
    import { Image, Pen } from "@lucide/svelte";
    import { user as me } from "$lib/stores";

    import { onMount } from "svelte";
    import { UserAPI } from "$lib/api";
    import { Specialities } from "$lib/types/speciality";
    import {
        validate,
        validation,
        type Fillable,
        extract,
    } from "$lib/validation";

    import Plan from "../dashboard/blocks/Plan.svelte";
    import View from "$lib/components/ui/View.svelte";
    import Input from "$lib/components/ui/Input.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import InputArray from "$lib/components/ui/InputArray.svelte";

    import {
        bloodTypes,
        type BloodType,
        type Patient,
    } from "$lib/types/users/patient";

    import type { Doctor } from "$lib/types/users/doctor";
    import type { IDoctor, IUser, User } from "$lib/types/users";
    import type { Admin } from "$lib/types/users/admin";
    import { scale } from "svelte/transition";

    const userId = Number(page.params.id);

    let {
        edit = false,
    }: {
        edit: boolean;
    } = $props();

    let user: IUser | null = $state(null);
    let isMe = $derived(!edit && user && (user as any).id === ($me as any).id);
    let isLoading = $state(true);
    let errorMessage = $state("");
    let avatarUrl = $state("");

    // Form data using Fillable pattern with validators
    let formData: Fillable = $state({
        firstName: {
            value: "",
            error: "",
            validator: validation.name,
        },
        lastName: {
            value: "",
            error: "",
            validator: validation.name,
        },
        email: {
            value: "",
            error: "",
            validator: validation.email,
        },
        phoneNumber: {
            value: "",
            error: "",
            validator: validation.phoneNumber,
        },
        address: {
            value: "",
            error: "",
            validator: (v: string) => "", // Optional field
        },
        dateOfBirth: {
            value: new Date().toISOString().split("T")[0],
            error: "",
            validator: (v: string) => "", // Optional field
        },
        gender: {
            value: "",
            error: "",
            validator: (v: string) => "", // Optional field
        },
        speciality: {
            value: "",
            error: "",
            validator: (v: string) => "", // Optional for patients
        },
        licenseNumber: {
            value: "",
            error: "",
            validator: (v: string) => "", // Optional for patients
        },
        yearsOfExperience: {
            value: "",
            error: "",
            validator: (v: string) => "", // Optional for patients
        },
        bloodType: {
            value: "",
            error: "",
            validator: (v: string) => "", // Optional for doctors
        },
        emergencyContact: {
            value: "",
            error: "",
            validator: (v: string) => "", // Optional for doctors
        },
        allergies: {
            value: [] as string[],
            error: "",
            validator: (v: string) => "", // Optional field
        },
        medicalHistory: {
            value: [] as string[],
            error: "",
            validator: (v: string) => "", // Optional field
        },
    });

    onMount(async () => {
        try {
            if (typeof userId !== "number")
                throw new Error("User ID is required");
            user = await UserAPI.getById(userId);
            if (user === null) return;

            formData.firstName.value = user.firstName;
            formData.lastName.value = user.lastName;
            formData.email.value = user.email;
            formData.phoneNumber.value = user.phoneNumber || "";
            avatarUrl = user.avatarUrl || "";
            formData.dateOfBirth.value = new Date(
                user.dateOfBirth || new Date(),
            )
                .toISOString()
                .split("T")[0];
            formData.gender.value = user.gender || "";
            formData.address.value = user.address || "";
            formData.medicalHistory.value =
                (user as Patient).medicalHistory || [];

            if (user.type === Users.Doctor) {
                const doctor = user as Doctor;
                formData.speciality.value = doctor.speciality || "";
                formData.yearsOfExperience.value =
                    doctor.getYearsOfExperience()?.toString() || "";
            } else if (user.type === Users.Patient) {
                const patient = user as Patient;
                formData.bloodType.value = patient.bloodType || "";
                formData.allergies.value = patient.allergies || [];
                formData.emergencyContact.value =
                    patient.emergencyContact || "";
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
                avatarUrl = reader.result as string;
            };
            reader.readAsDataURL(file);
        }
    }

    function validateField(fieldName: keyof typeof formData) {
        const field = formData[fieldName];
        const validationError = field.validator
            ? field.validator(String(field.value ?? ""))
            : "";
        field.error = validationError;
    }

    async function handleSubmit() {
        if (edit) {
            try {
                // Validate required fields before submission
                validateField("firstName");
                validateField("lastName");
                validateField("email");
                validateField("phoneNumber");

                // Check if any field has error
                const hasErrors = Object.keys(formData).some(
                    (key) =>
                        formData[key as keyof typeof formData].error !== "",
                );

                if (hasErrors) {
                    errorMessage = "Please correct the errors in the form.";
                    return;
                }

                isLoading = true;
                errorMessage = "";

                const cleanData = extract(formData) as any;
                const updateData: any = {
                    firstName: cleanData.firstName,
                    lastName: cleanData.lastName,
                    email: cleanData.email,
                    phoneNumber: cleanData.phoneNumber,
                    avatarUrl: avatarUrl || user?.avatarUrl,
                    dateOfBirth: cleanData.dateOfBirth,
                    gender: cleanData.gender,
                    address: cleanData.address,
                };

                if (user?.type === Users.Doctor) {
                    updateData.speciality = cleanData.speciality;
                    updateData.yearsOfExperience =
                        parseInt(cleanData.yearsOfExperience) || 0;
                } else if (user?.type === Users.Patient) {
                    updateData.bloodType = cleanData.bloodType;
                    updateData.allergies = cleanData.allergies;
                    updateData.emergencyContact = cleanData.emergencyContact;
                    updateData.medicalHistory = cleanData.medicalHistory;
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
                    {#if avatarUrl}
                        <img
                            src={avatarUrl}
                            alt="Profile"
                            class="avatar-image"
                        />
                    {:else}
                        <div class="avatar-placeholder">
                            {formData.firstName.value?.charAt(
                                0,
                            )}{formData.lastName.value?.charAt(0)}
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
                <div transition:scale>
                    {#if isMe && ($me as Admin).plan}
                        <div>
                            <Plan />
                            <br />
                        </div>
                    {/if}

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
                                        bind:value={formData.firstName.value}
                                        bind:error={formData.firstName.error}
                                        validation={formData.firstName
                                            .validator}
                                        required
                                    />
                                </div>

                                <div class="form-group">
                                    <Input
                                        label="Last Name"
                                        showLabel
                                        category={edit ? "input" : "display"}
                                        type="text"
                                        bind:value={formData.lastName.value}
                                        bind:error={formData.lastName.error}
                                        validation={formData.lastName.validator}
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
                                        bind:value={formData.email.value}
                                        bind:error={formData.email.error}
                                        validation={formData.email.validator}
                                        required
                                    />
                                </div>

                                <div class="form-group">
                                    <Input
                                        label="Phone Number"
                                        showLabel
                                        category={edit ? "input" : "display"}
                                        type="tel"
                                        bind:value={formData.phoneNumber.value}
                                        bind:error={formData.phoneNumber.error}
                                        validation={formData.phoneNumber
                                            .validator}
                                    />
                                </div>
                            </div>

                            <div class="form-row">
                                <div class="form-group">
                                    {#if !edit}
                                        <Input
                                            label="Date of Birth"
                                            showLabel
                                            category={edit
                                                ? "input"
                                                : "display"}
                                            type="date"
                                            value={formData.dateOfBirth.value}
                                        />
                                    {:else}
                                        <Input
                                            label="Date of Birth"
                                            showLabel
                                            type="date"
                                            bind:value={
                                                formData.dateOfBirth.value
                                            }
                                            bind:error={
                                                formData.dateOfBirth.error
                                            }
                                            validation={formData.dateOfBirth
                                                .validator}
                                        />
                                    {/if}
                                </div>

                                <div class="form-group">
                                    {#if edit}
                                        <Input
                                            category="select"
                                            label="Select Gender"
                                            bind:value={formData.gender.value}
                                            bind:error={formData.gender.error}
                                            validation={formData.gender
                                                .validator}
                                            options={[
                                                {
                                                    value: "male",
                                                    label: "Male",
                                                },
                                                {
                                                    value: "female",
                                                    label: "Female",
                                                },
                                            ]}
                                            placeholder="Select Gender"
                                        />
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
                                    bind:value={formData.address.value}
                                    bind:error={formData.address.error}
                                    validation={formData.address.validator}
                                />
                            </div>
                        </div>

                        {#if user.type === Users.Doctor}
                            <div class="form-section">
                                <h3>Professional Information</h3>
                                <div class="form-row">
                                    <div class="form-group">
                                        {#if edit}
                                            <Input
                                                category="select"
                                                label="Select Specialization"
                                                bind:value={
                                                    formData.speciality.value
                                                }
                                                bind:error={
                                                    formData.speciality.error
                                                }
                                                validation={formData.speciality
                                                    .validator}
                                                options={Specialities.map(
                                                    (s) => ({
                                                        value: s,
                                                        label: s,
                                                    }),
                                                )}
                                                placeholder="Select Specialization"
                                            />
                                        {:else}
                                            <Input
                                                label="Speciality"
                                                showLabel
                                                category="display"
                                                value={(user as IDoctor)
                                                    .speciality}
                                            />
                                        {/if}
                                    </div>

                                    <div class="form-group">
                                        <Input
                                            label="License Number"
                                            showLabel
                                            category={edit
                                                ? "input"
                                                : "display"}
                                            type="text"
                                            bind:value={
                                                formData.licenseNumber.value
                                            }
                                            bind:error={
                                                formData.licenseNumber.error
                                            }
                                            validation={formData.licenseNumber
                                                .validator}
                                        />
                                    </div>
                                </div>

                                <div class="form-group">
                                    <Input
                                        label="Years of Experience"
                                        showLabel
                                        category={edit ? "input" : "display"}
                                        type="number"
                                        bind:value={
                                            formData.yearsOfExperience.value
                                        }
                                        bind:error={
                                            formData.yearsOfExperience.error
                                        }
                                        validation={formData.yearsOfExperience
                                            .validator}
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
                                            <Input
                                                category="select"
                                                label="Select Blood Type"
                                                bind:value={
                                                    formData.bloodType.value
                                                }
                                                bind:error={
                                                    formData.bloodType.error
                                                }
                                                validation={formData.bloodType
                                                    .validator}
                                                options={bloodTypes.map(
                                                    (b) => ({
                                                        value: b,
                                                        label: b,
                                                    }),
                                                )}
                                                placeholder="Select Blood Type"
                                            />
                                        {:else}
                                            {(user as Patient).bloodType}
                                        {/if}
                                    </div>

                                    <div class="form-group">
                                        <Input
                                            label="Emergency Contact"
                                            showLabel
                                            category={edit
                                                ? "input"
                                                : "display"}
                                            type="text"
                                            bind:value={
                                                formData.emergencyContact.value
                                            }
                                            bind:error={
                                                formData.emergencyContact.error
                                            }
                                            validation={formData
                                                .emergencyContact.validator}
                                        />
                                    </div>

                                    <div class="form-group">
                                        <InputArray
                                            show={!edit}
                                            bind:value={
                                                formData.allergies.value
                                            }
                                        />
                                    </div>

                                    <div class="form-group">
                                        <InputArray
                                            show={!edit}
                                            bind:value={
                                                formData.medicalHistory.value
                                            }
                                        />
                                    </div>
                                </div>
                            </div>
                        {/if}

                        <div class="actions">
                            {#if edit}
                                <Button
                                    type="button"
                                    onClick={() => window.history.back()}
                                    category="secondary"
                                >
                                    Cancel
                                </Button>
                                <Button
                                    type="submit"
                                    disabled={isLoading}
                                    variant="primary"
                                >
                                    {isLoading ? "Saving..." : "Save Changes"}
                                </Button>
                            {/if}

                            {#if isMe}
                                <Button
                                    variant="primary"
                                    Icon={Pen}
                                    href="/dashboard/users/{($me as any)
                                        .id}/modify"
                                >
                                    Edit my profile
                                </Button>
                            {/if}
                        </div>
                    </form>
                </div>
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

        .form-row {
            grid-template-columns: 1fr;
        }

        .actions {
            flex-direction: column;
        }
    }
</style>

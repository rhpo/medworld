<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import Input from "$lib/components/ui/Input.svelte";
    import type { Cabinet } from "$lib/types/cabinet";
    import { getPermissionsFromUserType } from "$lib/types/permission";
    import type { Admin } from "$lib/types/users/admin";
    import type { SuperAdmin } from "$lib/types/users/superadmin";
    import {
        Building,
        Save,
        Trash,
        MapPin,
        Clock,
        WifiOff,
        Wifi,
        ParkingCircle,
        Coffee,
        Pill,
        FlaskConical,
        AlertTriangle,
        CreditCard,
        BadgeCheck,
        Calendar,
        Accessibility,
    } from "@lucide/svelte";
    import { fade, slide } from "svelte/transition";

    interface IProps {
        user: SuperAdmin | Admin;
        cabinet: Cabinet | null;
    }

    let { user, cabinet = null }: IProps = $props();
    let permissions = getPermissionsFromUserType(user.type);
    let activeTab = $state("general");
    let isSaving = $state(false);
    let showDeleteConfirm = $state(false);

    let days = [
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
        "Sunday",
    ];
    let workingHours = $state(cabinet?.openingHours || {});

    let showLocationPicker = $state(false);

    async function handleSave() {
        if (!cabinet || !permissions.includes("edit_cabinet")) return;

        isSaving = true;
        try {
            alert("Changes saved successfully!");
        } catch (error) {
            alert("Error saving changes");
        } finally {
            isSaving = false;
        }
    }

    async function handleDelete() {
        if (!cabinet || !permissions.includes("remove_cabinet")) return;

        if (!showDeleteConfirm) {
            showDeleteConfirm = true;
            return;
        }

        try {
            alert("Cabinet deleted successssfully!");
        } catch (error) {
            alert("Error deleting cabinet");
        }
    }

    function updateWorkingHours(
        day: string,
        type: "open" | "close",
        value: string,
    ) {
        if (!cabinet) return;

        workingHours = {
            ...workingHours,
            [day]: {
                ...workingHours[day],
                [type]: value,
            },
        };
    }
</script>

<Block
    Icon={Building}
    group="cabinets"
    title="Manage {cabinet ? cabinet.name : 'Cabinet'}"
>
    {#if !cabinet || !permissions.includes("view_cabinet")}
        <div class="empty-state" transition:fade>
            <Building size={48} />
            <h3>Cannot find/view cabinet</h3>
            <p>
                You don't have permission to view this cabinet or it doesn't
                exist.
            </p>
        </div>
    {:else}
        <div class="cabinet-manager">
            <nav class="tabs">
                <button
                    class="tab-button"
                    class:active={activeTab === "general"}
                    onclick={() => (activeTab = "general")}
                >
                    General
                </button>
                <button
                    class="tab-button"
                    class:active={activeTab === "features"}
                    onclick={() => (activeTab = "features")}
                >
                    Features
                </button>
                <button
                    class="tab-button"
                    class:active={activeTab === "schedule"}
                    onclick={() => (activeTab = "schedule")}
                >
                    Schedule
                </button>
            </nav>

            <div class="tab-content">
                {#if activeTab === "general"}
                    <div class="form-section" transition:fade>
                        <h3>Basic Information</h3>
                        <div class="form-grid">
                            <Input
                                placeholder="Cabinet name..."
                                label="Name"
                                showLabel
                                bind:value={cabinet.name}
                            />
                            <Input
                                placeholder="Phone number..."
                                label="Phone number"
                                showLabel
                                bind:value={cabinet.phone}
                                type="tel"
                            />
                        </div>

                        <div class="location-section">
                            <h4>Location</h4>
                            <div class="location-input">
                                <Input
                                    placeholder="Address..."
                                    label="Address"
                                    showLabel
                                    bind:value={cabinet.location.address}
                                />
                                <button
                                    class="map-picker-btn"
                                    onclick={() => (showLocationPicker = true)}
                                >
                                    Pick on map
                                </button>
                            </div>

                            {#if showLocationPicker}
                                <div class="map-container">
                                    <h4>
                                        I'm so tired it's 4AM I can't keep it
                                        up.
                                    </h4>
                                </div>
                            {/if}
                        </div>
                    </div>
                {:else if activeTab === "features"}
                    <div class="form-section" transition:fade>
                        <h3>Cabinet Features</h3>
                        <div class="features-grid">
                            <label class="feature-item">
                                <input
                                    type="checkbox"
                                    bind:checked={cabinet.acceptsInsurance}
                                />
                                <CreditCard size={20} />
                                <span>Accepts Insurance</span>
                            </label>

                            <label class="feature-item">
                                <input
                                    type="checkbox"
                                    bind:checked={cabinet.acceptsUrgent}
                                />
                                <AlertTriangle size={20} />
                                <span>Accepts Urgent Cases</span>
                            </label>

                            <label class="feature-item">
                                <input
                                    type="checkbox"
                                    bind:checked={cabinet.accessHandicap}
                                />
                                <Accessibility size={20} />
                                <span>Wheelchair Accessible</span>
                            </label>

                            <label class="feature-item">
                                <input
                                    type="checkbox"
                                    bind:checked={cabinet.hasParking}
                                />
                                <ParkingCircle size={20} />
                                <span>Has Parking</span>
                            </label>

                            <label class="feature-item">
                                <input
                                    type="checkbox"
                                    bind:checked={cabinet.hasWifi}
                                />
                                {#if cabinet.hasWifi}
                                    <Wifi size={20} />
                                {:else}
                                    <WifiOff size={20} />
                                {/if}
                                <span>Free Wi-Fi</span>
                            </label>
                        </div>
                    </div>
                {:else if activeTab === "schedule"}
                    <div class="form-section" transition:fade>
                        <h3>Working Hours</h3>
                        <div class="schedule-grid">
                            {#each days as day}
                                <div class="day-schedule">
                                    <h4>{day}</h4>
                                    <div class="time-inputs">
                                        <Input
                                            type="time"
                                            label="Opens at"
                                            showLabel
                                            value={workingHours[day]?.open ||
                                                "09:00"}
                                            onInput={(e: any) =>
                                                updateWorkingHours(
                                                    day,
                                                    "open",
                                                    e.target.value,
                                                )}
                                        />
                                        <Input
                                            type="time"
                                            label="Closes at"
                                            showLabel
                                            value={workingHours[day]?.close ||
                                                "17:00"}
                                            onInput={(e: any) =>
                                                updateWorkingHours(
                                                    day,
                                                    "close",
                                                    e.target.value,
                                                )}
                                        />
                                    </div>
                                </div>
                            {/each}
                        </div>
                    </div>
                {/if}
            </div>
        </div>

        <div class="actions">
            {#if permissions.includes("remove_cabinet")}
                <Button
                    category="error"
                    Icon={Trash}
                    label={showDeleteConfirm
                        ? "Confirm Delete"
                        : `Delete ${cabinet.name}`}
                    onclick={handleDelete}
                />
            {/if}

            {#if permissions.includes("edit_cabinet")}
                <Button
                    category="primary"
                    Icon={Save}
                    label={isSaving ? "Saving..." : `Save Changes`}
                    disabled={isSaving}
                    onclick={handleSave}
                />
            {/if}
        </div>
    {/if}
</Block>

<style>
    .cabinet-manager {
        display: flex;
        flex-direction: column;
        gap: 2rem;
        margin-bottom: 2rem;
    }

    .tabs {
        display: flex;
        gap: 1rem;
        padding-bottom: 1rem;
        border-bottom: 1px solid var(--border-color);
    }

    .tab-button {
        background: none;
        border: none;
        padding: 0.5rem 1rem;
        font-size: 0.9rem;
        color: var(--text-muted);
        cursor: pointer;
        border-radius: 0.5rem;
        transition: all 0.2s ease;
    }

    .tab-button:hover {
        background: var(--background-hover);
        color: var(--text-color);
    }

    .tab-button.active {
        background: var(--color-primary);
        color: white;
    }

    .tab-content {
        padding: 1rem 0;
    }

    .form-section {
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
    }

    .form-section h3 {
        margin: 0;
        font-size: 1.1rem;
        font-weight: 600;
        color: var(--text-color);
    }

    .form-section h4 {
        margin: 0 0 1rem 0;
        font-size: 0.9rem;
        font-weight: 600;
        color: var(--text-muted);
    }

    .form-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
        gap: 1rem;
    }

    .location-section {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .location-input {
        display: flex;
        gap: 1rem;
        align-items: flex-end;
    }

    .map-picker-btn {
        background: var(--background-secondary);
        border: 1px solid var(--border-color);
        padding: 0.75rem 1rem;
        border-radius: 0.5rem;
        color: var(--color-primary);
        font-size: 0.875rem;
        cursor: pointer;
        transition: all 0.2s ease;
        white-space: nowrap;
    }

    .map-picker-btn:hover {
        background: var(--background-hover);
        border-color: var(--color-primary);
    }

    .map-container {
        height: 300px;
        background: var(--background-secondary);
        border-radius: 1rem;
        overflow: hidden;
        margin-top: 1rem;
    }

    .features-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
        gap: 1rem;
    }

    .feature-item {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        padding: 1rem;
        background: var(--background-secondary);
        border: 1px solid var(--border-color);
        border-radius: 0.75rem;
        cursor: pointer;
        transition: all 0.2s ease;
    }

    .feature-item:hover {
        border-color: var(--color-primary);
        transform: translateY(-2px);
    }

    .feature-item input[type="checkbox"] {
        width: 1.125rem;
        height: 1.125rem;
        border-radius: 0.25rem;
        border: 2px solid var(--border-color);
        appearance: none;
        cursor: pointer;
        position: relative;
        transition: all 0.2s ease;
    }

    .feature-item input[type="checkbox"]:checked {
        background-color: var(--color-primary);
        border-color: var(--color-primary);
    }

    .feature-item input[type="checkbox"]:checked::after {
        content: "✓";
        color: white;
        font-size: 0.75rem;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }

    .feature-item span {
        font-size: 0.875rem;
        color: var(--text-color);
    }

    .schedule-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
        gap: 1.5rem;
    }

    .day-schedule {
        background: var(--background-secondary);
        padding: 1rem;
        border-radius: 0.75rem;
        border: 1px solid var(--border-color);
    }

    .day-schedule h4 {
        margin: 0 0 1rem 0;
        color: var(--text-color);
    }

    .time-inputs {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 1rem;
    }

    .empty-state {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        padding: 3rem;
        text-align: center;
        color: var(--text-muted);
    }

    .empty-state h3 {
        margin: 1rem 0 0.5rem;
        font-size: 1.1rem;
        font-weight: 600;
        color: var(--text-color);
    }

    .empty-state p {
        margin: 0;
        font-size: 0.9rem;
    }

    .actions {
        display: flex;
        justify-content: flex-end;
        gap: 1rem;
        margin-top: 1rem;
        padding-top: 1rem;
        border-top: 1px solid var(--border-color);
    }

    @media (max-width: 768px) {
        .tabs {
            overflow-x: auto;
            padding-bottom: 0.5rem;
            margin: 0 -1rem;
            padding: 0 1rem 0.5rem;
        }

        .tab-button {
            padding: 0.375rem 0.75rem;
            font-size: 0.875rem;
            white-space: nowrap;
        }

        .form-grid {
            grid-template-columns: 1fr;
        }

        .location-input {
            flex-direction: column;
            gap: 0.5rem;
            align-items: stretch;
        }

        .features-grid {
            grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
        }

        .schedule-grid {
            grid-template-columns: 1fr;
        }

        .actions {
            flex-direction: column-reverse;
            gap: 0.5rem;
        }
    }

    @media (prefers-reduced-motion: reduce) {
        .tab-button,
        .feature-item,
        .map-picker-btn {
            transition: none;
        }

        .feature-item:hover {
            transform: none;
        }
    }
</style>

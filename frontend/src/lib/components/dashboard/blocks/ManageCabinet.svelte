<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import Input from "$lib/components/ui/Input.svelte";

    import { fade } from "svelte/transition";
    import { getPermissionsFromUserType } from "$lib/types/permission";

    import type { Admin } from "$lib/types/users/admin";
    import type { Cabinet, CabinetFillable } from "$lib/types/cabinet";
    import type { SuperAdmin } from "$lib/types/users/superadmin";

    import {
        Building,
        Save,
        Trash,
        WifiOff,
        Wifi,
        ParkingCircle,
        AlertTriangle,
        CreditCard,
        Accessibility,
    } from "@lucide/svelte";

    import { extract, validate, validation } from "$lib/validation";

    import type { Fillable } from "$lib/validation";

    interface IProps {
        user: SuperAdmin | Admin;
        cabinet: Cabinet;
        [key: string]: any;
    }

    let { user, cabinet, ...rest }: IProps = $props();
    let permissions = getPermissionsFromUserType(user.type);

    let activeTab = $state("general");
    let isSaving = $state(false);
    let showDeleteConfirm = $state(false);

    const defaultHours: Record<string, { open: string; close: string }> = {
        monday: { open: "09:00", close: "17:00" },
        tuesday: { open: "09:00", close: "17:00" },
        wednesday: { open: "09:00", close: "17:00" },
        thursday: { open: "09:00", close: "17:00" },
        friday: { open: "09:00", close: "17:00" },
        saturday: { open: "09:00", close: "13:00" },
        sunday: { open: "Closed", close: "Closed" },
    };

    const mergedHours = {
        ...defaultHours,
        ...JSON.parse(JSON.stringify(cabinet.openingHours || {})),
    };

    const initializeOpeningHours = () => {
        const hours: Record<
            string,
            {
                value: { open: string; close: string };
                error: string;
                validator: (value: string) => string;
            }
        > = {};

        [
            "monday",
            "tuesday",
            "wednesday",
            "thursday",
            "friday",
            "saturday",
            "sunday",
        ].forEach((day) => {
            hours[day] = {
                value: mergedHours[day] || { open: "09:00", close: "17:00" },
                error: "",
                validator: validation.nothing,
            };
        });

        return hours;
    };

    let data: Fillable = $state({
        name: {
            value: "",
            error: "",
            validator: validation.name,
        },
        phone: {
            value: cabinet.phone,
            error: "",
            validator: validation.phoneNumber,
        },
        location: {
            value: { ...cabinet.location },
            error: "",
            validator: validation.nothing,
        },
        accessHandicap: {
            value: cabinet.accessHandicap,
            error: "",
            validator: validation.nothing,
        },
        hasParking: {
            value: cabinet.hasParking,
            error: "",
            validator: validation.nothing,
        },
        hasWifi: {
            value: cabinet.hasWifi,
            error: "",
            validator: validation.nothing,
        },
        acceptsUrgent: {
            value: cabinet.acceptsUrgent,
            error: "",
            validator: validation.nothing,
        },
        acceptsInsurance: {
            value: cabinet.acceptsInsurance,
            error: "",
            validator: validation.nothing,
        },
        openingHours: initializeOpeningHours(),
    });

    function isValidTime(time: string) {
        return /^([0-1]\d|2[0-3]):([0-5]\d)$/.test(time);
    }

    function isValidOpeningHours(hours: {
        open: string;
        close: string;
    }): boolean {
        if (hours.open === "Closed" && hours.close === "Closed") return true;
        if (!isValidTime(hours.open) || !isValidTime(hours.close)) return false;

        const [openHour, openMinute] = hours.open.split(":").map(Number);
        const [closeHour, closeMinute] = hours.close.split(":").map(Number);

        if (openHour > closeHour) return false;
        if (openHour === closeHour && openMinute >= closeMinute) return false;

        return true;
    }

    async function handleSave() {
        let error = validate(data);
        if (error) return alert(error);

        if (!permissions.includes("edit_cabinet")) return;

        isSaving = true;
        try {
            let readyData = extract(data);
            // do something with readyData
        } catch (error) {
            alert("Error saving changes");
        } finally {
            isSaving = false;
        }
    }

    async function handleDelete() {
        if (!permissions.includes("remove_cabinet")) return;

        if (!showDeleteConfirm) {
            showDeleteConfirm = true;
            return;
        }

        try {
            alert("Cabinet deleted successfully!");
        } catch (error) {
            alert("Error deleting cabinet");
        }
    }
</script>

<Block
    Icon={Building}
    group="manage_cabinet"
    title="Manage {cabinet?.name || 'Cabinet'}"
    {...rest}
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
        <nav class="tabs">
            <button
                class="tab-button"
                class:active={activeTab === "general"}
                onclick={() => (activeTab = "general")}>General</button
            >
            <button
                class="tab-button"
                class:active={activeTab === "features"}
                onclick={() => (activeTab = "features")}>Features</button
            >
            <button
                class="tab-button"
                class:active={activeTab === "schedule"}
                onclick={() => (activeTab = "schedule")}>Schedule</button
            >
        </nav>

        <div class="tab-content">
            {#if activeTab === "general"}
                <div class="form-section" transition:fade>
                    <h3>Basic Information</h3>
                    <Input
                        placeholder="Cabinet name..."
                        label="Name"
                        showLabel
                        bind:value={data.name.value}
                        bind:error={data.name.error}
                        validation={data.name.validator}
                        required
                    />
                    <Input
                        placeholder="Phone number..."
                        label="Phone"
                        showLabel
                        bind:value={data.phone.value}
                        bind:error={data.phone.error}
                        validation={data.phone.validator}
                        type="tel"
                        required
                    />
                    <h4>Location</h4>
                    <Input
                        placeholder="Address..."
                        label="Address"
                        showLabel
                        bind:value={data.location.value.address}
                        bind:error={data.location.error}
                        validation={data.location.validator}
                        required
                    />
                </div>
            {:else if activeTab === "features"}
                <div class="form-section" transition:fade>
                    <h3>Cabinet Features</h3>
                    <div class="features-grid">
                        <label class="feature-item">
                            <input
                                type="checkbox"
                                bind:checked={data.acceptsInsurance.value}
                            />
                            <CreditCard size={20} />
                            <span>Accepts Insurance</span>
                        </label>
                        <label class="feature-item">
                            <input
                                type="checkbox"
                                bind:checked={data.acceptsUrgent.value}
                            />
                            <AlertTriangle size={20} />
                            <span>Accepts Urgent Cases</span>
                        </label>
                        <label class="feature-item">
                            <input
                                type="checkbox"
                                bind:checked={data.accessHandicap.value}
                            />
                            <Accessibility size={20} />
                            <span>Wheelchair Accessible</span>
                        </label>
                        <label class="feature-item">
                            <input
                                type="checkbox"
                                bind:checked={data.hasParking.value}
                            />
                            <ParkingCircle size={20} />
                            <span>Has Parking</span>
                        </label>
                        <label class="feature-item">
                            <input
                                type="checkbox"
                                bind:checked={data.hasWifi.value}
                            />
                            {#if data.hasWifi}
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
                        {#each ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"] as day}
                            <div class="day-schedule">
                                <h4>{day}</h4>
                                <div class="time-inputs">
                                    <Input
                                        type="time"
                                        label="Opens"
                                        showLabel
                                        bind:value={
                                            data.openingHours[day.toLowerCase()]
                                                .value.open
                                        }
                                        bind:error={
                                            data.openingHours[day.toLowerCase()]
                                                .error
                                        }
                                        validation={(e: string) => {
                                            if (
                                                isValidOpeningHours({
                                                    open: e,
                                                    close: data.openingHours[
                                                        day.toLowerCase()
                                                    ].value.close,
                                                })
                                            ) {
                                                return "";
                                            }
                                            return "Invalid opening hours";
                                        }}
                                    />
                                    <Input
                                        type="time"
                                        label="Closes"
                                        showLabel
                                        bind:value={
                                            data.openingHours[day.toLowerCase()]
                                                .value.close
                                        }
                                        bind:error={
                                            data.openingHours[day.toLowerCase()]
                                                .error
                                        }
                                        validation={(e: string) => {
                                            if (
                                                isValidOpeningHours({
                                                    open: data.openingHours[
                                                        day.toLowerCase()
                                                    ].value.open,
                                                    close: e,
                                                })
                                            ) {
                                                return "";
                                            }
                                            return "Invalid opening hours";
                                        }}
                                    />
                                </div>
                            </div>
                        {/each}
                    </div>
                </div>
            {/if}
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
                    label={isSaving ? "Saving..." : "Save Changes"}
                    disabled={isSaving}
                    onclick={handleSave}
                />
            {/if}
        </div>
    {/if}
</Block>

<style>
    .tabs {
        display: flex;
        gap: 1rem;
        padding-bottom: 1rem;
        border-bottom: 1px solid var(--border-color);
        margin-bottom: 2rem;
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

    .form-section {
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
        margin-bottom: 2rem;
    }

    .form-section h3 {
        margin: 0;
        font-size: 1.1rem;
        font-weight: 600;
    }

    .form-section h4 {
        margin: 1rem 0 0.5rem;
        font-size: 0.9rem;
        font-weight: 600;
        color: var(--text-muted);
    }

    .features-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
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
        cursor: pointer;
    }

    .schedule-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
        gap: 1.5rem;
    }

    .day-schedule {
        background: var(--background-secondary);
        padding: 1rem;
        border-radius: 0.75rem;
        border: 1px solid var(--border-color);
    }

    .day-schedule h4 {
        margin: 0 0 1rem;
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
    }

    .empty-state p {
        margin: 0;
        font-size: 0.9rem;
    }

    .actions {
        display: flex;
        justify-content: flex-end;
        gap: 1rem;
        padding-top: 1rem;
        border-top: 1px solid var(--border-color);
    }

    @media (max-width: 768px) {
        .tabs {
            overflow-x: auto;
            margin: 0 -1rem;
            padding: 0 1rem 0.5rem;
        }

        .tab-button {
            padding: 0.375rem 0.75rem;
            font-size: 0.875rem;
            white-space: nowrap;
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
        .feature-item {
            transition: none;
        }

        .feature-item:hover {
            transform: none;
        }
    }
</style>

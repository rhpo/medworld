<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import type { Cabinet } from "$lib/types/cabinet";
    import { Users } from "$lib/types/users";
    import {
        MapPin,
        Phone,
        Users as UsersIcon,
        Clock,
        Shield,
        Edit,
        Settings,
    } from "@lucide/svelte";
    import { createEventDispatcher, onMount } from "svelte";

    interface IProps {
        cabinet?: Cabinet;
        compact?: boolean;
        permissions?: { canEdit?: boolean; canManageStaff?: boolean };
        mode?: "view" | "edit" | "create";
        onEdit?: (id: number) => void;
        onManageStaff?: (id: number) => void;
        onView?: (id: number) => void;
        onSave?: (cabinet: Cabinet) => void;
        onCancel?: () => void;
    }

    // Use the project's $props / $state / $derived patterns for consistency
    let {
        cabinet = {
            id: 0,
            name: "",
            address: "",
            phone: "",
            consultationDuration: 30,
            isClosed: false,
            doctors: [],
            assistants: [],
            openingHours: {},
        } as unknown as Cabinet,
        mode = "view",
        permissions = {},
        onEdit = () => {},
        onManageStaff = () => {},
        onView = () => {},
        onSave = () => {},
        onCancel = () => {},
    }: IProps = $props();

    const dispatch = createEventDispatcher();

    let showHours = $state(false);

    // derive today's name and hours
    const weekDays = [
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
    ];
    let todayIndex = new Date().getDay();
    let todayName = $derived(weekDays[todayIndex]);

    function formatTime(time: string) {
        try {
            return new Date(`2000-01-01T${time}`).toLocaleTimeString([], {
                hour: "2-digit",
                minute: "2-digit",
            });
        } catch (e) {
            return time;
        }
    }

    function handleEdit() {
        onEdit?.(cabinet.id);
        dispatch("edit", { id: cabinet.id });
    }

    function handleManageStaff() {
        onManageStaff?.(cabinet.id);
        dispatch("manageStaff", { id: cabinet.id });
    }

    function handleView() {
        onView?.(cabinet.id);
        dispatch("view", { id: cabinet.id });
    }

    // small defensive defaults with required Cabinet shape
    let cabinetWithDefaults = $derived(
        cabinet ??
            ({
                id: 0,
                name: "—",
                address: "",
                phone: "",
                createdAt: new Date(),
                admin: {
                    id: 0,
                    firstName: "",
                    lastName: "",
                    email: "",
                    phoneNumber: "",
                    type: Users.Admin,
                    getFullName: () => "—",
                    plan: {
                        id: 0,
                        name: "—",
                        price: 0,
                        features: [],
                    },
                },
                doctors: [],
                assistants: [],
                openingHours: {},
                consultationDuration: 30,
                isClosed: true,
                premiumPack: undefined,
            } as Cabinet),
    );

    // accessibility: compute today hours summary
    let todayHours = $derived(
        cabinetWithDefaults.openingHours?.[todayName.toLowerCase()],
    );

    onMount(() => {
        // recalc today index in case of SSR
        todayIndex = new Date().getDay();
    });
</script>

<Block title="Manage Cabinet">
    <div class="manage-header">
        <div class="title">
            <h3 class="cabinet-name">{cabinet.name}</h3>
            {#if cabinet.address}
                <div class="subtitle">
                    <MapPin size={14} /> <span>{cabinet.address}</span>
                </div>
            {/if}
        </div>

        <div class="actions">
            <Button
                category="third"
                onClick={handleView}
                aria-label="View cabinet"><Settings size={16} /></Button
            >
            {#if permissions.canManageStaff}
                <Button
                    category="third"
                    onClick={handleManageStaff}
                    aria-label="Manage staff"><UsersIcon size={16} /></Button
                >
            {/if}
            {#if permissions.canEdit}
                <Button
                    category="primary"
                    onClick={handleEdit}
                    aria-label="Edit cabinet"
                    Icon={Edit}>Edit</Button
                >
            {/if}
        </div>
    </div>

    <div class="grid">
        <div class="summary">
            <div class="line">
                <Phone size={16} /> <span>{cabinet.phone || "N/A"}</span>
            </div>
            <div class="line">
                <UsersIcon size={16} />
                <span>Admin: {cabinet.admin?.getFullName?.() ?? "—"}</span>
            </div>
            <div class="line stats">
                <span class="badge"
                    >Doctors: {cabinet.doctors?.length ?? 0}</span
                >
                <span class="badge"
                    >Assistants: {cabinet.assistants?.length ?? 0}</span
                >
                {#if cabinet.premiumPack}
                    <span class="badge premium">{cabinet.premiumPack}</span>
                {/if}
            </div>
        </div>

        <div class="hours">
            <div class="hours-head">
                <div class="today">
                    Today: {#if todayHours}{formatTime(todayHours.open)} - {formatTime(
                            todayHours.close,
                        )}{:else}<i>Closed</i>{/if}
                </div>
                <button
                    class="toggle"
                    onclick={() => (showHours = !showHours)}
                    aria-expanded={showHours}
                    aria-controls="hours-list"
                    >{showHours ? "Hide" : "View all"}</button
                >
            </div>

            {#if showHours}
                <div id="hours-list" class="hours-list">
                    {#each weekDays as day}
                        {#if cabinet.openingHours?.[day.toLowerCase()]}
                            <div class="hours-item">
                                <div class="day">{day}</div>
                                <div class="time">
                                    {formatTime(
                                        cabinet.openingHours[day.toLowerCase()]
                                            .open,
                                    )} - {formatTime(
                                        cabinet.openingHours[day.toLowerCase()]
                                            .close,
                                    )}
                                </div>
                            </div>
                        {:else}
                            <div class="hours-item closed">
                                <div class="day">{day}</div>
                                <div class="time"><i>Closed</i></div>
                            </div>
                        {/if}
                    {/each}
                </div>
            {/if}
        </div>

        <div class="settings">
            <div class="settings-item">
                <div class="icon"><Clock size={16} /></div>
                <div class="info">
                    <div class="label">Default Duration</div>
                    <div class="value">
                        {cabinet.consultationDuration ?? "—"} minutes
                    </div>
                </div>
            </div>

            <div class="settings-item">
                <div class="icon"><Shield size={16} /></div>
                <div class="info">
                    <div class="label">Status</div>
                    <div class="value">
                        {cabinet.isClosed ? "Closed" : "Open"}
                    </div>
                </div>
            </div>
        </div>
    </div>
</Block>

<style>
    .manage-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 1rem;
    }

    .title .cabinet-name {
        margin: 0;
        font-size: 1.05rem;
        font-weight: 700;
    }

    .subtitle {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        color: var(--text-muted);
        font-size: 0.9rem;
        margin-top: 0.25rem;
    }

    .actions {
        display: flex;
        gap: 0.5rem;
        align-items: center;
    }

    .grid {
        display: grid;
        grid-template-columns: 1fr 260px 200px;
        gap: 1rem;
        margin-top: 1rem;
        align-items: start;
    }

    .summary {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    .line {
        display: flex;
        gap: 0.6rem;
        align-items: center;
        color: var(--text-color);
        font-size: 0.95rem;
    }

    .stats .badge {
        display: inline-block;
        background: var(--background-alt);
        padding: 0.35rem 0.6rem;
        border-radius: 999px;
        margin-right: 0.5rem;
        font-size: 0.85rem;
    }

    .stats .premium {
        background: linear-gradient(
            90deg,
            var(--color-primary),
            var(--color-primary-dark)
        );
        color: white;
    }

    .hours {
        background: transparent;
    }

    .hours-head {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 1rem;
    }

    .hours-list {
        margin-top: 0.75rem;
        background: var(--background-alt);
        border-radius: 8px;
        overflow: hidden;
    }

    .hours-item {
        display: flex;
        justify-content: space-between;
        padding: 0.6rem 0.9rem;
        border-bottom: 1px solid var(--border-color);
        font-size: 0.95rem;
    }

    .hours-item.closed {
        color: var(--text-muted);
        background: var(--background-hover);
    }

    .settings {
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
    }

    .settings-item {
        display: flex;
        gap: 0.75rem;
        align-items: center;
        padding: 0.6rem;
        background: var(--background-alt);
        border-radius: 8px;
    }

    .settings-item .icon {
        display: inline-flex;
        width: 36px;
        height: 36px;
        align-items: center;
        justify-content: center;
        border-radius: 6px;
        background: var(--primary-color);
        color: white;
    }

    .settings-item .label {
        font-size: 0.85rem;
        color: var(--text-muted);
    }

    .settings-item .value {
        font-weight: 600;
    }

    .toggle {
        background: transparent;
        border: none;
        color: var(--color-primary);
        cursor: pointer;
        font-size: 0.9rem;
    }

    @media (max-width: 900px) {
        .grid {
            grid-template-columns: 1fr 200px;
        }
    }

    @media (max-width: 640px) {
        .grid {
            grid-template-columns: 1fr;
        }

        .manage-header {
            flex-direction: column;
            align-items: flex-start;
            gap: 0.5rem;
        }

        .actions {
            width: 100%;
            justify-content: flex-end;
        }

        .settings-item {
            padding: 0.5rem;
        }
    }
</style>

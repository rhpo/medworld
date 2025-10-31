<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import View from "$lib/components/ui/View.svelte";
    import Face from "$lib/components/ui/Face.svelte";
    import {
        Calendar,
        Users as UsersIcon,
        Clock,
        MapPin,
        Phone,
        Shield,
        Settings,
    } from "@lucide/svelte";

    import { fakeCabinets } from "$lib/types/fakedata";
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import type { Cabinet } from "$lib/types/cabinet";
    import type { User, Users } from "$lib/types/users";

    // get id from route params and load the matching cabinet
    const cabinetId = Number($page.params.id);
    let cabinet: Cabinet =
        fakeCabinets.find((c) => c.id === cabinetId) ?? fakeCabinets[0];

    // (optional) if you want to fetch from an API on mount, do it here.
    onMount(() => {
        // e.g. fetch(`/api/cabinets/${cabinetId}`).then(...)
    });

    const weekDays = [
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
    ];

    function formatTime(time: string): string {
        return new Date(`2000-01-01T${time}`).toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
        });
    }
</script>

<Face title={cabinet.name} description="View cabinet details and settings">
    <View>
        <div class="page-content">
            <div class="cabinet-header">
                <h1>{cabinet.name}</h1>
            </div>

            <div class="cabinet-grid">
                <!-- Basic Information -->
                <Block title="Basic Information">
                    <div class="info-list">
                        <div class="info-item">
                            <MapPin size={20} />
                            <span>{cabinet.address}</span>
                        </div>
                        <div class="info-item">
                            <Phone size={20} />
                            <span>{cabinet.phone}</span>
                        </div>
                        <div class="info-item">
                            <Shield size={20} />
                            <span
                                >Premium Pack: {cabinet.premiumPack ||
                                    "None"}</span
                            >
                        </div>
                    </div>
                </Block>

                <!-- Staff -->
                <Block title="Staff Members">
                    <div class="staff-list">
                        <div class="staff-section">
                            <h3>Admin</h3>
                            <div class="staff-member">
                                <UsersIcon size={16} />
                                <span>{cabinet.admin.getFullName()}</span>
                            </div>
                        </div>

                        <div class="staff-section">
                            <h3>Doctors ({cabinet.doctors.length})</h3>
                            {#each cabinet.doctors as doctor}
                                <div class="staff-member">
                                    <UsersIcon size={16} />
                                    <span>{doctor.getFullName()}</span>
                                </div>
                            {/each}
                        </div>

                        <div class="staff-section">
                            <h3>Assistants ({cabinet.assistants.length})</h3>
                            {#each cabinet.assistants as assistant}
                                <div class="staff-member">
                                    <UsersIcon size={16} />
                                    <span>{assistant.getFullName()}</span>
                                </div>
                            {/each}
                        </div>
                    </div>
                </Block>

                <!-- Opening Hours -->
                <Block title="Opening Hours">
                    <div class="hours-list">
                        {#each weekDays as day}
                            <div
                                class="hours-item"
                                class:closed={!cabinet.openingHours?.[
                                    day.toLowerCase()
                                ]}
                            >
                                <div class="day">{day}</div>
                                <div class="time">
                                    {#if cabinet.openingHours?.[day.toLowerCase()]}
                                        {formatTime(
                                            cabinet.openingHours[
                                                day.toLowerCase()
                                            ].open,
                                        )} - {formatTime(
                                            cabinet.openingHours[
                                                day.toLowerCase()
                                            ].close,
                                        )}
                                    {:else}
                                        <i>Closed</i>
                                    {/if}
                                </div>
                            </div>
                        {/each}
                    </div>
                </Block>

                <!-- Consultation Settings -->
                <Block title="Consultation Settings">
                    <div class="settings-list">
                        <div class="settings-item">
                            <div class="settings-icon">
                                <Clock size={20} />
                            </div>
                            <div class="settings-info">
                                <div class="settings-label">
                                    Default Duration
                                </div>
                                <div class="settings-value">
                                    {cabinet.consultationDuration} minutes
                                </div>
                            </div>
                        </div>
                        <div class="settings-item">
                            <div class="settings-icon">
                                <Shield size={20} />
                            </div>
                            <div class="settings-info">
                                <div class="settings-label">Status</div>
                                <div class="settings-value">
                                    {cabinet.isClosed ? "Closed" : "Open"}
                                </div>
                            </div>
                        </div>
                    </div>
                </Block>
            </div>
        </div>
    </View>
</Face>

<style>
    .page-content {
        flex: 1;
        width: 100%;
    }

    .cabinet-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 1.5rem;
        padding: 0 0.5rem;
    }

    .cabinet-header h1 {
        margin: 0;
        font-size: 1.75rem;
        line-height: 1.2;
    }

    .cabinet-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
        gap: 1rem;
        padding: 0 0.5rem;
    }

    .info-list {
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
    }

    .info-item {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        color: var(--text-color);
        font-size: 0.95rem;
    }

    .info-item :global(svg) {
        color: var(--primary-color);
        flex-shrink: 0;
    }

    .staff-list {
        display: flex;
        flex-direction: column;
        gap: 1.5rem;
    }

    .staff-section {
        padding: 0.5rem 0;
    }

    .staff-section h3 {
        font-size: 0.9rem;
        font-weight: 600;
        color: var(--text-muted);
        margin: 0 0 0.5rem 0;
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .staff-member {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        padding: 0.35rem 0;
        font-size: 0.95rem;
    }

    .staff-member :global(svg) {
        color: var(--primary-color);
        flex-shrink: 0;
    }

    .hours-list {
        display: flex;
        flex-direction: column;
        background: var(--background-alt);
        border-radius: 8px;
        overflow: hidden;
    }

    .hours-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0.75rem 1rem;
        border-bottom: 1px solid var(--border-color);
        font-size: 0.95rem;
    }

    .hours-item:last-child {
        border-bottom: none;
    }

    .hours-item.closed {
        color: var(--text-muted);
        background: var(--background-hover);
    }

    .day {
        font-weight: 500;
    }

    .time {
        color: var(--text-muted);
    }

    .time i {
        font-style: italic;
    }

    .settings-list {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .settings-item {
        display: flex;
        gap: 1rem;
        padding: 0.75rem;
        background: var(--background-alt);
        border-radius: 8px;
    }

    .settings-icon {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 40px;
        height: 40px;
        background: var(--primary-color);
        border-radius: 8px;
        color: white;
    }

    .settings-info {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
    }

    .settings-label {
        color: var(--text-muted);
        font-size: 0.9rem;
    }

    .settings-value {
        font-weight: 500;
    }

    @media (max-width: 768px) {
        .cabinet-header {
            margin-bottom: 1rem;
        }

        .cabinet-header h1 {
            font-size: 1.5rem;
        }

        .cabinet-grid {
            grid-template-columns: 1fr;
            gap: 0.75rem;
        }

        .info-item,
        .staff-member {
            font-size: 0.9rem;
        }

        .hours-item {
            padding: 0.6rem 0.75rem;
            font-size: 0.9rem;
        }

        .settings-item {
            padding: 0.6rem;
        }

        .settings-icon {
            width: 36px;
            height: 36px;
        }

        .settings-icon :global(svg) {
            width: 18px;
            height: 18px;
        }
    }

    @media (max-width: 480px) {
        .cabinet-header {
            padding: 0;
        }

        .cabinet-grid {
            padding: 0;
        }

        .hours-item {
            padding: 0.35rem 0;
        }
    }
</style>

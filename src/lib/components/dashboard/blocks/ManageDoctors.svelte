<script lang="ts">
    import { AllAPI } from "$lib/api";
    import Avatar from "$lib/components/ui/Avatar.svelte";
    import Block from "$lib/components/ui/Block.svelte";
    import Button from "$lib/components/ui/Button.svelte";
    import IconButton from "$lib/components/ui/IconButton.svelte";
    import type { Permission } from "$lib/types/permission";
    import type { IUser } from "$lib/types/users";
    import type { Admin } from "$lib/types/users/admin";
    import type { Doctor } from "$lib/types/users/doctor";
    import { Crown, Pen, Plus } from "@lucide/svelte";
    import { onMount } from "svelte";

    // props: accept any user shape to avoid strict Admin typing here
    let {
        user,
        permissions,
    }: {
        user: IUser;
        permissions: Permission[];
    } = $props();

    let doctors: Doctor[] = $state([]);

    onMount(async () => {
        // prefer server call for canonical list; if the user already has a doctors list, use it
        if (user && Array.isArray((user as any).doctors)) {
            doctors = (user as any).doctors;
        } else {
            doctors = await AllAPI.listAllDoctors();
        }
    });
</script>

<main>
    <Block title="Manage Doctors">
        <table>
            <thead>
                <tr>
                    <th>Avatar</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Phone</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each doctors as doctor}
                    <tr>
                        <td>
                            <Avatar
                                size="48px"
                                avatarUrl={doctor.avatarUrl}
                                alt={doctor.getFullName()}
                            />
                        </td>
                        <td>{doctor.getFullName()}</td>
                        <td>{doctor.email}</td>
                        <td>{doctor.phoneNumber || "N/A"}</td>
                        <td class="actions">
                            {#if permissions.includes("edit_doctor")}
                                <IconButton
                                    href="/dashboard/users/{doctor.id}/modify"
                                    label="Edit doctor account"
                                    target="_blank"
                                    Icon={Pen}
                                ></IconButton>
                            {/if}

                            {#if permissions.includes("set_admin_doctor")}
                                <IconButton
                                    Icon={Crown}
                                    label="Set as Admin"
                                    color="var(--gold-dark)"
                                ></IconButton>
                            {/if}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>

        {#if permissions.includes("add_doctor")}
            <div class="add-actions">
                <Button
                    label="Add a new Doctor"
                    Icon={Plus}
                    href="/dashboard/users/add"
                    category="primary"
                    style="margin-top: 1.5rem; width: 100%"
                ></Button>

                <Button
                    label="Add an existing Doctor"
                    Icon={Plus}
                    href="#"
                    category="secondary"
                    style="margin-top: 1.5rem; width: 100%"
                ></Button>
            </div>
        {/if}
    </Block>
</main>

<style>
    .add-actions {
        display: flex;
        gap: 1rem;
    }

    .actions {
        display: flex;
        gap: 0.5rem;
    }
</style>

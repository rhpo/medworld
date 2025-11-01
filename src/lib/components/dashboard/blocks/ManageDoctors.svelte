<script lang="ts">
    import { AllAPI } from "$lib/api";
    import Avatar from "$lib/components/ui/Avatar.svelte";
    import Block from "$lib/components/ui/Block.svelte";
    import IconButton from "$lib/components/ui/IconButton.svelte";
    import type { Permission } from "$lib/types/permission";
    import { Users, type User } from "$lib/types/users";
    import type { Admin } from "$lib/types/users/admin";
    import { Crown, Edit, Edit2, Pen } from "@lucide/svelte";
    import { onMount } from "svelte";

    // props: accept any user shape to avoid strict Admin typing here
    let {
        user,
        permissions,
    }: {
        user: Admin;
        permissions: Permission[];
    } = $props();

    let doctors: User<Users.Doctor>[] = $state([]);

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
                                    href="/routes/dashboard/users/edit"
                                    label="Edit doctor account"
                                    Icon={Pen}
                                ></IconButton>
                            {/if}
                            {#if permissions.includes("set_admin_doctor")}
                                <IconButton
                                    Icon={Crown}
                                    label="Set as Admin"
                                    color="#AA6C39"
                                ></IconButton>
                            {/if}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </Block>
</main>

<style>
    .actions {
        display: flex;
        gap: 0.5rem;
    }
</style>

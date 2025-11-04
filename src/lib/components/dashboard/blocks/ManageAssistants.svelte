<script lang="ts">
    import { AllAPI } from "$lib/api";
    import Avatar from "$lib/components/ui/Avatar.svelte";
    import Block from "$lib/components/ui/Block.svelte";
    import IconButton from "$lib/components/ui/IconButton.svelte";
    import type { Permission } from "$lib/types/permission";
    import { Users, type User } from "$lib/types/users";
    import type { Doctor } from "$lib/types/users/doctor";
    import { Unlink, Users2 } from "@lucide/svelte";
    import { onMount } from "svelte";

    let {
        user,
        permissions,
    }: {
        user: Doctor;
        permissions: Permission[];
    } = $props();

    let assistants: User<Users.Assistant>[] = $state([]);

    onMount(async () => {
        if (user.type === Users.Doctor) {
            assistants = user.assistants;
        } else {
            assistants = await AllAPI.listAllAssistants();
        }
    });
</script>

<Block group="assistants" title="My Assistants" Icon={Users2}>
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
            {#each assistants as assistant}
                <tr>
                    <td>
                        <Avatar
                            size="48px"
                            avatarUrl={assistant.avatarUrl}
                            alt={assistant.getFullName()}
                        />
                    </td>
                    <td>{assistant.getFullName()}</td>
                    <td>{assistant.email}</td>
                    <td>{assistant.phoneNumber || "N/A"}</td>
                    <td class="actions">
                        {#if permissions.includes("edit_assistant")}
                            <button>Edit</button>
                        {/if}

                        {#if permissions.includes("assign_assistant")}
                            <IconButton
                                Icon={Unlink}
                                type="error"
                                label="Unlink Assistant"
                            />
                        {/if}
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</Block>

<style>
    .actions {
        display: flex;
        gap: 0.5rem;
    }
</style>

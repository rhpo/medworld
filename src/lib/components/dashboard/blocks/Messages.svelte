<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import { user } from "$lib/stores";
    import type { Admin } from "$lib/types/users/admin";
    import type { Doctor } from "$lib/types/users/doctor";

    let messages = ($user as Doctor | Admin).messages || [];
</script>

<Block group="messages" title="Messages">
    {#if messages.length === 0}
        <p>No messages available.</p>
    {:else}
        <ul>
            {#each messages as message}
                <li>
                    <div>
                        <h3>
                            {message.sender.getFullName()}
                        </h3>

                        <span
                            >{new Date(message.date).toLocaleDateString()}</span
                        >
                    </div>

                    <p>{message.content}</p>
                </li>
            {/each}
        </ul>
    {/if}
</Block>

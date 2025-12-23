<script lang="ts">
    import Block from "$lib/components/ui/Block.svelte";
    import IconButton from "$lib/components/ui/IconButton.svelte";
    import Input from "$lib/components/ui/Input.svelte";
    import { user } from "$lib/stores";
    import { validate, validation, type Fillable } from "$lib/validation";
    import type { IDoctor, IUser } from "$lib/types/users";
    import type { Admin } from "$lib/types/users/admin";
    import type { Doctor } from "$lib/types/users/doctor";
    import { MessagesSquare, Send } from "@lucide/svelte";

    let messages = ($user as IDoctor).messages || [];
    let receiver: IDoctor = $state($user as IDoctor);

    const messageValidator = (msg: string) => {
        return msg.trim().length > 0 ? "" : "Message cannot be empty.";
    };

    let formData: Fillable = $state({
        message: {
            value: "",
            error: "",
            validator: messageValidator,
        },
    });

    function validateField(fieldName: keyof typeof formData) {
        const field = formData[fieldName];
        const validationError = field.validator
            ? field.validator(String(field.value ?? ""))
            : "";
        field.error = validationError;
    }

    function sendMessage() {
        validateField("message");

        if (formData.message.error) {
            return;
        }

        // Handle send message logic here
        console.log("Sending message:", formData.message.value);
        formData.message.value = "";
        formData.message.error = "";
    }
</script>

<Block group="messages" title="Messages" Icon={MessagesSquare}>
    {#if messages.length === 0}
        <p class="no-messages">No messages available.</p>
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

                    <p>
                        {message.content.message}
                        {#if message.content.attachment}
                            -
                            {message.content.attachment.notes}
                        {/if}
                    </p>
                </li>
            {/each}
        </ul>

        <div class="actions">
            <div class="friend">
                <Input
                    category="select"
                    bind:value={receiver}
                    options={[
                        { value: $user, label: "Manel Manel" },
                        { value: $user, label: "Said Saifi" },
                        ...($user as IDoctor).cabinet.doctors.map((doctor) => ({
                            value: doctor,
                            label: doctor.getFullName(),
                        })),
                    ]}
                />
            </div>
            <div class="message-input">
                <Input
                    bind:value={formData.message.value}
                    bind:error={formData.message.error}
                    validation={formData.message.validator}
                    placeholder="Send Message..."
                />
            </div>
            <div class="send-button">
                <IconButton
                    Icon={Send}
                    label="Send Message"
                    onclick={sendMessage}
                />
            </div>
        </div>
    {/if}
</Block>

<style>
    .actions {
        margin-top: 2rem;
        display: flex;
        width: 100%;
        gap: 0.5rem;
        height: fit-content;
    }

    .message-input {
        flex: 1;
    }

    ul {
        list-style: none;
        padding: 0;
        margin: 0;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    li {
        background: var(--background-secondary, #f9fafb);
        border: 1px solid var(--border-color, #e5e7eb);
        border-radius: 1rem;
        padding: 1rem 1.25rem;
        transition: all 0.2s ease;
        box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
    }

    li:hover {
        background: var(--background-hover, #f3f4f6);
        transform: translateY(-1px);
        box-shadow: 0 3px 6px rgba(0, 0, 0, 0.08);
    }

    li div {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 0.5rem;
    }

    h3 {
        font-size: 1rem;
        font-weight: 600;
        color: var(--text-primary, #111827);
        margin: 0;
    }

    span {
        font-size: 0.85rem;
        color: var(--text-muted, #6b7280);
    }

    p {
        font-size: 0.95rem;
        color: var(--text-secondary, #374151);
        line-height: 1.5;
        margin: 0;
    }

    p:empty::before {
        content: "No content";
        color: #9ca3af;
        font-style: italic;
    }

    .no-messages {
        text-align: center;
        padding: 2rem 0;
        color: var(--text-muted, #6b7280);
        font-style: italic;
    }
</style>

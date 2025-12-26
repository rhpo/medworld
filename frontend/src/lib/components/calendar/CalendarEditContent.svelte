<script lang="ts">
  import type { Calendar } from "$lib/types/calendar";
  import Button from "$lib/components/ui/Button.svelte";

  let {
    calendar,
    loading,
    error,
    editedAvailability,
    saving,
    saveSuccess,
    onAddSlot,
    onRemoveSlot,
    onUpdateSlot,
    onSave,
  }: {
    calendar: Calendar | null;
    loading: boolean;
    error: string | null;
    editedAvailability: Array<{ date: string; slots: string[] }>;
    saving: boolean;
    saveSuccess: boolean;
    onAddSlot: (dateIndex: number) => void;
    onRemoveSlot: (dateIndex: number, slotIndex: number) => void;
    onUpdateSlot: (dateIndex: number, slotIndex: number, value: string) => void;
    onSave: () => void;
  } = $props();
</script>

{#if loading}
  <p>Loading calendar...</p>
{:else if error}
  <p class="error">Error: {error}</p>
{:else if calendar}
  <div class="calendar-edit">
    {#if saveSuccess}
      <div class="success-message">Calendar saved successfully!</div>
    {/if}

    <div class="availability-section">
      <h4>Edit Available Slots</h4>
      {#each editedAvailability as availability, dateIndex}
        <div class="date-group">
          <div class="date-header">
            <label>Date: {new Date(availability.date).toDateString()}</label>
          </div>
          <div class="slots-container">
            {#each availability.slots as slot, slotIndex}
              <div class="slot-input-group">
                <input
                  type="time"
                  value={slot}
                  onchange={(e) =>
                    onUpdateSlot(dateIndex, slotIndex, e.currentTarget.value)}
                />
                <button
                  type="button"
                  class="remove-btn"
                  onclick={() => onRemoveSlot(dateIndex, slotIndex)}
                >
                  Remove
                </button>
              </div>
            {/each}
            <button
              type="button"
              class="add-btn"
              onclick={() => onAddSlot(dateIndex)}
            >
              + Add Time Slot
            </button>
          </div>
        </div>
      {/each}
    </div>

    <div class="actions">
      <Button onClick={onSave} disabled={saving} category="primary">
        {saving ? "Saving..." : "Save Calendar"}
      </Button>
    </div>
  </div>
{:else}
  <p>No calendar found</p>
{/if}

<style>
  .calendar-edit {
    padding: 1rem;
  }

  .info-section {
    margin-bottom: 2rem;
    padding: 1rem;
    background-color: #f8fafc;
    border-radius: 0.5rem;
  }

  .info-section h3 {
    margin: 0.5rem 0;
    color: #1e293b;
  }

  .availability-section {
    margin: 2rem 0;
  }

  .availability-section h4 {
    margin-bottom: 1rem;
    color: #1e293b;
  }

  .date-group {
    margin-bottom: 1.5rem;
    padding: 1rem;
    background-color: #f8fafc;
    border-radius: 0.5rem;
    border-left: 4px solid var(--color-primary);
  }

  .date-header {
    margin-bottom: 1rem;
  }

  .date-header label {
    font-weight: 600;
    color: #1e293b;
  }

  .slots-container {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .slot-input-group {
    display: flex;
    gap: 0.5rem;
    align-items: center;
  }

  input[type="time"] {
    flex: 1;
    padding: 0.5rem;
    border: 1px solid #cbd5e1;
    border-radius: 0.25rem;
    font-size: 1rem;
  }

  .remove-btn {
    padding: 0.5rem 1rem;
    background-color: #ef4444;
    color: white;
    border: none;
    border-radius: 0.25rem;
    cursor: pointer;
    font-size: 0.875rem;
  }

  .remove-btn:hover {
    background-color: #dc2626;
  }

  .add-btn {
    padding: 0.5rem 1rem;
    background-color: #10b981;
    color: white;
    border: none;
    border-radius: 0.25rem;
    cursor: pointer;
    font-size: 0.875rem;
    align-self: flex-start;
  }

  .add-btn:hover {
    background-color: #059669;
  }

  .actions {
    margin-top: 2rem;
    padding-top: 1rem;
    border-top: 1px solid #e2e8f0;
    display: flex;
    gap: 1rem;
  }

  .error {
    color: #dc2626;
    padding: 1rem;
    background-color: #fee2e2;
    border-radius: 0.5rem;
  }

  .success-message {
    color: #059669;
    padding: 1rem;
    background-color: #d1fae5;
    border-radius: 0.5rem;
    margin-bottom: 1rem;
  }
</style>

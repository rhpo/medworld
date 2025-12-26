<script lang="ts">
  import type { Calendar } from "$lib/types/calendar";
  import { apiClient } from "$lib/api/client";
  import { page } from "$app/stores";
  import CalendarEditContent from "$lib/components/calendar/CalendarEditContent.svelte";
  import View from "$lib/components/ui/View.svelte";

  let calendar: Calendar | null = $state(null);
  let loading = $state(true);
  let error = $state<string | null>(null);
  let saving = $state(false);
  let saveSuccess = $state(false);

  // Form state
  let editedAvailability = $state<
    Array<{
      date: string;
      slots: string[];
    }>
  >([]);

  async function loadCalendar() {
    try {
      const id = $page.params.id;
      calendar = await apiClient.get<Calendar>(`/calendars/${id}`);
      editedAvailability = JSON.parse(JSON.stringify(calendar.availability));
    } catch (err) {
      error = err instanceof Error ? err.message : "An error occurred";
      console.error("Calendar load error:", err);
    } finally {
      loading = false;
    }
  }

  async function saveCalendar() {
    saving = true;
    error = null;
    saveSuccess = false;

    try {
      const id = $page.params.id;
      await apiClient.put(`/calendars/${id}`, {
        availability: editedAvailability,
      });

      saveSuccess = true;
      setTimeout(() => {
        saveSuccess = false;
      }, 3000);
    } catch (err) {
      error = err instanceof Error ? err.message : "An error occurred";
      console.error("Calendar save error:", err);
    } finally {
      saving = false;
    }
  }

  function addSlot(dateIndex: number) {
    editedAvailability[dateIndex].slots.push("");
    editedAvailability = editedAvailability;
  }

  function removeSlot(dateIndex: number, slotIndex: number) {
    editedAvailability[dateIndex].slots.splice(slotIndex, 1);
    editedAvailability = editedAvailability;
  }

  function updateSlot(dateIndex: number, slotIndex: number, value: string) {
    editedAvailability[dateIndex].slots[slotIndex] = value;
    editedAvailability = editedAvailability;
  }

  $effect(() => {
    loadCalendar();
  });
</script>

<View>
  <CalendarEditContent
    {calendar}
    {loading}
    {error}
    {editedAvailability}
    {saving}
    {saveSuccess}
    onAddSlot={addSlot}
    onRemoveSlot={removeSlot}
    onUpdateSlot={updateSlot}
    onSave={saveCalendar}
  />
</View>

<script lang="ts">
  import type { Appointment } from "$lib/types/appointment";
  import Modal from "$lib/components/ui/Modal.svelte";

  let {
    appointment,
    onClose,
  }: { appointment: Appointment; onClose: () => void } = $props();
</script>

<Modal isOpen={true} {onClose} title="Appointment Details">
  <div class="detail-group">
    <label>Patient</label>
    <p>{appointment.patient.fullName}</p>
  </div>

  <div class="detail-group">
    <label>Doctor</label>
    <p>{appointment.doctor.fullName}</p>
  </div>

  <div class="detail-group">
    <label>Cabinet</label>
    <p>{appointment.cabinet.name}</p>
  </div>

  <div class="detail-group">
    <label>Date</label>
    <p>{new Date(appointment.date).toLocaleString()}</p>
  </div>

  <div class="detail-group">
    <label>Status</label>
    <p class="status {appointment.status.toLowerCase()}">
      {appointment.status}
    </p>
  </div>
</Modal>

<style>
  .detail-group {
    margin-bottom: 1.5rem;
  }

  .detail-group label {
    display: block;
    font-weight: 600;
    color: #374151;
    margin-bottom: 0.5rem;
    font-size: 0.875rem;
    text-transform: uppercase;
    letter-spacing: 0.05em;
  }

  .detail-group p {
    margin: 0;
    color: #1f2937;
    font-size: 1rem;
  }

  .status {
    border-radius: 9999px;
    border: 1px solid;
    width: fit-content;
    padding: 0.25rem 1rem;
    text-transform: capitalize;
    display: inline-block;
  }

  .status.completed {
    background: rgba(172, 255, 47, 0.128);
    border-color: rgba(172, 255, 47, 0.409);
    color: #4b8300;
  }

  .status.confirmed {
    background: rgba(47, 255, 158, 0.128);
    border-color: rgba(4, 135, 102, 0.409);
    color: #036b4d;
  }

  .status.in_progress {
    background: rgba(255, 245, 47, 0.128);
    border-color: rgba(214, 230, 0, 0.409);
    color: #858e00;
  }

  .status.cancelled {
    background: rgba(255, 47, 47, 0.128);
    border-color: rgba(218, 40, 40, 0.409);
    color: #c53030;
  }

  .status.scheduled {
    background: rgba(44, 139, 234, 0.128);
    border-color: rgba(44, 139, 234, 0.409);
    color: #2c8bea;
  }

  .status.no_show {
    background: rgba(107, 114, 128, 0.128);
    border-color: rgba(107, 114, 128, 0.409);
    color: #374151;
  }
</style>

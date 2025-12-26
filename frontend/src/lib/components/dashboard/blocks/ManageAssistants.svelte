<script lang="ts">
  import { AllAPI, CabinetAPI } from "$lib/api";
  import Avatar from "$lib/components/ui/Avatar.svelte";
  import Block from "$lib/components/ui/Block.svelte";
  import IconButton from "$lib/components/ui/IconButton.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import Input from "$lib/components/ui/Input.svelte";
  import UserSelectorModal from "$lib/components/dashboard/modals/UserSelectorModal.svelte";
  import type { Permission } from "$lib/types/permission";
  import { Users, type IDoctor, type User } from "$lib/types/users";
  import type { Doctor } from "$lib/types/users/doctor";
  import { Trash2, Users2, Plus, Pen } from "@lucide/svelte";
  import { onMount } from "svelte";
  import type { Assistant } from "$lib/types/users/assistant";
  import {
    extract,
    validate,
    validation,
    type Fillable,
  } from "$lib/validation";

  let {
    user,
    permissions,
  }: {
    user: Doctor;
    permissions: Permission[];
  } = $props();

  let assistants: Assistant[] = $state([]);
  let showAddModal = $state(false);
  let cabinetDoctors: Doctor[] = $state([]);
  let selectedDoctorId = $state<string>("");
  let showCreateForm = $state(false);
  let creating = $state(false);
  let createAssignedDoctorId = $state<string>("");

  let createFormData: Fillable = $state({
    firstName: { value: "", error: "", validator: validation.notEmpty },
    lastName: { value: "", error: "", validator: validation.notEmpty },
    email: { value: "", error: "", validator: validation.email },
    password: { value: "", error: "", validator: validation.password },
  });

  async function loadAssistants() {
    try {
      if (
        user.type === Users.Doctor ||
        (user.type === Users.Admin && (user as IDoctor).assistant)
      ) {
        // @ts-ignore
        assistants = [(user as IDoctor).assistant];
      } else if (
        (user.type === Users.Admin || user.type === Users.Doctor) &&
        (user as any).cabinetId
      ) {
        assistants = await CabinetAPI.getAssistants((user as any).cabinetId);
      } else {
        // If the user doesn't have a cabinet or is not a doctor/admin,
        // only superadmins should be able to list all assistants.
        if (user.type === Users.SuperAdmin) {
          assistants = await AllAPI.listAllAssistants();
        } else {
          console.log(
            "ManageAssistants: User is not authorized to list all assistants.",
          );
          assistants = [];
        }
      }
    } catch (e) {
      console.error("Failed to load assistants:", e);
      assistants = [];
    }
  }

  onMount(async () => {
    await loadAssistants();

    // For admin, allow choosing which cabinet doctor the assistant will be assigned to
    try {
      if (
        (user.type === Users.Admin || user.type === Users.Doctor) &&
        (user as any).cabinetId
      ) {
        cabinetDoctors = await CabinetAPI.getDoctors((user as any).cabinetId);
      }
    } catch (e) {
      cabinetDoctors = [];
    }

    if (user.type === Users.Doctor) {
      selectedDoctorId = String((user as any).doctorId || "");
    }
  });

  async function handleAddExisting(email: string) {
    try {
      const cabinetId = (user as any).cabinetId;
      const doctorId = selectedDoctorId ? Number(selectedDoctorId) : undefined;

      if (user.type === Users.Admin && !doctorId) {
        alert("Please select a doctor to assign this assistant to.");
        return;
      }

      await CabinetAPI.addAssistant(cabinetId, email, doctorId);
      await loadAssistants();
      showAddModal = false;
    } catch (e) {
      console.error(e);
      alert("Failed to add assistant. Check email and try again.");
    }
  }

  async function handleRemove(assistantId: number) {
    if (
      !confirm(
        "Are you sure you want to remove this assistant from the cabinet?",
      )
    )
      return;
    try {
      // @ts-ignore
      await CabinetAPI.removeAssistant((user as any).cabinetId, assistantId);
      await loadAssistants();
    } catch (e) {
      console.error(e);
      alert("Failed to remove assistant.");
    }
  }

  async function handleAssign(assistantId: number, doctorId: number) {
    try {
      const cabinetId = (user as any).cabinetId;
      await CabinetAPI.assignAssistant(cabinetId, assistantId, doctorId);
      await loadAssistants();
    } catch (e: any) {
      console.error(e);
      alert(e?.message || "Failed to update assistant assignment.");
    }
  }

  async function handleCreateAssistant() {
    try {
      const cabinetId = (user as any).cabinetId;
      const doctorId = createAssignedDoctorId
        ? Number(createAssignedDoctorId)
        : 0;

      const formError = validate(createFormData);
      if (formError) {
        alert(formError);
        return;
      }

      const cleanData = extract(createFormData) as any;
      const firstName = String(cleanData.firstName || "").trim();
      const lastName = String(cleanData.lastName || "").trim();
      const email = String(cleanData.email || "").trim();
      const password = String(cleanData.password || "");

      if (!doctorId) {
        alert("Please select a doctor to assign this assistant to.");
        return;
      }
      creating = true;
      await CabinetAPI.createAssistant(cabinetId, {
        firstName,
        lastName,
        email,
        password,
        doctorId,
      });
      await loadAssistants();
      showCreateForm = false;
      createFormData.firstName.value = "";
      createFormData.lastName.value = "";
      createFormData.email.value = "";
      createFormData.password.value = "";
      createAssignedDoctorId = "";
    } catch (e: any) {
      console.error(e);
      alert(e?.message || "Failed to create assistant.");
    } finally {
      creating = false;
    }
  }
</script>

<Block group="assistants" title="My Assistants" Icon={Users2}>
  {#if assistants.length > 0}
    <table>
      <thead>
        <tr>
          <th class="desktop">Avatar</th>
          <th>Name</th>
          <th>Email</th>
          <th>Phone</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each assistants as assistant}
          <tr>
            <td class="desktop">
              <Avatar
                size="48px"
                avatarUrl={assistant.avatarUrl}
                alt={assistant.fullName}
              />
            </td>
            <td>{assistant.fullName}</td>
            <td>{assistant.email}</td>
            <td>{assistant.phoneNumber || "N/A"}</td>
            {#if user.type === Users.Admin}
              <td>
                <select
                  style="width: 100%"
                  value={String((assistant as any)?.doctorId || "")}
                  onchange={(e) => {
                    const v = Number(
                      (e.target as HTMLSelectElement).value || "0",
                    );
                    handleAssign(assistant.id, v);
                  }}
                >
                  <option value="0">Unassigned</option>
                  {#each cabinetDoctors as d}
                    <option value={String((d as any).doctorId || d.id)}
                      >{d.fullName}</option
                    >
                  {/each}
                </select>
              </td>
            {/if}
            <td
              class:actions={!permissions.includes("edit_assistant") &&
                !(
                  permissions.includes("assign_assistant") &&
                  user.type === "admin"
                )}
            >
              {#if permissions.includes("edit_assistant")}
                <IconButton
                  href="/dashboard/users/{assistant.id}/modify"
                  label="Edit Assistant"
                  target="_blank"
                  Icon={Pen}
                />
              {/if}

              {#if permissions.includes("assign_assistant") && user.type === "admin"}
                <IconButton
                  Icon={Trash2}
                  type="error"
                  label="Unlink Assistant"
                  onClick={() => handleRemove(assistant.id)}
                />
              {/if}

              {#if !permissions.includes("edit_assistant") && !(permissions.includes("assign_assistant") && user.type === "admin")}
                Disallowed.
              {/if}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {:else}
    <h3>No assistants found</h3>
  {/if}

  {#if user.type === "admin"}
    <div class="add-actions">
      <div style="width: 100%">
        <Input
          category="input"
          theme="secondary"
          type="text"
          placeholder="Assign to doctorId (select a doctor first)"
          value={selectedDoctorId}
          disabled
        />
        <select
          style="margin-top: 0.5rem; width: 100%"
          bind:value={selectedDoctorId}
        >
          <option value="">Select a doctor</option>
          {#each cabinetDoctors as d}
            <option value={String((d as any).doctorId || d.id)}
              >{d.fullName}</option
            >
          {/each}
        </select>
      </div>

      <Button
        label="Add an existing Assistant"
        Icon={Plus}
        onClick={() => (showAddModal = true)}
        category="secondary"
        style="margin-top: 1.5rem; width: 100%"
      ></Button>

      <Button
        label={showCreateForm ? "Cancel" : "Create New Assistant"}
        Icon={Plus}
        onClick={() => (showCreateForm = !showCreateForm)}
        category="secondary"
        style="margin-top: 1.5rem; width: 100%"
      />
    </div>

    {#if showCreateForm}
      <div style="margin-top: 1rem">
        <Input
          label="First Name"
          showLabel
          validation={createFormData.firstName.validator}
          bind:value={createFormData.firstName.value}
        />
        <Input
          label="Last Name"
          showLabel
          validation={createFormData.lastName.validator}
          bind:value={createFormData.lastName.value}
        />
        <Input
          label="Email"
          showLabel
          type="email"
          validation={createFormData.email.validator}
          bind:value={createFormData.email.value}
        />
        <Input
          label="Password"
          showLabel
          type="password"
          bind:value={createFormData.password.value}
        />
        <select
          style="margin-top: 0.5rem; width: 100%"
          bind:value={createAssignedDoctorId}
        >
          <option value="">Select a doctor</option>
          {#each cabinetDoctors as d}
            <option value={String((d as any).doctorId || d.id)}
              >{d.fullName}</option
            >
          {/each}
        </select>
        <Button
          label={creating ? "Creating..." : "Create"}
          onClick={handleCreateAssistant}
          category="primary"
          style="margin-top: 0.75rem; width: 100%"
        />
      </div>
    {/if}
  {/if}
</Block>

<UserSelectorModal
  isOpen={showAddModal}
  type="assistant"
  onClose={() => (showAddModal = false)}
  onSelect={handleAddExisting}
/>

<style>
  .actions {
    display: flex;
    gap: 0.5rem;
  }

  .add-actions {
    display: flex;
    gap: 1rem;
  }
</style>

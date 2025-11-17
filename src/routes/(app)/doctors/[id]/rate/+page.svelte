<script lang="ts">
  import { page } from "$app/stores";
  import Stars from "$lib/components/ui/Stars.svelte";
  import Input from "$lib/components/ui/Input.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import { DoctorAPI } from "$lib/api";

  import type {
    RatingDoctor,
    UserExperience,
    EquippementRating,
  } from "$lib/types/rating";
  import { DEFAULT_WEIGHTS } from "$lib/types/rating";
  import View from "$lib/components/ui/View.svelte";
  import Switch from "$lib/components/ui/Switch.svelte";
  import { goto } from "$app/navigation";

  let doctorId = $page.params.id;
  let doctor = DoctorAPI.getById(Number(doctorId));

  // -----------------------------
  // Full rating state
  // -----------------------------
  let form: RatingDoctor = $state({
    id: 0,
    patient: null as any,
    cabinet: null as any,
    date: new Date(),
    equippement: {
      disponibility: false,
      rating: 0,
    } as EquippementRating,
    userExperience: {
      reception: 0,
      hygiene: 0,
      waitTime: 0,
      communication: 0,
      professionalism: 0,
      emplacement: 0,
    } as UserExperience,
    review: "",
  });

  // -----------------------------
  // Generate a step for each field
  // -----------------------------
  const steps: { key: string; type: string; label: string }[] = [];

  // Equippement
  steps.push({
    key: "equippement.disponibility",
    type: "boolean",
    label: "Is the Equipment available?",
  });
  steps.push({
    key: "equippement.rating",
    type: "stars",
    label: "Rate the Equipment (1–5)",
  });

  // UserExperience
  for (const k of Object.keys(DEFAULT_WEIGHTS.userExperienceWeights)) {
    steps.push({
      key: `userExperience.${k}`,
      type: "stars",
      label: `Rate ${capitalisationFull(k)}`,
    });
  }

  // Review
  steps.push({
    key: "review",
    type: "textarea",
    label: "Write your review (optional)",
  });

  function capitalisation(name: string) {
    return name[0].toUpperCase() + name.substring(1);
  }

  function capitalisationFull(name: string) {
    return name.split(" ").map(capitalisation);
  }

  // -----------------------------
  // Step navigation
  // -----------------------------
  let currentStep = $state(0);

  function getFieldValue(fieldKey: string) {
    const [obj, key] = fieldKey.split(".");
    return key ? form[obj][key] : form[fieldKey];
  }

  function updateField(path: string, value: any) {
    const parts = path.split(".");
    if (parts.length === 2) {
      form[parts[0]][parts[1]] = value;
    } else {
      form[path] = value;
    }
  }

  function submitReview() {
    console.log("Submitting:", form);
    alert("Submitted!");
    goto("/dashboard");
  }
</script>

{#await doctor}
  Loading doctor...
{:then doctor}
  {#if !doctor}
    Doctor Not Found
  {:else}
    <h2>Rate Dr. {doctor.getFullName()}</h2>
    <main class="rate-doctor-page">
      {#if currentStep < steps.length}
        <div class="step">
          <h2>{steps[currentStep].label}</h2>

          {#if steps[currentStep].type === "boolean"}
            <div style="margin-top: 1rem">
              <Switch
                onChange={(e) => updateField(steps[currentStep].key, e)}
              />
            </div>
          {:else if steps[currentStep].type === "stars"}
            <Stars
              max={5}
              value={getFieldValue(steps[currentStep].key)}
              onChange={(e) => updateField(steps[currentStep].key, e.detail)}
            />
          {:else if steps[currentStep].type === "textarea"}
            <Input
              category="textarea"
              placeholder="Write here..."
              bind:value={form.review}
            />
          {/if}

          <div style="margin-top:1rem;">
            {#if currentStep > 0}
              <Button
                label="Back"
                category="secondary"
                onClick={() => currentStep--}
              />
            {/if}

            {#if currentStep < steps.length - 1}
              <Button label="Next" onClick={() => currentStep++} />
            {:else}
              <Button label="Submit" onClick={submitReview} />
            {/if}
          </div>
        </div>
      {/if}
    </main>
  {/if}
{/await}

<style>
  .rate-doctor-page {
    max-width: 500px;
    margin: 3rem auto;
    padding: 2rem;
    background: #fefefe;
    border-radius: 1rem;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
    font-family: "Segoe UI", sans-serif;
  }

  h2 {
    text-align: center;
    font-size: 1.5rem;
    margin-bottom: 1.5rem;
    color: #333;
  }

  .step {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .step h2 {
    margin-bottom: 0;
  }

  input[type="checkbox"] {
    width: 20px;
    height: 20px;
    cursor: pointer;
  }

  .navigation {
    display: flex;
    justify-content: space-between;
    width: 100%;
    margin-top: 2rem;
  }

  Button {
    min-width: 120px;
  }

  /* Progress bar */
  .progress-container {
    width: 100%;
    background: #eee;
    border-radius: 0.5rem;
    height: 8px;
    margin-bottom: 2rem;
    overflow: hidden;
  }

  .progress-bar {
    height: 100%;
    background: #4f46e5;
    transition: width 0.3s ease;
  }
</style>

<script lang="ts">
  // bindable value
  export let value = 0;
  export let max = 5;

  export let onChange = (e) => {};
  function setRating(n: number) {
    value = n;
    onChange(n);
  }
</script>

<div class="rating">
  {#each Array(max) as _, i}
    <span
      class="star"
      class:active={i < value}
      on:click={() => setRating(i + 1)}>★</span
    >
  {/each}
</div>

<style>
  .rating {
    display: inline-flex;
    gap: 4px;
    cursor: pointer;
  }

  .star {
    font-size: 24px;
    opacity: 0.4;
    transition: opacity 0.15s ease;
  }

  .star.active {
    color: goldenrod;
    opacity: 1;
  }
</style>

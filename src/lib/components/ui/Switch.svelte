<script lang="ts">
  export let checked: boolean = false;
  export let disabled: boolean = false;
  export let onChange = (e: boolean) => {};

  function toggle() {
    if (!disabled) checked = !checked;
    onChange(checked);
  }
</script>

<div
  class="switch {checked ? 'on' : 'off'} {disabled ? 'disabled' : ''}"
  on:click={toggle}
  role="switch"
  aria-checked={checked}
  aria-disabled={disabled}
>
  <div class="thumb"></div>
</div>

<style>
  .switch {
    width: 50px;
    height: 28px;
    border-radius: 14px;
    background-color: #ccc;
    position: relative;
    cursor: pointer;
    transition: background-color 0.25s;
    display: inline-flex;
    align-items: center;
    padding: 2px;
  }

  .switch.on {
    /* background-color: #4cd964; */
    background-color: var(--color-primary-dark);
  }

  .switch.disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .thumb {
    width: 24px;
    height: 24px;
    background-color: white;
    border-radius: 50%;
    transition: transform 0.25s;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
  }

  .switch.on .thumb {
    transform: translateX(22px);
  }

  .switch.off .thumb {
    transform: translateX(0);
  }
</style>

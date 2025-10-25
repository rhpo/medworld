<script lang="ts">
  import "../../app.css";

  import logo from "$lib/assets/logo.svg";
  import Logo from "$lib/components/Logo.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import View from "$lib/components/ui/View.svelte";
  import TopNotification from "$lib/components/TopNotification.svelte";
  import { onMount } from "svelte";
  import { Lightbulb } from "@lucide/svelte";
  import { SITE_NAME } from "$lib";

  let { children } = $props();

  let isFull = $state(false);

  onMount(() => {
    const handleScroll = () => {
      isFull = window.scrollY > 28;
    };

    window.addEventListener("scroll", handleScroll);

    handleScroll();

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  });
</script>

<svelte:head>
  <link rel="icon" href={logo} />
  <title>{SITE_NAME}</title>
</svelte:head>

<main class:full={isFull}>
  <TopNotification
    title={"Lorem ipsum dolor sit amet, consectetur adipiscing elit."}
  />

  <View>
    <nav>
      <a class="logo" href="/">
        <div class="icon">
          <Logo rotate />
        </div>
        <h1 class:invisible={!isFull} class="title">{SITE_NAME}</h1>
      </a>

      <ul class="items" class:invisible={!isFull}>
        <li><a href="/">Features</a></li>
        <li><a href="/about">About</a></li>
        <li><a href="/pricing">Pricing</a></li>
        <li><a href="/contact">Contact</a></li>
      </ul>

      <div class="actions">
        <div class:invisible={!isFull}>
          <Button category="third" href="/login">Log in</Button>
        </div>
        <Button Icon={Lightbulb} large href="/get-started">Get started</Button>
      </div>
    </nav>
  </View>
</main>

{@render children?.()}

<style>
  * {
    transition: opacity var(--transition-duration) ease-in-out;
  }

  .logo * {
    transition: all var(--transition-duration) var(--transition-easing);
  }

  .logo:hover h1 {
    color: var(--color-primary-dark);
  }

  .logo:hover .icon {
    transform: rotate(45deg) !important;
  }

  .logo .icon {
    height: 100%;
  }

  .actions {
    display: flex;
    flex-direction: row;
    gap: 1rem;

    align-items: center;
  }

  .title {
    font-size: 2.3rem;
    font-family: var(--font-serif);
    font-weight: 300;
  }

  .items {
    display: flex;
    flex-direction: row;
    gap: 2rem;

    list-style: none;
  }

  .items li {
    font-size: 1.2rem;
  }

  .logo {
    height: 60px;

    display: flex;
    align-items: center;
    gap: 1rem;
  }

  :global(body) {
    min-height: 300vh;
  }

  main {
    position: sticky;
    top: 0;
    left: 0;
    width: 100%;
    min-height: var(--nav-height);
    background-color: rgba(255, 255, 255, 0);

    z-index: 1000;

    margin: 0;

    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;

    transition: all 0.4s ease-in-out;
  }

  main.full {
    background-color: rgba(255, 255, 255, 0.51);
    box-shadow: 0 2px 15px rgba(0, 0, 0, 0.17);
    backdrop-filter: blur(15px);
  }

  nav {
    /* padding: 1rem 2rem; */
    padding: 1rem 0;

    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;

    width: 100%;
  }
</style>

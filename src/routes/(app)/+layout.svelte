<script lang="ts">
  import "../../app.css";

  import { onMount } from "svelte";
  import { Lightbulb } from "@lucide/svelte";
  import { links, SITE_DESCRIPTION, SITE_NAME } from "$lib";
  import { slide } from "svelte/transition";

  import Logo from "$lib/components/Logo.svelte";
  import logo from "$lib/assets/logo.svg";
  import View from "$lib/components/ui/View.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import TopNotification from "$lib/components/TopNotification.svelte";

  import { Hamburger } from "svelte-hamburgers";
  let { children } = $props();

  let isFull = $state(false);
  let isMenuOpen = $state(false);

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

<main class:full={isMenuOpen}>
  <div class="container" class:full={isFull}>
    <TopNotification
      title={"Offering the most modern cabinet management technology."}
    />

    <View>
      <nav>
        <a class="logo" href="/">
          <div class="icon">
            <Logo rotate />
          </div>
          <h1 class:invisible={!isFull} class="title">{SITE_NAME}</h1>
        </a>

        <ul class="desktop items" class:invisible={!isFull}>
          {#each links as link}
            <li><a href={link.url}>{link.name}</a></li>
          {/each}
        </ul>

        <div class="actions">
          <div class="desktop desktop-actions">
            <div class:invisible={!isFull}>
              <Button category="third" href="/login">Log in</Button>
            </div>
            <Button Icon={Lightbulb} large href="/get-started"
              >Get started</Button
            >
          </div>

          <div class="mobile hamburger">
            <Hamburger bind:open={isMenuOpen} />
          </div>
        </div>
      </nav>
    </View>

    <menu class:open={isMenuOpen} transition:slide>
      <ul class="mobile-links">
        {#each links as link}
          <li><a href={link.url}>{link.name}</a></li>
        {/each}
      </ul>

      <div class="desktop-actions">
        <Button Icon={Lightbulb} large href="/get-started">Get started</Button>
        <div class:invisible={!isFull}>
          <Button category="third" href="/login">Log in</Button>
        </div>
      </div>
    </menu>
  </div>
</main>

{@render children?.()}

<footer>
  <View center>
    <div class="footer-content">
      <div class="footer-brand">
        <div class="footer-logo">
          <Logo />
        </div>
        <h1 class="footer-name">
          {SITE_NAME}
        </h1>
        <p class="slogan">
          {SITE_DESCRIPTION}
        </p>
      </div>

      <div class="footer-links">
        <h3>Quick Links</h3>
        <ul>
          {#each links as link}
            <li><a href={link.url}>{link.name}</a></li>
          {/each}
        </ul>
      </div>

      <div class="footer-contact">
        <h3>Contact Us</h3>
        <ul>
          <li>
            <a href="tel:+213553238410">+213 (0) 553 238 410</a>
          </li>

          <li>
            <a href="mailto:contact@medworld.com">contact@medworld.com</a>
          </li>

          <li>
            <a href="https://wa.me/+213553238410">WhatsApp</a>
          </li>
        </ul>
      </div>
    </div>
  </View>
</footer>

<style>
  .mobile-links {
    list-style: none;
    padding: 0;
  }

  .mobile-links > li {
    width: fit-content;
    font-size: 1.7rem;
    font-weight: 200;
    font-family: var(--font-sans);
  }

  menu {
    height: 0;
    opacity: 0;
    overflow: hidden;
    padding: 0 2rem;
    transition: all var(--transition-duration) var(--transition-easing);

    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  menu.open {
    height: auto;
    opacity: 1;
    padding: 1rem 2rem;
    padding-top: 0.5rem;
  }

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

  .desktop-actions {
    display: flex;
    flex-direction: row;
    gap: 1rem;

    align-items: center;
  }

  .title {
    font-size: 2.3rem;
    font-family: var(--font-cool);
    font-weight: 500;
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

  main {
    position: sticky;
    top: 0;
    left: 0;
    width: 100%;
    min-height: var(--nav-height);

    z-index: 1000;
    margin: 0;

    transition: all 0.4s ease-in-out;
  }

  .container {
    position: relative;
    background-color: rgba(255, 255, 255, 0);
  }

  .container.full {
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

  .mobile {
    display: none;
  }

  @media screen and (max-width: 1050px) {
    .mobile {
      display: block;
    }

    .desktop {
      display: none;
    }

    nav {
      padding: 1rem 0.5rem 1rem 0;
    }
  }

  footer {
    background: var(--color-primary-darker);
    padding: 4rem 0;
    margin-top: 4rem;
    height: 50vh;
  }

  footer * {
    color: var(--background-primary);
  }

  .footer-content {
    width: 100%;

    display: grid;
    grid-template-columns: 2fr 1fr 1fr;
    gap: 4rem;
  }

  .footer-logo {
    width: 120px;
    margin-bottom: 1rem;
  }

  .footer-name {
    margin-bottom: 1rem;
    font-weight: 100;
    font-family: var(--font-cool);
    font-weight: 100;
  }

  .slogan {
    font-size: 1.2rem;
    max-width: 300px;
    line-height: 1.6;
  }

  .footer-links,
  .footer-contact {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .footer-links h3,
  .footer-contact h3 {
    font-size: 1.4rem;
    margin-bottom: 1.5rem;
    font-weight: 500;
  }

  .footer-links ul,
  .footer-contact ul {
    list-style: none;
    padding: 0;
  }

  .footer-links li,
  .footer-contact li {
    margin-bottom: 0.8rem;
  }

  .footer-links a,
  .footer-contact a {
    color: var(--background-primary);

    text-decoration: none;
    transition: color 0.2s ease;
  }

  .footer-links a:hover,
  .footer-contact a:hover {
    color: var(--color-primary);
  }

  @media screen and (max-width: 768px) {
    .footer-content {
      grid-template-columns: 1fr;
      gap: 2rem;
    }
  }
</style>

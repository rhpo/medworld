<script>
  import { SITE_DESCRIPTION, SITE_NAME } from "$lib";
  import Button from "$lib/components/ui/Button.svelte";
  import Face from "$lib/components/ui/Face.svelte";
  import View from "$lib/components/ui/View.svelte";
  import { ArrowRight, Brain, Lightbulb } from "@lucide/svelte";
</script>

<Face title="Landing Page" description={SITE_DESCRIPTION}>
  <main>
    <video src="./hero.mp4" autoplay muted loop poster="./hero-cover.png"
    ></video>

    <div class="hero">
      <View style="justify-content: end; padding: 2rem;">
        <h1 style="color: white; font-size: 3rem; margin-bottom: 1rem;">
          Digitally taking care
        </h1>
        <p style="color: white; font-size: 1.5rem; max-width: 600px;">
          {SITE_NAME} - Empowering your cabinets with smart digital technology.
        </p>

        <div class="actions">
          <Button category="third" href="/about" style="margin-top: 1rem;"
            >Learn more</Button
          >

          <Button
            large
            blink
            category="primary"
            href="/get-started"
            Icon={ArrowRight}
            iconPosition="right"
            style="margin-top: 1rem;">Get Started</Button
          >
        </div>
      </View>
    </div>
  </main>
</Face>

<style>
  .hero {
    z-index: 100;
    position: relative;
    height: 100%;
    background-image: linear-gradient(transparent, rgba(0, 0, 0, 0.776));
  }

  .hero h1 {
    font-weight: 400;
  }

  .actions {
    display: flex;
    gap: 1rem;
    flex-wrap: wrap;
    align-items: center;
  }

  main {
    background-image: url("./hero-cover.png");
    background-size: cover;
    background-position: center;

    height: calc(70vh - var(--nav-height));

    overflow: hidden;
    position: relative;
  }

  main video {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
    object-position: center;
  }
</style>

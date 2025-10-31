import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [sveltekit()],
  optimizeDeps: {
    include: ["@lucide/svelte"],
  },
  server: {
    allowedHosts: ["localhost", "d096ab60cc8e.ngrok-free.app"]
  }
});

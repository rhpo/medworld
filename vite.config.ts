import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [sveltekit()],
  optimizeDeps: {
    include: ["@lucide/svelte"],
  },
  server: {
    allowedHosts: ["localhost", "184570eeef3b.ngrok-free.app"]
  }
});

import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd());
  return {
    plugins: [react()],
    optimizeDeps: {
      include: ["recharts"],
    },
    server: {
      port: 8080,
    },
    define: {
      BACKEND_ADDRESS: JSON.stringify(env.BACKEND_ADDRESS),
    },
  };
});

import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd());
  return {
    plugins: [react()],
    optimizeDeps: {
      include: ['recharts'],
    },
    server: {
      port: 80,
      host: true,
      watch: {
        usePolling: true,
      },
      hmr: {
	clientHost: "energy.bubliny.at",
	clientPort: 80,
      },
    },
    build: {
      commonjsOptions: {
        include: [/recharts/, /node_modules/],
      },
    },
    define: {
      VITE_APP_BACKEND_ADDRESS: JSON.stringify(env.VITE_APP_BACKEND_ADDRESS),
    },
  };
});

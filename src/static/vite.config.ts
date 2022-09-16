import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [svelte()],
	server: {
		port: 5007,
		hmr: {
			port: 5007,
			path: 'vite-hmr-ws',
		},
	},
});

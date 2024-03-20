import path from 'path';
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      '@sass': path.resolve(__dirname, 'src/assets/sass'),
      '@images': path.resolve(__dirname, 'src/assets/images'),
    }
  },
  plugins: [react()],
  preview: {
    port: 3000,
    strictPort: true,
  },
  server: {
    host: '0.0.0.0',
    port: 3000,
  },
})

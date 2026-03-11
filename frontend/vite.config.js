import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'
import { resolve } from 'path'

export default defineConfig({
  base: '/110claw/',
  plugins: [vue(), tailwindcss()],
  resolve: {
    alias: { '@': resolve(__dirname, 'src') }
  },
  server: {
    port: 8090,
    host: '0.0.0.0',
    proxy: {
      '/110claw/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      }
    }
  }
})

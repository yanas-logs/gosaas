// nuxt.config.ts
export default defineNuxtConfig({
  compatibilityDate: '2026-06-16', // Mengikuti waktu sistem Anda saat ini

  modules: ['@nuxt/ui'],
  css: ['./app/asset/css/main.css'],
  devtools: { enabled: true },
  
  future: {
    compatibilityVersion: 4
  }
})
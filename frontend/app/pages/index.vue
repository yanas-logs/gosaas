<template>
  <div class="bg-gray-50 dark:bg-gray-950 min-h-screen pb-16">

    <section class="mx-auto max-w-7xl px-4 pt-8 sm:px-6 lg:px-8">
      <div class="w-full overflow-hidden rounded-2xl shadow-sm border border-gray-100 dark:border-gray-800">
        <UCarousel
          v-model="activeSlide"
          :items="featuresBanners"
          indicators
          arrows
          class="w-full"
        >
          <template #default="{ item }">
            <div class="relative w-full aspect-[21/9] flex items-center justify-center bg-gradient-to-br from-gray-900 via-indigo-950 to-slate-900 text-white p-6 group cursor-pointer">
              <div class="text-center z-10 max-w-xl px-4">
                <span class="inline-block text-[10px] font-bold uppercase tracking-widest bg-primary-600 text-white px-2 py-0.5 rounded mb-3">SaaS Platform Feature</span>
                <h2 class="text-xl sm:text-3xl font-black uppercase tracking-wide group-hover:text-primary-400 transition mb-2">{{ item.title }}</h2>
                <p class="text-xs sm:text-sm text-gray-300">{{ item.description }}</p>
                <div class="mt-4 flex justify-center gap-3 opacity-0 group-hover:opacity-100 transition duration-300">
                  <UButton size="sm" color="primary" to="/register">Mulai Sekarang</UButton>
                </div>
              </div>
              <div class="absolute inset-0 bg-black/40"></div>
            </div>
          </template>
        </UCarousel>
      </div>
    </section>

    <section class="mx-auto max-w-7xl px-4 pt-10 sm:px-6 lg:px-8">
      <div class="bg-white dark:bg-gray-900 border border-gray-200/60 dark:border-gray-800/60 rounded-2xl p-6 shadow-sm">
        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 lg:grid-cols-10 gap-4">
          <div v-for="(category, idx) in globalCategories" :key="idx" class="flex flex-col items-center justify-center group cursor-pointer p-2 rounded-xl hover:bg-gray-50 dark:hover:bg-gray-800/50 transition">
            <div class="w-12 h-12 rounded-xl bg-gray-50 dark:bg-gray-800 group-hover:bg-primary-50 dark:group-hover:bg-primary-950/40 flex items-center justify-center mb-2 shadow-sm transition">
              <UIcon :name="category.icon" class="w-6 h-6 text-gray-700 dark:text-gray-300 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition" />
            </div>
            <span class="text-xs font-bold text-gray-900 dark:text-white text-center group-hover:text-primary-600 dark:group-hover:text-primary-400 transition truncate w-full px-1">
              {{ category.name }}
            </span>
          </div>
        </div>
      </div>
    </section>

    <section class="mx-auto max-w-5xl px-4 pt-16 sm:px-6 lg:px-8 text-center">
      <UCard class="w-full shadow-md bg-white dark:bg-gray-900 border border-gray-100 dark:border-gray-800">
        <div class="flex flex-col items-center py-8 px-4">
          <div class="p-3 bg-primary-50 dark:bg-primary-950/50 rounded-full mb-4">
            <UIcon name="i-heroicons-bolt" class="w-10 h-10 text-primary-600 dark:text-primary-400" />
          </div>

          <h1 class="text-3xl font-extrabold text-gray-900 dark:text-white mb-2">
            Brand
          </h1>
          <p class="text-sm text-gray-600 dark:text-gray-400 mb-8 max-w-md">
            Selamat datang di platform gateway top-up multi-tenant. Mulai bisnis top-up game Anda sendiri dengan mudah.
          </p>
          
          <div class="flex flex-col sm:flex-row gap-3 w-full justify-center max-w-xs">
            <UButton to="/register" color="primary" variant="solid" size="lg" icon="i-heroicons-rocket-launch" block>
              Mulai Sekarang
            </UButton>
            <UButton to="/login" color="gray" variant="outline" size="lg" icon="i-heroicons-arrow-right-end-on-rectangle" block>
              Masuk Dashboard
            </UButton>
          </div>
        </div>
      </UCard>
    </section>

  </div>
</template>

<script setup>
import { ref } from 'vue'

const activeSlide = ref(0)

const featuresBanners = [
  { title: 'Safe & Reliable Platform', description: 'Infrastruktur gateway top-up berkecepatan tinggi dengan jaminan keamanan server lapis tiga dan proteksi transaksi enkripsi penuh.' },
  { title: 'Sistem Kemitraan Multi-Tenant', description: 'Buat, atur, dan kelola branding toko top-up Anda sendiri hanya dalam waktu 5 menit tanpa perlu memahami baris kode backend.' },
  { title: 'Metode Otomasi API Gateway', description: 'Dukungan penuh untuk sinkronisasi inventaris item game secara real-time langsung ke supplier utama Anda.' }
]

const globalCategories = [
  { name: 'Top Up', icon: 'i-heroicons-credit-card' },
  { name: 'Game Key', icon: 'i-heroicons-key' },
  { name: 'Akun', icon: 'i-heroicons-user-circle' },
  { name: 'Voucher', icon: 'i-heroicons-ticket' },
  { name: 'Roblox Games', icon: 'i-heroicons-cube' },
  { name: 'Item', icon: 'i-heroicons-shield-check' },
  { name: 'Koin Game', icon: 'i-heroicons-circle-stack' },
  { name: 'RPG Games', icon: 'i-heroicons-sparkles' },
  { name: 'Steam Key', icon: 'i-heroicons-computer-desktop' },
  { name: 'Simulation', icon: 'i-heroicons-globe-alt' }
]

useHead({
  title: 'GoSaas Topup - Home'
})
</script>
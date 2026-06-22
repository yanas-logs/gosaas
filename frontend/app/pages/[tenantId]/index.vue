<template>
  <div class="bg-gray-50 dark:bg-gray-950 min-h-screen pb-12">

    <section class="bg-white dark:bg-gray-900 border-b border-gray-200/70 dark:border-gray-800/70 py-8">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 flex flex-col items-center text-center">
        <h1 class="text-2xl font-extrabold text-gray-950 dark:text-white capitalize tracking-tight mb-2 flex items-center gap-2">
          <UIcon name="i-heroicons-shopping-bag" class="text-primary-600 dark:text-primary-400 w-6 h-6" />
          Toko {{ tenantName }}
        </h1>
        <p class="text-xs text-gray-500 dark:text-gray-400 max-w-md mb-6">
          Selamat datang di etalase resmi toko top-up milik {{ tenantName }}. Silakan pilih layanan game favorit Anda di bawah ini dengan proses instan 24 jam.
        </p>
        <div class="w-full max-w-2xl">
          <UInput
            v-model="searchGameQuery"
            icon="i-heroicons-magnifying-glass"
            placeholder="Cari Game, Diamond, Voucher..."
            size="xl"
            class="w-full shadow-sm"
            ui="rounded-full"
          />
        </div>
      </div>
    </section>

    <section class="mx-auto max-w-7xl px-4 pt-8 sm:px-6 lg:px-8">
      <div class="w-full overflow-hidden rounded-2xl shadow-sm border border-gray-100 dark:border-gray-800">
        <UCarousel
          v-model="activeSlide"
          :items="banners"
          indicators
          arrows
          class="w-full"
        >
          <template #default="{ item }">
            <div class="relative w-full aspect-21/9 flex items-center justify-center bg-linear-to-r from-gray-900 to-slate-800 text-white p-6 group cursor-pointer">
              <div class="text-center z-10">
                <span class="inline-block text-[10px] font-bold uppercase tracking-widest bg-primary-600 text-white px-2 py-0.5 rounded mb-2">Special Event</span>
                <h3 class="text-xl sm:text-2xl font-black uppercase tracking-wider group-hover:text-primary-400 transition mb-1">{{ item.title }}</h3>
                <p class="text-xs sm:text-sm text-gray-300 max-w-md mx-auto">{{ item.description }}</p>
              </div>
              <div class="absolute inset-0 bg-black/30"></div>
            </div>
          </template>
        </UCarousel>
      </div>
    </section>

    <section class="mx-auto max-w-7xl px-4 pt-10 sm:px-6 lg:px-8">
      <div class="flex items-center gap-2 mb-4">
        <UIcon name="i-heroicons-bolt" class="w-5 h-5 text-amber-500 animate-pulse" />
        <h2 class="text-lg font-bold text-gray-950 dark:text-white tracking-tight">Flash Sale Terbatas</h2>
      </div>
      <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4">
        <UCard v-for="(item, idx) in flashSales" :key="idx" class="hover:ring-1 hover:ring-primary-500 transition shadow-sm bg-white dark:bg-gray-900 border border-gray-100 dark:border-gray-800">
          <div class="flex flex-col h-full justify-between">
            <div>
              <span class="text-[10px] font-medium text-gray-400 block mb-1">{{ item.game }}</span>
              <h4 class="text-xs font-bold text-gray-900 dark:text-white line-clamp-2 mb-2">{{ item.name }}</h4>
            </div>
            <div class="mt-2">
              <span class="text-xs text-red-500 line-through block font-medium">{{ item.originalPrice }}</span>
              <span class="text-sm font-extrabold text-primary-600 dark:text-primary-400 block">{{ item.promoPrice }}</span>
              <div class="w-full bg-gray-100 dark:bg-gray-800 h-1.5 rounded-full mt-3 overflow-hidden">
                <div class="bg-primary-500 h-full rounded-full" :style="{ width: item.progress + '%' }"></div>
              </div>
              <span class="text-[9px] text-gray-400 block mt-1 text-right">{{ item.sold }} Terjual</span>
            </div>
          </div>
        </UCard>
      </div>
    </section>

    <section class="mx-auto max-w-7xl px-4 pt-10 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between border-b border-gray-200 dark:border-gray-800 pb-3 mb-6">
        <div class="flex items-center gap-2">
          <UIcon name="i-heroicons-squares-2x2" class="w-5 h-5 text-primary-600 dark:text-primary-400" />
          <h2 class="text-lg font-bold text-gray-950 dark:text-white tracking-tight">Semua Kategori Game</h2>
        </div>
      </div>
      
      <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-4">
        <UCard v-for="(game, idx) in games" :key="idx" class="hover:ring-1 hover:ring-primary-500 cursor-pointer transition shadow-sm bg-white dark:bg-gray-900 border border-gray-100 dark:border-gray-800 group">
          <div class="flex flex-col items-center py-4">
            <div class="p-3 bg-gray-50 dark:bg-gray-800 rounded-2xl mb-3 group-hover:bg-primary-50 dark:group-hover:bg-primary-950/40 transition">
              <UIcon :name="game.icon" class="w-8 h-8 text-gray-700 dark:text-gray-300 group-hover:text-primary-600 dark:group-hover:text-primary-400 transition" />
            </div>
            <span class="font-bold text-xs text-gray-900 dark:text-white text-center group-hover:text-primary-600 dark:group-hover:text-primary-400 transition px-1">
              {{ game.name }}
            </span>
          </div>
        </UCard>
      </div>
    </section>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const route = useRoute()
const tenantName = computed(() => route.params.tenantId)

const searchGameQuery = ref('')
const activeSlide = ref(0)

const banners = [
  { title: 'Promo Top-Up Kilat', description: 'Nikmati diskon gila hingga 20% untuk semua item game malam ini!' },
  { title: 'Event Diamond Murah', description: 'Proses otomatis langsung masuk ke ID game Anda dalam hitungan detik.' },
  { title: 'Metode Bayar Lengkap', description: 'Mendukung QRIS, e-Wallet, dan Virtual Account transfer bank otomatis.' },
  { title: 'Steam Summer Sale', description: 'Dapatkan Steam Wallet Code termurah dan instan di sini sekarang juga.' }
]

const flashSales = [
  { game: 'Mobile Legends', name: '2010 (1708+302) Diamonds - (Limited)', originalPrice: 'Rp 500.000', promoPrice: 'Rp 473.999', progress: 85, sold: '84/99' },
  { game: 'PUBG Mobile', name: '60 UC - (Limited)', originalPrice: 'Rp 15.999', promoPrice: 'Rp 15.899', progress: 30, sold: '15/50' },
  { game: 'PUBG Mobile', name: '8100 UC - (Limited Edition)', originalPrice: 'Rp 1.630.000', promoPrice: 'Rp 1.589.999', progress: 10, sold: '1/10' },
  { game: 'Roblox', name: '300 Robux Gift Card - (Instan)', originalPrice: 'Rp 79.000', promoPrice: 'Rp 64.099', progress: 60, sold: '6/10' },
  { game: 'Free Fire', name: '720 Diamonds - Garena Instan', originalPrice: 'Rp 100.000', promoPrice: 'Rp 91.500', progress: 45, sold: '45/100' }
]

const games = [
  { name: 'Mobile Legends', icon: 'i-heroicons-device-phone-mobile' },
  { name: 'Free Fire', icon: 'i-heroicons-fire' },
  { name: 'PUBG Mobile', icon: 'i-heroicons-shield-check' },
  { name: 'Genshin Impact', icon: 'i-heroicons-sparkles' },
  { name: 'Valorant', icon: 'i-heroicons-computer-desktop' },
  { name: 'Roblox', icon: 'i-heroicons-cube' }
]

useHead({
  title: `Toko ${tenantName.value} - GoSaas Topup`
})
</script>
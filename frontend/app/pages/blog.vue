<template>
  <div class="bg-gray-50 dark:bg-gray-950 min-h-[calc(100vh-4rem)] py-12 px-4 sm:px-6 lg:px-8">
    <div class="mx-auto max-w-7xl">

      <div class="text-center max-w-2xl mx-auto mb-12">
        <div class="inline-flex p-3 bg-primary-50 dark:bg-primary-950/40 rounded-full mb-4">
          <UIcon name="i-heroicons-document-text" class="w-10 h-10 text-primary-600 dark:text-primary-400" />
        </div>
        <h1 class="text-3xl font-extrabold text-gray-950 dark:text-white tracking-tight">
          Artikel & Blog Terbaru
        </h1>
        <p class="text-sm text-gray-500 dark:text-gray-400 mt-2">
          Temukan tips, trik seputar game, informasi event promo, dan berita terbaru dari ekosistem platform kami.
        </p>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <UCard v-for="post in blogPosts" :key="post.id" class="hover:ring-1 hover:ring-primary-500 transition shadow-sm bg-white dark:bg-gray-900 border border-gray-200/60 dark:border-gray-800/60 group flex flex-col h-full overflow-hidden">
          <div class="flex flex-col h-full justify-between">
            <div>
              <div class="flex items-center gap-2 mb-3">
                <span class="text-[10px] font-bold uppercase tracking-widest bg-primary-50 dark:bg-primary-950/50 text-primary-600 dark:text-primary-400 px-2 py-0.5 rounded">
                  {{ post.category }}
                </span>
                <span class="text-xs text-gray-400 font-medium">
                  {{ post.date }}
                </span>
              </div>

              <h3 class="text-lg font-bold text-gray-950 dark:text-white group-hover:text-primary-600 dark:group-hover:text-primary-400 transition line-clamp-2 mb-2">
                {{ post.title }}
              </h3>
              <p class="text-xs text-gray-500 dark:text-gray-400 line-clamp-3 mb-4">
                {{ post.excerpt }}
              </p>
            </div>

            <div class="flex items-center justify-between pt-4 border-t border-gray-100 dark:border-gray-800 mt-auto">
              <div class="flex items-center gap-2">
                <div class="w-6 h-6 rounded-full bg-gray-100 dark:bg-gray-800 flex items-center justify-center text-[10px] font-bold text-gray-600 dark:text-gray-400 border border-gray-200 dark:border-gray-700">
                  {{ post.authorInitials }}
                </div>
                <span class="text-xs font-semibold text-gray-700 dark:text-gray-300">
                  {{ post.author }}
                </span>
              </div>
              
              <UButton variant="ghost" color="primary" size="xs" suffix-icon="i-heroicons-arrow-right" class="font-bold">
                Baca Selengkapnya
              </UButton>
            </div>
          </div>
        </UCard>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const blogPosts = ref([
  {
    id: 1,
    category: 'Promo',
    date: '24 Juni 2026',
    title: 'Cara Mendapatkan Diskon Maksimal Event Starlight Mobile Legends',
    excerpt: 'Simak panduan lengkap mengoptimalkan pengisian diamond pada platform kami untuk mengamankan skin Starlight bulan ini dengan harga termurah.',
    author: 'Admin Jef',
    authorInitials: 'AJ'
  },
  {
    id: 2,
    category: 'Update',
    date: '20 Juni 2026',
    title: 'Integrasi Otomatis Pembayaran QRIS Berkecepatan Tinggi',
    excerpt: 'Kami baru saja memperbarui infrastruktur gateway pembayaran. Kini deteksi mutasi saldo dari e-wallet ke akun tenant Anda berjalan kurang dari 2 detik.',
    author: 'Tech Team',
    authorInitials: 'TT'
  },
  {
    id: 3,
    category: 'Tips & Trik',
    date: '15 Juni 2026',
    title: 'Strategi Mengelola Stok Voucher Game untuk Pemilik Toko Baru',
    excerpt: 'Bagi Anda yang baru saja mendaftar sistem multi-tenant, ini trik mengalokasikan modal balance agar tidak mengalami kehabisan stok saat flash sale berlangsung.',
    author: 'Sultan Ciruas',
    authorInitials: 'SC'
  }
])

useHead({
  title: 'Blog & Berita - GoSaas Topup'
})
</script>
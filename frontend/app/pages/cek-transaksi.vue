<template>
  <div class="flex min-h-[calc(100vh-4rem)] items-center justify-center bg-gray-50 dark:bg-gray-950 p-4">
    <UCard class="w-full max-w-md bg-white dark:bg-gray-900 shadow-sm border border-gray-100 dark:border-gray-800">
      <h2 class="text-xl font-bold text-gray-950 dark:text-white mb-4">Cek Transaksi</h2>
      <div class="space-y-4">
        <UInput placeholder="Masukkan Invoice ID Anda..." size="lg" class="w-full" />
        <UButton color="primary" block size="lg">Periksa Status</UButton>
      </div>
    </UCard>
  </div>
</template>

<script setup>
useHead({
  title: 'Cek Transaksi - GoSaas Topup'
})
</script>
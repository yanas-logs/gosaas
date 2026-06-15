<!-- app/pages/index.vue -->
<template>
  <div class="flex min-h-screen flex-col items-center justify-center bg-gray-50 p-4 text-center">
    <UCard class="w-full max-w-sm shadow-md">
      <h1 class="text-2xl font-bold text-gray-900 mb-2">ID Topup SaaS</h1>
      <p class="text-sm text-gray-600 mb-6">Selamat datang di platform gateway top-up multi-tenant.</p>
      
      <div class="flex flex-col gap-3">
        <UButton to="/login" color="primary" variant="solid" block>
          Masuk ke Dashboard
        </UButton>
        <UButton to="/register" color="gray" variant="ghost" block>
          Belum punya akun? Daftar
        </UButton>
      </div>
    </UCard>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: [
    function (to, from) {
      return navigateTo('/login')
    }
  ]
})
</script>
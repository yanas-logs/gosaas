<template>
  <div class="flex min-h-screen items-center justify-center px-4 py-12 sm:px-6 lg:px-8 bg-gray-100">
    <div class="w-full max-w-md space-y-8 bg-white p-8 rounded-xl shadow-md">
      <div>
        <h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">Buat Akun Baru</h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          Atau
          <NuxtLink to="/login" class="font-medium text-indigo-600 hover:text-indigo-500">Sudah punya akun? Login</NuxtLink>
        </p>
      </div>
      
      <form class="mt-8 space-y-6" @submit.prevent="handleRegister">
        <div class="space-y-4 rounded-md shadow-sm">
          <div>
            <label class="block text-sm font-medium text-gray-700">Tenant ID (UUID)</label>
            <input v-model="form.tenant_id" type="text" required class="relative block w-full rounded-md border border-gray-300 p-2 text-gray-900 focus:ring-2 focus:ring-indigo-600 sm:text-sm" placeholder="Masukkan Tenant UUID dari DBeaver" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Nama Lengkap</label>
            <input v-model="form.name" type="text" required class="relative block w-full rounded-md border border-gray-300 p-2 text-gray-900 focus:ring-2 focus:ring-indigo-600 sm:text-sm" placeholder="Nama Anda" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Email</label>
            <input v-model="form.email" type="email" required class="relative block w-full rounded-md border border-gray-300 p-2 text-gray-900 focus:ring-2 focus:ring-indigo-600 sm:text-sm" placeholder="email@contoh.com" />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Password</label>
            <input v-model="form.password" type="password" required class="relative block w-full rounded-md border border-gray-300 p-2 text-gray-900 focus:ring-2 focus:ring-indigo-600 sm:text-sm" placeholder="••••••••" />
          </div>
        </div>

        <div v-if="errorMessage" class="text-red-500 text-sm text-center bg-red-50 p-2 rounded">
          {{ errorMessage }}
        </div>
        <div v-if="successMessage" class="text-green-500 text-sm text-center bg-green-50 p-2 rounded">
          {{ successMessage }}
        </div>

        <div>
          <button type="submit" :disabled="loading" class="group relative flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white hover:bg-indigo-500 focus:ring-2 focus:ring-indigo-600 disabled:bg-indigo-400">
            {{ loading ? 'Memproses...' : 'Daftar' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const form = ref({
  tenant_id: '',
  name: '',
  email: '',
  password: ''
})

const loading = ref(false)
const errorMessage = ref('')
const successMessage = ref('')

const handleRegister = async () => {
  loading.value = true
  errorMessage.value = ''
  successMessage.value = ''
  
  try {
    await $fetch('http://localhost:8080/register', {
      method: 'POST',
      body: form.value
    })
    
    successMessage.value = 'Registrasi berhasil! Mengalihkan ke halaman login...'
    setTimeout(() => {
      navigateTo('/login')
    }, 2000)
  } catch (error) {
    errorMessage.value = error.data?.error || 'Terjadi kesalahan sistem.'
  } finally {
    loading.value = false
  }
}
</script>
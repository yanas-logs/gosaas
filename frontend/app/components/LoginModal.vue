<template>
  <UModal v-model:open="isOpen" title="Masuk ke Dashboard" description="Silakan masuk untuk melanjutkan." class="z-50">
    <template #body>
      <div class="space-y-4">
        <div>
          <label class="mb-1 block text-sm font-medium text-gray-700 dark:text-gray-200">Email</label>
          <UInput v-model="form.email" type="email" size="lg" class="w-full" placeholder="Email" />
        </div>

        <div>
          <label class="mb-1 block text-sm font-medium text-gray-700 dark:text-gray-200">Password</label>
          <UInput v-model="form.password" type="password" size="lg" class="w-full" placeholder="Password" />
        </div>

        <UAlert v-if="errorMessage" color="error" variant="soft" :title="errorMessage" />

        <UButton color="primary" variant="solid" block :loading="loading" @click="handleLogin">
          {{ loading ? 'Memasuki...' : 'Sign In' }}
        </UButton>
      </div>
    </template>
  </UModal>
</template>

<script setup>
const { isOpen, close } = useLoginModal()

const form = ref({
  email: '',
  password: ''
})

const loading = ref(false)
const errorMessage = ref('')

const handleLogin = async () => {
  loading.value = true
  errorMessage.value = ''

  try {
    const response = await $fetch('http://localhost:8080/', {
      method: 'POST',
      body: form.value
    })

    localStorage.setItem('auth_token', response.token)
    close()
    alert('Login sukses! Token JWT berhasil disimpan.')
  } catch (error) {
    errorMessage.value = error.data?.error || 'Email atau password salah.'
  } finally {
    loading.value = false
  }
}
</script>

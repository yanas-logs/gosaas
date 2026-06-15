<template>
  <section class="flex min-h-[calc(100vh-4rem)] items-center justify-center px-4 py-12 sm:px-6 lg:px-8">
    <UCard class="w-full max-w-md">
      <template #header>
        <div class="text-center">
          <h2 class="text-2xl font-semibold text-gray-900 dark:text-white">{{ t('loginTitle') }}</h2>
          <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">
            {{ t('loginSubtitle') }}
            <NuxtLink to="/register" class="font-medium text-primary-600 hover:text-primary-500">{{ t('loginRegisterLink') }}</NuxtLink>
          </p>
        </div>
      </template>

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
          {{ loading ? t('loginLoading') : t('loginButton') }}
        </UButton>
      </div>
    </UCard>
  </section>
</template>

<script setup>
const { t } = useI18n()

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
    const response = await $fetch('http://localhost:8080/login', {
      method: 'POST',
      body: form.value
    })

    localStorage.setItem('auth_token', response.token)
    alert('Login sukses! Token JWT berhasil disimpan.')
  } catch (error) {
    errorMessage.value = error.data?.error || 'Email atau password salah.'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="flex min-h-[calc(100vh-4rem)] flex-col items-center justify-center bg-gray-50 dark:bg-gray-950 p-4 text-center">
    <UCard class="w-full max-w-md shadow-md bg-white dark:bg-gray-900">
      <div class="flex flex-col items-center py-6">
        <div class="p-3 bg-primary-50 dark:bg-primary-950/50 rounded-full mb-4">
          <UIcon name="i-heroicons-bolt" class="w-10 h-10 text-primary-600 dark:text-primary-400" />
        </div>

        <h1 class="text-3xl font-extrabold text-gray-900 dark:text-white mb-2">
          Brand
        </h1>
        <p class="text-sm text-gray-600 dark:text-gray-400 mb-8 max-w-xs">
          {{ t('welcome') }}
        </p>
        
        <div class="flex flex-col sm:flex-row gap-3 w-full justify-center">
          <UButton to="/register" color="primary" variant="solid" size="lg" icon="i-heroicons-rocket-launch">
            {{ t('startNow') }}
          </UButton>
          <UButton to="/login" color="gray" variant="outline" size="lg" icon="i-heroicons-arrow-right-end-on-rectangle">
            {{ t('enterDashboard') }}
          </UButton>
        </div>
      </div>
    </UCard>
  </div>
</template>

<script setup>
import { useLanguage } from '@/composables/useLanguage.js'

const { t } = useLanguage()

useHead({
  title: 'GoSaas Topup - Home'
})
</script>
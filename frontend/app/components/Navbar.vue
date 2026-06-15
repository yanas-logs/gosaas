<template>
  <header class="sticky top-0 z-40 border-b border-gray-200/70 bg-white/75 shadow-sm backdrop-blur-xl dark:border-gray-800/70 dark:bg-gray-950/70">
    <nav class="mx-auto flex h-16 w-full max-w-7xl items-center justify-between gap-4 px-4 sm:px-6 lg:px-8">
      <NuxtLink to="/" class="flex items-center gap-2 text-lg font-semibold text-gray-900 dark:text-white">
        <span></span>
        Logo Brand
      </NuxtLink>

      <div class="hidden items-center gap-6 md:flex">
        <UButton v-for="item in links[0]" :key="item.to" :to="item.to" color="neutral" variant="ghost" size="sm">
          {{ item.label }}
        </UButton>
      </div>

      <div class="flex items-center gap-3">
        <UInput
          v-model="searchQuery"
          icon="i-heroicons-magnifying-glass"
          placeholder="Cari produk, fitur, atau topik..."
          size="sm"
          class="hidden w-56 lg:flex"
          @keyup.enter="handleSearch"
        />

        <USelect
          v-model="currentLang"
          :options="languages"
          option-attribute="label"
          value-attribute="value"
          size="sm"
          class="w-24"
          @update:model-value="handleLanguageChange"
        />

        <UButton color="primary" variant="solid" size="sm" @click="openLoginModal">
          {{ t('navSignIn') }}
        </UButton>

        <UButton color="neutral" variant="ghost" size="sm" icon="i-heroicons-moon" @click="isDark = !isDark" />
      </div>
    </nav>
  </header>
</template>

<script setup>
import { useNavbarLogic } from '@/composables/useNavbarLogic.js'

const { isDark, searchQuery, handleSearch, currentLang, languages, handleLanguageChange, links } = useNavbarLogic()
const { t } = useI18n()
const { open } = useLoginModal()

const openLoginModal = () => {
  open()
}
</script>

<style src="@/asset/css/navbar.css" scoped></style>
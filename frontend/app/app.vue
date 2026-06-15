<template>
  <div class="min-h-screen bg-gray-50 text-gray-900 dark:bg-gray-950 dark:text-gray-50">
    <Navbar />

    <main class="pt-4 md:pt-6">
      <div class="mx-auto w-full max-w-7xl px-4 pb-12 sm:px-6 lg:px-8">
        <NuxtPage />
      </div>
    </main>

    <LoginModal />
  </div>
</template>
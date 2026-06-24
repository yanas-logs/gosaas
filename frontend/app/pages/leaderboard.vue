<template>
  <div class="bg-gray-50 dark:bg-gray-950 min-h-[calc(100vh-4rem)] py-12 px-4 sm:px-6 lg:px-8">
    <div class="mx-auto max-w-4xl">
      
      <div class="text-center mb-10">
        <div class="inline-flex p-3 bg-amber-50 dark:bg-amber-950/40 rounded-full mb-4">
          <UIcon name="i-heroicons-trophy" class="w-10 h-10 text-amber-500" />
        </div>
        <h1 class="text-3xl font-extrabold text-gray-950 dark:text-white tracking-tight">
          Top Spender Leaderboard
        </h1>
        <p class="text-sm text-gray-500 dark:text-gray-400 mt-2">
          Apresiasi bagi para pelanggan setia dengan akumulasi transaksi tertinggi bulan ini.
        </p>
      </div>

      <div class="bg-white dark:bg-gray-900 rounded-2xl border border-gray-200/70 dark:border-gray-800/70 shadow-sm overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="bg-gray-50 dark:bg-gray-800/50 border-b border-gray-200 dark:border-gray-800 text-xs font-bold text-gray-500 dark:text-gray-400 uppercase tracking-wider">
                <th class="py-4 px-6 text-center w-20">Peringkat</th>
                <th class="py-4 px-6">Nama Pengguna</th>
                <th class="py-4 px-6">Total Transaksi</th>
                <th class="py-4 px-6 text-center">Status</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200 dark:divide-gray-800 text-sm">
              <tr v-for="(user, index) in leaderboardData" :key="user.id" class="hover:bg-gray-50/50 dark:hover:bg-gray-800/30 transition">
                
                <td class="py-4 px-6 text-center font-bold">
                  <span v-if="index === 0" class="flex items-center justify-center text-xl">🥇</span>
                  <span v-else-if="index === 1" class="flex items-center justify-center text-xl">🥈</span>
                  <span v-else-if="index === 2" class="flex items-center justify-center text-xl">🥉</span>
                  <span v-else class="text-gray-500 dark:text-gray-400">#{{ index + 1 }}</span>
                </td>

                <td class="py-4 px-6 font-semibold text-gray-900 dark:text-white">
                  {{ user.username }}
                </td>

                <td class="py-4 px-6 font-bold text-primary-600 dark:text-primary-400">
                  {{ formatCurrency(user.totalSpend) }}
                </td>

                <td class="py-4 px-6 text-center">
                  <span :class="[
                    'inline-flex items-center px-2.5 py-1 rounded-full text-xs font-semibold',
                    index < 3 ? 'bg-amber-50 dark:bg-amber-950/40 text-amber-700 dark:text-amber-400 border border-amber-200 dark:border-amber-800' : 'bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-400'
                  ]">
                    {{ index < 3 ? 'Elite Spender' : 'Active Member' }}
                  </span>
                </td>

              </tr>
            </tbody>
          </table>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const leaderboardData = ref([
  { id: 1, username: 'Alexandria99', totalSpend: 15450000 },
  { id: 2, username: 'BintangTimur', totalSpend: 12200000 },
  { id: 3, username: 'Zexth_Gaming', totalSpend: 9850000 },
  { id: 4, username: 'Wawan_Topup', totalSpend: 7400000 },
  { id: 5, username: 'Sultan_Ciruas', totalSpend: 5900000 },
  { id: 6, username: 'RianGamerz', totalSpend: 4200000 },
  { id: 7, username: 'Dinda_Cute', totalSpend: 3100000 }
])

const formatCurrency = (value) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(value)
}

useHead({
  title: 'Leaderboard - GoSaas Topup'
})
</script>
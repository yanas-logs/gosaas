import { ref, computed } from 'vue'
import { useLanguage } from './useLanguage.js'

export function useNavbarLogic() {
  const colorMode = useColorMode()
  const { locale } = useLanguage()
  
  const isDark = computed({
    get () { return colorMode.value === 'dark' },
    set () { colorMode.preference = colorMode.value === 'dark' ? 'light' : 'dark' }
  })

  const searchQuery = ref('')
  const handleSearch = () => {
    if (searchQuery.value.trim()) {
      alert(`Mencari: ${searchQuery.value}`)
    }
  }

  const currentLang = computed({
    get: () => locale.value,
    set: (value) => { locale.value = value }
  })

  const languages = [
    { label: 'ID', value: 'ID' },
    { label: 'EN', value: 'EN' }
  ]

  const handleLanguageChange = (value) => {
    currentLang.value = value
  }

  const links = [
    [{
      label: 'Home',
      icon: 'i-heroicons-home',
      to: '/'
    }, {
      label: 'Cek Transaksi',
      icon: 'i-heroicons-magnifying-glass-circle',
      to: '/cek-transaksi'
    }, {
      label: 'Leaderboard',
      icon: 'i-heroicons-trophy',
      to: '/leaderboard'
    }, {
      label: 'Blog',
      icon: 'i-heroicons-document-text',
      to: '/blog'
    }, {
      label: 'Kalkulator',
      icon: 'i-heroicons-calculator',
      to: '/kalkulator'
    }]
  ]

  const moreOptions = [
    [
      { label: 'Option A', slot: 'option-a', to: '#option-a' },
      { label: 'Option B', slot: 'option-b', to: '#option-b' },
      { label: 'Option C', slot: 'option-c', to: '#option-c' }
    ]
  ]

  return {
    isDark,
    searchQuery,
    handleSearch,
    currentLang,
    languages,
    handleLanguageChange,
    links,
    moreOptions
  }
}
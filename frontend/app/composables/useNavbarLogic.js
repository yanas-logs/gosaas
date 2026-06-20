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
      label: 'Fitur',
      icon: 'i-heroicons-sparkles',
      to: '#features'
    }, {
      label: 'Harga',
      icon: 'i-heroicons-credit-card',
      to: '#pricing'
    }]
  ]

  return {
    isDark,
    searchQuery,
    handleSearch,
    currentLang,
    languages,
    handleLanguageChange,
    links
  }
}
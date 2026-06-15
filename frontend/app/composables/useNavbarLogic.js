import { ref, computed } from 'vue'

export function useNavbarLogic() {
  const { locale } = useI18n()
  const colorMode = useColorMode()
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

  const currentLang = ref('ID')
  const languages = [
    { label: 'ID', value: 'ID' },
    { label: 'EN', value: 'EN' }
  ]
  const handleLanguageChange = (value) => {
    currentLang.value = value
    locale.value = value
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
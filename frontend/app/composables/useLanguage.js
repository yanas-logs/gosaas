import { useState } from '#app'

const translations = {
  ID: {
    welcome: 'Selamat datang di platform gateway top-up multi-tenant.',
    startNow: 'Mulai Sekarang',
    enterDashboard: 'Masuk Dashboard',
    navSignIn: 'Sign In',
    loginTitle: 'Masuk ke Dashboard',
    noAccount: 'Belum punya akun? Daftar',
    emailLabel: 'Email',
    passwordLabel: 'Password',
    searchPlaceholder: 'Cari game atau layanan...'
  },
  EN: {
    welcome: 'Welcome to the multi-tenant top-up gateway platform.',
    startNow: 'Start Now',
    enterDashboard: 'Enter Dashboard',
    navSignIn: 'Sign In',
    loginTitle: 'Sign In to Dashboard',
    noAccount: "Don't have an account? Register",
    emailLabel: 'Email',
    passwordLabel: 'Password',
    searchPlaceholder: 'Search games or services...'
  }
}

export function useLanguage() {
  const locale = useState('locale', () => 'ID')

  const t = (key) => {
    return translations[locale.value]?.[key] || key
  }

  return {
    locale,
    t
  }
}
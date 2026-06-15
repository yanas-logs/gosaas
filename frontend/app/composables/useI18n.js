export function useI18n() {
  const locale = useState('app-locale', () => 'ID')

  const messages = {
    ID: {
      navSignIn: 'Masuk',
      loginTitle: 'Masuk ke Dashboard',
      loginSubtitle: 'Atau',
      loginRegisterLink: 'Belum punya akun? Daftar',
      loginButton: 'Sign In',
      loginLoading: 'Memasuki...'
    },
    EN: {
      navSignIn: 'Sign In',
      loginTitle: 'Go to Dashboard',
      loginSubtitle: 'Or',
      loginRegisterLink: 'Don’t have an account? Register',
      loginButton: 'Sign In',
      loginLoading: 'Signing in...'
    }
  }

  const t = (key) => messages[locale.value]?.[key] || key

  return {
    locale,
    t
  }
}

import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref<Record<string, any> | null>(null)
  const isDark = ref(localStorage.getItem('isDark') === 'true')

  const isLoggedIn = computed(() => !!token.value)
  const username = computed(() => userInfo.value?.username || '')

  function setToken(newToken: string) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  function setUserInfo(info: Record<string, any>) {
    userInfo.value = info
  }

  function logout() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
  }

  function toggleDark() {
    isDark.value = !isDark.value
    localStorage.setItem('isDark', String(isDark.value))
    document.documentElement.classList.toggle('dark', isDark.value)
  }

  function initTheme() {
    document.documentElement.classList.toggle('dark', isDark.value)
  }

  return { token, userInfo, isDark, isLoggedIn, username, setToken, setUserInfo, logout, toggleDark, initTheme }
})

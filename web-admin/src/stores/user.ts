import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as loginApi } from '@/api'
import router from '@/router'

interface LoginResult {
  token: string
  user: {
    id: number
    username: string
    nickname: string
    avatar: string
    role: string
  }
}

export const useUserStore = defineStore('user', () => {
  const token = ref<string>(localStorage.getItem('admin_token') || '')
  const userInfo = ref<LoginResult['user'] | null>(
    JSON.parse(localStorage.getItem('admin_user') || 'null')
  )

  const isLoggedIn = computed(() => !!token.value)
  const username = computed(() => userInfo.value?.username || '')
  const nickname = computed(() => userInfo.value?.nickname || '')
  const avatar = computed(() => userInfo.value?.avatar || '')
  const role = computed(() => userInfo.value?.role || '')
  const isAdmin = computed(() => userInfo.value?.role === 'admin')

  async function login(usernameInput: string, password: string) {
    const res = await loginApi({ username: usernameInput, password })
    const data = res.data as unknown as LoginResult
    token.value = data.token
    userInfo.value = data.user
    localStorage.setItem('admin_token', data.token)
    localStorage.setItem('admin_user', JSON.stringify(data.user))
  }

  function logout() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('admin_token')
    localStorage.removeItem('admin_user')
    router.push('/login')
  }

  return {
    token,
    userInfo,
    isLoggedIn,
    username,
    nickname,
    avatar,
    role,
    isAdmin,
    login,
    logout,
  }
})
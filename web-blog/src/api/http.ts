import axios from 'axios'
import type { AxiosResponse, InternalAxiosRequestConfig } from 'axios'
import { useUserStore } from '@/stores/user'

const http = axios.create({
  baseURL: '/blog-api/v1',
  timeout: 15000,
  headers: { 'Content-Type': 'application/json' }
})

http.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers.Authorization = `Bearer ${userStore.token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

http.interceptors.response.use(
  (response: AxiosResponse) => {
    const { data } = response
    if (data.code !== 0) {
      return Promise.reject(new Error(data.message || data.msg || '请求失败'))
    }
    return data
  },
  (error) => {
    if (error.response?.status === 401) {
      const userStore = useUserStore()
      userStore.logout()
      window.location.hash = '#/login'
    }
    return Promise.reject(error.response?.data?.message || error.response?.data?.msg || error.message || '网络错误')
  }
)

export default http

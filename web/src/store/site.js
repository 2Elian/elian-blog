import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getSiteConfig } from '@/api'

export const useSiteStore = defineStore('site', () => {
  const config = ref({})
  const loaded = ref(false)

  async function loadConfig() {
    if (loaded.value) return
    try {
      const data = await getSiteConfig()
      const map = {}
      for (const item of data || []) {
        map[item.key] = item.value
      }
      config.value = map
      loaded.value = true
    } catch {
      // ignore
    }
  }

  return { config, loadConfig }
})
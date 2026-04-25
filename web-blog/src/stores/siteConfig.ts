import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getSiteConfig } from '@/api'

interface WebsiteInfo {
  website_name: string
  website_avatar: string
  website_intro: string
  website_notice: string | string[]
  website_record_no: string
}

interface SocialAccountInfo {
  name: string
  platform: string
  link_url: string
  enabled: boolean
}

interface WebsiteFeature {
  is_chat_room: number
  is_ai_assistant: number
  is_music_player: number
  is_comment_review: number
  is_email_notice: number
  is_message_review: number
  is_reward: number
}

export const useSiteConfigStore = defineStore('siteConfig', () => {
  const config = ref<Record<string, any> | null>(null)
  const loaded = ref(false)

  const websiteInfo = computed<WebsiteInfo>(() => {
    return config.value?.website_info || {} as WebsiteInfo
  })

  const siteName = computed(() => websiteInfo.value.website_name || 'Elian Blog')
  const siteIntro = computed(() => websiteInfo.value.website_intro || '记录生活，分享技术')
  const siteNotice = computed<string[]>(() => {
    const raw = websiteInfo.value.website_notice
    if (!raw) return []
    if (Array.isArray(raw)) return raw.filter((s: string) => s.trim())
    if (typeof raw === 'string' && raw.trim()) return [raw]
    return []
  })
  const siteRecordNo = computed(() => websiteInfo.value.website_record_no || '')
  const siteAvatar = computed(() => websiteInfo.value.website_avatar || '')

  const heroName = computed(() => (websiteInfo.value as any).website_hero_name || '')
  const heroDesc = computed(() => (websiteInfo.value as any).website_hero_desc || '')
  const heroSubtitleTexts = computed<string[]>(() => {
    const raw = (websiteInfo.value as any).website_hero_subtitles
    if (!raw) return []
    if (Array.isArray(raw)) return raw.filter((s: string) => s.trim())
    return []
  })

  const socialUrlList = computed<SocialAccountInfo[]>(() => {
    return config.value?.social_url_list || []
  })

  const websiteFeature = computed<WebsiteFeature>(() => {
    return config.value?.website_feature || {} as WebsiteFeature
  })

  async function fetchConfig() {
    if (loaded.value) return
    try {
      const res = await getSiteConfig() as any
      config.value = res.data
      loaded.value = true
    } catch {
      config.value = {}
      loaded.value = true
    }
  }

  return {
    config, loaded,
    websiteInfo, siteName, siteIntro, siteNotice,
    siteRecordNo, siteAvatar,
    heroName, heroDesc, heroSubtitleTexts,
    socialUrlList, websiteFeature,
    fetchConfig
  }
})

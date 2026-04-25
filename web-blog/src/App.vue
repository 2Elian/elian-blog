<template>
  <n-config-provider :theme="isDark ? darkTheme : undefined" :theme-overrides="themeOverrides">
    <n-message-provider>
      <n-dialog-provider>
        <n-notification-provider>
          <router-view />
        </n-notification-provider>
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import {
  NConfigProvider,
  NMessageProvider,
  NDialogProvider,
  NNotificationProvider,
  darkTheme,
  type GlobalThemeOverrides
} from 'naive-ui'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const isDark = computed(() => userStore.isDark)

const themeOverrides = computed<GlobalThemeOverrides>(() => ({
  common: {
    primaryColor: '#1a1a1a',
    primaryColorHover: '#333333',
    primaryColorPressed: '#000000',
    primaryColorSuppl: '#1a1a1a',
    borderRadius: '8px',
    borderRadiusSmall: '6px',
    fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "PingFang SC", "Microsoft YaHei", sans-serif'
  },
  Button: {
    borderRadiusMedium: '20px',
    borderRadiusLarge: '24px',
    borderRadiusSmall: '16px'
  },
  Tag: {
    borderRadius: '14px'
  },
  Card: {
    borderRadius: '12px',
    boxShadow: '0 2px 12px rgba(0, 0, 0, 0.06)'
  }
}))

onMounted(() => {
  userStore.initTheme()
})
</script>

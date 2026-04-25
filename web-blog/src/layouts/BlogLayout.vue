<template>
  <div class="blog-layout" :class="{ dark: isDark }">
    <AppHeader />
    <main class="main-content">
      <router-view v-slot="{ Component }">
        <transition name="page" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
    <AppFooter />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import AppHeader from '@/components/AppHeader.vue'
import AppFooter from '@/components/AppFooter.vue'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const isDark = computed(() => userStore.isDark)
</script>

<style scoped lang="scss">
.blog-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
  transition: background-color var(--transition-normal);
}

.main-content {
  flex: 1;
  padding-top: calc(var(--header-height) + 20px);
  padding-bottom: 40px;
}
</style>

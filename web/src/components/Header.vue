<template>
  <header class="header">
    <div class="header-inner container">
      <router-link to="/" class="logo">
        <img src="/favicon.svg" alt="logo" class="logo-icon" />
        <span class="site-name">{{ siteName }}</span>
      </router-link>

      <nav class="nav">
        <router-link to="/" :class="{ active: $route.path === '/' }">首页</router-link>
        <router-link to="/blog" :class="{ active: $route.path.startsWith('/blog') || $route.path.startsWith('/article') }">博客</router-link>
        <router-link to="/learn" :class="{ active: $route.path === '/learn' }">学习</router-link>
        <router-link to="/friend-link" :class="{ active: $route.path === '/friend-link' }">友链</router-link>
        <router-link to="/messages" :class="{ active: $route.path === '/messages' }">留言板</router-link>
      </nav>

      <div class="header-right">
        <template v-if="userStore.isLoggedIn">
          <el-dropdown trigger="click">
            <span class="user-avatar">
              <el-avatar :size="32" :src="userStore.userInfo?.avatar">
                {{ userStore.userInfo?.username?.charAt(0) || 'U' }}
              </el-avatar>
              <span class="username">{{ userStore.userInfo?.username }}</span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="userStore.logout()">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
        <template v-else>
          <router-link to="/login">
            <el-button type="primary" size="small" round>登录</el-button>
          </router-link>
        </template>
      </div>
    </div>
  </header>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useUserStore } from '@/store/user'
import { useSiteStore } from '@/store/site'

const userStore = useUserStore()
const siteStore = useSiteStore()

const siteName = computed(() => siteStore.config['site_name'] || 'Elian Blog')

onMounted(() => {
  if (userStore.isLoggedIn && !userStore.userInfo) {
    userStore.fetchUserInfo()
  }
})
</script>

<style lang="scss" scoped>
.header {
  background: var(--bg-card);
  border-bottom: 1px solid var(--border-color);
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-inner {
  display: flex;
  align-items: center;
  height: 60px;
  gap: 40px;
}

.logo {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  flex-shrink: 0;

  &:hover { color: var(--text-primary); }
}

.logo-icon {
  width: 32px;
  height: 32px;
}

.nav {
  display: flex;
  gap: 8px;
  flex: 1;

  a {
    padding: 6px 16px;
    border-radius: 20px;
    font-size: 14px;
    color: var(--text-secondary);
    transition: all 0.2s;

    &:hover {
      color: var(--primary);
      background: rgba(64, 158, 255, 0.06);
    }

    &.active {
      color: var(--primary);
      background: rgba(64, 158, 255, 0.1);
      font-weight: 500;
    }
  }
}

.header-right {
  flex-shrink: 0;
}

.user-avatar {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.username {
  font-size: 14px;
  color: var(--text-secondary);
}
</style>
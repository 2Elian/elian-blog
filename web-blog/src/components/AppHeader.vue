<template>
  <header class="app-header" :class="{ dark: isDark }">
    <div class="header-container">
      <router-link to="/" class="logo">
        <span class="logo-text">Elian Blog</span>
      </router-link>

      <nav class="nav-links" :class="{ active: menuActive }">
        <router-link
          v-for="link in navLinks"
          :key="link.path"
          :to="link.path"
          class="nav-link"
          @click="menuActive = false"
        >
          <n-icon size="18"><component :is="link.icon" /></n-icon>
          <span>{{ link.text }}</span>
        </router-link>
      </nav>

      <div class="header-actions">
        <n-button quaternary circle class="action-btn" @click="showSearch = true">
          <template #icon>
            <n-icon><SearchOutline /></n-icon>
          </template>
        </n-button>

        <n-button quaternary circle class="action-btn" @click="toggleDark">
          <template #icon>
            <n-icon><component :is="isDark ? SunnyOutline : MoonOutline" /></n-icon>
          </template>
        </n-button>

        <n-button
          v-if="!isLoggedIn"
          class="login-btn"
          @click="$router.push('/login')"
        >
          登录
        </n-button>

        <n-dropdown v-else :options="userOptions" @select="handleUserAction">
          <n-avatar round :size="32" class="user-avatar">
            {{ username?.charAt(0)?.toUpperCase() || 'U' }}
          </n-avatar>
        </n-dropdown>

        <button class="menu-toggle" @click="menuActive = !menuActive">
          <n-icon size="24"><MenuOutline /></n-icon>
        </button>
      </div>
    </div>

    <!-- Search Modal -->
    <n-modal v-model:show="showSearch" preset="card" style="width: 500px; max-width: 90vw;">
      <div class="search-modal">
        <n-input
          v-model:value="searchKeyword"
          placeholder="搜索文章..."
          size="large"
          clearable
          @keyup.enter="handleSearch"
        >
          <template #prefix>
            <n-icon><SearchOutline /></n-icon>
          </template>
        </n-input>
        <n-button type="primary" style="margin-top: 16px;" block @click="handleSearch">
          搜索
        </n-button>
      </div>
    </n-modal>
  </header>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import {
  NButton,
  NIcon,
  NAvatar,
  NDropdown,
  NModal,
  NInput
} from 'naive-ui'
import {
  SearchOutline,
  MoonOutline,
  SunnyOutline,
  HomeOutline,
  BookOutline,
  TimeOutline,
  PricetagOutline,
  CubeOutline,
  MenuOutline
} from '@vicons/ionicons5'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const isDark = computed(() => userStore.isDark)
const isLoggedIn = computed(() => userStore.isLoggedIn)
const username = computed(() => userStore.username)

const menuActive = ref(false)
const showSearch = ref(false)
const searchKeyword = ref('')

const navLinks = [
  { path: '/', text: '首页', icon: HomeOutline },
  { path: '/blog', text: '博客', icon: BookOutline },
  { path: '/archive', text: '归档', icon: TimeOutline },
  { path: '/tags', text: '标签', icon: PricetagOutline },
  { path: '/products', text: '产品', icon: CubeOutline }
]

const userOptions = [
  { label: '个人中心', key: 'profile' },
  { label: '退出登录', key: 'logout' }
]

function toggleDark() {
  userStore.toggleDark()
}

function handleSearch() {
  if (searchKeyword.value.trim()) {
    router.push({ path: '/blog', query: { search: searchKeyword.value.trim() } })
    showSearch.value = false
    searchKeyword.value = ''
  }
}

function handleUserAction(key: string) {
  if (key === 'logout') {
    userStore.logout()
    router.push('/')
  }
}
</script>

<style scoped lang="scss">
.app-header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  background: var(--bg-card);
  border-bottom: 1px solid var(--border-color);
  transition: background-color var(--transition-normal), border-color var(--transition-normal);

  &::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    height: 2px;
    background: var(--accent-gradient);
  }
}

.header-container {
  max-width: 1400px;
  margin: 0 auto;
  height: var(--header-height);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 32px;
}

.logo {
  text-decoration: none;
  flex-shrink: 0;
}

.logo-text {
  font-size: 22px;
  font-weight: 700;
  letter-spacing: -0.5px;
  color: var(--text-primary);
}

.nav-links {
  display: flex;
  gap: 4px;
  flex-wrap: nowrap;

  @media (max-width: 768px) {
    position: fixed;
    top: var(--header-height);
    left: 0;
    right: 0;
    background: var(--bg-card);
    flex-direction: column;
    padding: 16px 20px;
    gap: 4px;
    box-shadow: var(--shadow-md);
    transform: translateY(-100%);
    opacity: 0;
    pointer-events: none;
    transition: transform var(--transition-normal), opacity var(--transition-normal);

    &.active {
      transform: translateY(0);
      opacity: 1;
      pointer-events: auto;
    }
  }
}

.nav-link {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 14px;
  color: var(--text-secondary);
  font-size: 14px;
  border-radius: 8px;
  transition: all var(--transition-fast);
  white-space: nowrap;
  flex-shrink: 0;

  &:hover,
  &.router-link-active {
    color: var(--text-primary);
    background: rgba(0, 0, 0, 0.06);
  }
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.action-btn {
  color: var(--text-secondary);

  &:hover {
    color: var(--text-primary);
    background: rgba(0, 0, 0, 0.06);
  }
}

.login-btn {
  margin-left: 8px;
}

.user-avatar {
  cursor: pointer;
  background: var(--accent-gradient);
  color: white;
  font-weight: 600;
  transition: transform var(--transition-fast);

  &:hover {
    transform: scale(1.05);
  }
}

.menu-toggle {
  display: none;
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-primary);
  padding: 4px;

  @media (max-width: 768px) {
    display: flex;
    align-items: center;
    justify-content: center;
  }
}

.search-modal {
  padding: 8px;
}

@media (max-width: 768px) {
  .nav-links {
    display: flex;
  }

  .login-btn {
    display: none;
  }
}
</style>

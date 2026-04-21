<template>
  <div class="admin-layout">
    <!-- 侧边栏 -->
    <div class="sidebar-container" :class="{ 'is-collapsed': isCollapsed }">
      <div class="sidebar-logo">
        <el-icon class="logo-img" :size="32" color="#fff"><Promotion /></el-icon>
        <span v-show="!isCollapsed" class="logo-title">Elian Blog</span>
      </div>
      <el-menu
        class="sidebar-menu"
        :default-active="activeMenu"
        :collapse="isCollapsed"
        :collapse-transition="false"
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409eff"
        router
      >
        <template v-for="route in menuRoutes" :key="route.path">
          <el-menu-item :index="'/' + route.path">
            <el-icon v-if="route.meta?.icon">
              <component :is="route.meta.icon" />
            </el-icon>
            <template #title>{{ route.meta?.title }}</template>
          </el-menu-item>
        </template>
      </el-menu>
    </div>

    <!-- 主内容区 -->
    <div class="main-container" :class="{ 'is-collapsed': isCollapsed }">
      <!-- 导航栏 -->
      <div class="navbar">
        <div class="navbar-left">
          <div
            class="hamburger"
            :class="{ 'is-active': isCollapsed }"
            @click="toggleSidebar"
          >
            <el-icon :size="20"><Fold /></el-icon>
          </div>
          <el-breadcrumb separator="/" class="breadcrumb">
            <el-breadcrumb-item :to="{ path: '/dashboard' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item v-if="currentRoute.meta?.title && currentRoute.name !== 'Dashboard'">
              {{ currentRoute.meta.title }}
            </el-breadcrumb-item>
          </el-breadcrumb>
        </div>

        <div class="navbar-right">
          <!-- 全屏 -->
          <el-tooltip content="全屏" placement="bottom">
            <div class="navbar-action" @click="toggleFullscreen">
              <el-icon><FullScreen /></el-icon>
            </div>
          </el-tooltip>

          <!-- 用户下拉 -->
          <el-dropdown trigger="click" @command="handleUserCommand">
            <div class="user-dropdown">
              <div class="user-avatar">
                {{ userStore.nickname?.charAt(0)?.toUpperCase() || 'A' }}
              </div>
              <span class="user-name">{{ userStore.nickname || userStore.username || 'Admin' }}</span>
              <el-icon style="margin-left: 4px"><ArrowDown /></el-icon>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item disabled>
                  <el-icon><User /></el-icon>
                  {{ userStore.role === 'admin' ? '管理员' : '编辑者' }}
                </el-dropdown-item>
                <el-dropdown-item divided command="logout">
                  <el-icon><SwitchButton /></el-icon>
                  退出登录
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>

      <!-- 页面主体 -->
      <div class="app-main">
        <router-view v-slot="{ Component }">
          <keep-alive>
            <component :is="Component" />
          </keep-alive>
        </router-view>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessageBox } from 'element-plus'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const isCollapsed = ref(false)

// 获取当前路由的子路由作为菜单项（过滤有权限的）
const menuRoutes = computed(() => {
  const mainRoute = router.options.routes.find((r) => r.path === '/')
  if (!mainRoute?.children) return []
  const userRole = userStore.role
  return mainRoute.children.filter((child) => {
    if (!child.meta?.roles) return true
    return (child.meta.roles as string[]).includes(userRole)
  })
})

const activeMenu = computed(() => route.path)
const currentRoute = computed(() => route)

function toggleSidebar() {
  isCollapsed.value = !isCollapsed.value
}

function toggleFullscreen() {
  if (!document.fullscreenElement) {
    document.documentElement.requestFullscreen()
  } else {
    document.exitFullscreen()
  }
}

function handleUserCommand(command: string) {
  if (command === 'logout') {
    ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }).then(() => {
      userStore.logout()
    }).catch(() => {})
  }
}
</script>
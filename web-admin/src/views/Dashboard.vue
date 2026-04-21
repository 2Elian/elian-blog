<template>
  <div class="dashboard">
    <div class="page-header">
      <h2>仪表盘</h2>
    </div>

    <!-- 统计卡片 -->
    <el-row :gutter="16" class="stat-cards">
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-content">
            <div class="stat-info">
              <div class="stat-label">文章总数</div>
              <div class="stat-number">{{ stats.article_count }}</div>
            </div>
            <div class="stat-icon blue">
              <el-icon :size="28"><Document /></el-icon>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-content">
            <div class="stat-info">
              <div class="stat-label">用户总数</div>
              <div class="stat-number">{{ stats.user_count }}</div>
            </div>
            <div class="stat-icon green">
              <el-icon :size="28"><User /></el-icon>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-content">
            <div class="stat-info">
              <div class="stat-label">评论总数</div>
              <div class="stat-number">{{ stats.comment_count }}</div>
            </div>
            <div class="stat-icon orange">
              <el-icon :size="28"><ChatDotRound /></el-icon>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="6">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-content">
            <div class="stat-info">
              <div class="stat-label">访问总量</div>
              <div class="stat-number">{{ formatNumber(stats.view_count) }}</div>
            </div>
            <div class="stat-icon red">
              <el-icon :size="28"><View /></el-icon>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 欢迎卡片 -->
    <el-card class="welcome-card" style="margin-top: 16px;">
      <div style="text-align: center; padding: 20px 0;">
        <el-icon :size="64" color="#409eff"><Promotion /></el-icon>
        <h2 style="margin: 16px 0 8px; color: #303133;">
          欢迎回来，{{ userStore.nickname || userStore.username || 'Admin' }}！
        </h2>
        <p style="color: #909399; font-size: 14px;">
          欢迎使用 Elian Blog 管理后台，祝您工作愉快！
        </p>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted } from 'vue'
import { Document, User, ChatDotRound, View, Promotion } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'
import { getDashboardStats } from '@/api'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()

const stats = reactive({
  article_count: 0,
  user_count: 0,
  comment_count: 0,
  view_count: 0,
})

function formatNumber(num: number): string {
  if (num >= 10000) {
    return (num / 10000).toFixed(1) + 'w'
  }
  return num.toLocaleString()
}

async function fetchStats() {
  try {
    const res = await getDashboardStats()
    stats.article_count = res.data.article_count || 0
    stats.user_count = res.data.user_count || 0
    stats.comment_count = res.data.comment_count || 0
    stats.view_count = res.data.view_count || 0
  } catch {
    ElMessage.warning('获取统计数据失败')
  }
}

onMounted(() => {
  fetchStats()
})
</script>

<style scoped>
.stat-cards {
  margin-bottom: 16px;
}

.stat-cards .el-col {
  margin-bottom: 16px;
}
</style>
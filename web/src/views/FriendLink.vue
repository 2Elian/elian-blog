<template>
  <div class="friend-link-page container">
    <div class="page-header">
      <h1>友情链接</h1>
      <p>优秀的博主和技术网站推荐</p>
    </div>

    <LoadingState :loading="loading" :error="error" emptyText="暂无友链" @retry="loadFriends">
      <div class="friend-grid">
        <a class="friend-card" v-for="friend in friends" :key="friend.id" :href="friend.url" target="_blank" rel="noopener">
          <el-avatar :size="48" :src="friend.logo">{{ friend.name.charAt(0) }}</el-avatar>
          <div class="friend-info">
            <h3>{{ friend.name }}</h3>
            <p>{{ friend.description }}</p>
          </div>
          <el-icon class="link-icon"><Link /></el-icon>
        </a>
      </div>
    </LoadingState>

    <div class="apply-section">
      <h2>申请友链</h2>
      <p>欢迎优质博主申请友链互换，请通过留言板提交您的网站信息。</p>
      <router-link to="/messages">
        <el-button type="primary">前往留言板</el-button>
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import LoadingState from '@/components/LoadingState.vue'
import { getFriendLinks } from '@/api'

const loading = ref(true)
const error = ref('')
const friends = ref([])

onMounted(() => loadFriends())

async function loadFriends() {
  loading.value = true
  error.value = ''
  try {
    friends.value = await getFriendLinks() || []
  } catch {
    error.value = '加载失败'
  } finally {
    loading.value = false
  }
}
</script>

<style lang="scss" scoped>
.friend-link-page { padding-bottom: 40px; }

.page-header {
  text-align: center;
  padding: 40px 0;
  h1 { font-size: 32px; margin-bottom: 12px; }
  p { font-size: 16px; color: var(--text-muted); }
}

.friend-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
  margin-bottom: 40px;

  @media (max-width: 600px) { grid-template-columns: 1fr; }
}

.friend-card {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  box-shadow: var(--shadow-sm);
  transition: all 0.3s;

  &:hover { box-shadow: var(--shadow-md); transform: translateY(-2px); color: var(--text-primary); }
}

.friend-info {
  flex: 1;
  h3 { font-size: 16px; margin-bottom: 6px; }
  p { font-size: 13px; color: var(--text-muted); line-height: 1.5; }
}

.link-icon {
  color: var(--text-muted);
  font-size: 20px;
}

.apply-section {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 32px;
  text-align: center;
  box-shadow: var(--shadow-sm);

  h2 { font-size: 20px; margin-bottom: 12px; }
  p { font-size: 14px; color: var(--text-muted); margin-bottom: 20px; }
}
</style>
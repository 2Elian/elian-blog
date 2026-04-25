<template>
  <div class="friends-page">
    <div class="page-header">
      <h1 class="page-title gradient-text">友情链接</h1>
      <p class="page-desc">那些优秀的博主们</p>
    </div>

    <div class="friends-grid">
      <a
        v-for="friend in friends"
        :key="friend.id"
        :href="friend.url"
        target="_blank"
        rel="noopener noreferrer"
        class="friend-card"
      >
        <div class="card-cover">
          <img v-if="friend.avatar" :src="friend.avatar" :alt="friend.name" />
          <div v-else class="avatar-placeholder">
            {{ friend.name?.charAt(0) }}
          </div>
        </div>
        <div class="card-content">
          <h3 class="friend-name">{{ friend.name }}</h3>
          <p class="friend-desc" v-if="friend.description">{{ friend.description }}</p>
          <div class="friend-link">
            <n-icon><OpenOutline /></n-icon>
            <span>访问</span>
          </div>
        </div>
      </a>
    </div>

    <n-empty v-if="!loading && friends.length === 0" description="暂无友链" style="padding: 60px 0;" />

    <!-- Apply Section -->
    <div class="apply-section">
      <n-card>
        <h3>申请友链</h3>
        <p>欢迎交换友链！请在留言板留言，格式如下：</p>
        <ul>
          <li>网站名称：您的网站名称</li>
          <li>网站地址：您的网站URL</li>
          <li>网站描述：简单介绍（50字以内）</li>
          <li>头像链接：您的头像URL</li>
        </ul>
        <n-button type="primary" round @click="$router.push('/messages')">
          前往留言
        </n-button>
      </n-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { NIcon, NEmpty, NCard, NButton } from 'naive-ui'
import { OpenOutline } from '@vicons/ionicons5'
import { getFriendLinks } from '@/api'

interface Friend {
  id: number
  name: string
  url: string
  avatar?: string
  description?: string
}

const friends = ref<Friend[]>([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await getFriendLinks() as any
    friends.value = res.data || []
  } catch (e) {
    console.error('Failed to load friends:', e)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped lang="scss">
.friends-page {
  animation: fadeInUp 0.5s ease;
  max-width: 900px;
  margin: 0 auto;
  padding: 0 32px;

  @media (max-width: 640px) {
    padding: 0 20px;
  }
}

.page-header {
  text-align: center;
  margin-bottom: 40px;
}

.page-title {
  font-size: 32px;
  font-weight: 700;
  margin-bottom: 10px;
}

.page-desc {
  color: var(--text-muted);
  font-size: 16px;
}

.friends-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 24px;
  margin-bottom: 40px;
}

.friend-card {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  overflow: hidden;
  text-decoration: none;
  box-shadow: var(--shadow-sm);
  transition: all var(--transition-normal);
  display: flex;
  flex-direction: column;

  &:hover {
    transform: translateY(-6px);
    box-shadow: var(--shadow-lg);

    .friend-name {
      color: var(--primary-color);
    }

    .card-cover img,
    .avatar-placeholder {
      transform: scale(1.05);
    }
  }
}

.card-cover {
  height: 100px;
  background: linear-gradient(135deg, #1a1a1a, #3a3a3a);
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform var(--transition-slow);
  }

  .avatar-placeholder {
    font-size: 36px;
    font-weight: 700;
    color: white;
    transition: transform var(--transition-slow);
  }
}

.card-content {
  padding: 20px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.friend-name {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
  transition: color var(--transition-fast);
}

.friend-desc {
  font-size: 14px;
  color: var(--text-muted);
  line-height: 1.5;
  margin-bottom: 16px;
  flex: 1;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.friend-link {
  display: flex;
  align-items: center;
  gap: 6px;
  color: var(--text-secondary);
  font-size: 14px;
  font-weight: 500;
}

.apply-section {
  margin-top: 40px;

  :deep(.n-card) {
    background: var(--bg-card);
    border-radius: var(--radius-md);
    box-shadow: var(--shadow-sm);
  }

  :deep(.n-card__content) {
    text-align: center;
  }

  h3 {
    font-size: 20px;
    font-weight: 600;
    color: var(--text-primary);
    margin-bottom: 16px;
  }

  p {
    color: var(--text-secondary);
    margin-bottom: 12px;
  }

  ul {
    list-style: none;
    color: var(--text-muted);
    font-size: 14px;
    margin-bottom: 20px;
    text-align: left;
    display: inline-block;

    li {
      margin-bottom: 4px;
    }
  }
}
</style>
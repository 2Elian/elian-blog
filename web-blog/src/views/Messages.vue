<template>
  <div class="messages-page">
    <div class="page-header">
      <h1 class="page-title gradient-text">留言板</h1>
      <p class="page-desc">欢迎留下你的足迹</p>
    </div>

    <!-- Message Form -->
    <div class="message-form-card">
      <n-input
        v-model:value="messageContent"
        type="textarea"
        placeholder="写下你想说的话..."
        :rows="4"
        maxlength="500"
        show-count
      />
      <div class="form-actions">
        <n-button
          type="primary"
          round
          :disabled="!messageContent.trim()"
          :loading="submitting"
          @click="submitMessage"
        >
          发表留言
        </n-button>
      </div>
    </div>

    <!-- Messages List -->
    <div class="messages-list">
      <div
        v-for="msg in messages"
        :key="msg.id"
        class="message-item"
      >
        <div class="message-avatar">
          <n-avatar round :size="48">
            {{ (msg.username || 'U').charAt(0).toUpperCase() }}
          </n-avatar>
        </div>
        <div class="message-content">
          <div class="message-header">
            <span class="message-author">{{ msg.username || '匿名用户' }}</span>
            <span class="message-time">{{ formatTime(msg.created_at) }}</span>
          </div>
          <p class="message-text">{{ msg.content }}</p>
        </div>
      </div>
    </div>

    <n-empty v-if="!loading && messages.length === 0" description="暂无留言，快来抢沙发吧~" style="padding: 60px 0;" />

    <!-- Pagination -->
    <div class="pagination" v-if="total > pageSize">
      <n-pagination
        v-model:page="currentPage"
        :page-count="Math.ceil(total / pageSize)"
        @update:page="loadMessages"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { NInput, NButton, NAvatar, NEmpty, NPagination, useMessage } from 'naive-ui'
import { getMessages, postMessage } from '@/api'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'

interface Message {
  id: number
  content: string
  username?: string
  created_at: string
}

const router = useRouter()
const message = useMessage()
const userStore = useUserStore()

const messages = ref<Message[]>([])
const messageContent = ref('')
const submitting = ref(false)
const loading = ref(true)
const currentPage = ref(1)
const total = ref(0)
const pageSize = 10

function formatTime(date: string) {
  const d = new Date(date)
  const now = new Date()
  const diff = now.getTime() - d.getTime()

  const minutes = Math.floor(diff / 60000)
  const hours = Math.floor(diff / 3600000)
  const days = Math.floor(diff / 86400000)

  if (minutes < 1) return '刚刚'
  if (minutes < 60) return `${minutes}分钟前`
  if (hours < 24) return `${hours}小时前`
  if (days < 30) return `${days}天前`

  return d.toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' })
}

async function loadMessages() {
  loading.value = true
  try {
    const res = await getMessages({ page: currentPage.value, page_size: pageSize }) as any
    messages.value = res.data?.list || []
    total.value = res.data?.total || 0
  } catch (e) {
    console.error('Failed to load messages:', e)
  } finally {
    loading.value = false
  }
}

async function submitMessage() {
  if (!userStore.isLoggedIn) {
    message.warning('请先登录')
    router.push('/login')
    return
  }

  if (!messageContent.value.trim()) return

  submitting.value = true
  try {
    await postMessage({ content: messageContent.value.trim() })
    message.success('留言成功')
    messageContent.value = ''
    currentPage.value = 1
    loadMessages()
  } catch (e: any) {
    message.error(e.message || '留言失败')
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  loadMessages()
})
</script>

<style scoped lang="scss">
.messages-page {
  animation: fadeInUp 0.5s ease;
  max-width: 800px;
  margin: 0 auto;
}

.page-header {
  text-align: center;
  margin-bottom: 30px;
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

.message-form-card {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 24px;
  margin-bottom: 30px;
  box-shadow: var(--shadow-sm);
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 16px;
}

.messages-list {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-sm);
  overflow: hidden;
}

.message-item {
  display: flex;
  gap: 16px;
  padding: 24px;
  border-bottom: 1px solid var(--border-color);
  transition: background-color var(--transition-fast);

  &:last-child {
    border-bottom: none;
  }

  &:hover {
    background: rgba(233, 84, 107, 0.02);
  }
}

.message-avatar {
  flex-shrink: 0;

  :deep(.n-avatar) {
    background: var(--accent-gradient);
    color: white;
    font-weight: 600;
    font-size: 18px;
  }
}

.message-content {
  flex: 1;
  min-width: 0;
}

.message-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 10px;
}

.message-author {
  font-weight: 600;
  color: var(--text-primary);
  font-size: 15px;
}

.message-time {
  font-size: 13px;
  color: var(--text-muted);
}

.message-text {
  color: var(--text-secondary);
  font-size: 15px;
  line-height: 1.7;
  margin: 0;
  white-space: pre-wrap;
  word-break: break-word;
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 30px;
}
</style>
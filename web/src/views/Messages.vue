<template>
  <div class="messages-page container">
    <div class="page-header">
      <h1>留言板</h1>
      <p>欢迎留言交流，分享想法</p>
    </div>

    <!-- 发表留言 -->
    <div class="message-form" v-if="userStore.isLoggedIn">
      <el-input v-model="messageContent" type="textarea" :rows="4" placeholder="写下你想说的话..." />
      <el-button type="primary" :loading="submitting" @click="submitMessage" style="margin-top: 12px;">发表留言</el-button>
    </div>
    <div class="login-tip" v-else>
      <router-link to="/login">登录</router-link> 后参与留言
    </div>

    <!-- 留言列表 -->
    <LoadingState :loading="loading" :error="error" emptyText="暂无留言，快来发表第一条吧" @retry="loadMessages">
      <div class="message-list">
        <div class="message-item" v-for="msg in messages" :key="msg.id">
          <el-avatar :size="40" :src="msg.user?.avatar">{{ msg.user?.username?.charAt(0) }}</el-avatar>
          <div class="message-body">
            <div class="message-header">
              <span class="username">{{ msg.user?.username || '匿名用户' }}</span>
              <span class="date">{{ formatDate(msg.created_at) }}</span>
            </div>
            <p class="message-text">{{ msg.content }}</p>
            <div class="message-images" v-if="msg.images?.length">
              <el-image v-for="(img, i) in msg.images" :key="i" :src="img" fit="cover" class="msg-image" :preview-src-list="msg.images" />
            </div>
          </div>
        </div>
      </div>
    </LoadingState>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import LoadingState from '@/components/LoadingState.vue'
import { getMessages, createMessage } from '@/api'
import { useUserStore } from '@/store/user'

const userStore = useUserStore()

const loading = ref(true)
const error = ref('')
const messages = ref([])
const messageContent = ref('')
const submitting = ref(false)
const page = ref(1)

onMounted(() => loadMessages())

async function loadMessages() {
  loading.value = true
  error.value = ''
  try {
    const data = await getMessages({ page: page.value, page_size: 20 })
    messages.value = data?.list || []
  } catch {
    error.value = '加载失败'
  } finally {
    loading.value = false
  }
}

async function submitMessage() {
  if (!messageContent.value.trim()) return
  submitting.value = true
  try {
    await createMessage({ content: messageContent.value, type: 'guestbook' })
    messageContent.value = ''
    await loadMessages()
  } catch {}
  finally { submitting.value = false }
}

function formatDate(date) {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-CN')
}
</script>

<style lang="scss" scoped>
.messages-page { max-width: 800px; padding-bottom: 40px; }

.page-header {
  text-align: center;
  padding: 40px 0;
  h1 { font-size: 32px; margin-bottom: 12px; }
  p { font-size: 16px; color: var(--text-muted); }
}

.message-form {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: var(--shadow-sm);
}

.login-tip {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 24px;
  text-align: center;
  margin-bottom: 24px;
  a { color: var(--primary); }
}

.message-list {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.message-item {
  display: flex;
  gap: 12px;
}

.message-body { flex: 1; }

.message-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  .username { font-weight: 500; color: var(--text-primary); }
  .date { font-size: 12px; color: var(--text-muted); }
}

.message-text {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.6;
}

.message-images {
  display: flex;
  gap: 8px;
  margin-top: 12px;
}

.msg-image {
  width: 120px;
  height: 120px;
  border-radius: var(--radius);
  cursor: pointer;
}
</style>
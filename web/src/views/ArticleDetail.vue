<template>
  <div class="article-detail container">
    <LoadingState :loading="loading" :error="error" emptyText="文章不存在" @retry="loadArticle">
      <article class="article-content">
        <!-- 头部 -->
        <header class="article-header">
          <h1 class="title">{{ article.title }}</h1>
          <div class="meta">
            <span><el-icon><Calendar /></el-icon> {{ formatDate(article.created_at) }}</span>
            <span><el-icon><View /></el-icon> {{ article.view_count || 0 }} 阅读</span>
            <span><el-icon><ChatDotRound /></el-icon> {{ article.comment_count || 0 }} 评论</span>
            <span v-if="article.category"><el-icon><Folder /></el-icon> {{ article.category.name }}</span>
          </div>
          <div class="tags" v-if="article.tags?.length">
            <el-tag v-for="tag in article.tags" :key="tag.id" :color="tag.color" effect="dark" round size="small">{{ tag.name }}</el-tag>
          </div>
        </header>

        <!-- 封面 -->
        <div class="cover" v-if="article.cover">
          <img :src="article.cover" alt="封面" />
        </div>

        <!-- 正文 -->
        <div class="content-body" v-html="article.content"></div>
      </article>

      <!-- 评论区 -->
      <section class="comment-section">
        <h3>评论 <span class="count">({{ comments.length }})</span></h3>

        <!-- 发表评论 -->
        <div class="comment-form" v-if="userStore.isLoggedIn">
          <el-input v-model="commentContent" type="textarea" :rows="3" placeholder="写下你的评论..." />
          <el-button type="primary" :loading="submitting" @click="submitComment" style="margin-top: 12px;">发表评论</el-button>
        </div>
        <div class="login-tip" v-else>
          <router-link to="/login">登录</router-link> 后参与讨论
        </div>

        <!-- 评论列表 -->
        <div class="comment-list">
          <div class="comment-item" v-for="comment in comments" :key="comment.id">
            <el-avatar :size="40" :src="comment.user?.avatar">{{ comment.user?.username?.charAt(0) }}</el-avatar>
            <div class="comment-body">
              <div class="comment-header">
                <span class="username">{{ comment.user?.username || '匿名用户' }}</span>
                <span class="date">{{ formatDate(comment.created_at) }}</span>
              </div>
              <p class="comment-text">{{ comment.content }}</p>
            </div>
          </div>
        </div>
      </section>
    </LoadingState>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import LoadingState from '@/components/LoadingState.vue'
import { getArticle, getComments, createComment } from '@/api'
import { useUserStore } from '@/store/user'

const route = useRoute()
const userStore = useUserStore()

const loading = ref(true)
const error = ref('')
const article = ref({})
const comments = ref([])
const commentContent = ref('')
const submitting = ref(false)

onMounted(() => loadArticle())
watch(() => route.params.id, () => loadArticle())

async function loadArticle() {
  loading.value = true
  error.value = ''
  try {
    article.value = await getArticle(route.params.id)
    await loadComments()
  } catch {
    error.value = '加载失败'
  } finally {
    loading.value = false
  }
}

async function loadComments() {
  try {
    const data = await getComments(route.params.id, { page: 1, page_size: 50 })
    comments.value = data?.list || []
  } catch {}
}

async function submitComment() {
  if (!commentContent.value.trim()) return
  submitting.value = true
  try {
    await createComment({
      article_id: Number(route.params.id),
      content: commentContent.value,
      type: 'article'
    })
    commentContent.value = ''
    await loadComments()
  } catch {}
  finally { submitting.value = false }
}

function formatDate(date) {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-CN')
}
</script>

<style lang="scss" scoped>
.article-detail { max-width: 800px; }

.article-content {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 32px;
  box-shadow: var(--shadow-sm);
}

.article-header {
  margin-bottom: 24px;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border-color);
}

.title {
  font-size: 28px;
  font-weight: 600;
  line-height: 1.4;
  margin-bottom: 12px;
}

.meta {
  display: flex;
  gap: 16px;
  font-size: 14px;
  color: var(--text-muted);
  margin-bottom: 12px;

  span { display: flex; align-items: center; gap: 4px; }
}

.tags { display: flex; gap: 8px; }

.cover {
  margin-bottom: 24px;
  border-radius: var(--radius);
  overflow: hidden;
  img { width: 100%; }
}

.content-body {
  font-size: 16px;
  line-height: 1.8;
  color: var(--text-primary);

  h1, h2, h3 { margin: 24px 0 12px; font-weight: 600; }
  p { margin-bottom: 16px; }
  img { max-width: 100%; border-radius: var(--radius); }
  pre { background: #1e1e1e; color: #d4d4d4; padding: 16px; border-radius: var(--radius); overflow-x: auto; }
  code { background: rgba(0,0,0,0.06); padding: 2px 6px; border-radius: 4px; font-size: 14px; }
}

.comment-section {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 24px;
  margin-top: 24px;

  h3 { font-size: 18px; margin-bottom: 16px; }
  .count { color: var(--text-muted); font-size: 14px; }
}

.comment-form { margin-bottom: 24px; }

.login-tip {
  padding: 16px;
  background: rgba(64, 158, 255, 0.06);
  border-radius: var(--radius);
  text-align: center;
  margin-bottom: 24px;
  a { color: var(--primary); }
}

.comment-list { display: flex; flex-direction: column; gap: 16px; }

.comment-item {
  display: flex;
  gap: 12px;
}

.comment-body { flex: 1; }

.comment-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  .username { font-weight: 500; color: var(--text-primary); }
  .date { font-size: 12px; color: var(--text-muted); }
}

.comment-text {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.6;
}
</style>
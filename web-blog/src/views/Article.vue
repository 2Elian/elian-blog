<template>
  <div class="article-page" v-if="article">
    <!-- Article Header -->
    <div class="article-header">
      <div class="header-cover" v-if="article.cover">
        <img :src="article.cover" alt="cover" class="cover-image" />
      </div>
      <div class="header-content">
        <div class="header-meta">
          <span class="meta-category" v-if="article.category">
            <n-tag :bordered="false" size="small">{{ article.category.name }}</n-tag>
          </span>
          <span class="meta-date">{{ formatDate(article.created_at) }}</span>
          <span class="meta-views" v-if="article.views">
            <n-icon><EyeOutline /></n-icon> {{ article.views }} 阅读
          </span>
        </div>
        <h1 class="article-title">{{ article.title }}</h1>
        <div class="article-tags" v-if="article.tags?.length">
          <n-tag
            v-for="tag in article.tags"
            :key="tag.id"
            size="small"
            :bordered="false"
            class="tag-item"
          >
            # {{ tag.name }}
          </n-tag>
        </div>
      </div>
    </div>

    <!-- Article Body -->
    <div class="article-body">
      <div class="article-wrapper">
        <div class="article-content" v-html="renderedContent"></div>

        <!-- Prev / Next Navigation -->
        <div class="article-nav" v-if="prevArticle || nextArticle">
          <div class="nav-item prev" v-if="prevArticle" @click="goToArticle(prevArticle.id)">
            <span class="nav-label">&larr; 上一篇</span>
            <span class="nav-title">{{ prevArticle.title }}</span>
          </div>
          <div v-else></div>
          <div class="nav-item next" v-if="nextArticle" @click="goToArticle(nextArticle.id)">
            <span class="nav-label">下一篇 &rarr;</span>
            <span class="nav-title">{{ nextArticle.title }}</span>
          </div>
        </div>

        <!-- Comments Section -->
        <div class="comments-section">
          <h3 class="section-title">评论 ({{ comments.length }})</h3>

          <!-- Comment Form -->
          <div class="comment-form">
            <n-input
              v-model:value="commentContent"
              type="textarea"
              placeholder="写下你的评论..."
              :rows="3"
            />
            <n-button
              type="primary"
              round
              :disabled="!commentContent.trim()"
              :loading="submitting"
              @click="submitComment"
              style="margin-top: 12px;"
            >
              发表评论
            </n-button>
          </div>

          <!-- Comments List -->
          <div class="comments-list">
            <div v-for="comment in comments" :key="comment.id" class="comment-item">
              <div class="comment-avatar">
                <n-avatar round :size="36">
                  {{ (comment.username || 'U').charAt(0).toUpperCase() }}
                </n-avatar>
              </div>
              <div class="comment-body">
                <div class="comment-header">
                  <span class="comment-author">{{ comment.username || '匿名用户' }}</span>
                  <span class="comment-time">{{ formatDate(comment.created_at) }}</span>
                </div>
                <p class="comment-text">{{ comment.content }}</p>
              </div>
            </div>
          </div>
          <n-empty v-if="comments.length === 0" description="暂无评论，快来抢沙发吧~" style="padding: 30px 0;" />
        </div>
      </div>
    </div>
  </div>

  <div class="loading-state" v-else>
    <n-spin size="large" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, computed, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { NTag, NInput, NButton, NAvatar, NIcon, NEmpty, NSpin, useMessage } from 'naive-ui'
import { EyeOutline } from '@vicons/ionicons5'
import mermaid from 'mermaid'
import { renderMarkdown } from '@/utils/markdown'
import { getArticle, getComments, postComment } from '@/api'
import { useUserStore } from '@/stores/user'

mermaid.initialize({ startOnLoad: false, theme: 'default' })

interface Article {
  id: number
  title: string
  content: string
  cover?: string
  created_at: string
  views?: number
  category?: { id: number; name: string }
  tags?: { id: number; name: string }[]
  prev_article?: { id: number; title: string }
  next_article?: { id: number; title: string }
}

interface Comment {
  id: number
  content: string
  username?: string
  created_at: string
}

const route = useRoute()
const router = useRouter()
const message = useMessage()
const userStore = useUserStore()

const article = ref<Article | null>(null)
const comments = ref<Comment[]>([])
const commentContent = ref('')
const submitting = ref(false)

const renderedContent = computed(() => {
  if (!article.value?.content) return ''
  return renderMarkdown(article.value.content)
})

const prevArticle = computed(() => article.value?.prev_article)
const nextArticle = computed(() => article.value?.next_article)

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

function goToArticle(id: number) {
  router.push(`/article/${id}`)
}

async function loadArticle(id: number) {
  try {
    const res = await getArticle(id) as any
    article.value = res.data

    // Load comments
    const commentRes = await getComments(id) as any
    comments.value = commentRes.data?.list || commentRes.data || []

    // Render mermaid diagrams
    await nextTick()
    mermaid.run()
  } catch (e) {
    console.error('Failed to load article:', e)
    message.error('加载文章失败')
  }
}

async function submitComment() {
  if (!userStore.isLoggedIn) {
    message.warning('请先登录')
    router.push('/login')
    return
  }

  if (!commentContent.value.trim() || !article.value) return

  submitting.value = true
  try {
    await postComment({
      article_id: article.value.id,
      content: commentContent.value.trim()
    })
    message.success('评论成功')
    commentContent.value = ''

    // Reload comments
    const commentRes = await getComments(article.value.id) as any
    comments.value = commentRes.data?.list || commentRes.data || []
  } catch (e: any) {
    message.error(e.message || '评论失败')
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  const id = Number(route.params.id)
  if (id) loadArticle(id)
})

watch(() => route.params.id, (newId) => {
  if (newId) {
    loadArticle(Number(newId))
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
})
</script>

<style scoped lang="scss">
.article-page {
  animation: fadeInUp 0.5s ease;
  max-width: 1100px;
  margin: 0 auto;
  padding: 0 32px;

  @media (max-width: 640px) {
    padding: 0 20px;
  }
}

.article-header {
  background: linear-gradient(135deg, #0a0a0a 0%, #2a2a2a 100%);
  padding: 0;
  border-radius: var(--radius-lg);
  margin-bottom: 30px;
  color: white;
  position: relative;
  overflow: hidden;

  .cover-image {
    width: 100%;
    height: 300px;
    object-fit: cover;
    display: block;
  }

  .header-content {
    padding: 50px 40px 60px;
    position: relative;
    z-index: 1;
  }

  &::before {
    content: '';
    position: absolute;
    top: -50px;
    right: -50px;
    width: 200px;
    height: 200px;
    background: rgba(255, 255, 255, 0.1);
    border-radius: 50%;
  }

  &::after {
    content: '';
    position: absolute;
    bottom: -30px;
    left: 10%;
    width: 120px;
    height: 120px;
    background: rgba(255, 255, 255, 0.05);
    border-radius: 50%;
  }

  @media (max-width: 640px) {
    .header-content {
      padding: 30px 20px 40px;
    }
    .cover-image {
      height: 180px;
    }
  }
}

.header-meta {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 16px;
  font-size: 14px;
  opacity: 0.9;

  .meta-category {
    :deep(.n-tag) {
      background: rgba(255, 255, 255, 0.2);
      color: white;
    }
  }

  .meta-views {
    display: flex;
    align-items: center;
    gap: 4px;
  }
}

.article-title {
  font-size: 32px;
  font-weight: 700;
  line-height: 1.3;
  margin-bottom: 16px;

  @media (max-width: 640px) {
    font-size: 24px;
  }
}

.article-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;

  .tag-item {
    background: rgba(255, 255, 255, 0.15);
    color: white;
    cursor: pointer;
    transition: background-color var(--transition-fast);

    &:hover {
      background: rgba(255, 255, 255, 0.25);
    }
  }
}

.article-body {
  max-width: 1060px;
  margin: 0 auto;
}

.article-wrapper {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 40px;
  box-shadow: var(--shadow-sm);

  @media (max-width: 640px) {
    padding: 24px 20px;
  }
}

.article-content {
  font-size: 16px;
  line-height: 1.8;
  color: var(--text-secondary);
  word-break: break-word;
}

// Prev / Next Navigation
.article-nav {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-top: 50px;
  padding-top: 30px;
  border-top: 1px solid var(--border-color);

  @media (max-width: 640px) {
    grid-template-columns: 1fr;
  }
}

.nav-item {
  padding: 16px;
  border-radius: var(--radius-sm);
  background: rgba(0, 0, 0, 0.02);
  cursor: pointer;
  transition: all var(--transition-fast);

  &:hover {
    background: rgba(0, 0, 0, 0.05);
  }

  &.next {
    text-align: right;
  }
}

.nav-label {
  font-size: 13px;
  color: var(--text-muted);
  display: block;
  margin-bottom: 6px;
}

.nav-title {
  font-size: 15px;
  color: var(--primary-color);
  font-weight: 500;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

// Comments Section
.comments-section {
  margin-top: 50px;
  padding-top: 30px;
  border-top: 1px solid var(--border-color);
}

.comment-form {
  margin-bottom: 30px;
}

.comments-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.comment-item {
  display: flex;
  gap: 14px;
  padding: 16px;
  border-radius: var(--radius-sm);
  transition: background-color var(--transition-fast);

  &:hover {
    background: rgba(0, 0, 0, 0.02);
  }
}

.comment-avatar {
  flex-shrink: 0;

  :deep(.n-avatar) {
    background: var(--text-primary);
    color: var(--bg-card);
    font-weight: 600;
  }
}

.comment-body {
  flex: 1;
  min-width: 0;
}

.comment-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.comment-author {
  font-weight: 600;
  color: var(--text-primary);
  font-size: 14px;
}

.comment-time {
  font-size: 12px;
  color: var(--text-muted);
}

.comment-text {
  color: var(--text-secondary);
  font-size: 14px;
  line-height: 1.6;
  margin: 0;
}

.loading-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
}
</style>
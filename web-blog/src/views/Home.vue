<template>
  <div class="home-page">
    <!-- Hero Section (Full Width) -->
    <section class="hero">
      <div class="hero-bg">
        <div class="hero-grid"></div>
        <div class="hero-shape shape-1"></div>
        <div class="hero-shape shape-2"></div>
        <div class="hero-shape shape-3"></div>
      </div>
      <div class="hero-content">
        <div class="hero-badge">Welcome to my blog</div>
        <h1 class="hero-title">
<!--          <span class="greeting">欢迎来到我的博客</span>-->
          <span class="name">Elian</span>
        </h1>
        <p class="hero-subtitle">
          {{ typedText }}<span class="cursor">|</span>
        </p>
        <p class="hero-desc">
          Agent算法工程师 / LLM后训练工程师 / 后端开发爱好者
        </p>
        <div class="hero-actions">
          <button class="hero-btn primary" @click="$router.push('/blog')">
            <span>开始阅读</span>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M5 12h14M12 5l7 7-7 7"/></svg>
          </button>
          <button class="hero-btn ghost" @click="$router.push('/about')">
            关于作者
          </button>
        </div>
        <div class="hero-stats">
          <div class="stat-item">
            <span class="stat-value">{{ articleCount }}</span>
            <span class="stat-label">篇文章</span>
          </div>
          <div class="stat-divider"></div>
          <div class="stat-item">
            <span class="stat-value">{{ categoryCount }}</span>
            <span class="stat-label">个分类</span>
          </div>
          <div class="stat-divider"></div>
          <div class="stat-item">
            <span class="stat-value">{{ tagCount }}</span>
            <span class="stat-label">个标签</span>
          </div>
        </div>
      </div>
    </section>

    <!-- Main Content -->
    <div class="home-main">
      <div class="main-content">
        <!-- Latest Articles -->
        <section class="section">
          <div class="section-header">
            <h2 class="section-title">最新文章</h2>
            <button class="view-all-btn" @click="$router.push('/blog')">
              查看全部
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M5 12h14M12 5l7 7-7 7"/></svg>
            </button>
          </div>
          <div class="articles-grid">
            <ArticleCard
              v-for="article in articles"
              :key="article.id"
              :article="article"
            />
          </div>
          <n-empty v-if="!loading && articles.length === 0" description="暂无文章" />
        </section>
      </div>

      <Sidebar />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { NEmpty } from 'naive-ui'
import ArticleCard from '@/components/ArticleCard.vue'
import Sidebar from '@/components/Sidebar.vue'
import { getArticles, getCategories, getTags } from '@/api'

interface Article {
  id: number
  title: string
  summary?: string
  cover?: string
  created_at: string
  views?: number
  comments_count?: number
  category?: { id: number; name: string }
  tags?: { id: number; name: string }[]
}

const articles = ref<Article[]>([])
const loading = ref(true)
const articleCount = ref(0)
const categoryCount = ref(0)
const tagCount = ref(0)

// Typing effect
const subtitleTexts = [
  'AutoModelForCauseModel.from_pretrained(elian)',
  'elian-cli init agent',
  'https://github.com/2Elian',
  '欢迎来到Elian博客, 在这里您能看到LLM与Agent相关的技术分享',
  '励志称为一名优秀的Agent算法工程师'
]
const typedText = ref('')
let currentTextIndex = 0
let charIndex = 0
let isDeleting = false
let typingTimer: ReturnType<typeof setTimeout> | null = null

function typeEffect() {
  const currentText = subtitleTexts[currentTextIndex]

  if (!isDeleting) {
    typedText.value = currentText.substring(0, charIndex + 1)
    charIndex++

    if (charIndex === currentText.length) {
      isDeleting = true
      typingTimer = setTimeout(typeEffect, 2000)
      return
    }
  } else {
    typedText.value = currentText.substring(0, charIndex - 1)
    charIndex--

    if (charIndex === 0) {
      isDeleting = false
      currentTextIndex = (currentTextIndex + 1) % subtitleTexts.length
    }
  }

  typingTimer = setTimeout(typeEffect, isDeleting ? 40 : 80)
}

onMounted(async () => {
  typingTimer = setTimeout(typeEffect, 500)

  try {
    const [articlesRes, categoriesRes, tagsRes] = await Promise.all([
      getArticles({ page: 1, page_size: 6 }) as any,
      getCategories() as any,
      getTags() as any
    ])
    const data = articlesRes.data
    articles.value = data?.list || data || []
    articleCount.value = data?.total || articles.value.length
    categoryCount.value = (categoriesRes.data || []).length
    tagCount.value = (tagsRes.data || []).length
  } catch (e) {
    console.error('Failed to load articles:', e)
  } finally {
    loading.value = false
  }
})

onUnmounted(() => {
  if (typingTimer) clearTimeout(typingTimer)
})
</script>

<style scoped lang="scss">
.home-page {
  animation: fadeInUp 0.5s ease;
}

// ===== Hero Section (Full Width) =====
.hero {
  position: relative;
  width: 100vw;
  margin-left: calc(-50vw + 50%);
  min-height: 520px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  background: linear-gradient(135deg, #0a0a0a 0%, #1a1a1a 50%, #111111 100%);
  padding: 80px 32px;

  @media (max-width: 640px) {
    min-height: 440px;
    padding: 60px 20px;
  }
}

.hero-bg {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

.hero-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.03) 1px, transparent 1px);
  background-size: 60px 60px;
}

.hero-shape {
  position: absolute;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.06), transparent 70%);

  &.shape-1 {
    width: 500px;
    height: 500px;
    top: -200px;
    right: -100px;
  }

  &.shape-2 {
    width: 350px;
    height: 350px;
    bottom: -150px;
    left: -100px;
  }

  &.shape-3 {
    width: 200px;
    height: 200px;
    top: 40%;
    left: 15%;
  }
}
/*
属性	值	作用
position	relative	相对定位，让内部绝对定位的子元素相对于这个容器定位，同时不脱离正常文档流
text-align	center	内部文本/行内元素水平居中显示
color	white	文字颜色为白色
max-width	640px	最大宽度为 640 像素，超过这个宽度不再变宽，小于时自适应
*/
.hero-content {
  position: relative;
  text-align: center;
  color: white;
  max-width: 640px;
}

.hero-badge {
  display: inline-block;
  padding: 6px 16px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 24px;
  backdrop-filter: blur(10px);
  letter-spacing: 0.5px;
}

.hero-title {
  margin-bottom: 16px;

  .greeting {
    display: block;
    font-size: 18px;
    font-weight: 400;
    opacity: 0.6;
    margin-bottom: 8px;
    letter-spacing: 2px;
  }

  .name {
    display: block;
    font-size: 56px;
    font-weight: 800;
    letter-spacing: -1px;
    color: #ffffff;

    @media (max-width: 640px) {
      font-size: 40px;
    }
  }
}

.hero-subtitle {
  font-size: 18px;
  opacity: 0.8;
  margin-bottom: 8px;
  min-height: 28px;

  .cursor {
    animation: pulse 1s infinite;
    font-weight: 100;
  }
}

.hero-desc {
  font-size: 14px;
  opacity: 0.5;
  margin-bottom: 32px;
  letter-spacing: 1px;
}

.hero-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
  margin-bottom: 40px;
}

.hero-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 28px;
  border-radius: 8px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: all var(--transition-fast);
  border: none;

  &.primary {
    background: #ffffff;
    color: #0a0a0a;

    &:hover {
      background: #f0f0f0;
      transform: translateY(-2px);
      box-shadow: 0 4px 20px rgba(255, 255, 255, 0.2);
    }
  }

  &.ghost {
    background: transparent;
    color: rgba(255, 255, 255, 0.8);
    border: 1px solid rgba(255, 255, 255, 0.25);

    &:hover {
      border-color: rgba(255, 255, 255, 0.5);
      background: rgba(255, 255, 255, 0.05);
    }
  }
}

.hero-stats {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 24px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.stat-value {
  font-size: 24px;
  font-weight: 700;
  color: #ffffff;
}

.stat-label {
  font-size: 12px;
  color: rgba(255, 255, 255, 0.5);
}

.stat-divider {
  width: 1px;
  height: 30px;
  background: rgba(255, 255, 255, 0.15);
}

// ===== Main Content =====
.home-main {
  max-width: 1400px;
  margin: 40px auto 0;
  padding: 0 32px;
  display: grid;
  grid-template-columns: 1fr 300px;
  gap: 24px;

  @media (max-width: 1024px) {
    grid-template-columns: 1fr;
    padding: 0 20px;
  }
}

.section {
  margin-bottom: 50px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;

  .section-title {
    margin-bottom: 0;
  }
}

.view-all-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  border: none;
  background: none;
  color: var(--text-secondary);
  font-size: 14px;
  cursor: pointer;
  border-radius: 6px;
  transition: all var(--transition-fast);

  &:hover {
    color: var(--text-primary);
    background: rgba(0, 0, 0, 0.05);
  }
}

.articles-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;

  @media (max-width: 600px) {
    grid-template-columns: 1fr;
  }
}
</style>

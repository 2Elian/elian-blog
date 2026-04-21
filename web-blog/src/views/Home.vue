<template>
  <div class="home-page">
    <!-- Hero Section -->
    <section class="hero">
      <div class="hero-bg">
        <div class="hero-shape shape-1"></div>
        <div class="hero-shape shape-2"></div>
        <div class="hero-shape shape-3"></div>
      </div>
      <div class="hero-content">
        <h1 class="hero-title">
          <span class="greeting">Hello, I'm</span>
          <span class="gradient-text name">Elian</span>
        </h1>
        <p class="hero-subtitle">
          {{ typedText }}<span class="cursor">|</span>
        </p>
        <div class="hero-actions">
          <n-button type="primary" size="large" round @click="$router.push('/blog')">
            开始阅读
          </n-button>
          <n-button size="large" round ghost @click="$router.push('/about')">
            关于我
          </n-button>
        </div>
      </div>
    </section>

    <!-- Latest Articles Section -->
    <section class="section">
      <div class="section-header">
        <h2 class="section-title">最新文章</h2>
        <n-button text type="primary" @click="$router.push('/blog')">
          查看全部 &rarr;
        </n-button>
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

    <!-- Two Column Layout for Bottom Section -->
    <div class="home-bottom">
      <div class="bottom-main">
        <!-- Recent Comments or other content -->
      </div>
      <Sidebar />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { NButton, NEmpty } from 'naive-ui'
import ArticleCard from '@/components/ArticleCard.vue'
import Sidebar from '@/components/Sidebar.vue'
import { getArticles } from '@/api'

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

// Typing effect
const subtitleTexts = ['热爱编程，热爱生活', '用代码改变世界', '记录技术与生活点滴', 'Stay hungry, Stay foolish']
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
    const res = await getArticles({ page: 1, page_size: 6 }) as any
    articles.value = res.data?.list || res.data || []
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

// Hero Section
.hero {
  position: relative;
  min-height: 420px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: var(--radius-lg);
  overflow: hidden;
  margin-bottom: 50px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 60px 40px;

  @media (max-width: 640px) {
    min-height: 320px;
    padding: 40px 24px;
  }
}

.hero-bg {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

.hero-shape {
  position: absolute;
  border-radius: 50%;
  opacity: 0.1;

  &.shape-1 {
    width: 300px;
    height: 300px;
    background: white;
    top: -80px;
    right: -60px;
  }

  &.shape-2 {
    width: 200px;
    height: 200px;
    background: white;
    bottom: -40px;
    left: -40px;
  }

  &.shape-3 {
    width: 120px;
    height: 120px;
    background: white;
    top: 50%;
    left: 20%;
  }
}

.hero-content {
  position: relative;
  text-align: center;
  color: white;
}

.hero-title {
  margin-bottom: 16px;

  .greeting {
    display: block;
    font-size: 18px;
    font-weight: 400;
    opacity: 0.9;
    margin-bottom: 8px;
  }

  .name {
    font-size: 52px;
    font-weight: 800;
    background: linear-gradient(to right, #fff, #ffd1dc);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;

    @media (max-width: 640px) {
      font-size: 36px;
    }
  }
}

.hero-subtitle {
  font-size: 18px;
  opacity: 0.9;
  margin-bottom: 30px;
  min-height: 28px;

  .cursor {
    animation: pulse 1s infinite;
    font-weight: 100;
  }
}

.hero-actions {
  display: flex;
  gap: 16px;
  justify-content: center;

  :deep(.n-button--primary-type) {
    background: rgba(255, 255, 255, 0.2);
    border-color: rgba(255, 255, 255, 0.5);
    backdrop-filter: blur(10px);

    &:hover {
      background: rgba(255, 255, 255, 0.3);
      border-color: rgba(255, 255, 255, 0.7);
    }
  }

  :deep(.n-button--default-type) {
    color: white;
    border-color: rgba(255, 255, 255, 0.4);

    &:hover {
      border-color: rgba(255, 255, 255, 0.6);
    }
  }
}

// Section
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

.articles-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;

  @media (max-width: 900px) {
    grid-template-columns: repeat(2, 1fr);
  }

  @media (max-width: 600px) {
    grid-template-columns: 1fr;
  }
}

.home-bottom {
  display: grid;
  grid-template-columns: 1fr 320px;
  gap: 24px;

  @media (max-width: 1024px) {
    grid-template-columns: 1fr;
  }
}
</style>
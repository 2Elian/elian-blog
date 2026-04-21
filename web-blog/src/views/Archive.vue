<template>
  <div class="archive-page">
    <div class="page-header">
      <h1 class="page-title gradient-text">文章归档</h1>
      <p class="page-desc">共 {{ totalArticles }} 篇文章</p>
    </div>

    <div class="archive-content">
      <div class="timeline">
        <div
          v-for="(group, index) in archiveGroups"
          :key="group.year"
          class="timeline-year"
          :style="{ animationDelay: index * 0.1 + 's' }"
        >
          <div class="year-header">
            <span class="year-dot"></span>
            <h2 class="year-title">{{ group.year }}</h2>
            <span class="year-count">{{ group.articles.length }} 篇</span>
          </div>

          <div class="year-articles">
            <div
              v-for="article in group.articles"
              :key="article.id"
              class="article-item"
              @click="goToArticle(article.id)"
            >
              <span class="article-date">{{ formatDateShort(article.created_at) }}</span>
              <span class="article-title">{{ article.title }}</span>
              <span class="article-category" v-if="article.category">{{ article.category.name }}</span>
            </div>
          </div>
        </div>
      </div>

      <n-empty v-if="!loading && archiveGroups.length === 0" description="暂无文章" style="padding: 60px 0;" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { NEmpty } from 'naive-ui'
import { getArticles } from '@/api'

interface Article {
  id: number
  title: string
  created_at: string
  category?: { id: number; name: string }
}

interface ArchiveGroup {
  year: string
  articles: Article[]
}

const router = useRouter()
const articles = ref<Article[]>([])
const loading = ref(true)

const totalArticles = computed(() => articles.value.length)

const archiveGroups = computed(() => {
  const groups: Record<string, Article[]> = {}

  articles.value.forEach(article => {
    const year = new Date(article.created_at).getFullYear().toString()
    if (!groups[year]) groups[year] = []
    groups[year].push(article)
  })

  return Object.entries(groups)
    .sort((a, b) => b[0].localeCompare(a[0]))
    .map(([year, arts]) => ({
      year,
      articles: arts.sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
    }))
})

function formatDateShort(date: string) {
  const d = new Date(date)
  return `${d.getMonth() + 1}-${String(d.getDate()).padStart(2, '0')}`
}

function goToArticle(id: number) {
  router.push(`/article/${id}`)
}

onMounted(async () => {
  try {
    const res = await getArticles({ page: 1, page_size: 1000 }) as any
    articles.value = res.data?.list || res.data || []
  } catch (e) {
    console.error('Failed to load archive:', e)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped lang="scss">
.archive-page {
  animation: fadeInUp 0.5s ease;
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

.archive-content {
  max-width: 800px;
  margin: 0 auto;
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 30px;
  box-shadow: var(--shadow-sm);
}

.timeline {
  position: relative;
  padding-left: 20px;

  &::before {
    content: '';
    position: absolute;
    left: 6px;
    top: 0;
    bottom: 0;
    width: 2px;
    background: linear-gradient(to bottom, var(--primary-color), #764ba2);
    opacity: 0.3;
  }
}

.timeline-year {
  animation: fadeInUp 0.5s ease forwards;
  opacity: 0;
  margin-bottom: 30px;

  &:last-child {
    margin-bottom: 0;
  }
}

.year-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 16px;
}

.year-dot {
  position: relative;
  width: 14px;
  height: 14px;
  background: var(--primary-color);
  border-radius: 50%;
  z-index: 1;
  box-shadow: 0 0 0 4px var(--bg-card), 0 0 0 6px rgba(233, 84, 107, 0.2);

  &::before {
    content: '';
    position: absolute;
    inset: -4px;
    border-radius: 50%;
    border: 2px solid var(--primary-color);
    animation: pulse 2s infinite;
  }
}

.year-title {
  font-size: 24px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0;
}

.year-count {
  font-size: 14px;
  color: var(--text-muted);
  background: rgba(233, 84, 107, 0.08);
  padding: 2px 10px;
  border-radius: 10px;
}

.year-articles {
  padding-left: 30px;
}

.article-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  margin-bottom: 8px;
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all var(--transition-fast);
  border-left: 2px solid transparent;

  &:hover {
    background: rgba(233, 84, 107, 0.05);
    border-left-color: var(--primary-color);

    .article-title {
      color: var(--primary-color);
    }
  }

  &:last-child {
    margin-bottom: 0;
  }
}

.article-date {
  font-size: 13px;
  color: var(--text-muted);
  font-family: monospace;
  min-width: 40px;
}

.article-title {
  flex: 1;
  color: var(--text-primary);
  font-size: 15px;
  transition: color var(--transition-fast);
}

.article-category {
  font-size: 12px;
  color: var(--primary-color);
  background: rgba(233, 84, 107, 0.1);
  padding: 2px 10px;
  border-radius: 10px;
}
</style>
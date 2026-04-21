<template>
  <div class="tags-page">
    <div class="page-header">
      <h1 class="page-title gradient-text">标签云</h1>
      <p class="page-desc">点击标签查看相关文章</p>
    </div>

    <div class="tags-container" v-if="!selectedTag">
      <div
        v-for="tag in tags"
        :key="tag.id"
        class="tag-card"
        :style="{ fontSize: getTagSize(tag.count) + 'px' }"
        @click="selectTag(tag)"
      >
        <span class="tag-hash">#</span>
        {{ tag.name }}
        <span class="tag-count" v-if="tag.count">{{ tag.count }}</span>
      </div>
    </div>

    <div class="selected-tag" v-else>
      <div class="tag-header">
        <n-button text @click="clearSelection">
          &larr; 返回
        </n-button>
        <h2 class="tag-title">
          #{{ selectedTag.name }}
          <span class="count">{{ filteredArticles.length }} 篇文章</span>
        </h2>
      </div>

      <div class="articles-list">
        <div
          v-for="article in filteredArticles"
          :key="article.id"
          class="article-row"
          @click="goToArticle(article.id)"
        >
          <span class="article-date">{{ formatDate(article.created_at) }}</span>
          <span class="article-title">{{ article.title }}</span>
        </div>
      </div>

      <n-empty v-if="filteredArticles.length === 0" description="该标签下暂无文章" />
    </div>

    <n-empty v-if="!loading && tags.length === 0" description="暂无标签" style="padding: 60px 0;" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { NButton, NEmpty } from 'naive-ui'
import { getTags, getArticles } from '@/api'

interface Tag {
  id: number
  name: string
  count?: number
}

interface Article {
  id: number
  title: string
  created_at: string
}

const router = useRouter()
const tags = ref<Tag[]>([])
const articles = ref<Article[]>([])
const selectedTag = ref<Tag | null>(null)
const filteredArticles = ref<Article[]>([])
const loading = ref(true)

function getTagSize(count?: number) {
  if (!count) return 16
  return Math.min(24, 14 + count * 1.5)
}

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

async function selectTag(tag: Tag) {
  selectedTag.value = tag
  try {
    const res = await getArticles({ tag_id: tag.id, page: 1, page_size: 100 }) as any
    filteredArticles.value = res.data?.list || []
  } catch (e) {
    console.error('Failed to load articles by tag:', e)
  }
}

function clearSelection() {
  selectedTag.value = null
  filteredArticles.value = []
}

function goToArticle(id: number) {
  router.push(`/article/${id}`)
}

onMounted(async () => {
  try {
    const [tagsRes, articlesRes] = await Promise.all([
      getTags() as any,
      getArticles({ page: 1, page_size: 1000 }) as any
    ])
    tags.value = tagsRes.data || []
    articles.value = articlesRes.data?.list || []
  } catch (e) {
    console.error('Failed to load tags:', e)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped lang="scss">
.tags-page {
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

.tags-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 16px;
  padding: 20px;
  background: var(--bg-card);
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-sm);
}

.tag-card {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 12px 24px;
  background: linear-gradient(135deg, rgba(237, 110, 160, 0.1) 0%, rgba(236, 140, 105, 0.1) 100%);
  border-radius: 24px;
  cursor: pointer;
  transition: all var(--transition-normal);
  color: var(--text-primary);
  font-weight: 500;

  &:hover {
    background: linear-gradient(135deg, rgba(237, 110, 160, 0.2) 0%, rgba(236, 140, 105, 0.2) 100%);
    transform: translateY(-4px);
    box-shadow: var(--shadow-md);
  }

  .tag-hash {
    color: var(--primary-color);
    font-weight: 600;
  }

  .tag-count {
    font-size: 0.7em;
    background: var(--primary-color);
    color: white;
    padding: 2px 8px;
    border-radius: 10px;
    margin-left: 4px;
  }
}

.selected-tag {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 30px;
  box-shadow: var(--shadow-sm);
}

.tag-header {
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--border-color);
}

.tag-title {
  font-size: 24px;
  font-weight: 600;
  color: var(--text-primary);
  margin-top: 12px;
  display: flex;
  align-items: center;
  gap: 12px;

  .count {
    font-size: 14px;
    font-weight: 400;
    color: var(--text-muted);
  }
}

.articles-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.article-row {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 14px 16px;
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: all var(--transition-fast);

  &:hover {
    background: rgba(233, 84, 107, 0.05);

    .article-title {
      color: var(--primary-color);
    }
  }
}

.article-date {
  font-size: 13px;
  color: var(--text-muted);
  font-family: monospace;
  min-width: 80px;
}

.article-title {
  flex: 1;
  color: var(--text-primary);
  font-size: 15px;
  transition: color var(--transition-fast);
}
</style>
<template>
  <div class="blog-page">
    <div class="blog-layout">
      <!-- Left Sidebar: Categories -->
      <aside class="blog-sidebar-left">
        <div class="sidebar-section">
          <h4 class="sidebar-title">分类</h4>
          <div class="sidebar-cat-list">
            <div
              class="sidebar-cat-item"
              :class="{ active: !selectedCategory }"
              @click="selectCategory(null)"
            >全部</div>
            <div
              v-for="cat in categories"
              :key="cat.id"
              class="sidebar-cat-item"
              :class="{ active: selectedCategory === cat.id }"
              @click="selectCategory(cat.id)"
            >{{ cat.name }}</div>
          </div>
        </div>
      </aside>

      <!-- Main Content -->
      <div class="blog-main">
        <!-- Search Bar -->
        <div class="search-bar">
          <n-input
            v-model:value="keyword"
            placeholder="搜索文章..."
            clearable
            round
            size="large"
            @keyup.enter="handleSearch"
          >
            <template #prefix>
              <n-icon><SearchOutline /></n-icon>
            </template>
          </n-input>
        </div>

        <!-- Article List -->
        <div class="article-list">
          <ArticleCard
            v-for="article in articles"
            :key="article.id"
            :article="article"
            class="article-list-card"
          />
        </div>

        <n-empty v-if="!loading && articles.length === 0" description="暂无文章" style="padding: 60px 0;" />

        <!-- Pagination -->
        <div class="pagination" v-if="total > pageSize">
          <n-pagination
            v-model:page="currentPage"
            :page-count="Math.ceil(total / pageSize)"
            :page-size="pageSize"
            show-quick-jumper
            @update:page="loadArticles"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { NInput, NIcon, NEmpty, NPagination } from 'naive-ui'
import { SearchOutline } from '@vicons/ionicons5'
import ArticleCard from '@/components/ArticleCard.vue'
import { getArticles, searchArticles, getCategories } from '@/api'

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

interface Category {
  id: number
  name: string
  count?: number
}

const route = useRoute()
const articles = ref<Article[]>([])
const categories = ref<Category[]>([])
const loading = ref(true)
const keyword = ref('')
const selectedCategory = ref<number | null>(null)
const currentPage = ref(1)
const total = ref(0)
const pageSize = 10

function selectCategory(id: number | null) {
  selectedCategory.value = id
  currentPage.value = 1
  loadArticles()
}

function handleSearch() {
  currentPage.value = 1
  loadArticles()
}

async function loadArticles() {
  loading.value = true
  try {
    let res: any
    const params: any = { page: currentPage.value, page_size: pageSize }

    if (selectedCategory.value) {
      params.category_id = selectedCategory.value
    }

    if (keyword.value.trim()) {
      res = await searchArticles({ keyword: keyword.value.trim(), ...params })
    } else {
      res = await getArticles(params)
    }

    articles.value = res.data?.list || []
    total.value = res.data?.total || 0
  } catch (e) {
    console.error('Failed to load articles:', e)
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  // Check query params
  if (route.query.search) {
    keyword.value = route.query.search as string
  }
  if (route.query.category) {
    selectedCategory.value = Number(route.query.category)
  }
  if (route.query.tag) {
    // Will be handled by articles API with tag filter
  }

  try {
    const catRes = await getCategories() as any
    categories.value = (catRes.data || []).map((c: any) => ({ id: c.id, name: c.name }))
  } catch (e) {
    console.error('Failed to load categories:', e)
  }

  loadArticles()
})

watch(() => route.query, () => {
  if (route.query.search) keyword.value = route.query.search as string
  if (route.query.category) selectedCategory.value = Number(route.query.category)
  loadArticles()
})
</script>

<style scoped lang="scss">
.blog-page {
  animation: fadeInUp 0.5s ease;
  max-width: 1100px;
  margin: 0 auto;
  padding: 0 32px;

  @media (max-width: 1024px) {
    padding: 0 20px;
  }
}

.blog-layout {
  display: grid;
  grid-template-columns: 180px 1fr;
  gap: 28px;

  @media (max-width: 900px) {
    grid-template-columns: 1fr;
  }
}

.blog-sidebar-left {
  position: sticky;
  top: calc(var(--header-height, 64px) + 20px);
  height: fit-content;

  @media (max-width: 900px) {
    position: static;
  }
}

.sidebar-section {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 16px;
  box-shadow: var(--shadow-sm);
}

.sidebar-title {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 12px;
  padding-bottom: 10px;
  border-bottom: 1px solid var(--border-color);
}

.sidebar-cat-list {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.sidebar-cat-item {
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  color: var(--text-secondary);
  transition: all var(--transition-fast);

  &:hover {
    background: rgba(0, 0, 0, 0.04);
    color: var(--text-primary);
  }

  &.active {
    background: var(--primary-color);
    color: white;
    font-weight: 500;
  }
}

html.dark .sidebar-cat-item:hover {
  background: rgba(255, 255, 255, 0.08);
}

.blog-main {
  min-width: 0;
}

.search-bar {
  margin-bottom: 20px;

  :deep(.n-input) {
    border-radius: 24px;
  }
}

.article-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.article-list-card {
  :deep(.article-card) {
    --card-cover-height: 100%;
    display: grid;
    grid-template-columns: 180px 1fr;
    border-radius: var(--radius-md);

    @media (max-width: 640px) {
      grid-template-columns: 1fr;
      --card-cover-height: 140px;
    }
  }

  :deep(.card-cover) {
    min-height: 80px;
    border-radius: var(--radius-md) 0 0 var(--radius-md);

    @media (max-width: 640px) {
      border-radius: var(--radius-md) var(--radius-md) 0 0;
    }
  }

  :deep(.card-content) {
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 14px 16px;
  }

  :deep(.card-desc) {
    -webkit-line-clamp: 1;
    margin-bottom: 8px;
  }

  :deep(.card-title) {
    font-size: 15px;
    margin-bottom: 6px;
  }

  :deep(.card-meta) {
    font-size: 12px;
    gap: 12px;
  }
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 40px;
}
</style>
<template>
  <div class="blog-page">
    <div class="blog-layout">
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

        <!-- Category Filter -->
        <div class="category-filter" v-if="categories.length">
          <n-button
            :type="!selectedCategory ? 'primary' : 'default'"
            size="small"
            round
            @click="selectCategory(null)"
          >
            全部
          </n-button>
          <n-button
            v-for="cat in categories"
            :key="cat.id"
            :type="selectedCategory === cat.id ? 'primary' : 'default'"
            size="small"
            round
            @click="selectCategory(cat.id)"
          >
            {{ cat.name }}
          </n-button>
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

      <!-- Sidebar -->
      <Sidebar />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { NInput, NIcon, NButton, NEmpty, NPagination } from 'naive-ui'
import { SearchOutline } from '@vicons/ionicons5'
import ArticleCard from '@/components/ArticleCard.vue'
import Sidebar from '@/components/Sidebar.vue'
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
    categories.value = catRes.data || []
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
}

.blog-layout {
  display: grid;
  grid-template-columns: 1fr 320px;
  gap: 24px;

  @media (max-width: 1024px) {
    grid-template-columns: 1fr;
  }
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

.category-filter {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 24px;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border-color);
}

.article-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.article-list-card {
  :deep(.article-card) {
    display: grid;
    grid-template-columns: 240px 1fr;
    border-radius: var(--radius-md);

    @media (max-width: 640px) {
      grid-template-columns: 1fr;
    }
  }

  :deep(.card-cover) {
    height: 100%;
    min-height: 180px;
    border-radius: var(--radius-md) 0 0 var(--radius-md);

    @media (max-width: 640px) {
      border-radius: var(--radius-md) var(--radius-md) 0 0;
      height: 180px;
    }
  }
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 40px;
}
</style>
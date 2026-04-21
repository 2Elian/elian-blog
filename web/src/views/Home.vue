<template>
  <div class="home container">
    <!-- 置顶文章 -->
    <section class="section pinned-section" v-if="pinnedArticles.length">
      <h2 class="section-title">置顶推荐</h2>
      <div class="pinned-grid">
        <ArticleCard v-for="article in pinnedArticles" :key="article.id" :article="article" />
      </div>
    </section>

    <!-- 最新文章 -->
    <section class="section">
      <div class="section-header">
        <h2 class="section-title">最新文章</h2>
        <router-link to="/blog" class="view-more">
          查看更多 <el-icon><ArrowRight /></el-icon>
        </router-link>
      </div>
      <LoadingState :loading="loading" :error="error" :empty="!articles.length" emptyText="暂无文章" @retry="loadArticles">
        <div class="article-grid">
          <ArticleCard v-for="article in articles" :key="article.id" :article="article" />
        </div>
      </LoadingState>
    </section>

    <!-- 分类概览 -->
    <section class="section categories-section" v-if="categories.length">
      <h2 class="section-title">分类导航</h2>
      <div class="category-grid">
        <div class="category-card" v-for="cat in categories" :key="cat.id" @click="$router.push(`/blog?category=${cat.id}`)">
          <div class="cat-icon">
            <el-icon size="28"><Folder /></el-icon>
          </div>
          <div class="cat-info">
            <h3>{{ cat.name }}</h3>
            <p>{{ cat.article_count || 0 }} 篇文章</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import ArticleCard from '@/components/ArticleCard.vue'
import LoadingState from '@/components/LoadingState.vue'
import { getArticles, getCategories } from '@/api'

const loading = ref(true)
const error = ref('')
const articles = ref([])
const pinnedArticles = ref([])
const categories = ref([])

onMounted(async () => {
  await loadArticles()
  await loadCategories()
})

async function loadArticles() {
  loading.value = true
  error.value = ''
  try {
    const data = await getArticles({ page: 1, page_size: 6 })
    articles.value = data?.list || []
    pinnedArticles.value = articles.value.filter(a => a.is_top).slice(0, 3)
  } catch (e) {
    error.value = '加载失败'
  } finally {
    loading.value = false
  }
}

async function loadCategories() {
  try {
    const data = await getCategories()
    categories.value = data || []
  } catch {}
}
</script>

<style lang="scss" scoped>
.home {
  padding-bottom: 40px;
}

.section {
  margin-bottom: 40px;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 16px;
}

.view-more {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 14px;
  color: var(--primary);
}

.pinned-grid, .article-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;

  @media (max-width: 900px) {
    grid-template-columns: repeat(2, 1fr);
  }
  @media (max-width: 600px) {
    grid-template-columns: 1fr;
  }
}

.category-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;

  @media (max-width: 900px) {
    grid-template-columns: repeat(2, 1fr);
  }
}

.category-card {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: var(--shadow-sm);

  &:hover {
    box-shadow: var(--shadow-md);
    transform: translateY(-2px);
  }
}

.cat-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, var(--primary), #67c23a);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
}

.cat-info h3 {
  font-size: 15px;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.cat-info p {
  font-size: 13px;
  color: var(--text-muted);
}
</style>
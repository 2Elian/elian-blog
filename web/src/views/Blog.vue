<template>
  <div class="blog-page container">
    <div class="blog-layout">
      <!-- 左侧栏 -->
      <aside class="sidebar">
        <div class="sidebar-card">
          <h4>文章分类</h4>
          <ul class="cat-list">
            <li :class="{ active: !selectedCategory }" @click="selectCategory(null)">全部文章</li>
            <li v-for="cat in categories" :key="cat.id" :class="{ active: selectedCategory === cat.id }" @click="selectCategory(cat.id)">
              {{ cat.name }}
              <span class="count">{{ cat.article_count }}</span>
            </li>
          </ul>
        </div>

        <div class="sidebar-card">
          <h4>热门标签</h4>
          <div class="tag-list">
            <el-tag v-for="tag in tags" :key="tag.id" size="small" :color="tag.color" effect="dark" round class="tag-item">
              {{ tag.name }}
            </el-tag>
          </div>
        </div>

        <div class="sidebar-card user-card" v-if="userStore.isLoggedIn">
          <el-avatar :size="48" :src="userStore.userInfo?.avatar">
            {{ userStore.userInfo?.username?.charAt(0) }}
          </el-avatar>
          <div class="user-info">
            <h4>{{ userStore.userInfo?.username }}</h4>
            <p>欢迎回来</p>
          </div>
        </div>
        <div class="sidebar-card" v-else>
          <h4>登录</h4>
          <p class="login-tip">登录后可评论和留言</p>
          <router-link to="/login">
            <el-button type="primary" size="small" block>立即登录</el-button>
          </router-link>
        </div>
      </aside>

      <!-- 主内容 -->
      <div class="main-content">
        <div class="search-bar">
          <el-input v-model="searchKeyword" placeholder="搜索文章..." clearable @keyup.enter="handleSearch">
            <template #suffix>
              <el-icon class="search-icon" @click="handleSearch"><Search /></el-icon>
            </template>
          </el-input>
        </div>

        <div class="tabs-bar">
          <el-radio-group v-model="sortType" size="small">
            <el-radio-button value="latest">最新</el-radio-button>
            <el-radio-button value="hot">最热</el-radio-button>
            <el-radio-button value="featured">精选</el-radio-button>
          </el-radio-group>
        </div>

        <LoadingState :loading="loading" :error="error" :empty="!articles.length" emptyText="暂无文章" @retry="loadArticles">
          <div class="article-list">
            <ArticleCard v-for="article in articles" :key="article.id" :article="article" />
          </div>
          <el-pagination v-if="total > pageSize" :total="total" :page-size="pageSize" :current-page="page" layout="prev, pager, next" @current-change="onPageChange" background />
        </LoadingState>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import ArticleCard from '@/components/ArticleCard.vue'
import LoadingState from '@/components/LoadingState.vue'
import { getArticles, getCategories, getTags, searchArticles } from '@/api'
import { useUserStore } from '@/store/user'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const loading = ref(true)
const error = ref('')
const articles = ref([])
const categories = ref([])
const tags = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = 9
const selectedCategory = ref(null)
const sortType = ref('latest')
const searchKeyword = ref('')

onMounted(async () => {
  if (route.query.category) selectedCategory.value = Number(route.query.category)
  await loadCategories()
  await loadTags()
  await loadArticles()
})

watch([selectedCategory, sortType], () => loadArticles())

function selectCategory(catId) {
  selectedCategory.value = catId
  page.value = 1
}

async function loadArticles() {
  loading.value = true
  error.value = ''
  try {
    let data
    if (searchKeyword.value) {
      data = await searchArticles({ keyword: searchKeyword.value, page: page.value, page_size: pageSize })
    } else {
      const params = { page: page.value, page_size: pageSize }
      if (selectedCategory.value) params.category_id = selectedCategory.value
      if (sortType.value === 'hot') params.order = 'view_count'
      if (sortType.value === 'featured') params.is_featured = true
      data = await getArticles(params)
    }
    articles.value = data?.list || []
    total.value = data?.total || 0
  } catch (e) {
    error.value = '加载失败'
  } finally {
    loading.value = false
  }
}

async function loadCategories() {
  try {
    categories.value = await getCategories() || []
  } catch {}
}

async function loadTags() {
  try {
    tags.value = await getTags() || []
  } catch {}
}

function handleSearch() {
  page.value = 1
  loadArticles()
}

function onPageChange(p) {
  page.value = p
  loadArticles()
}
</script>

<style lang="scss" scoped>
.blog-page { min-height: calc(100vh - 140px); }

.blog-layout {
  display: grid;
  grid-template-columns: 260px 1fr;
  gap: 24px;
}

.sidebar { display: flex; flex-direction: column; gap: 16px; }

.sidebar-card {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 16px;
  box-shadow: var(--shadow-sm);

  h4 {
    font-size: 16px;
    margin-bottom: 12px;
    color: var(--text-primary);
  }
}

.cat-list {
  list-style: none;
  li {
    padding: 10px 12px;
    border-radius: 6px;
    cursor: pointer;
    font-size: 14px;
    color: var(--text-secondary);
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 4px;
    transition: all 0.2s;

    &:hover { background: rgba(64, 158, 255, 0.06); color: var(--primary); }
    &.active { background: rgba(64, 158, 255, 0.1); color: var(--primary); font-weight: 500; }
  }
  .count { font-size: 12px; color: var(--text-muted); }
}

.tag-list { display: flex; flex-wrap: wrap; gap: 8px; }
.tag-item { cursor: pointer; }

.login-tip { font-size: 13px; color: var(--text-muted); margin-bottom: 12px; }

.user-card {
  display: flex;
  align-items: center;
  gap: 12px;
  .user-info h4 { margin-bottom: 2px; font-size: 16px; }
  .user-info p { font-size: 13px; color: var(--text-muted); margin: 0; }
}

.main-content { display: flex; flex-direction: column; gap: 16px; }

.search-bar {
  .search-icon { cursor: pointer; }
}

.tabs-bar { margin-bottom: 8px; }

.article-list {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;

  @media (max-width: 900px) { grid-template-columns: repeat(2, 1fr); }
  @media (max-width: 600px) { grid-template-columns: 1fr; }
}

.el-pagination {
  margin-top: 24px;
  justify-content: center;
}
</style>
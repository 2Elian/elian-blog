<template>
  <aside class="sidebar">
    <!-- Author Card -->
    <div class="widget author-card">
      <div class="author-avatar">
        <n-avatar round :size="80" class="avatar">
          E
        </n-avatar>
      </div>
      <h3 class="author-name">Elian</h3>
      <p class="author-bio">热爱编程，热爱生活</p>
      <div class="author-stats">
        <div class="stat-item">
          <span class="stat-value">{{ stats.articles }}</span>
          <span class="stat-label">文章</span>
        </div>
        <div class="stat-item">
          <span class="stat-value">{{ stats.categories }}</span>
          <span class="stat-label">分类</span>
        </div>
        <div class="stat-item">
          <span class="stat-value">{{ stats.tags }}</span>
          <span class="stat-label">标签</span>
        </div>
      </div>
    </div>

    <!-- Categories Widget -->
    <div class="widget categories-widget" v-if="categories.length">
      <h4 class="widget-title">分类</h4>
      <div class="category-list">
        <div
          v-for="category in categories"
          :key="category.id"
          class="category-item"
          @click="filterByCategory(category.id)"
        >
          <span class="category-name">{{ category.name }}</span>
          <span class="category-count">{{ category.count || 0 }}</span>
        </div>
      </div>
    </div>

    <!-- Tags Cloud Widget -->
    <div class="widget tags-widget" v-if="tags.length">
      <h4 class="widget-title">标签云</h4>
      <div class="tags-cloud">
        <n-tag
          v-for="tag in tags"
          :key="tag.id"
          :bordered="false"
          class="tag-item"
          :style="{ fontSize: getTagSize(tag.count) + 'px' }"
          @click="filterByTag(tag.id)"
        >
          {{ tag.name }}
        </n-tag>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { NAvatar, NTag } from 'naive-ui'
import { getCategories, getTags } from '@/api'

interface Category {
  id: number
  name: string
  count?: number
}

interface Tag {
  id: number
  name: string
  count?: number
}

const router = useRouter()

const categories = ref<Category[]>([])
const tags = ref<Tag[]>([])
const stats = ref({
  articles: 0,
  categories: 0,
  tags: 0
})

onMounted(async () => {
  try {
    const [catRes, tagRes] = await Promise.all([
      getCategories() as any,
      getTags() as any
    ])
    categories.value = catRes.data || []
    tags.value = tagRes.data || []
    stats.value.categories = categories.value.length
    stats.value.tags = tags.value.length
  } catch (e) {
    console.error('Failed to load sidebar data:', e)
  }
})

function getTagSize(count?: number) {
  if (!count) return 13
  return Math.min(16, 12 + count)
}

function filterByCategory(id: number) {
  router.push({ path: '/blog', query: { category: String(id) } })
}

function filterByTag(id: number) {
  router.push({ path: '/blog', query: { tag: String(id) } })
}
</script>

<style scoped lang="scss">
.sidebar {
  display: flex;
  flex-direction: column;
  gap: 20px;
  position: sticky;
  top: calc(var(--header-height) + 20px);
  height: fit-content;
}

.widget {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 20px;
  box-shadow: var(--shadow-sm);
  transition: box-shadow var(--transition-normal);

  &:hover {
    box-shadow: var(--shadow-md);
  }
}

.widget-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 16px;
  padding-bottom: 10px;
  border-bottom: 1px solid var(--border-color);
}

// Author Card
.author-card {
  text-align: center;
  background: linear-gradient(180deg, rgba(233, 84, 107, 0.05) 0%, var(--bg-card) 100%);
}

.author-avatar {
  margin-bottom: 12px;
}

.avatar {
  background: var(--accent-gradient);
  color: white;
  font-size: 32px;
  font-weight: 700;
}

.author-name {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 6px;
}

.author-bio {
  font-size: 14px;
  color: var(--text-muted);
  margin-bottom: 16px;
}

.author-stats {
  display: flex;
  justify-content: space-around;
  padding-top: 16px;
  border-top: 1px solid var(--border-color);
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.stat-value {
  font-size: 20px;
  font-weight: 700;
  background: var(--accent-gradient);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.stat-label {
  font-size: 12px;
  color: var(--text-muted);
}

// Categories
.category-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.category-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 12px;
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: background-color var(--transition-fast);

  &:hover {
    background: rgba(233, 84, 107, 0.06);
  }

  .category-name {
    color: var(--text-secondary);
    font-size: 14px;
  }

  .category-count {
    background: rgba(233, 84, 107, 0.1);
    color: var(--primary-color);
    font-size: 12px;
    padding: 2px 8px;
    border-radius: 10px;
  }
}

// Tags Cloud
.tags-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tag-item {
  background: rgba(233, 84, 107, 0.08);
  color: var(--primary-color);
  cursor: pointer;
  transition: all var(--transition-fast);

  &:hover {
    background: rgba(233, 84, 107, 0.15);
    transform: translateY(-2px);
  }
}

@media (max-width: 1024px) {
  .sidebar {
    display: none;
  }
}
</style>
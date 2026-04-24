<template>
  <aside class="sidebar">
    <!-- Author Card (Commented out per user request)
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
    -->

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
import { NTag } from 'naive-ui'
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
    categories.value = (catRes.data || []).map((c: any) => ({ id: c.id, name: c.name, count: c.article_count }))
    tags.value = (tagRes.data || []).map((t: any) => ({ id: t.id, name: t.name, count: t.article_count }))
    stats.value.categories = categories.value.length
    stats.value.tags = tags.value.length
  } catch (e) {
    console.error('Failed to load sidebar data:', e)
  }
})

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
  font-size: 15px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 16px;
  padding-bottom: 10px;
  border-bottom: 1px solid var(--border-color);
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
  border-radius: 6px;
  cursor: pointer;
  transition: background-color var(--transition-fast);

  &:hover {
    background: rgba(0, 0, 0, 0.04);
  }

  .category-name {
    color: var(--text-secondary);
    font-size: 14px;
  }

  .category-count {
    background: rgba(0, 0, 0, 0.06);
    color: var(--text-secondary);
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
  background: rgba(0, 0, 0, 0.04);
  color: var(--text-secondary);
  cursor: pointer;
  transition: all var(--transition-fast);

  &:hover {
    background: rgba(0, 0, 0, 0.08);
    color: var(--text-primary);
  }
}

@media (max-width: 1024px) {
  .sidebar {
    display: none;
  }
}
</style>
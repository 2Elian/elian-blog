<template>
  <aside class="sidebar">
    <!-- Categories Widget -->
    <div class="widget categories-widget" v-if="categories && categories.length">
      <h4 class="widget-title">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="vertical-align: -2px; margin-right: 4px"><path d="M4 4h6v6H4zM14 4h6v6h-6zM4 14h6v6H4zM14 14h6v6h-6z"/></svg>
        分类
      </h4>
      <div class="category-list">
        <div
          class="category-item"
          :class="{ active: selectedCategory === null }"
          @click="$emit('selectCategory', null)"
        >
          <span class="category-name">全部</span>
        </div>
        <div
          v-for="cat in categories"
          :key="cat.id"
          class="category-item"
          :class="{ active: selectedCategory === cat.id }"
          @click="$emit('selectCategory', cat.id)"
        >
          <span class="category-name">{{ cat.name }}</span>
        </div>
      </div>
    </div>

    <!-- Notice Widget -->
    <div class="widget notice-widget" v-if="notices.length">
      <h4 class="widget-title">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="vertical-align: -2px; margin-right: 4px"><path d="M18 8A6 6 0 006 8c0 7-3 9-3 9h18s-3-2-3-9"/><path d="M13.73 21a2 2 0 01-3.46 0"/></svg>
        通知公告
      </h4>
      <div class="notice-list">
        <div v-for="(notice, idx) in notices" :key="idx" class="notice-item">
          <span class="notice-dot"></span>
          <span class="notice-text">{{ notice }}</span>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useSiteConfigStore } from '@/stores/siteConfig'

interface Category {
  id: number
  name: string
  count?: number
}

const props = defineProps<{
  categories?: Category[]
  selectedCategory?: number | null
}>()

defineEmits<{
  selectCategory: [id: number | null]
}>()

const siteConfig = useSiteConfigStore()
const notices = computed(() => siteConfig.siteNotice)
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

.notice-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.notice-item {
  display: flex;
  align-items: flex-start;
  gap: 10px;
}

.notice-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: var(--primary-color);
  margin-top: 7px;
  flex-shrink: 0;
}

.notice-text {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.6;
}

.category-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.category-item {
  display: flex;
  align-items: center;
  gap: 8px;
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

html.dark .category-item:hover {
  background: rgba(255, 255, 255, 0.08);
}

@media (max-width: 1024px) {
  .sidebar {
    display: none;
  }
}
</style>

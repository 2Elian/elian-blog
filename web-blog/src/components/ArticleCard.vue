<template>
  <div class="article-card" @click="navigateToArticle">
    <div class="card-cover">
      <img v-if="article.cover" :src="article.cover" :alt="article.title" loading="lazy" />
      <div v-else class="cover-placeholder">
        <span>{{ article.title?.charAt(0) }}</span>
      </div>
      <div class="cover-overlay"></div>
      <div class="card-category" v-if="article.category">
        <span>{{ article.category.name }}</span>
      </div>
    </div>

    <div class="card-content">
      <h3 class="card-title">{{ article.title }}</h3>
      <p class="card-desc" v-if="article.summary">{{ article.summary }}</p>

      <div class="card-tags" v-if="article.tags?.length">
        <n-tag
          v-for="tag in article.tags.slice(0, 3)"
          :key="tag.id"
          size="small"
          :bordered="false"
          class="tag-item"
        >
          {{ tag.name }}
        </n-tag>
      </div>

      <div class="card-meta">
        <span class="meta-item">
          <n-icon><TimeOutline /></n-icon>
          {{ formatDate(article.created_at) }}
        </span>
        <span class="meta-item" v-if="article.views">
          <n-icon><EyeOutline /></n-icon>
          {{ article.views }}
        </span>
        <span class="meta-item" v-if="article.comments_count">
          <n-icon><ChatbubbleOutline /></n-icon>
          {{ article.comments_count }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { NTag, NIcon } from 'naive-ui'
import { TimeOutline, EyeOutline, ChatbubbleOutline } from '@vicons/ionicons5'

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

const props = defineProps<{
  article: Article
}>()

const router = useRouter()

function navigateToArticle() {
  router.push(`/article/${props.article.id}`)
}

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}
</script>

<style scoped lang="scss">
.article-card {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  overflow: hidden;
  cursor: pointer;
  transition: transform var(--transition-normal), box-shadow var(--transition-normal);
  box-shadow: var(--shadow-sm);

  &:hover {
    transform: translateY(-4px);
    box-shadow: var(--shadow-lg);

    .cover-overlay {
      opacity: 0.3;
    }

    .card-title {
      color: var(--primary-color);
    }
  }
}

.card-cover {
  position: relative;
  height: 180px;
  overflow: hidden;
  background: linear-gradient(135deg, #1a1a1a 0%, #333333 100%);

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform var(--transition-slow);
  }

  .cover-placeholder {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(135deg, #1a1a1a, #2a2a2a);
    font-size: 48px;
    font-weight: 700;
    color: rgba(255, 255, 255, 0.8);
  }

  .cover-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(to top, rgba(0, 0, 0, 0.5), transparent);
    opacity: 0;
    transition: opacity var(--transition-normal);
  }
}

.card-category {
  position: absolute;
  top: 12px;
  left: 12px;

  span {
    background: rgba(255, 255, 255, 0.9);
    color: #1a1a1a;
    padding: 4px 12px;
    border-radius: 20px;
    font-size: 12px;
    font-weight: 500;
  }
}

.card-content {
  padding: 16px 20px 20px;
}

.card-title {
  font-size: 17px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 10px;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  transition: color var(--transition-fast);
}

.card-desc {
  font-size: 14px;
  color: var(--text-muted);
  line-height: 1.6;
  margin-bottom: 12px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 14px;

  .tag-item {
    background: rgba(0, 0, 0, 0.05);
    color: var(--text-secondary);
    font-size: 12px;
  }
}

.card-meta {
  display: flex;
  gap: 16px;
  font-size: 13px;
  color: var(--text-muted);
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;

  .n-icon {
    font-size: 14px;
  }
}
</style>

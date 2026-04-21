<template>
  <div class="article-card" @click="$router.push(`/article/${article.id}`)">
    <div class="card-cover" v-if="article.cover">
      <img :src="article.cover" :alt="article.title" />
    </div>
    <div class="card-body">
      <div class="card-tags" v-if="article.tags?.length">
        <el-tag v-for="tag in article.tags.slice(0, 3)" :key="tag.id" size="small" :color="tag.color" effect="dark" round>
          {{ tag.name }}
        </el-tag>
        <el-tag v-if="article.category" size="small" type="info" round>{{ article.category.name }}</el-tag>
      </div>
      <h3 class="card-title">{{ article.title }}</h3>
      <p class="card-desc">{{ article.description || stripHtml(article.content) }}</p>
      <div class="card-meta">
        <span class="meta-item">
          <el-icon><View /></el-icon>
          {{ article.view_count || 0 }}
        </span>
        <span class="meta-item">
          <el-icon><ChatDotRound /></el-icon>
          {{ article.comment_count || 0 }}
        </span>
        <span class="meta-date">{{ formatDate(article.created_at) }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({ article: { type: Object, required: true } })

function stripHtml(html) {
  if (!html) return ''
  return html.replace(/<[^>]*>/g, '').slice(0, 120) + '...'
}

function formatDate(date) {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-CN')
}
</script>

<style lang="scss" scoped>
.article-card {
  background: var(--bg-card);
  border-radius: var(--radius);
  overflow: hidden;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: var(--shadow-sm);

  &:hover {
    box-shadow: var(--shadow-md);
    transform: translateY(-2px);
  }
}

.card-cover {
  height: 180px;
  overflow: hidden;

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.3s;
  }

  .article-card:hover & img {
    transform: scale(1.05);
  }
}

.card-body {
  padding: 16px;
}

.card-tags {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
  margin-bottom: 8px;
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  line-height: 1.5;
  margin-bottom: 8px;
  color: var(--text-primary);
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-desc {
  font-size: 13px;
  color: var(--text-muted);
  line-height: 1.6;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  margin-bottom: 12px;
}

.card-meta {
  display: flex;
  align-items: center;
  gap: 16px;
  font-size: 12px;
  color: var(--text-muted);
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.meta-date {
  margin-left: auto;
}
</style>
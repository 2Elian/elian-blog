<template>
  <div class="learn-page container">
    <div class="page-header">
      <h1>学习路线</h1>
      <p>精心整理的学习路线，助你快速成长</p>
    </div>

    <div class="learn-grid">
      <div class="learn-card" v-for="item in learnItems" :key="item.id">
        <div class="card-icon" :style="{ background: item.color }">
          <el-icon size="32">{{ item.icon }}</el-icon>
        </div>
        <div class="card-content">
          <h3>{{ item.title }}</h3>
          <p>{{ item.desc }}</p>
          <div class="card-footer">
            <span class="articles">{{ item.articles }} 篇文章</span>
            <el-button type="primary" size="small" text @click="viewArticles(item)">查看路线</el-button>
          </div>
        </div>
      </div>
    </div>

    <section class="featured-section" v-if="featuredArticles.length">
      <h2>精选文章</h2>
      <div class="article-grid">
        <ArticleCard v-for="article in featuredArticles" :key="article.id" :article="article" />
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import ArticleCard from '@/components/ArticleCard.vue'
import { getArticles } from '@/api'

const learnItems = ref([
  { id: 1, title: '前端开发', desc: 'HTML/CSS/JS → Vue/React → Node.js', icon: 'Monitor', color: 'linear-gradient(135deg, #409eff, #79bbff)', articles: 15 },
  { id: 2, title: '后端开发', desc: 'Java/Go/Python → Spring/Gin/Django', icon: 'DataAnalysis', color: 'linear-gradient(135deg, #67c23a, #95d475)', articles: 12 },
  { id: 3, title: '数据库', desc: 'MySQL → Redis → MongoDB → Elasticsearch', icon: 'Coin', color: 'linear-gradient(135deg, #e6a23c, #eebe77)', articles: 8 },
  { id: 4, title: '运维部署', desc: 'Linux → Docker → Kubernetes → CI/CD', icon: 'SetUp', color: 'linear-gradient(135deg, #f56c6c, #fab6b6)', articles: 6 },
  { id: 5, title: '系统设计', desc: '架构模式 → 分布式系统 → 高并发', icon: 'Share', color: 'linear-gradient(135deg, #909399, #c0c4cc)', articles: 10 },
  { id: 6, title: 'AI & ML', desc: '机器学习 → 深度学习 → NLP/CV', icon: 'MagicStick', color: 'linear-gradient(135deg, #9b59b6, #b97dd8)', articles: 5 },
])

const featuredArticles = ref([])

onMounted(async () => {
  try {
    const data = await getArticles({ page: 1, page_size: 6, is_featured: true })
    featuredArticles.value = data?.list || []
  } catch {}
})

function viewArticles(item) {
  window.location.href = `/blog?keyword=${item.title}`
}
</script>

<style lang="scss" scoped>
.learn-page { padding-bottom: 40px; }

.page-header {
  text-align: center;
  padding: 40px 0;
  h1 { font-size: 32px; margin-bottom: 12px; }
  p { font-size: 16px; color: var(--text-muted); }
}

.learn-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  margin-bottom: 40px;

  @media (max-width: 900px) { grid-template-columns: repeat(2, 1fr); }
  @media (max-width: 600px) { grid-template-columns: 1fr; }
}

.learn-card {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 24px;
  display: flex;
  gap: 16px;
  box-shadow: var(--shadow-sm);
  transition: all 0.3s;

  &:hover { box-shadow: var(--shadow-md); transform: translateY(-2px); }
}

.card-icon {
  width: 64px;
  height: 64px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  flex-shrink: 0;
}

.card-content {
  flex: 1;
  h3 { font-size: 18px; margin-bottom: 8px; }
  p { font-size: 14px; color: var(--text-muted); margin-bottom: 12px; }
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  .articles { font-size: 13px; color: var(--text-secondary); }
}

.featured-section {
  h2 { font-size: 20px; margin-bottom: 20px; }
}

.article-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;

  @media (max-width: 900px) { grid-template-columns: repeat(2, 1fr); }
  @media (max-width: 600px) { grid-template-columns: 1fr; }
}
</style>
<template>
  <div class="products-page">
    <div class="page-header">
      <h1 class="page-title gradient-text">产品展示</h1>
      <p class="page-desc">我的作品与项目</p>
    </div>

    <div class="showcase-list">
      <div
        v-for="(product, idx) in products"
        :key="product.id"
        class="showcase-card"
        :class="{ 'showcase-reverse': idx % 2 === 1 }"
        @click="handleClick(product)"
      >
        <div v-if="product.cover" class="showcase-image">
          <img :src="getImageUrl(product.cover)" :alt="product.name" />
        </div>
        <div v-else class="showcase-visual">
          <div class="showcase-number">{{ String(idx + 1).padStart(2, '0') }}</div>
        </div>
        <div class="showcase-info">
          <div class="showcase-meta">
            <span class="showcase-badge" :class="getTypeClass(product.type)">{{ getTypeName(product.type) }}</span>
            <span v-if="product.price > 0" class="showcase-price">{{ product.price }} 元</span>
          </div>
          <h3 class="showcase-title">{{ product.name }}</h3>
          <p class="showcase-desc" v-if="product.description">{{ product.description }}</p>
          <div class="showcase-footer">
            <span class="showcase-link-text">了解更多 →</span>
          </div>
        </div>
      </div>
    </div>

    <n-empty v-if="!loading && products.length === 0" description="暂无产品" style="padding: 60px 0;" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { NEmpty } from 'naive-ui'
import { getProducts } from '@/api'

interface Product {
  id: number
  name: string
  description?: string
  price: number
  cover?: string
  status: number
  sort: number
  type: string
  created_at: string
}

const products = ref<Product[]>([])
const loading = ref(true)
const router = useRouter()

function handleClick(product: Product) {
  router.push(`/product/${product.id}`)
}

function getImageUrl(path: string): string {
  if (!path) return ''
  if (path.startsWith('http')) return path
  return `${window.location.origin}${path.startsWith('/') ? '' : '/'}${path}`
}

function getTypeName(type: string): string {
  return type || '其他'
}

function getTypeClass(type: string): string {
  const map: Record<string, string> = { 'AI产品': 'badge-ai', '工具': 'badge-tool' }
  return map[type] || 'badge-other'
}

async function loadProducts() {
  try {
    const res = await getProducts() as any
    products.value = res.data?.list || []
  } catch (e) {
    console.error('Failed to load products:', e)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadProducts()
})
</script>

<style scoped lang="scss">
.products-page {
  animation: fadeInUp 0.5s ease;
  max-width: 900px;
  margin: 0 auto;
  padding: 0 32px;

  @media (max-width: 640px) {
    padding: 0 20px;
  }
}

.page-header {
  text-align: center;
  margin-bottom: 40px;
}

.page-title {
  font-size: 32px;
  font-weight: 700;
  margin-bottom: 10px;
}

.page-desc {
  color: var(--text-muted);
  font-size: 16px;
}

// ===== Showcase Cards (same as About page) =====
.showcase-list {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.showcase-card {
  display: grid;
  grid-template-columns: 180px 1fr;
  gap: 0;
  border: 1px solid var(--border-color);
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  background: var(--bg-card);
  cursor: pointer;

  &:hover {
    border-color: var(--primary-color);
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
    transform: translateY(-4px);

    .showcase-image img {
      transform: scale(1.05);
    }
  }

  &.showcase-reverse {
    grid-template-columns: 1fr 180px;

    .showcase-visual { order: 2; }
    .showcase-image { order: 2; }
    .showcase-info { order: 1; }
  }

  @media (max-width: 640px) {
    grid-template-columns: 1fr !important;

    .showcase-visual,
    .showcase-image { order: 0 !important; min-height: 120px; }
    .showcase-info { order: 0 !important; }
  }
}

.showcase-image {
  overflow: hidden;

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.5s ease;
  }
}

.showcase-visual {
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, var(--primary-color), #7c3aed);
  min-height: 160px;
}

.showcase-number {
  font-size: 48px;
  font-weight: 900;
  color: rgba(255, 255, 255, 0.2);
  letter-spacing: -2px;
}

.showcase-info {
  padding: 24px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.showcase-meta {
  margin-bottom: 8px;
  display: flex;
  gap: 8px;
  align-items: center;
}

.showcase-badge {
  display: inline-block;
  padding: 3px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;

  &.badge-ai {
    background: rgba(102, 126, 234, 0.1);
    border: 1px solid rgba(102, 126, 234, 0.15);
    color: #667eea;
  }

  &.badge-tool {
    background: rgba(72, 187, 120, 0.1);
    border: 1px solid rgba(72, 187, 120, 0.15);
    color: #48bb78;
  }

  &.badge-other {
    background: rgba(237, 137, 54, 0.08);
    border: 1px solid rgba(237, 137, 54, 0.15);
    color: #d97706;
  }
}

.showcase-price {
  font-size: 14px;
  font-weight: 600;
  color: var(--primary-color);
}

.showcase-title {
  font-size: 16px;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 10px;
  line-height: 1.4;
}

.showcase-desc {
  font-size: 14px;
  line-height: 1.7;
  color: var(--text-secondary);
  margin-bottom: 14px;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.showcase-footer {
  display: flex;
  align-items: center;
}

.showcase-link-text {
  font-size: 13px;
  color: var(--primary-color);
  font-weight: 500;
}
</style>

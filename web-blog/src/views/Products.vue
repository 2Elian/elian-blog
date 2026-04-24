<template>
  <div class="products-page">
    <div class="page-header">
      <h1 class="page-title gradient-text">产品展示</h1>
      <p class="page-desc">我的作品与项目</p>
    </div>

    <div class="products-grid">
      <div
        v-for="product in products"
        :key="product.id"
        class="product-card"
        @click="router.push(`/product/${product.id}`)"
      >
        <div class="card-cover">
          <img v-if="product.cover" :src="getImageUrl(product.cover)" :alt="product.name" />
          <div v-else class="cover-placeholder">
            <n-icon size="48"><CubeOutline /></n-icon>
          </div>
        </div>
        <div class="card-content">
          <div class="product-type">
            <span :class="getTypeClass(product.type)">{{ getTypeName(product.type) }}</span>
          </div>
          <h3 class="product-name">{{ product.name }}</h3>
          <p class="product-desc" v-if="product.description">{{ product.description }}</p>
          <div class="product-footer">
            <span class="product-price" v-if="product.price > 0">{{ product.price }} 元</span>
            <span class="product-price free" v-else>免费</span>
            <n-icon v-if="product.link" size="16"><OpenOutline /></n-icon>
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
import { NIcon, NEmpty } from 'naive-ui'
import { CubeOutline, OpenOutline } from '@vicons/ionicons5'
import { getProducts } from '@/api'

interface Product {
  id: number
  name: string
  description?: string
  price: number
  cover?: string
  status: number
  sort: number
  type: number
  link?: string
  created_at: string
}

const products = ref<Product[]>([])
const loading = ref(true)
const router = useRouter()

function getImageUrl(path: string): string {
  if (!path) return ''
  if (path.startsWith('http')) return path
  return `http://localhost:8080${path.startsWith('/') ? '' : '/'}${path}`
}

function getTypeName(type: number): string {
  const types: Record<number, string> = {
    1: 'AI产品',
    2: '工具',
    3: '其他'
  }
  return types[type] || '其他'
}

function getTypeClass(type: number): string {
  const classes: Record<number, string> = {
    1: 'type-ai',
    2: 'type-tool',
    3: 'type-other'
  }
  return classes[type] || 'type-other'
}

function openLink(url: string) {
  window.open(url, '_blank', 'noopener,noreferrer')
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
  max-width: 1100px;
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

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 24px;
}

.product-card {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  overflow: hidden;
  text-decoration: none;
  box-shadow: var(--shadow-sm);
  transition: all var(--transition-normal);
  display: flex;
  flex-direction: column;

  &:hover {
    transform: translateY(-6px);
    box-shadow: var(--shadow-lg);

    .product-name {
      color: var(--primary-color);
    }

    .card-cover img {
      transform: scale(1.05);
    }
  }
}

.card-cover {
  height: 180px;
  background: linear-gradient(135deg, #1a1a1a, #2a2a2a);
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform var(--transition-slow);
  }

  .cover-placeholder {
    color: rgba(255, 255, 255, 0.3);
  }
}

.card-content {
  padding: 20px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.product-type {
  margin-bottom: 12px;

  span {
    display: inline-block;
    padding: 4px 10px;
    border-radius: 4px;
    font-size: 12px;
    font-weight: 500;
  }

  .type-ai {
    background: rgba(102, 126, 234, 0.15);
    color: #667eea;
  }

  .type-tool {
    background: rgba(72, 187, 120, 0.15);
    color: #48bb78;
  }

  .type-other {
    background: rgba(237, 137, 54, 0.15);
    color: #ed8936;
  }
}

.product-name {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 8px;
  transition: color var(--transition-fast);
}

.product-desc {
  font-size: 14px;
  color: var(--text-muted);
  line-height: 1.6;
  margin-bottom: 16px;
  flex: 1;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.product-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.product-price {
  font-size: 16px;
  font-weight: 600;
  color: var(--primary-color);

  &.free {
    color: var(--text-secondary);
    font-weight: 500;
  }
}
</style>

<template>
  <div class="product-detail-page" v-if="product">
    <div class="detail-header">
      <div class="header-cover" v-if="product.cover">
        <img :src="product.cover" :alt="product.name" />
      </div>
      <div class="header-overlay">
        <div class="header-content">
          <button class="back-btn" @click="router.back()">&larr; 返回产品列表</button>
          <div class="product-type">
            <span :class="getTypeClass(product.type)">{{ getTypeName(product.type) }}</span>
          </div>
          <h1 class="product-name">{{ product.name }}</h1>
          <div class="product-meta">
            <span class="product-price" v-if="product.price > 0">{{ product.price }} 元</span>
            <span class="product-price free" v-else>免费</span>
            <span class="product-date">{{ product.created_at }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="detail-body">
      <div class="body-card">
        <div class="description-content" v-if="product.description">
          <h2 class="desc-title">产品介绍</h2>
          <div class="desc-text" v-html="renderedDescription"></div>
        </div>
        <n-empty v-else description="暂无详细介绍" style="padding: 40px 0;" />

        <div class="action-bar" v-if="product.link">
          <a :href="product.link" target="_blank" rel="noopener noreferrer" class="action-btn">
            访问链接
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M18 13v6a2 2 0 01-2 2H5a2 2 0 01-2-2V8a2 2 0 012-2h6"/><polyline points="15 3 21 3 21 9"/><line x1="10" y1="14" x2="21" y2="3"/></svg>
          </a>
        </div>
      </div>
    </div>
  </div>

  <div class="loading-state" v-else>
    <n-spin size="large" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { NEmpty, NSpin } from 'naive-ui'
import { marked } from 'marked'
import { getProduct } from '@/api'

interface Product {
  id: number
  name: string
  description: string
  price: number
  cover: string
  status: number
  sort: number
  type: number
  link: string
  created_at: string
}

const route = useRoute()
const router = useRouter()
const product = ref<Product | null>(null)

const renderedDescription = computed(() => {
  if (!product.value?.description) return ''
  return marked(product.value.description)
})

function getTypeName(type: number): string {
  const types: Record<number, string> = { 1: 'AI产品', 2: '工具', 3: '其他' }
  return types[type] || '其他'
}

function getTypeClass(type: number): string {
  const classes: Record<number, string> = { 1: 'type-ai', 2: 'type-tool', 3: 'type-other' }
  return classes[type] || 'type-other'
}

async function loadProduct(id: number) {
  try {
    const res = await getProduct(id) as any
    product.value = res.data
  } catch (e) {
    console.error('Failed to load product:', e)
  }
}

onMounted(() => {
  const id = Number(route.params.id)
  if (id) loadProduct(id)
})

watch(() => route.params.id, (newId) => {
  if (newId) {
    loadProduct(Number(newId))
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
})
</script>

<style scoped lang="scss">
.product-detail-page {
  animation: fadeInUp 0.5s ease;
  max-width: 900px;
  margin: 0 auto;
  padding: 0 32px;

  @media (max-width: 640px) {
    padding: 0 20px;
  }
}

.detail-header {
  position: relative;
  border-radius: var(--radius-lg);
  overflow: hidden;
  margin-bottom: 30px;
  min-height: 300px;
  background: linear-gradient(135deg, #0a0a0a 0%, #2a2a2a 100%);

  .header-cover img {
    width: 100%;
    height: 300px;
    object-fit: cover;
    display: block;
  }

  .header-overlay {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    padding: 40px;
    background: linear-gradient(transparent, rgba(0, 0, 0, 0.85));
    color: white;

    @media (max-width: 640px) {
      padding: 20px;
    }
  }
}

.back-btn {
  background: none;
  border: none;
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
  cursor: pointer;
  margin-bottom: 16px;
  padding: 0;

  &:hover {
    color: white;
  }
}

.product-type {
  margin-bottom: 12px;

  span {
    display: inline-block;
    padding: 4px 12px;
    border-radius: 4px;
    font-size: 13px;
    font-weight: 500;
  }

  .type-ai {
    background: rgba(102, 126, 234, 0.3);
    color: #a0b0ff;
  }

  .type-tool {
    background: rgba(72, 187, 120, 0.3);
    color: #8eeca8;
  }

  .type-other {
    background: rgba(237, 137, 54, 0.3);
    color: #f0b060;
  }
}

.product-name {
  font-size: 32px;
  font-weight: 700;
  margin-bottom: 12px;

  @media (max-width: 640px) {
    font-size: 24px;
  }
}

.product-meta {
  display: flex;
  align-items: center;
  gap: 16px;
}

.product-price {
  font-size: 20px;
  font-weight: 700;

  &.free {
    font-size: 16px;
    font-weight: 500;
    opacity: 0.7;
  }
}

.product-date {
  font-size: 14px;
  opacity: 0.6;
}

.detail-body {
  max-width: 860px;
  margin: 0 auto;
}

.body-card {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 40px;
  box-shadow: var(--shadow-sm);

  @media (max-width: 640px) {
    padding: 24px 20px;
  }
}

.desc-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 20px;
  padding-bottom: 14px;
  border-bottom: 1px solid var(--border-color);
}

.desc-text {
  font-size: 15px;
  line-height: 1.8;
  color: var(--text-secondary);

  :deep(h1), :deep(h2), :deep(h3) {
    color: var(--text-primary);
    margin: 24px 0 12px;
  }

  :deep(p) {
    margin-bottom: 12px;
  }

  :deep(code) {
    background: rgba(0, 0, 0, 0.06);
    padding: 2px 6px;
    border-radius: 4px;
    font-size: 0.9em;
  }

  :deep(pre) {
    background: rgba(0, 0, 0, 0.06);
    padding: 16px;
    border-radius: 8px;
    overflow-x: auto;
  }

  :deep(ul), :deep(ol) {
    padding-left: 24px;
    margin-bottom: 12px;
  }
}

.action-bar {
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid var(--border-color);
}

.action-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background: var(--text-primary);
  color: var(--bg-card);
  border-radius: 8px;
  font-size: 15px;
  font-weight: 500;
  text-decoration: none;
  transition: opacity var(--transition-fast);

  &:hover {
    opacity: 0.85;
  }
}

.loading-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
}
</style>

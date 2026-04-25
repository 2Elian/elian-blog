<template>
  <div class="product-detail-page" v-if="product">
    <div class="detail-header">
      <div class="header-cover" v-if="product.cover">
        <img :src="getImageUrl(product.cover)" :alt="product.name" />
      </div>
      <div class="header-overlay">
        <div class="header-content">
          <button class="back-btn" @click="router.back()">&larr; 返回产品列表</button>
          <div class="product-type">
            <span class="type-badge">{{ product.type || '其他' }}</span>
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
        <div class="content-section" v-if="product.content" v-html="renderedContent"></div>
        <div class="content-section" v-else-if="product.description">
          <p>{{ product.description }}</p>
        </div>
        <n-empty v-else description="暂无详细介绍" style="padding: 40px 0;" />
      </div>
    </div>
  </div>

  <div class="loading-state" v-else>
    <n-spin size="large" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { NEmpty, NSpin } from 'naive-ui'
import mermaid from 'mermaid'
import { renderMarkdown } from '@/utils/markdown'
import { getProduct } from '@/api'

mermaid.initialize({ startOnLoad: false, theme: 'default' })

interface Product {
  id: number
  name: string
  description: string
  content: string
  price: number
  cover: string
  status: number
  sort: number
  type: string
  created_at: string
}

const route = useRoute()
const router = useRouter()
const product = ref<Product | null>(null)

const renderedContent = computed(() => {
  if (!product.value?.content) return ''
  return renderMarkdown(product.value.content)
})

function getImageUrl(path: string): string {
  if (!path) return ''
  if (path.startsWith('http')) return path
  return `http://localhost:8080${path.startsWith('/') ? '' : '/'}${path}`
}

async function loadProduct(id: number) {
  try {
    const res = await getProduct(id) as any
    product.value = res.data
    await nextTick()
    mermaid.run()
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
  max-width: 1100px;
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

  .type-badge {
    display: inline-block;
    padding: 4px 12px;
    border-radius: 4px;
    font-size: 13px;
    font-weight: 500;
    background: rgba(255, 255, 255, 0.15);
    color: rgba(255, 255, 255, 0.9);
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
  max-width: 1060px;
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

.content-section {
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

  :deep(img) {
    max-width: 100%;
    border-radius: 8px;
  }
}

.loading-state {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 300px;
}
</style>
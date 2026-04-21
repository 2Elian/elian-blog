<template>
  <div class="page-view container">
    <LoadingState :loading="loading" :error="error" emptyText="页面不存在" @retry="loadPage">
      <article class="page-content">
        <h1>{{ page.title }}</h1>
        <div class="content-body" v-html="page.content"></div>
      </article>
    </LoadingState>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import LoadingState from '@/components/LoadingState.vue'
import { getPageBySlug } from '@/api'

const route = useRoute()
const loading = ref(true)
const error = ref('')
const page = ref({})

onMounted(() => loadPage())
watch(() => route.params.slug, () => loadPage())

async function loadPage() {
  loading.value = true
  error.value = ''
  try {
    page.value = await getPageBySlug(route.params.slug)
  } catch {
    error.value = '加载失败'
  } finally {
    loading.value = false
  }
}
</script>

<style lang="scss" scoped>
.page-view { max-width: 800px; min-height: calc(100vh - 140px); }

.page-content {
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 32px;
  box-shadow: var(--shadow-sm);

  h1 {
    font-size: 28px;
    font-weight: 600;
    margin-bottom: 24px;
    padding-bottom: 16px;
    border-bottom: 1px solid var(--border-color);
  }
}

.content-body {
  font-size: 16px;
  line-height: 1.8;
  color: var(--text-primary);

  h2, h3 { margin: 24px 0 12px; font-weight: 600; }
  p { margin-bottom: 16px; }
  img { max-width: 100%; border-radius: var(--radius); }
}
</style>
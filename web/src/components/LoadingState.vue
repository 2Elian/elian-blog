<template>
  <el-skeleton :rows="rows" animated :loading="loading">
    <template v-if="error">
      <div class="state-box">
        <el-empty :description="error" />
        <el-button type="primary" @click="$emit('retry')" text>重试</el-button>
      </div>
    </template>
    <template v-else-if="!loading && empty">
      <el-empty :description="emptyText" />
    </template>
    <template v-else>
      <slot />
    </template>
  </el-skeleton>
</template>

<script setup>
defineProps({
  loading: Boolean,
  error: { type: String, default: '' },
  empty: Boolean,
  emptyText: { type: String, default: '暂无数据' },
  rows: { type: Number, default: 5 }
})
defineEmits(['retry'])
</script>

<style scoped>
.state-box {
  text-align: center;
  padding: 40px 0;
}
</style>
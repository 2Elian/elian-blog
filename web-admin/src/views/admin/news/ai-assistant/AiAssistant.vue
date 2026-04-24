<template>
  <div class="ai-assistant-container">
    <el-card>
      <template #header>
        <div class="flex items-center justify-between">
          <span class="text-lg font-bold">AI 助手</span>
          <el-tag type="info">即将推出</el-tag>
        </div>
      </template>
      <el-empty description="AI 助手功能正在开发中，敬请期待...">
        <template #image>
          <el-icon :size="80" color="#909399">
            <svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
              <path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z" fill="currentColor"/>
              <path d="M464 336a48 48 0 1 0 96 0 48 48 0 1 0-96 0zm72 112h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V456c0-4.4-3.6-8-8-8z" fill="currentColor"/>
            </svg>
          </el-icon>
        </template>
      </el-empty>
    </el-card>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.ai-assistant-container {
  padding: 20px;
}
</style>

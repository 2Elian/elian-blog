<template>
  <div class="state-page">
    <div class="state-header">
      <h2 class="page-title">服务器监控</h2>
      <p class="page-desc">实时系统状态 · 每10秒自动刷新</p>
    </div>

    <!-- Runtime Info -->
    <div class="info-grid" v-if="state.os">
      <div class="info-card runtime">
        <div class="card-icon">
          <el-icon :size="28"><Monitor /></el-icon>
        </div>
        <div class="card-body">
          <h3 class="card-title">运行环境</h3>
          <div class="info-rows">
            <div class="info-row">
              <span class="label">操作系统</span>
              <span class="value">{{ state.os.goos }}</span>
            </div>
            <div class="info-row">
              <span class="label">Go 版本</span>
              <span class="value">{{ state.os.goVersion }}</span>
            </div>
            <div class="info-row">
              <span class="label">CPU 核心数</span>
              <span class="value highlight">{{ state.os.numCpu }}</span>
            </div>
            <div class="info-row">
              <span class="label">协程数量</span>
              <span class="value highlight">{{ state.os.numGoroutine }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- CPU -->
      <div class="info-card cpu" v-if="state.cpu">
        <div class="card-icon">
          <el-icon :size="28"><Cpu /></el-icon>
        </div>
        <div class="card-body">
          <h3 class="card-title">CPU 使用率</h3>
          <div class="cpu-summary">
            <div class="gauge-ring" :style="{ '--pct': cpuPercent }">
              <span class="gauge-value">{{ cpuPercent }}%</span>
            </div>
            <div class="cpu-cores" v-if="state.cpu.cpus">
              <div v-for="(item, index) in state.cpu.cpus.slice(0, 8)" :key="index" class="core-bar">
                <span class="core-label">Core {{ index }}</span>
                <div class="core-track">
                  <div class="core-fill" :style="{ width: item.toFixed(0) + '%', background: getBarColor(item) }"></div>
                </div>
                <span class="core-pct">{{ item.toFixed(0) }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Disk & Memory -->
    <div class="metric-grid">
      <div class="metric-card" v-if="state.disk">
        <div class="metric-header">
          <div class="metric-title">
            <el-icon :size="22"><Coin /></el-icon>
            <span>磁盘</span>
          </div>
          <span class="metric-badge" :style="{ background: getBarColor(state.disk.usedPercent) }">
            {{ state.disk.usedPercent?.toFixed(1) || 0 }}%
          </span>
        </div>
        <el-progress
          :percentage="state.disk.usedPercent || 0"
          :stroke-width="10"
          :color="progressColors"
          :show-text="false"
        />
        <div class="metric-details">
          <div class="metric-row">
            <span>已使用</span>
            <span>{{ state.disk.usedGb || 0 }} GB</span>
          </div>
          <div class="metric-row">
            <span>总容量</span>
            <span>{{ state.disk.totalGb || 0 }} GB</span>
          </div>
        </div>
      </div>

      <div class="metric-card" v-if="state.ram">
        <div class="metric-header">
          <div class="metric-title">
            <el-icon :size="22"><Memo /></el-icon>
            <span>内存</span>
          </div>
          <span class="metric-badge" :style="{ background: getBarColor(state.ram.usedPercent) }">
            {{ state.ram.usedPercent?.toFixed(1) || 0 }}%
          </span>
        </div>
        <el-progress
          :percentage="state.ram.usedPercent || 0"
          :stroke-width="10"
          :color="progressColors"
          :show-text="false"
        />
        <div class="metric-details">
          <div class="metric-row">
            <span>已使用</span>
            <span>{{ (state.ram.usedMb / 1024)?.toFixed(2) || 0 }} GB</span>
          </div>
          <div class="metric-row">
            <span>总容量</span>
            <span>{{ (state.ram.totalMb / 1024)?.toFixed(2) || 0 }} GB</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onUnmounted } from 'vue'
import { Monitor, Cpu, Coin, Memo } from '@element-plus/icons-vue'
import { WebsiteAPI } from '@/api/website'

const state = ref<any>({})
const timer = ref<any>(null)

const progressColors = [
  { color: '#67c23a', percentage: 40 },
  { color: '#e6a23c', percentage: 70 },
  { color: '#f56c6c', percentage: 100 },
]

const cpuPercent = computed(() => {
  if (!state.value.cpu?.cpus?.length) return 0
  const avg = state.value.cpu.cpus.reduce((a: number, b: number) => a + b, 0) / state.value.cpu.cpus.length
  return avg.toFixed(0)
})

function getBarColor(pct: number) {
  if (pct < 40) return '#67c23a'
  if (pct < 70) return '#e6a23c'
  return '#f56c6c'
}

const reload = async () => {
  try {
    const res = await WebsiteAPI.getSystemStateApi()
    state.value = res.data || {}
  } catch (e) {
    console.error('Failed to load system state:', e)
  }
}

reload()
timer.value = setInterval(reload, 10000)

onUnmounted(() => {
  clearInterval(timer.value)
})
</script>

<style scoped>
.state-page {
  padding: 20px;
  max-width: 1200px;
}

.state-header {
  margin-bottom: 24px;
}

.page-title {
  font-size: 22px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  margin: 0 0 6px;
}

.page-desc {
  font-size: 14px;
  color: var(--el-text-color-secondary);
}

.info-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 16px;
}

@media (max-width: 768px) {
  .info-grid {
    grid-template-columns: 1fr;
  }
}

.info-card {
  background: var(--el-bg-color);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08);
  display: flex;
  gap: 20px;
}

.card-icon {
  width: 56px;
  height: 56px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
}

.cpu .card-icon {
  background: linear-gradient(135deg, #f093fb, #f5576c);
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  margin: 0 0 14px;
}

.info-rows {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.info-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 14px;
}

.info-row .label {
  color: var(--el-text-color-secondary);
}

.info-row .value {
  color: var(--el-text-color-primary);
  font-weight: 500;
}

.info-row .value.highlight {
  color: #667eea;
  font-weight: 600;
}

.cpu-summary {
  display: flex;
  gap: 20px;
  align-items: flex-start;
}

.gauge-ring {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: conic-gradient(#667eea calc(var(--pct) * 1%), #e8e8e8 0);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.gauge-ring::before {
  content: '';
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background: var(--el-bg-color);
  position: absolute;
}

.gauge-value {
  position: relative;
  z-index: 1;
  font-size: 16px;
  font-weight: 700;
  color: #667eea;
}

.cpu-cores {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.core-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
}

.core-label {
  width: 48px;
  color: var(--el-text-color-secondary);
  flex-shrink: 0;
}

.core-track {
  flex: 1;
  height: 6px;
  background: #f0f0f0;
  border-radius: 3px;
  overflow: hidden;
}

.core-fill {
  height: 100%;
  border-radius: 3px;
  transition: width 0.5s ease;
}

.core-pct {
  width: 36px;
  text-align: right;
  color: var(--el-text-color-secondary);
}

.metric-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

@media (max-width: 768px) {
  .metric-grid {
    grid-template-columns: 1fr;
  }
}

.metric-card {
  background: var(--el-bg-color);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08);
}

.metric-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.metric-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 16px;
  font-weight: 600;
  color: var(--el-text-color-primary);
}

.metric-badge {
  padding: 4px 12px;
  border-radius: 12px;
  color: white;
  font-size: 13px;
  font-weight: 600;
}

.metric-details {
  margin-top: 16px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.metric-row {
  display: flex;
  justify-content: space-between;
  font-size: 14px;
  color: var(--el-text-color-secondary);
}
</style>

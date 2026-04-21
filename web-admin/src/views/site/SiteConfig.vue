<template>
  <div class="site-config">
    <div class="page-header">
      <h2>站点设置</h2>
      <el-button type="primary" :loading="saveLoading" @click="handleSave">
        <el-icon><Check /></el-icon> 保存设置
      </el-button>
    </div>

    <el-card v-loading="loading">
      <el-form label-width="160px" label-position="right">
        <el-form-item v-for="(value, key) in configForm" :key="key" :label="getConfigLabel(key as string)">
          <el-input
            v-if="isLongText(value)"
            v-model="configForm[key]"
            type="textarea"
            :rows="3"
            :placeholder="`请输入${getConfigLabel(key as string)}`"
          />
          <el-input
            v-else
            v-model="configForm[key]"
            :placeholder="`请输入${getConfigLabel(key as string)}`"
          />
        </el-form-item>

        <!-- 添加新配置项 -->
        <el-divider />
        <div style="display: flex; gap: 12px; align-items: center; margin-bottom: 16px;">
          <el-input v-model="newKey" placeholder="配置项键名（英文）" style="width: 200px" />
          <el-input v-model="newValue" placeholder="配置项值" style="width: 200px" />
          <el-button type="success" @click="handleAddConfig">
            <el-icon><Plus /></el-icon> 添加配置
          </el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Check, Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { getSiteConfig, setSiteConfig } from '@/api'

const loading = ref(false)
const saveLoading = ref(false)
const newKey = ref('')
const newValue = ref('')

const configForm = reactive<Record<string, string>>({})

// 常见配置项标签映射
const labelMap: Record<string, string> = {
  site_name: '站点名称',
  site_description: '站点描述',
  site_keywords: '站点关键词',
  site_url: '站点URL',
  site_logo: '站点Logo',
  site_favicon: 'Favicon',
  site_footer: '页脚信息',
  site_icp: 'ICP备案号',
  site_copyright: '版权信息',
  admin_email: '管理员邮箱',
  comment_moderation: '评论审核',
  posts_per_page: '每页文章数',
  social_github: 'GitHub',
  social_twitter: 'Twitter',
  social_email: '邮箱',
  social_wechat: '微信',
}

function getConfigLabel(key: string): string {
  return labelMap[key] || key
}

function isLongText(value: string): boolean {
  return value && value.length > 60
}

async function fetchConfig() {
  loading.value = true
  try {
    const res = await getSiteConfig()
    const data = res.data
    // 清空表单重新填充
    Object.keys(configForm).forEach((key) => delete configForm[key])
    if (data && typeof data === 'object') {
      Object.keys(data).forEach((key) => {
        configForm[key] = String(data[key] || '')
      })
    }
  } catch {
    ElMessage.error('获取站点配置失败')
  } finally {
    loading.value = false
  }
}

async function handleSave() {
  saveLoading.value = true
  try {
    await setSiteConfig({ ...configForm })
    ElMessage.success('保存成功')
  } catch {
    // handled
  } finally {
    saveLoading.value = false
  }
}

function handleAddConfig() {
  const key = newKey.value.trim()
  const value = newValue.value.trim()
  if (!key) {
    ElMessage.warning('请输入配置项键名')
    return
  }
  if (configForm.hasOwnProperty(key)) {
    ElMessage.warning('该配置项已存在')
    return
  }
  configForm[key] = value
  newKey.value = ''
  newValue.value = ''
  ElMessage.success('已添加配置项，请点击保存')
}

onMounted(() => fetchConfig())
</script>
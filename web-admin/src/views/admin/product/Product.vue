<template>
  <div class="app-container product-page">
    <page-content
      ref="contentRef"
      :content-config="contentConfig"
      @add-click="handleAddClick"
      @edit-click="handleEditClick"
      @search-click="handleSearchClick"
      @toolbar-click="handleToolbarClick"
      @operate-click="handleOperateClick"
      @filter-change="handleFilterChange"
    >
      <template #status="scope">
        <el-tag :type="scope.row.status == 1 ? 'success' : 'info'">
          {{ scope.row.status == 1 ? "上架" : "下架" }}
        </el-tag>
      </template>
      <template #type="scope">
        <el-tag>{{ scope.row.type || '其他' }}</el-tag>
      </template>
      <template #description="scope">
        <span style="color: #999; font-size: 13px">{{ scope.row.description ? scope.row.description.substring(0, 40) + (scope.row.description.length > 40 ? '...' : '') : '-' }}</span>
      </template>
    </page-content>

    <!-- Add Modal -->
    <page-modal
      ref="addModalRef"
      :modal-config="addModalConfig"
      @submit-click="handleSubmitClick"
    >
      <template #cover="{ formData }">
        <div class="cover-upload-wrap">
          <el-upload
            :show-file-list="false"
            :http-request="onUpload"
            :before-upload="beforeUpload"
            class="cover-uploader"
            :on-success="(resp: any) => formData.cover = resp.data.file_url"
          >
            <img v-if="formData.cover" :src="formData.cover" class="cover-preview" />
            <el-icon v-else class="cover-uploader-icon"><Plus /></el-icon>
          </el-upload>
          <el-input v-model="formData.cover" size="small" placeholder="或输入URL" style="width: 200px; margin-left: 10px" />
        </div>
      </template>
    </page-modal>

    <!-- Edit Modal -->
    <page-modal
      ref="editModalRef"
      :modal-config="editModalConfig"
      @submit-click="handleSubmitClick"
    >
      <template #cover="{ formData }">
        <div class="cover-upload-wrap">
          <el-upload
            :show-file-list="false"
            :http-request="onUpload"
            :before-upload="beforeUpload"
            class="cover-uploader"
            :on-success="(resp: any) => formData.cover = resp.data.file_url"
          >
            <img v-if="formData.cover" :src="formData.cover" class="cover-preview" />
            <el-icon v-else class="cover-uploader-icon"><Plus /></el-icon>
          </el-upload>
          <el-input v-model="formData.cover" size="small" placeholder="或输入URL" style="width: 200px; margin-left: 10px" />
        </div>
      </template>
    </page-modal>

    <!-- Detail Content Editor Dialog -->
    <el-dialog v-model="detailDialogVisible" title="编辑产品详情" width="1100px" destroy-on-close>
      <MdEditor
        v-model="detailContent"
        :auto-detect-code="true"
        placeholder="请输入产品详情内容（支持Markdown、Mermaid、图片上传）"
        style="height: calc(100vh - 260px)"
        @on-upload-img="uploadImg"
      />
      <template #footer>
        <el-button @click="detailDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveDetailContent" :loading="detailSaving">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import type { UploadRawFile, UploadRequestOptions } from 'element-plus'
import { ElMessage } from 'element-plus'
import { MdEditor } from 'md-editor-v3'
import usePage from '@/components/CURD/usePage'
import addModalConfig from './config/add'
import contentConfig from './config/content'
import editModalConfig from './config/edit'
import PageModal from '@/components/CURD/PageModal.vue'
import PageContent from '@/components/CURD/PageContent.vue'
import { uploadFile } from '@/utils/file'
import { ProductAPI } from '@/api/product'
import type { IOperateData } from '@/components/CURD/types'

const {
  contentRef,
  addModalRef,
  editModalRef,
  handleAddClick,
  handleEditClick,
  handleSubmitClick,
  handleSearchClick,
  handleFilterChange,
} = usePage()

// Detail content editor state
const detailDialogVisible = ref(false)
const detailContent = ref('')
const detailProductId = ref<number>(0)
const detailSaving = ref(false)

function handleToolbarClick(name: string) {
  console.log(name)
}

function handleOperateClick(data: IOperateData) {
  switch (data.name) {
    case 'edit':
      handleEditClick(data.row)
      break
    case 'editDetail':
      detailProductId.value = data.row.id
      detailContent.value = data.row.content || ''
      detailDialogVisible.value = true
      break
    default:
      break
  }
}

async function saveDetailContent() {
  detailSaving.value = true
  try {
    await ProductAPI.updateProductApi({
      id: detailProductId.value,
      content: detailContent.value,
    })
    ElMessage.success('保存成功')
    detailDialogVisible.value = false
    // Refresh list
    const contentComp = contentRef.value
    if (contentComp?.fetchPageData) {
      contentComp.fetchPageData({}, true)
    }
  } catch {
    ElMessage.error('保存失败')
  } finally {
    detailSaving.value = false
  }
}

function beforeUpload(rawFile: UploadRawFile) {
  if (!rawFile.type.startsWith('image/')) {
    ElMessage.error('只能上传图片文件')
    return false
  }
  return true
}

function onUpload(options: UploadRequestOptions) {
  return uploadFile(options.file, 'blog/product/')
}

async function uploadImg(files: Array<File>, callback: (urls: string[]) => void) {
  const res = await Promise.all(
    files.map((file) => {
      return new Promise((rev, rej) => {
        uploadFile(file, 'blog/product/')
          .then((res) => rev(res))
          .catch((error) => rej(error))
      })
    })
  )
  callback(res.map((item: any) => item.data.file_url))
}
</script>

<style scoped>
.product-page {
  width: 100%;
}

.cover-upload-wrap {
  display: flex;
  align-items: center;
}

.cover-uploader :deep(.el-upload) {
  position: relative;
  overflow: hidden;
  cursor: pointer;
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  transition: border-color 0.3s;
}

.cover-uploader :deep(.el-upload:hover) {
  border-color: #409eff;
}

.cover-uploader-icon {
  width: 80px;
  height: 60px;
  font-size: 28px;
  color: #8c939d;
  display: flex;
  align-items: center;
  justify-content: center;
}

.cover-preview {
  width: 80px;
  height: 60px;
  object-fit: cover;
  display: block;
}
</style>
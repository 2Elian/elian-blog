<template>
  <div class="page-list">
    <div class="page-header">
      <h2>页面管理</h2>
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon> 新建页面
      </el-button>
    </div>

    <el-card>
      <div class="filter-bar">
        <el-input
          v-model="searchForm.keyword"
          placeholder="搜索页面标题"
          clearable
          style="width: 240px"
          @clear="fetchList"
          @keyup.enter="fetchList"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-select v-model="searchForm.status" placeholder="状态" clearable style="width: 120px" @change="fetchList">
          <el-option label="已发布" value="published" />
          <el-option label="草稿" value="draft" />
        </el-select>
        <el-button type="primary" @click="fetchList">搜索</el-button>
        <el-button @click="resetSearch">重置</el-button>
      </div>

      <el-table :data="tableData" v-loading="loading" stripe border>
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column prop="title" label="页面标题" min-width="200" />
        <el-table-column prop="slug" label="别名" width="150" />
        <el-table-column prop="status" label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.status === 'published' ? 'success' : 'warning'" size="small">
              {{ row.status === 'published' ? '已发布' : '草稿' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="view_count" label="浏览" width="80" align="center" />
        <el-table-column prop="updated_at" label="更新时间" width="170" align="center" />
        <el-table-column label="操作" width="160" align="center" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" link size="small" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-wrapper">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.page_size"
          :page-sizes="[10, 20, 50]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="fetchList"
          @current-change="fetchList"
        />
      </div>
    </el-card>

    <el-dialog
      v-model="dialogVisible"
      :title="dialogType === 'create' ? '新建页面' : '编辑页面'"
      width="720px"
      destroy-on-close
    >
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="80px">
        <el-form-item label="标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入页面标题" />
        </el-form-item>
        <el-form-item label="别名" prop="slug">
          <el-input v-model="form.slug" placeholder="请输入页面别名（URL友好）" />
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input v-model="form.content" type="textarea" :rows="12" placeholder="请输入页面内容" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="form.status">
            <el-radio value="draft">草稿</el-radio>
            <el-radio value="published">发布</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Plus, Search } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import { getPageList, createPage, updatePage, deletePage } from '@/api'

interface PageItem {
  id: number
  title: string
  slug: string
  content: string
  status: string
  view_count: number
  updated_at: string
}

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const dialogType = ref<'create' | 'edit'>('create')
const tableData = ref<PageItem[]>([])
const formRef = ref<FormInstance>()

const searchForm = reactive({ keyword: '', status: '' })
const pagination = reactive({ page: 1, page_size: 10, total: 0 })

const form = reactive({
  id: 0,
  title: '',
  slug: '',
  content: '',
  status: 'draft',
})

const formRules: FormRules = {
  title: [{ required: true, message: '请输入页面标题', trigger: 'blur' }],
  slug: [{ required: true, message: '请输入页面别名', trigger: 'blur' }],
  content: [{ required: true, message: '请输入页面内容', trigger: 'blur' }],
}

async function fetchList() {
  loading.value = true
  try {
    const params: Record<string, any> = { page: pagination.page, page_size: pagination.page_size }
    if (searchForm.keyword) params.keyword = searchForm.keyword
    if (searchForm.status) params.status = searchForm.status
    const res = await getPageList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch {
    ElMessage.error('获取页面列表失败')
  } finally {
    loading.value = false
  }
}

function resetSearch() {
  searchForm.keyword = ''
  searchForm.status = ''
  pagination.page = 1
  fetchList()
}

function handleCreate() {
  dialogType.value = 'create'
  Object.assign(form, { id: 0, title: '', slug: '', content: '', status: 'draft' })
  dialogVisible.value = true
}

function handleEdit(row: PageItem) {
  dialogType.value = 'edit'
  Object.assign(form, {
    id: row.id,
    title: row.title,
    slug: row.slug,
    content: row.content,
    status: row.status,
  })
  dialogVisible.value = true
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    if (dialogType.value === 'create') {
      await createPage(form)
      ElMessage.success('创建成功')
    } else {
      await updatePage(form.id, form)
      ElMessage.success('更新成功')
    }
    dialogVisible.value = false
    fetchList()
  } catch {
    // handled
  } finally {
    submitLoading.value = false
  }
}

async function handleDelete(row: PageItem) {
  try {
    await ElMessageBox.confirm(`确定要删除页面「${row.title}」吗？`, '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    })
    await deletePage(row.id)
    ElMessage.success('删除成功')
    fetchList()
  } catch {
    // cancelled
  }
}

onMounted(() => fetchList())
</script>
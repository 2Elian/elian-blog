<template>
  <div class="article-list">
    <div class="page-header">
      <h2>文章管理</h2>
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon> 新建文章
      </el-button>
    </div>

    <!-- 搜索筛选 -->
    <el-card>
      <div class="filter-bar">
        <el-input
          v-model="searchForm.keyword"
          placeholder="搜索文章标题"
          clearable
          style="width: 240px"
          @clear="fetchList"
          @keyup.enter="fetchList"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-select v-model="searchForm.category_id" placeholder="选择分类" clearable style="width: 160px" @change="fetchList">
          <el-option v-for="cat in categoryOptions" :key="cat.id" :label="cat.name" :value="cat.id" />
        </el-select>
        <el-select v-model="searchForm.status" placeholder="状态" clearable style="width: 120px" @change="fetchList">
          <el-option label="已发布" value="published" />
          <el-option label="草稿" value="draft" />
        </el-select>
        <el-button type="primary" @click="fetchList">
          <el-icon><Search /></el-icon> 搜索
        </el-button>
        <el-button @click="resetSearch">重置</el-button>
      </div>

      <!-- 表格 -->
      <el-table :data="tableData" v-loading="loading" stripe border>
        <el-table-column prop="id" label="ID" width="70" align="center" />
        <el-table-column prop="title" label="标题" min-width="200" show-overflow-tooltip />
        <el-table-column prop="category_name" label="分类" width="120" align="center">
          <template #default="{ row }">
            <el-tag v-if="row.category_name" type="info">{{ row.category_name }}</el-tag>
            <span v-else style="color: #c0c4cc">未分类</span>
          </template>
        </el-table-column>
        <el-table-column label="标签" width="200">
          <template #default="{ row }">
            <el-tag v-for="tag in (row.tags || []).slice(0, 3)" :key="tag.id" size="small" style="margin: 2px">{{ tag.name }}</el-tag>
            <el-tag v-if="(row.tags || []).length > 3" size="small" type="info" style="margin: 2px">+{{ row.tags.length - 3 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.status === 'published' ? 'success' : 'warning'" size="small">
              {{ row.status === 'published' ? '已发布' : '草稿' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="view_count" label="浏览" width="80" align="center" />
        <el-table-column prop="created_at" label="创建时间" width="170" align="center" />
        <el-table-column label="操作" width="180" align="center" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" link size="small" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
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

    <!-- 新建/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogType === 'create' ? '新建文章' : '编辑文章'"
      width="720px"
      destroy-on-close
    >
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="80px">
        <el-form-item label="标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入文章标题" />
        </el-form-item>
        <el-form-item label="分类" prop="category_id">
          <el-select v-model="form.category_id" placeholder="请选择分类" clearable style="width: 100%">
            <el-option v-for="cat in categoryOptions" :key="cat.id" :label="cat.name" :value="cat.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input v-model="form.content" type="textarea" :rows="10" placeholder="请输入文章内容" />
        </el-form-item>
        <el-form-item label="摘要">
          <el-input v-model="form.summary" type="textarea" :rows="3" placeholder="请输入文章摘要" />
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
import {
  getArticleList,
  createArticle,
  updateArticle,
  deleteArticle,
  getCategoryList,
} from '@/api'

interface ArticleItem {
  id: number
  title: string
  content: string
  summary: string
  category_id: number
  category_name: string
  tags: { id: number; name: string }[]
  status: string
  view_count: number
  created_at: string
}

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const dialogType = ref<'create' | 'edit'>('create')
const tableData = ref<ArticleItem[]>([])
const categoryOptions = ref<{ id: number; name: string }[]>([])
const formRef = ref<FormInstance>()

const searchForm = reactive({
  keyword: '',
  category_id: undefined as number | undefined,
  status: '',
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0,
})

const form = reactive({
  id: 0,
  title: '',
  content: '',
  summary: '',
  category_id: undefined as number | undefined,
  status: 'draft',
})

const formRules: FormRules = {
  title: [{ required: true, message: '请输入文章标题', trigger: 'blur' }],
  content: [{ required: true, message: '请输入文章内容', trigger: 'blur' }],
}

async function fetchList() {
  loading.value = true
  try {
    const params: Record<string, any> = {
      page: pagination.page,
      page_size: pagination.page_size,
    }
    if (searchForm.keyword) params.keyword = searchForm.keyword
    if (searchForm.category_id) params.category_id = searchForm.category_id
    if (searchForm.status) params.status = searchForm.status

    const res = await getArticleList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch {
    ElMessage.error('获取文章列表失败')
  } finally {
    loading.value = false
  }
}

async function fetchCategoryOptions() {
  try {
    const res = await getCategoryList({ page_size: 100 })
    categoryOptions.value = res.data.list || []
  } catch {
    // 静默失败
  }
}

function resetSearch() {
  searchForm.keyword = ''
  searchForm.category_id = undefined
  searchForm.status = ''
  pagination.page = 1
  fetchList()
}

function handleCreate() {
  dialogType.value = 'create'
  form.id = 0
  form.title = ''
  form.content = ''
  form.summary = ''
  form.category_id = undefined
  form.status = 'draft'
  dialogVisible.value = true
}

function handleEdit(row: ArticleItem) {
  dialogType.value = 'edit'
  form.id = row.id
  form.title = row.title
  form.content = row.content
  form.summary = row.summary || ''
  form.category_id = row.category_id
  form.status = row.status
  dialogVisible.value = true
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return

  submitLoading.value = true
  try {
    const data = { ...form }
    if (dialogType.value === 'create') {
      await createArticle(data)
      ElMessage.success('创建成功')
    } else {
      await updateArticle(form.id, data)
      ElMessage.success('更新成功')
    }
    dialogVisible.value = false
    fetchList()
  } catch {
    // 错误已在拦截器中处理
  } finally {
    submitLoading.value = false
  }
}

async function handleDelete(row: ArticleItem) {
  try {
    await ElMessageBox.confirm(`确定要删除文章「${row.title}」吗？`, '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    })
    await deleteArticle(row.id)
    ElMessage.success('删除成功')
    fetchList()
  } catch {
    // 用户取消
  }
}

onMounted(() => {
  fetchList()
  fetchCategoryOptions()
})
</script>
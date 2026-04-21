<template>
  <div class="category-list">
    <div class="page-header">
      <h2>分类管理</h2>
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon> 新建分类
      </el-button>
    </div>

    <el-card>
      <!-- 搜索 -->
      <div class="filter-bar">
        <el-input
          v-model="searchForm.keyword"
          placeholder="搜索分类名称"
          clearable
          style="width: 240px"
          @clear="fetchList"
          @keyup.enter="fetchList"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-button type="primary" @click="fetchList">搜索</el-button>
        <el-button @click="resetSearch">重置</el-button>
      </div>

      <!-- 表格 -->
      <el-table :data="tableData" v-loading="loading" stripe border>
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column prop="name" label="分类名称" min-width="200" />
        <el-table-column prop="slug" label="别名" min-width="150" />
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="article_count" label="文章数" width="100" align="center" />
        <el-table-column prop="sort_order" label="排序" width="80" align="center" />
        <el-table-column prop="created_at" label="创建时间" width="170" align="center" />
        <el-table-column label="操作" width="160" align="center" fixed="right">
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
      :title="dialogType === 'create' ? '新建分类' : '编辑分类'"
      width="500px"
      destroy-on-close
    >
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="别名" prop="slug">
          <el-input v-model="form.slug" placeholder="请输入分类别名（URL友好）" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入分类描述" />
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="form.sort_order" :min="0" :max="999" />
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
import { getCategoryList, createCategory, updateCategory, deleteCategory } from '@/api'

interface CategoryItem {
  id: number
  name: string
  slug: string
  description: string
  article_count: number
  sort_order: number
  created_at: string
}

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const dialogType = ref<'create' | 'edit'>('create')
const tableData = ref<CategoryItem[]>([])
const formRef = ref<FormInstance>()

const searchForm = reactive({ keyword: '' })

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0,
})

const form = reactive({
  id: 0,
  name: '',
  slug: '',
  description: '',
  sort_order: 0,
})

const formRules: FormRules = {
  name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
  slug: [{ required: true, message: '请输入分类别名', trigger: 'blur' }],
}

async function fetchList() {
  loading.value = true
  try {
    const params: Record<string, any> = { page: pagination.page, page_size: pagination.page_size }
    if (searchForm.keyword) params.keyword = searchForm.keyword
    const res = await getCategoryList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch {
    ElMessage.error('获取分类列表失败')
  } finally {
    loading.value = false
  }
}

function resetSearch() {
  searchForm.keyword = ''
  pagination.page = 1
  fetchList()
}

function handleCreate() {
  dialogType.value = 'create'
  Object.assign(form, { id: 0, name: '', slug: '', description: '', sort_order: 0 })
  dialogVisible.value = true
}

function handleEdit(row: CategoryItem) {
  dialogType.value = 'edit'
  Object.assign(form, {
    id: row.id,
    name: row.name,
    slug: row.slug,
    description: row.description || '',
    sort_order: row.sort_order || 0,
  })
  dialogVisible.value = true
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    if (dialogType.value === 'create') {
      await createCategory(form)
      ElMessage.success('创建成功')
    } else {
      await updateCategory(form.id, form)
      ElMessage.success('更新成功')
    }
    dialogVisible.value = false
    fetchList()
  } catch {
    // handled by interceptor
  } finally {
    submitLoading.value = false
  }
}

async function handleDelete(row: CategoryItem) {
  try {
    await ElMessageBox.confirm(`确定要删除分类「${row.name}」吗？`, '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    })
    await deleteCategory(row.id)
    ElMessage.success('删除成功')
    fetchList()
  } catch {
    // cancelled
  }
}

onMounted(() => {
  fetchList()
})
</script>
<template>
  <div class="menu-list">
    <div class="page-header">
      <h2>菜单管理</h2>
      <el-button type="primary" @click="handleCreate(0)">
        <el-icon><Plus /></el-icon> 新建顶级菜单
      </el-button>
    </div>

    <el-card>
      <el-table
        :data="tableData"
        v-loading="loading"
        row-key="id"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        border
        default-expand-all
      >
        <el-table-column prop="name" label="菜单名称" min-width="200" />
        <el-table-column prop="icon" label="图标" width="100" align="center">
          <template #default="{ row }">
            <el-icon v-if="row.icon"><component :is="row.icon" /></el-icon>
            <span v-else style="color: #c0c4cc">-</span>
          </template>
        </el-table-column>
        <el-table-column prop="path" label="路由路径" width="160" />
        <el-table-column prop="sort_order" label="排序" width="80" align="center" />
        <el-table-column prop="visible" label="可见" width="80" align="center">
          <template #default="{ row }">
            <el-tag :type="row.visible !== false ? 'success' : 'danger'" size="small">
              {{ row.visible !== false ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="240" align="center">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click="handleCreate(row.id)">添加子菜单</el-button>
            <el-button type="warning" link size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" link size="small" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 新建/编辑菜单对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogType === 'create' ? '新建菜单' : '编辑菜单'"
      width="560px"
      destroy-on-close
    >
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="90px">
        <el-form-item label="上级菜单">
          <el-tree-select
            v-model="form.parent_id"
            :data="menuTreeOptions"
            :props="{ label: 'name', value: 'id', children: 'children' }"
            placeholder="请选择上级菜单（空为顶级）"
            clearable
            check-strictly
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="菜单名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入菜单名称" />
        </el-form-item>
        <el-form-item label="图标">
          <el-input v-model="form.icon" placeholder="请输入图标名称（如 Document）" />
        </el-form-item>
        <el-form-item label="路由路径" prop="path">
          <el-input v-model="form.path" placeholder="请输入路由路径" />
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="form.sort_order" :min="0" :max="999" />
        </el-form-item>
        <el-form-item label="是否可见">
          <el-switch v-model="form.visible" active-text="显示" inactive-text="隐藏" />
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
import { Plus } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import { getMenuList, createMenu, updateMenu, deleteMenu } from '@/api'

interface MenuItem {
  id: number
  parent_id: number
  name: string
  icon: string
  path: string
  sort_order: number
  visible: boolean
  children?: MenuItem[]
}

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const dialogType = ref<'create' | 'edit'>('create')
const tableData = ref<MenuItem[]>([])
const menuTreeOptions = ref<MenuItem[]>([])
const formRef = ref<FormInstance>()

const form = reactive({
  id: 0,
  parent_id: 0,
  name: '',
  icon: '',
  path: '',
  sort_order: 0,
  visible: true,
})

const formRules: FormRules = {
  name: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }],
  path: [{ required: true, message: '请输入路由路径', trigger: 'blur' }],
}

async function fetchList() {
  loading.value = true
  try {
    const res = await getMenuList()
    tableData.value = res.data.list || []
    // 构建菜单选项（添加顶级选项）
    menuTreeOptions.value = [
      { id: 0, name: '顶级菜单', parent_id: 0, icon: '', path: '', sort_order: 0, visible: true, children: res.data.list || [] },
    ]
  } catch {
    ElMessage.error('获取菜单列表失败')
  } finally {
    loading.value = false
  }
}

function handleCreate(parentId: number) {
  dialogType.value = 'create'
  Object.assign(form, {
    id: 0,
    parent_id: parentId,
    name: '',
    icon: '',
    path: '',
    sort_order: 0,
    visible: true,
  })
  dialogVisible.value = true
}

function handleEdit(row: MenuItem) {
  dialogType.value = 'edit'
  Object.assign(form, {
    id: row.id,
    parent_id: row.parent_id || 0,
    name: row.name,
    icon: row.icon || '',
    path: row.path || '',
    sort_order: row.sort_order || 0,
    visible: row.visible !== false,
  })
  dialogVisible.value = true
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    if (dialogType.value === 'create') {
      await createMenu(form)
      ElMessage.success('创建成功')
    } else {
      await updateMenu(form.id, form)
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

async function handleDelete(row: MenuItem) {
  if (row.children && row.children.length > 0) {
    ElMessage.warning('请先删除子菜单')
    return
  }
  try {
    await ElMessageBox.confirm(`确定要删除菜单「${row.name}」吗？`, '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    })
    await deleteMenu(row.id)
    ElMessage.success('删除成功')
    fetchList()
  } catch {
    // cancelled
  }
}

onMounted(() => fetchList())
</script>
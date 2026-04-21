<template>
  <div class="role-list">
    <div class="page-header">
      <h2>角色管理</h2>
      <el-button type="primary" @click="handleCreate">
        <el-icon><Plus /></el-icon> 新建角色
      </el-button>
    </div>

    <el-card>
      <div class="filter-bar">
        <el-input
          v-model="searchForm.keyword"
          placeholder="搜索角色名称"
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

      <el-table :data="tableData" v-loading="loading" stripe border>
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column prop="name" label="角色名称" min-width="150" />
        <el-table-column prop="code" label="角色标识" width="140" />
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column prop="created_at" label="创建时间" width="170" align="center" />
        <el-table-column label="操作" width="240" align="center" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="warning" link size="small" @click="handleAssignMenus(row)">分配菜单</el-button>
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

    <!-- 新建/编辑角色对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogType === 'create' ? '新建角色' : '编辑角色'"
      width="500px"
      destroy-on-close
    >
      <el-form ref="formRef" :model="form" :rules="formRules" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入角色名称" />
        </el-form-item>
        <el-form-item label="标识" prop="code">
          <el-input v-model="form.code" placeholder="请输入角色标识（如 admin, editor）" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入角色描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- 分配菜单对话框 -->
    <el-dialog
      v-model="menuDialogVisible"
      title="分配菜单权限"
      width="500px"
      destroy-on-close
    >
      <el-tree
        ref="menuTreeRef"
        :data="menuTreeData"
        show-checkbox
        node-key="id"
        :default-checked-keys="checkedMenuIds"
        :props="menuTreeProps"
        default-expand-all
      />
      <template #footer>
        <el-button @click="menuDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="menuSubmitLoading" @click="handleMenuSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Plus, Search } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import type { ElTree } from 'element-plus'
import {
  getRoleList,
  createRole,
  updateRole,
  deleteRole,
  assignRoleMenus,
  getRoleMenus,
  getMenuList,
} from '@/api'

interface RoleItem {
  id: number
  name: string
  code: string
  description: string
  created_at: string
}

interface MenuItem {
  id: number
  name: string
  children?: MenuItem[]
}

const loading = ref(false)
const submitLoading = ref(false)
const menuSubmitLoading = ref(false)
const dialogVisible = ref(false)
const menuDialogVisible = ref(false)
const dialogType = ref<'create' | 'edit'>('create')
const tableData = ref<RoleItem[]>([])
const formRef = ref<FormInstance>()
const menuTreeRef = ref<InstanceType<typeof ElTree>>()

const searchForm = reactive({ keyword: '' })
const pagination = reactive({ page: 1, page_size: 10, total: 0 })

const form = reactive({
  id: 0,
  name: '',
  code: '',
  description: '',
})

const formRules: FormRules = {
  name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入角色标识', trigger: 'blur' }],
}

// 菜单分配相关
const menuTreeData = ref<MenuItem[]>([])
const checkedMenuIds = ref<number[]>([])
const currentRoleId = ref(0)

const menuTreeProps = {
  label: 'name',
  children: 'children',
}

async function fetchList() {
  loading.value = true
  try {
    const params: Record<string, any> = { page: pagination.page, page_size: pagination.page_size }
    if (searchForm.keyword) params.keyword = searchForm.keyword
    const res = await getRoleList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch {
    ElMessage.error('获取角色列表失败')
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
  Object.assign(form, { id: 0, name: '', code: '', description: '' })
  dialogVisible.value = true
}

function handleEdit(row: RoleItem) {
  dialogType.value = 'edit'
  Object.assign(form, {
    id: row.id,
    name: row.name,
    code: row.code,
    description: row.description || '',
  })
  dialogVisible.value = true
}

async function handleSubmit() {
  const valid = await formRef.value?.validate().catch(() => false)
  if (!valid) return
  submitLoading.value = true
  try {
    if (dialogType.value === 'create') {
      await createRole(form)
      ElMessage.success('创建成功')
    } else {
      await updateRole(form.id, form)
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

async function handleDelete(row: RoleItem) {
  try {
    await ElMessageBox.confirm(`确定要删除角色「${row.name}」吗？`, '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    })
    await deleteRole(row.id)
    ElMessage.success('删除成功')
    fetchList()
  } catch {
    // cancelled
  }
}

async function handleAssignMenus(row: RoleItem) {
  currentRoleId.value = row.id
  menuDialogVisible.value = true

  // 加载菜单树
  try {
    const menuRes = await getMenuList()
    menuTreeData.value = menuRes.data.list || []
  } catch {
    menuTreeData.value = []
  }

  // 加载角色已分配的菜单
  try {
    const res = await getRoleMenus(row.id)
    checkedMenuIds.value = res.data.menu_ids || []
  } catch {
    checkedMenuIds.value = []
  }
}

async function handleMenuSubmit() {
  if (!menuTreeRef.value) return
  menuSubmitLoading.value = true
  try {
    const checkedKeys = menuTreeRef.value.getCheckedKeys(false) as number[]
    const halfCheckedKeys = menuTreeRef.value.getHalfCheckedKeys() as number[]
    const menuIds = [...checkedKeys, ...halfCheckedKeys]
    await assignRoleMenus(currentRoleId.value, { menu_ids: menuIds })
    ElMessage.success('菜单权限分配成功')
    menuDialogVisible.value = false
  } catch {
    // handled
  } finally {
    menuSubmitLoading.value = false
  }
}

onMounted(() => fetchList())
</script>
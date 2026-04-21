<template>
  <div class="comment-list">
    <div class="page-header">
      <h2>评论管理</h2>
    </div>

    <el-card>
      <div class="filter-bar">
        <el-input
          v-model="searchForm.keyword"
          placeholder="搜索评论内容"
          clearable
          style="width: 240px"
          @clear="fetchList"
          @keyup.enter="fetchList"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-select v-model="searchForm.status" placeholder="状态" clearable style="width: 140px" @change="fetchList">
          <el-option label="待审核" value="pending" />
          <el-option label="已通过" value="approved" />
          <el-option label="已拒绝" value="rejected" />
        </el-select>
        <el-button type="primary" @click="fetchList">搜索</el-button>
        <el-button @click="resetSearch">重置</el-button>
      </div>

      <el-table :data="tableData" v-loading="loading" stripe border>
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column prop="user_name" label="评论者" width="140" />
        <el-table-column prop="content" label="评论内容" min-width="250" show-overflow-tooltip />
        <el-table-column prop="article_title" label="所属文章" width="180" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="getStatusTagType(row.status)" size="small">
              {{ getStatusLabel(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="评论时间" width="170" align="center" />
        <el-table-column label="操作" width="200" align="center" fixed="right">
          <template #default="{ row }">
            <el-button
              v-if="row.status !== 'approved'"
              type="success"
              link
              size="small"
              @click="handleStatusChange(row, 'approved')"
            >通过</el-button>
            <el-button
              v-if="row.status !== 'rejected'"
              type="warning"
              link
              size="small"
              @click="handleStatusChange(row, 'rejected')"
            >拒绝</el-button>
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getCommentList, updateCommentStatus, deleteComment } from '@/api'

interface CommentItem {
  id: number
  user_name: string
  content: string
  article_title: string
  status: string
  created_at: string
}

const loading = ref(false)
const tableData = ref<CommentItem[]>([])

const searchForm = reactive({ keyword: '', status: '' })
const pagination = reactive({ page: 1, page_size: 10, total: 0 })

function getStatusTagType(status: string) {
  const map: Record<string, string> = { pending: 'warning', approved: 'success', rejected: 'danger' }
  return map[status] || 'info'
}

function getStatusLabel(status: string) {
  const map: Record<string, string> = { pending: '待审核', approved: '已通过', rejected: '已拒绝' }
  return map[status] || status
}

async function fetchList() {
  loading.value = true
  try {
    const params: Record<string, any> = { page: pagination.page, page_size: pagination.page_size }
    if (searchForm.keyword) params.keyword = searchForm.keyword
    if (searchForm.status) params.status = searchForm.status
    const res = await getCommentList(params)
    tableData.value = res.data.list || []
    pagination.total = res.data.total || 0
  } catch {
    ElMessage.error('获取评论列表失败')
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

async function handleStatusChange(row: CommentItem, status: string) {
  try {
    await updateCommentStatus(row.id, { status })
    ElMessage.success('状态更新成功')
    fetchList()
  } catch {
    // handled
  }
}

async function handleDelete(row: CommentItem) {
  try {
    await ElMessageBox.confirm('确定要删除该评论吗？', '删除确认', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    })
    await deleteComment(row.id)
    ElMessage.success('删除成功')
    fetchList()
  } catch {
    // cancelled
  }
}

onMounted(() => fetchList())
</script>
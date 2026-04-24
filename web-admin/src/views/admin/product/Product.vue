<template>
  <div class="app-container">
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
        <el-tag :type="scope.row.type === 1 ? 'primary' : scope.row.type === 2 ? 'warning' : 'info'">
          {{ scope.row.type === 1 ? "AI产品" : scope.row.type === 2 ? "工具" : "其他" }}
        </el-tag>
      </template>
    </page-content>

    <page-modal
      ref="addModalRef"
      :modal-config="addModalConfig"
      @submit-click="handleSubmitClick"
    />
    <page-modal
      ref="editModalRef"
      :modal-config="editModalConfig"
      @submit-click="handleSubmitClick"
    />
  </div>
</template>

<script setup lang="ts">
import type { IOperateData } from "@/components/CURD/types";
import usePage from "@/components/CURD/usePage";
import addModalConfig from "./config/add";
import contentConfig from "./config/content";
import editModalConfig from "./config/edit";
import PageModal from "@/components/CURD/PageModal.vue";
import PageContent from "@/components/CURD/PageContent.vue";

const {
  contentRef,
  addModalRef,
  editModalRef,
  handleAddClick,
  handleEditClick,
  handleSubmitClick,
  handleSearchClick,
  handleFilterChange,
} = usePage();

function handleToolbarClick(name: string) {
  console.log(name);
}

function handleOperateClick(data: IOperateData) {
  switch (data.name) {
    case "edit":
      handleEditClick(data.row);
      break;
    default:
      break;
  }
}
</script>

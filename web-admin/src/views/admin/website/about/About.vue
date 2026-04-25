<template>
  <div class="app-container">
    <el-card class="main-card">
      <div class="table-title">关于页面配置</div>
      <el-tabs v-model="activeTab">
        <!-- 基本信息 -->
        <el-tab-pane label="个人信息" name="hero">
          <el-form label-position="left" label-width="100px" style="max-width: 700px; margin-top: 1rem">
            <el-form-item label="姓名">
              <el-input v-model="aboutData.hero.name" size="small" />
            </el-form-item>
            <el-form-item label="身份标签">
              <el-input v-model="aboutData.hero.role" size="small" />
            </el-form-item>
            <el-form-item label="个人简介">
              <el-input v-model="aboutData.hero.description" type="textarea" :rows="3" />
            </el-form-item>
            <el-divider content-position="left">信息卡片</el-divider>
            <div v-for="(item, idx) in aboutData.hero.info_card" :key="idx" style="display: flex; gap: 8px; margin-bottom: 8px; align-items: center">
              <el-input v-model="item.label" placeholder="标签" size="small" style="width: 120px" />
              <el-input v-model="item.value" placeholder="值" size="small" style="flex: 1" />
              <el-button size="small" type="danger" @click="aboutData.hero.info_card.splice(idx, 1)">删除</el-button>
            </div>
            <el-button size="small" @click="aboutData.hero.info_card.push({ label: '', value: '' })">+ 添加信息</el-button>
            <el-divider content-position="left">社交链接</el-divider>
            <div v-for="(link, idx) in aboutData.hero.links" :key="idx" style="display: flex; gap: 8px; margin-bottom: 8px; align-items: center">
              <el-input v-model="link.name" placeholder="名称" size="small" style="width: 100px" />
              <el-input v-model="link.url" placeholder="链接" size="small" style="flex: 1" />
              <el-select v-model="link.icon" size="small" style="width: 100px">
                <el-option label="GitHub" value="github" />
                <el-option label="邮箱" value="email" />
                <el-option label="主页" value="website" />
              </el-select>
              <el-button size="small" type="danger" @click="aboutData.hero.links.splice(idx, 1)">删除</el-button>
            </div>
            <el-button size="small" @click="aboutData.hero.links.push({ name: '', url: '', icon: 'website' })">+ 添加链接</el-button>
          </el-form>
        </el-tab-pane>

        <!-- 技术栈 -->
        <el-tab-pane label="技术栈" name="tech">
          <div style="margin-top: 1rem">
            <div v-for="(group, gIdx) in aboutData.tech_stack" :key="gIdx" style="margin-bottom: 16px">
              <div style="display: flex; gap: 8px; align-items: center; margin-bottom: 8px">
                <el-input v-model="group.title" placeholder="分组名称" size="small" style="width: 200px" />
                <el-button size="small" type="danger" @click="aboutData.tech_stack.splice(gIdx, 1)">删除分组</el-button>
              </div>
              <div style="display: flex; flex-wrap: wrap; gap: 6px">
                <el-tag v-for="(tag, tIdx) in group.tags" :key="tIdx" closable @close="group.tags.splice(tIdx, 1)">{{ tag }}</el-tag>
              </div>
              <div style="display: flex; gap: 6px; margin-top: 6px">
                <el-input v-model="newTags[gIdx]" placeholder="输入标签，逗号分隔" size="small" style="width: 300px" @keyup.enter="addTags(gIdx)" />
                <el-button size="small" @click="addTags(gIdx)">添加</el-button>
              </div>
            </div>
            <el-button type="primary" size="small" @click="aboutData.tech_stack.push({ title: '', tags: [] })">+ 添加分组</el-button>
          </div>
        </el-tab-pane>

        <!-- 代表项目 -->
        <el-tab-pane label="代表项目" name="projects">
          <div style="margin-top: 1rem">
            <div v-for="(proj, pIdx) in aboutData.projects" :key="pIdx" style="border: 1px solid #ebeef5; border-radius: 8px; padding: 16px; margin-bottom: 16px">
              <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 12px">
                <strong>项目 {{ pIdx + 1 }}</strong>
                <el-button size="small" type="danger" @click="aboutData.projects.splice(pIdx, 1)">删除</el-button>
              </div>
              <el-form label-position="left" label-width="80px">
                <el-form-item label="项目名称">
                  <el-input v-model="proj.title" size="small" />
                </el-form-item>
                <el-form-item label="荣誉标签">
                  <el-input v-model="proj.badge" size="small" placeholder="如: KDD Cup Top1" />
                </el-form-item>
                <el-form-item label="项目描述">
                  <el-input v-model="proj.desc" type="textarea" :rows="3" />
                </el-form-item>
                <el-form-item label="技术标签">
                  <div style="display: flex; flex-wrap: wrap; gap: 6px; margin-bottom: 6px">
                    <el-tag v-for="(t, tIdx) in proj.tech" :key="tIdx" closable @close="proj.tech.splice(tIdx, 1)">{{ t }}</el-tag>
                  </div>
                  <el-input v-model="projNewTech[pIdx]" placeholder="逗号分隔" size="small" style="width: 300px" @keyup.enter="addProjTech(pIdx)" />
                  <el-button size="small" style="margin-left: 6px" @click="addProjTech(pIdx)">添加</el-button>
                </el-form-item>
                <el-form-item label="项目链接">
                  <el-input v-model="proj.link" size="small" placeholder="https://github.com/..." />
                </el-form-item>
                <el-form-item label="封面图">
                  <div style="display: flex; gap: 8px; align-items: center">
                    <el-upload
                      :show-file-list="false"
                      :http-request="onUpload"
                      :before-upload="beforeUpload"
                      class="avatar-uploader"
                      :on-success="(resp: any) => proj.image = resp.data.file_url"
                    >
                      <img v-if="proj.image" :src="proj.image" style="width: 80px; height: 80px; object-fit: cover; border-radius: 6px" />
                      <el-button v-else size="small">上传图片</el-button>
                    </el-upload>
                    <el-input v-model="proj.image" size="small" placeholder="/uploads/..." style="width: 260px" />
                  </div>
                </el-form-item>
              </el-form>
            </div>
            <el-button type="primary" size="small" @click="aboutData.projects.push({ title: '', badge: '', desc: '', tech: [], link: '', image: '' })">+ 添加项目</el-button>
          </div>
        </el-tab-pane>

        <!-- 科研与竞赛 -->
        <el-tab-pane label="科研与竞赛" name="awards">
          <div style="margin-top: 1rem">
            <div v-for="(award, aIdx) in aboutData.awards" :key="aIdx" style="display: flex; gap: 8px; margin-bottom: 8px; align-items: center">
              <el-input v-model="aboutData.awards[aIdx]" size="small" style="flex: 1" />
              <el-button size="small" type="danger" @click="aboutData.awards.splice(aIdx, 1)">删除</el-button>
            </div>
            <el-button size="small" @click="aboutData.awards.push('')">+ 添加</el-button>
          </div>
        </el-tab-pane>
      </el-tabs>

      <el-button class="edit-btn" size="default" type="primary" @click="saveAbout">保存全部</el-button>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, reactive } from "vue";
import type { UploadRawFile, UploadRequestOptions } from "element-plus";
import { ElMessage, ElMessageBox } from "element-plus";
import { uploadFile } from "@/utils/file";

function beforeUpload(rawFile: UploadRawFile) { return rawFile; }
function onUpload(options: UploadRequestOptions) { return uploadFile(options.file, "blog/about/"); }

const activeTab = ref("hero");

const defaultData = () => ({
  hero: {
    name: "李子茉(Elian)",
    role: "AI 算法工程师 / Agent 全栈开发",
    description:
      "硕士就读于南京信息工程大学数学专业，研究方向为深度学习。\n目前我正在寻找读cs/数学phd的机会, 如果您愿意推荐我, 请联系我的邮箱！",
    links: [
      { name: "GitHub", url: "https://github.com/2Elian", icon: "github" },
      { name: "邮箱", url: "mailto:lizimo@nuist.edu.cn", icon: "email" },
      { name: "主页", url: "https://2elian.github.io", icon: "website" },
    ],
    info_card: [
      { label: "学历", value: "硕士（双一流）" },
      { label: "院校", value: "南京信息工程大学" },
      { label: "专业", value: "数学" },
      { label: "方向", value: "LLM持续学习" },
      { label: "邮箱", value: "lizimo@nuist.edu.cn" },
    ],
  },
  tech_stack: [
    {
      title: "后端开发",
      tags: [
        "Python", "Go", "Java", "FastAPI", "Gin", "Go-Zero", "SpringBoot",
        "MySQL", "Redis", "Neo4j", "MongoDB",
      ],
    },
    {
      title: "AI / LLM",
      tags: [
        "LangChain", "LangGraph", "GraphRAG", "Eino", "SpringAI",
        "LoRA", "SFT", "RLHF", "Agentic-RL",
        "transformers", "llamafactory", "verl",
      ],
    },
    {
      title: "基础设施",
      tags: ["Linux", "Docker", "Git", "Milvus", "Vite"],
    },
  ],
  projects: [
    {
      title: "elian — Agent 全链路开发工具",
      badge: "",
      desc: "一个功能强大的 Agent 开发 Python 框架，包括 SDK、CLI 脚手架、数据治理到垂类微调的 Factory、以及 to-Go 工具链。比 LangChain 更稳定，比 LlamaIndex 更易扩展。",
      tech: ["Python", "Agent", "LLM", "CLI"],
      link: "",
      image: "",
    },
    {
      title: "基于 Plan-Execute-RePlan 的多源异构数据分析 Agent",
      badge: "KDD Cup 2026 Stage1-Top1",
      desc: "构建能分解复杂分析问题、协调多步骤推理的数据 Agent。设计了全局 Planner + ReAct 子任务 + RePlanner + 验证 Agent 的完整架构，采用 DAG 并发执行无依赖节点。",
      tech: ["Multi-Agent", "ReAct", "Agent-RL", "verl"],
      link: "",
      image: "",
    },
    {
      title: "基于 LangGraph 的图文广告合规审核 Agent",
      badge: "腾讯开悟法律AI大赛 Top1",
      desc: "一站式 AI 合规工作台，集成多 Agent 协作。通过异步改造、缓存优化、Prompt 精简将成本降低 86.67%，单次审核成本从 1 元降至 0.1 元。",
      tech: ["LangGraph", "Multi-Agent", "RAG", "FastAPI"],
      link: "",
      image: "",
    },
    {
      title: "基于 Eino 与 Gin 的智能客服 Agent",
      badge: "",
      desc: "以多模态电商智能客服比赛为背景，使用 Go 重构 Python Agent 核心链路，对比高并发场景下两种语言的性能表现与工程成本。",
      tech: ["Go", "Gin", "Eino", "FastAPI", "Redis"],
      link: "",
      image: "",
    },
  ],
  awards: [
    "KDD Cup 2026 多源异构数据分析 — Stage1 Top1",
    "腾讯开悟法律 AI 大赛 — Top1",
    "2025 华为杯全国研究生数学建模大赛 — 国家级三等奖",
    "2024 华为杯全国研究生数学建模大赛 — 国家级二等奖",
    "江苏省应用数学中心研究生论坛 — 墙报入围",
    "发表论文：V-PEP (Arxiv 2026)、Occ-MoE (IJPRAI 2026)、ATD-TBPS (TVC 2025)",
  ],
});

const aboutData = reactive(defaultData());
const newTags = ref<Record<number, string>>({});
const projNewTech = ref<Record<number, string>>({});

function addTags(gIdx: number) {
  const input = newTags.value[gIdx] || "";
  const tags = input.split(",").map((t) => t.trim()).filter(Boolean);
  aboutData.tech_stack[gIdx].tags.push(...tags);
  newTags.value[gIdx] = "";
}

function addProjTech(pIdx: number) {
  const input = projNewTech.value[pIdx] || "";
  const tags = input.split(",").map((t) => t.trim()).filter(Boolean);
  aboutData.projects[pIdx].tech.push(...tags);
  projNewTech.value[pIdx] = "";
}

async function fetchAbout() {
  try {
    const res = await fetch("/admin-api/v1/admin/get_about_me", {
      headers: { Authorization: `Bearer ${localStorage.getItem("token") || ""}` },
    });
    const json = await res.json();
    if (json.data?.content) {
      try {
        const parsed = JSON.parse(json.data.content);
        Object.assign(aboutData, parsed);
      } catch {
        // old markdown content, ignore
      }
    }
  } catch {
    // use defaults
  }
}

async function saveAbout() {
  const content = JSON.stringify(aboutData);
  try {
    const res = await fetch("/admin-api/v1/admin/update_about_me", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token") || ""}`,
      },
      body: JSON.stringify({ content }),
    });
    const json = await res.json();
    if (json.code === 200) {
      ElMessage.success("保存成功");
    } else {
      ElMessage.error(json.msg || "保存失败");
    }
  } catch {
    ElMessage.error("保存失败");
  }
}

// suppress unused import warning
var _ = ElMessageBox;

onMounted(() => {
  fetchAbout();
});
</script>

<style scoped>
.edit-btn {
  float: right;
  margin: 1rem 0;
}
</style>

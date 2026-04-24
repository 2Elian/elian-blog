<template>
  <div class="about-page">
    <!-- Hero Banner (Full Width) -->
    <section class="about-hero">
      <div class="hero-bg">
        <div class="hero-grid"></div>
      </div>
      <div class="hero-inner">
        <div class="hero-left">
          <div class="hero-badge">关于作者</div>
          <h1 class="hero-name">李子茉(Elian)</h1>
          <p class="hero-role">AI 算法工程师 / Agent 全栈开发</p>
          <p class="hero-desc">
            硕士就读于南京信息工程大学数学专业，研究方向为深度学习。<br />
            目前我正在寻找读cs/数学phd的机会, 如果您愿意推荐我, 请联系我的邮箱！
          </p>
          <div class="hero-links">
            <a href="https://github.com/2Elian" target="_blank" class="link-btn">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"/></svg>
              GitHub
            </a>
            <a href="mailto:lizimo@nuist.edu.cn" class="link-btn">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="2" y="4" width="20" height="16" rx="2"/><path d="M22 7l-10 7L2 7"/></svg>
              邮箱
            </a>
            <a href="https://2elian.github.io" target="_blank" class="link-btn">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="2" y1="12" x2="22" y2="12"/><path d="M12 2a15.3 15.3 0 014 10 15.3 15.3 0 01-4 10 15.3 15.3 0 01-4-10A15.3 15.3 0 0112 2z"/></svg>
              主页
            </a>
          </div>
        </div>
        <div class="hero-right">
          <div class="info-card">
            <div class="info-row">
              <span class="info-label">学历</span>
              <span class="info-value">硕士（双一流）</span>
            </div>
            <div class="info-row">
              <span class="info-label">院校</span>
              <span class="info-value">南京信息工程大学</span>
            </div>
            <div class="info-row">
              <span class="info-label">专业</span>
              <span class="info-value">数学</span>
            </div>
            <div class="info-row">
              <span class="info-label">方向</span>
              <span class="info-value">LLM持续学习</span>
            </div>
            <div class="info-row">
              <span class="info-label">邮箱</span>
              <span class="info-value">lizimo@nuist.edu.cn</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Content -->
    <div class="about-body">
      <div class="body-inner">
        <!-- Dynamic About Content from Backend -->
        <div v-if="aboutContent" class="about-dynamic" v-html="renderedAbout"></div>

        <!-- Default Static Content (shown when no backend content) -->
        <template v-else>
        <!-- Tech Stack -->
        <section class="content-block">
          <h2 class="block-title">
            <span class="title-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/></svg>
            </span>
            技术栈
          </h2>
          <div class="tech-sections">
            <div class="tech-group">
              <h3 class="tech-group-title">后端开发</h3>
              <div class="tech-tags">
                <span class="tech-tag" v-for="t in backendTech" :key="t">{{ t }}</span>
              </div>
            </div>
            <div class="tech-group">
              <h3 class="tech-group-title">AI / LLM</h3>
              <div class="tech-tags">
                <span class="tech-tag" v-for="t in aiTech" :key="t">{{ t }}</span>
              </div>
            </div>
            <div class="tech-group">
              <h3 class="tech-group-title">基础设施</h3>
              <div class="tech-tags">
                <span class="tech-tag" v-for="t in infraTech" :key="t">{{ t }}</span>
              </div>
            </div>
          </div>
        </section>

        <!-- Projects -->
        <section class="content-block">
          <h2 class="block-title">
            <span class="title-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 19a2 2 0 01-2 2H4a2 2 0 01-2-2V5a2 2 0 012-2h5l2 3h9a2 2 0 012 2z"/></svg>
            </span>
            代表项目
          </h2>
          <div class="project-list">
            <div class="project-card" v-for="project in projects" :key="project.title">
              <div class="project-header">
                <h3 class="project-title">{{ project.title }}</h3>
                <span class="project-badge" v-if="project.badge">{{ project.badge }}</span>
              </div>
              <p class="project-desc">{{ project.desc }}</p>
              <div class="project-tech">
                <span v-for="t in project.tech" :key="t" class="project-tech-tag">{{ t }}</span>
              </div>
            </div>
          </div>
        </section>

        <!-- Awards -->
        <section class="content-block">
          <h2 class="block-title">
            <span class="title-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M12 15l-2 5l2-1l2 1l-2-5z"/><circle cx="12" cy="9" r="6"/></svg>
            </span>
            科研与竞赛
          </h2>
          <div class="award-list">
            <div class="award-item" v-for="award in awards" :key="award">
              <span class="award-dot"></span>
              <span class="award-text">{{ award }}</span>
            </div>
          </div>
        </section>

        <!-- About Blog -->
        <section class="content-block">
          <h2 class="block-title">
            <span class="title-icon">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/></svg>
            </span>
            关于博客
          </h2>
          <p class="block-text">
            这个博客使用 Go-Zero + Vue3 构建，后端采用 Go-Zero rest 框架，前端使用 Naive UI 组件库。
            主要用于记录学习笔记、技术分享和项目经验。如果你对文章内容有任何疑问或建议，欢迎留言交流。
          </p>
        </section>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { marked } from 'marked'
import { getAboutMe } from '@/api'

const aboutContent = ref('')
const renderedAbout = computed(() => {
  if (!aboutContent.value) return ''
  return marked(aboutContent.value)
})

onMounted(async () => {
  try {
    const res = await getAboutMe() as any
    if (res.data?.content) {
      aboutContent.value = res.data.content
    }
  } catch (e) {
    // Use default content
  }
})

const backendTech = [
  'Python', 'Go', 'Java',
  'FastAPI', 'Gin', 'Go-Zero', 'SpringBoot',
  'MySQL', 'Redis', 'Neo4j', 'MongoDB'
]

const aiTech = [
  'LangChain', 'LangGraph', 'GraphRAG', 'Eino', 'SpringAI',
  'LoRA', 'SFT', 'RLHF', 'Agentic-RL',
  'transformers', 'llamafactory', 'verl'
]

const infraTech = [
  'Linux', 'Docker', 'Git', 'Milvus', 'Vite'
]

const projects = [
  {
    title: 'elian — Agent 全链路开发工具',
    badge: '',
    desc: '一个功能强大的 Agent 开发 Python 框架，包括 SDK、CLI 脚手架、数据治理到垂类微调的 Factory、以及 to-Go 工具链。比 LangChain 更稳定，比 LlamaIndex 更易扩展。',
    tech: ['Python', 'Agent', 'LLM', 'CLI']
  },
  {
    title: '基于 Plan-Execute-RePlan 的多源异构数据分析 Agent',
    badge: 'KDD Cup 2026 Stage1-Top1',
    desc: '构建能分解复杂分析问题、协调多步骤推理的数据 Agent。设计了全局 Planner + ReAct 子任务 + RePlanner + 验证 Agent 的完整架构，采用 DAG 并发执行无依赖节点。',
    tech: ['Multi-Agent', 'ReAct', 'Agent-RL', 'verl']
  },
  {
    title: '基于 LangGraph 的图文广告合规审核 Agent',
    badge: '腾讯开悟法律AI大赛 Top1',
    desc: '一站式 AI 合规工作台，集成多 Agent 协作。通过异步改造、缓存优化、Prompt 精简将成本降低 86.67%，单次审核成本从 1 元降至 0.1 元。',
    tech: ['LangGraph', 'Multi-Agent', 'RAG', 'FastAPI']
  },
  {
    title: '基于 Eino 与 Gin 的智能客服 Agent',
    badge: '',
    desc: '以多模态电商智能客服比赛为背景，使用 Go 重构 Python Agent 核心链路，对比高并发场景下两种语言的性能表现与工程成本。',
    tech: ['Go', 'Gin', 'Eino', 'FastAPI', 'Redis']
  }
]

const awards = [
  'KDD Cup 2026 多源异构数据分析 — Stage1 Top1',
  '腾讯开悟法律 AI 大赛 — Top1',
  '2025 华为杯全国研究生数学建模大赛 — 国家级三等奖',
  '2024 华为杯全国研究生数学建模大赛 — 国家级二等奖',
  '江苏省应用数学中心研究生论坛 — 墙报入围',
  '发表论文：V-PEP (Arxiv 2026)、Occ-MoE (IJPRAI 2026)、ATD-TBPS (TVC 2025)'
]
</script>

<style scoped lang="scss">
.about-page {
  animation: fadeInUp 0.5s ease;
}

// ===== Hero Banner =====
.about-hero {
  position: relative;
  width: 100vw;
  margin-left: calc(-50vw + 50%);
  padding: 80px 32px;
  background: linear-gradient(135deg, #0a0a0a 0%, #1a1a1a 100%);
  overflow: hidden;

  @media (max-width: 768px) {
    padding: 60px 20px;
  }
}

.hero-bg {
  position: absolute;
  inset: 0;
}

.hero-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.025) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.025) 1px, transparent 1px);
  background-size: 50px 50px;
}

.hero-inner {
  position: relative;
  max-width: 1100px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 340px;
  gap: 60px;
  align-items: center;

  @media (max-width: 900px) {
    grid-template-columns: 1fr;
    gap: 40px;
  }
}

.hero-left {
  color: white;
}

.hero-badge {
  display: inline-block;
  padding: 6px 16px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.6);
  margin-bottom: 24px;
  letter-spacing: 1px;
}

.hero-name {
  font-size: 48px;
  font-weight: 800;
  margin-bottom: 8px;
  letter-spacing: -0.5px;

  @media (max-width: 640px) {
    font-size: 36px;
  }
}

.hero-role {
  font-size: 16px;
  opacity: 0.7;
  margin-bottom: 20px;
}

.hero-desc {
  font-size: 14px;
  line-height: 1.8;
  opacity: 0.55;
  margin-bottom: 28px;
}

.hero-links {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

.link-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  color: rgba(255, 255, 255, 0.8);
  font-size: 14px;
  text-decoration: none;
  transition: all var(--transition-fast);

  &:hover {
    border-color: rgba(255, 255, 255, 0.5);
    background: rgba(255, 255, 255, 0.08);
    color: white;
  }
}

.info-card {
  background: rgba(255, 255, 255, 0.06);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  padding: 24px;
  backdrop-filter: blur(10px);
}

.info-row {
  display: flex;
  justify-content: space-between;
  padding: 10px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);

  &:last-child {
    border-bottom: none;
  }
}

.info-label {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.4);
}

.info-value {
  font-size: 13px;
  color: rgba(255, 255, 255, 0.8);
  font-weight: 500;
}

// ===== Body Content =====
.about-body {
  max-width: 900px;
  margin: 0 auto;
  padding: 40px 32px 60px;

  @media (max-width: 640px) {
    padding: 30px 20px 40px;
  }
}

.body-inner {
  display: flex;
  flex-direction: column;
  gap: 40px;
}

.content-block {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 28px;
  box-shadow: var(--shadow-sm);
}

.block-title {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 18px;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 20px;
  padding-bottom: 14px;
  border-bottom: 1px solid var(--border-color);
}

.title-icon {
  display: flex;
  align-items: center;
  color: var(--text-secondary);
}

.block-text {
  font-size: 14px;
  line-height: 1.8;
  color: var(--text-secondary);
}

// ===== Tech Stack =====
.tech-sections {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.tech-group-title {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 10px;
}

.tech-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tech-tag {
  display: inline-block;
  padding: 6px 14px;
  background: rgba(0, 0, 0, 0.05);
  color: var(--text-secondary);
  border-radius: 6px;
  font-size: 13px;
  transition: all var(--transition-fast);

  &:hover {
    background: rgba(0, 0, 0, 0.1);
    color: var(--text-primary);
  }
}

html.dark .tech-tag {
  background: rgba(255, 255, 255, 0.06);
  color: var(--text-secondary);

  &:hover {
    background: rgba(255, 255, 255, 0.1);
  }
}

// ===== Projects =====
.project-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.project-card {
  padding: 20px;
  border: 1px solid var(--border-color);
  border-radius: 10px;
  transition: all var(--transition-fast);

  &:hover {
    border-color: rgba(0, 0, 0, 0.15);
    box-shadow: var(--shadow-sm);
  }
}

.project-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

.project-title {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-primary);
}

.project-badge {
  padding: 3px 10px;
  background: rgba(0, 0, 0, 0.08);
  color: var(--text-primary);
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}

.project-desc {
  font-size: 14px;
  line-height: 1.7;
  color: var(--text-secondary);
  margin-bottom: 14px;
}

.project-tech {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.project-tech-tag {
  padding: 3px 10px;
  background: rgba(0, 0, 0, 0.04);
  color: var(--text-muted);
  border-radius: 4px;
  font-size: 12px;
}

// ===== Awards =====
.award-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.award-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
}

.award-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: var(--text-primary);
  margin-top: 8px;
  flex-shrink: 0;
}

.award-text {
  font-size: 14px;
  color: var(--text-secondary);
  line-height: 1.6;
}

.about-dynamic {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 28px;
  box-shadow: var(--shadow-sm);
  font-size: 15px;
  line-height: 1.8;
  color: var(--text-secondary);

  :deep(h1), :deep(h2), :deep(h3) {
    color: var(--text-primary);
    margin: 24px 0 12px;
  }

  :deep(p) {
    margin-bottom: 12px;
  }

  :deep(code) {
    background: rgba(0, 0, 0, 0.06);
    padding: 2px 6px;
    border-radius: 4px;
    font-size: 0.9em;
  }

  :deep(pre) {
    background: rgba(0, 0, 0, 0.06);
    padding: 16px;
    border-radius: 8px;
    overflow-x: auto;
  }

  :deep(ul), :deep(ol) {
    padding-left: 24px;
    margin-bottom: 12px;
  }

  :deep(img) {
    max-width: 100%;
    border-radius: 8px;
  }
}
</style>

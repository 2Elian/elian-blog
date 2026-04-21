<template>
  <div class="about-page">
    <div class="about-hero">
      <div class="hero-bg"></div>
      <div class="hero-content">
        <n-avatar round :size="120" class="hero-avatar">E</n-avatar>
        <h1 class="hero-title">Elian</h1>
        <p class="hero-subtitle">全栈开发者 / 技术博主</p>
        <div class="hero-social">
          <a href="#" class="social-link" title="GitHub">
            <n-icon size="20"><LogoGithub /></n-icon>
          </a>
          <a href="#" class="social-link" title="Email">
            <n-icon size="20"><MailOutline /></n-icon>
          </a>
        </div>
      </div>
    </div>

    <div class="about-content">
      <div class="content-section">
        <h2 class="section-title">关于我</h2>
        <p>
          你好，我是 Elian，一名热爱技术的开发者。这个博客是我的个人空间，用来记录学习笔记、技术分享和生活感悟。
        </p>
        <p>
          我专注于 Web 开发领域，熟悉 Vue、React、Go 等技术栈。我相信技术可以改变世界，也乐于分享自己的经验和见解。
        </p>
      </div>

      <div class="content-section">
        <h2 class="section-title">技术栈</h2>
        <div class="tech-tags">
          <n-tag v-for="tech in techStack" :key="tech" :bordered="false" size="large">
            {{ tech }}
          </n-tag>
        </div>
      </div>

      <div class="content-section">
        <h2 class="section-title">关于博客</h2>
        <p>
          这个博客使用 Vue 3 + Naive UI + Vite 构建，后端使用 Go 语言开发。支持 Markdown 写作、代码高亮、评论互动等功能。
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { NAvatar, NIcon, NTag } from 'naive-ui'
import { LogoGithub, MailOutline } from '@vicons/ionicons5'

const techStack = [
  'Vue 3', 'TypeScript', 'Naive UI', 'Vite',
  'React', 'Node.js', 'Go', 'PostgreSQL',
  'Redis', 'Docker', 'Git'
]
</script>

<style scoped lang="scss">
.about-page {
  animation: fadeInUp 0.5s ease;
  max-width: 800px;
  margin: 0 auto;
}

.about-hero {
  position: relative;
  padding: 60px 40px;
  border-radius: var(--radius-lg);
  overflow: hidden;
  margin-bottom: 30px;
}

.hero-bg {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  z-index: 0;

  &::before {
    content: '';
    position: absolute;
    inset: 0;
    background: radial-gradient(circle at 30% 50%, rgba(255, 255, 255, 0.1), transparent 50%);
  }
}

.hero-content {
  position: relative;
  z-index: 1;
  text-align: center;
  color: white;
}

.hero-avatar {
  background: rgba(255, 255, 255, 0.2) !important;
  backdrop-filter: blur(10px);
  font-size: 48px;
  font-weight: 700;
  margin-bottom: 16px;
}

.hero-title {
  font-size: 32px;
  font-weight: 700;
  margin-bottom: 8px;
}

.hero-subtitle {
  font-size: 16px;
  opacity: 0.9;
  margin-bottom: 20px;
}

.hero-social {
  display: flex;
  justify-content: center;
  gap: 12px;
}

.social-link {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(255, 255, 255, 0.15);
  border-radius: 50%;
  color: white;
  transition: all var(--transition-fast);

  &:hover {
    background: rgba(255, 255, 255, 0.25);
    transform: translateY(-2px);
  }
}

.about-content {
  background: var(--bg-card);
  border-radius: var(--radius-md);
  padding: 30px;
  box-shadow: var(--shadow-sm);
}

.content-section {
  margin-bottom: 30px;

  &:last-child {
    margin-bottom: 0;
  }

  p {
    color: var(--text-secondary);
    line-height: 1.8;
    margin-bottom: 12px;

    &:last-child {
      margin-bottom: 0;
    }
  }
}

.section-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 16px;
  padding-left: 12px;
  border-left: 3px solid var(--primary-color);
}

.tech-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;

  :deep(.n-tag) {
    background: rgba(233, 84, 107, 0.08);
    color: var(--primary-color);
    border-radius: 16px;
    padding: 4px 14px;
  }
}
</style>
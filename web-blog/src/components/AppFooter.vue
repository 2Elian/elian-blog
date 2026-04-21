<template>
  <footer class="app-footer">
    <div class="footer-content">
      <div class="footer-gradient"></div>
      <div class="footer-main">
        <div class="footer-brand">
          <span class="brand-name">Elian Blog</span>
          <p class="brand-desc">记录生活，分享技术</p>
        </div>

        <div class="footer-links">
          <div class="link-group">
            <h4>快速导航</h4>
            <router-link to="/">首页</router-link>
            <router-link to="/blog">博客</router-link>
            <router-link to="/archive">归档</router-link>
          </div>
          <div class="link-group">
            <h4>更多</h4>
            <router-link to="/tags">标签</router-link>
            <router-link to="/friends">友链</router-link>
            <router-link to="/messages">留言</router-link>
          </div>
        </div>
      </div>

      <div class="footer-bottom">
        <p>&copy; {{ currentYear }} Elian Blog. All rights reserved.</p>
        <p>Powered by Vue 3 + Naive UI</p>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>

<style scoped lang="scss">
.app-footer {
  margin-top: auto;
  background: var(--bg-card);
  border-top: 1px solid var(--border-color);
  overflow: hidden;
}

.footer-content {
  position: relative;
}

.footer-gradient {
  height: 3px;
  background: var(--accent-gradient);
}

.footer-main {
  max-width: var(--max-width);
  margin: 0 auto;
  padding: 40px 20px 30px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;

  @media (max-width: 640px) {
    flex-direction: column;
    gap: 30px;
  }
}

.footer-brand {
  .brand-name {
    font-size: 24px;
    font-weight: 700;
    background: var(--accent-gradient);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }

  .brand-desc {
    margin-top: 8px;
    color: var(--text-muted);
    font-size: 14px;
  }
}

.footer-links {
  display: flex;
  gap: 60px;

  @media (max-width: 640px) {
    gap: 40px;
  }
}

.link-group {
  h4 {
    font-size: 14px;
    font-weight: 600;
    color: var(--text-primary);
    margin-bottom: 12px;
  }

  a {
    display: block;
    color: var(--text-muted);
    font-size: 14px;
    margin-bottom: 8px;
    transition: color var(--transition-fast);

    &:hover {
      color: var(--primary-color);
    }
  }
}

.footer-bottom {
  max-width: var(--max-width);
  margin: 0 auto;
  padding: 16px 20px;
  border-top: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  font-size: 13px;
  color: var(--text-muted);

  p {
    margin: 0;
  }

  @media (max-width: 640px) {
    flex-direction: column;
    gap: 8px;
    text-align: center;
  }
}
</style>
<template>
  <footer class="footer">
    <div class="footer-inner container">
      <div class="footer-columns">
        <div class="footer-col">
          <h3>{{ siteName }}</h3>
          <p class="footer-desc">{{ siteDesc }}</p>
        </div>
        <div class="footer-col">
          <h4>快速链接</h4>
          <div class="footer-links">
            <router-link to="/">首页</router-link>
            <router-link to="/blog">博客</router-link>
            <router-link to="/learn">学习</router-link>
            <router-link to="/friend-link">友链</router-link>
          </div>
        </div>
        <div class="footer-col">
          <h4>关于</h4>
          <div class="footer-links">
            <router-link to="/page/about">关于我</router-link>
            <a href="https://github.com" target="_blank">GitHub</a>
          </div>
        </div>
      </div>
      <div class="footer-bottom">
        <p>Copyright &copy; {{ year }} {{ siteName }}. All rights reserved.</p>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { computed } from 'vue'
import { useSiteStore } from '@/store/site'

const siteStore = useSiteStore()
const siteName = computed(() => siteStore.config['site_name'] || 'Elian Blog')
const siteDesc = computed(() => siteStore.config['site_description'] || '个人技术博客')
const year = new Date().getFullYear()
</script>

<style lang="scss" scoped>
.footer {
  background: var(--bg-card);
  border-top: 1px solid var(--border-color);
  padding: 40px 0 24px;
  margin-top: 40px;
}

.footer-columns {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr;
  gap: 40px;
  padding-bottom: 24px;
  border-bottom: 1px solid var(--border-color);
}

.footer-col {
  h3 {
    font-size: 20px;
    margin-bottom: 8px;
    color: var(--text-primary);
  }

  h4 {
    font-size: 16px;
    margin-bottom: 12px;
    color: var(--text-primary);
  }
}

.footer-desc {
  color: var(--text-muted);
  font-size: 14px;
  line-height: 1.8;
}

.footer-links {
  display: flex;
  flex-direction: column;
  gap: 8px;

  a {
    color: var(--text-muted);
    font-size: 14px;
    &:hover { color: var(--primary); }
  }
}

.footer-bottom {
  padding-top: 20px;
  text-align: center;
  color: var(--text-muted);
  font-size: 13px;
}
</style>
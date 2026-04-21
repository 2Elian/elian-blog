import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/BlogLayout.vue'),
    children: [
      { path: '', name: 'Home', component: () => import('@/views/Home.vue') },
      { path: 'blog', name: 'Blog', component: () => import('@/views/Blog.vue') },
      { path: 'article/:id', name: 'ArticleDetail', component: () => import('@/views/ArticleDetail.vue') },
      { path: 'learn', name: 'Learn', component: () => import('@/views/Learn.vue') },
      { path: 'friend-link', name: 'FriendLink', component: () => import('@/views/FriendLink.vue') },
      { path: 'messages', name: 'Messages', component: () => import('@/views/Messages.vue') },
      { path: 'page/:slug', name: 'Page', component: () => import('@/views/Page.vue') },
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  }
})

export default router
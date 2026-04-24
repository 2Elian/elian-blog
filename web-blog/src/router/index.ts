import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('@/layouts/BlogLayout.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
        meta: { title: '首页' }
      },
      {
        path: 'blog',
        name: 'Blog',
        component: () => import('@/views/Blog.vue'),
        meta: { title: '博客' }
      },
      {
        path: 'article/:id',
        name: 'Article',
        component: () => import('@/views/Article.vue'),
        meta: { title: '文章详情' }
      },
      {
        path: 'archive',
        name: 'Archive',
        component: () => import('@/views/Archive.vue'),
        meta: { title: '归档' }
      },
      {
        path: 'tags',
        name: 'Tags',
        component: () => import('@/views/Tags.vue'),
        meta: { title: '标签' }
      },
      {
        path: 'products',
        name: 'Products',
        component: () => import('@/views/Products.vue'),
        meta: { title: '产品' }
      },
      {
        path: 'product/:id',
        name: 'ProductDetail',
        component: () => import('@/views/ProductDetail.vue'),
        meta: { title: '产品详情' }
      },
      {
        path: 'about',
        name: 'About',
        component: () => import('@/views/About.vue'),
        meta: { title: '关于' }
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { title: '登录' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(_to, _from, savedPosition) {
    if (savedPosition) return savedPosition
    return { top: 0, behavior: 'smooth' }
  }
})

router.beforeEach((to, _from, next) => {
  const title = (to.meta.title as string) || 'Elian Blog'
  document.title = `${title} - Elian Blog`

  // 公开页面不需要认证
  const publicPages = ['/', '/login', '/blog', '/archive', '/tags', '/products', '/about']
  if (publicPages.includes(to.path) || to.path.startsWith('/article/') || to.path.startsWith('/product/')) {
    next()
    return
  }

  // 其他页面需要登录
  const userStore = useUserStore()
  if (!userStore.isLoggedIn) {
    next(`/login?redirect=${to.path}`)
    return
  }

  next()
})

export default router

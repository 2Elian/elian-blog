import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import { useUserStore } from '@/stores/user'

const staticRoutes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { title: '登录', hidden: true },
  },
  {
    path: '/',
    component: () => import('@/layouts/AdminLayout.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard.vue'),
        meta: { title: '仪表盘', icon: 'Odometer', roles: ['admin', 'editor'] },
      },
      {
        path: 'article',
        name: 'ArticleList',
        component: () => import('@/views/article/ArticleList.vue'),
        meta: { title: '文章管理', icon: 'Document', roles: ['admin', 'editor'] },
      },
      {
        path: 'category',
        name: 'CategoryList',
        component: () => import('@/views/category/CategoryList.vue'),
        meta: { title: '分类管理', icon: 'FolderOpened', roles: ['admin', 'editor'] },
      },
      {
        path: 'tag',
        name: 'TagList',
        component: () => import('@/views/tag/TagList.vue'),
        meta: { title: '标签管理', icon: 'PriceTag', roles: ['admin', 'editor'] },
      },
      {
        path: 'page',
        name: 'PageList',
        component: () => import('@/views/page/PageList.vue'),
        meta: { title: '页面管理', icon: 'Notebook', roles: ['admin', 'editor'] },
      },
      {
        path: 'comment',
        name: 'CommentList',
        component: () => import('@/views/comment/CommentList.vue'),
        meta: { title: '评论管理', icon: 'ChatDotRound', roles: ['admin', 'editor'] },
      },
      {
        path: 'user',
        name: 'UserList',
        component: () => import('@/views/user/UserList.vue'),
        meta: { title: '用户管理', icon: 'User', roles: ['admin'] },
      },
      {
        path: 'friendlink',
        name: 'FriendLinkList',
        component: () => import('@/views/friendlink/FriendLinkList.vue'),
        meta: { title: '友链管理', icon: 'Link', roles: ['admin', 'editor'] },
      },
      {
        path: 'role',
        name: 'RoleList',
        component: () => import('@/views/role/RoleList.vue'),
        meta: { title: '角色管理', icon: 'UserFilled', roles: ['admin'] },
      },
      {
        path: 'menu',
        name: 'MenuList',
        component: () => import('@/views/menu/MenuList.vue'),
        meta: { title: '菜单管理', icon: 'Menu', roles: ['admin'] },
      },
      {
        path: 'site',
        name: 'SiteConfig',
        component: () => import('@/views/site/SiteConfig.vue'),
        meta: { title: '站点设置', icon: 'Setting', roles: ['admin'] },
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/dashboard',
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes: staticRoutes,
})

// 路由守卫
router.beforeEach((to, _from, next) => {
  document.title = `${to.meta.title || 'Elian Blog'} - Elian Blog Admin`

  const userStore = useUserStore()
  const token = userStore.token

  if (to.path === '/login') {
    if (token) {
      next('/dashboard')
    } else {
      next()
    }
    return
  }

  if (!token) {
    next(`/login?redirect=${to.path}`)
    return
  }

  // RBAC 权限检查
  const requiredRoles = to.meta.roles as string[] | undefined
  if (requiredRoles && requiredRoles.length > 0) {
    const userRole = userStore.role
    if (!userRole || !requiredRoles.includes(userRole)) {
      next('/dashboard')
      return
    }
  }

  next()
})

export default router
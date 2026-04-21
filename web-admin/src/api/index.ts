import { get, post, put, del } from './http'

// ==================== 认证 ====================
export const login = (data: { username: string; password: string }) =>
  post<{ token: string; user: any }>('/login', data)

export const getProfile = () => get<any>('/profile')

// ==================== 文章 ====================
export const getArticleList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/articles', params)

export const getArticle = (id: number) => get<any>(`/articles/${id}`)

export const createArticle = (data: any) => post<any>('/articles', data)

export const updateArticle = (id: number, data: any) => put<any>(`/articles/${id}`, data)

export const deleteArticle = (id: number) => del<any>(`/articles/${id}`)

// ==================== 分类 ====================
export const getCategoryList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/categories', params)

export const getCategory = (id: number) => get<any>(`/categories/${id}`)

export const createCategory = (data: any) => post<any>('/categories', data)

export const updateCategory = (id: number, data: any) => put<any>(`/categories/${id}`, data)

export const deleteCategory = (id: number) => del<any>(`/categories/${id}`)

// ==================== 标签 ====================
export const getTagList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/tags', params)

export const getTag = (id: number) => get<any>(`/tags/${id}`)

export const createTag = (data: any) => post<any>('/tags', data)

export const updateTag = (id: number, data: any) => put<any>(`/tags/${id}`, data)

export const deleteTag = (id: number) => del<any>(`/tags/${id}`)

// ==================== 用户 ====================
export const getUserList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/users', params)

export const deleteUser = (id: number) => del<any>(`/users/${id}`)

// ==================== 评论 ====================
export const getCommentList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/comments', params)

export const updateCommentStatus = (id: number, data: { status: string }) =>
  put<any>(`/comments/${id}/status`, data)

export const deleteComment = (id: number) => del<any>(`/comments/${id}`)

// ==================== 友链 ====================
export const getFriendLinkList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/friend-links', params)

export const getFriendLink = (id: number) => get<any>(`/friend-links/${id}`)

export const createFriendLink = (data: any) => post<any>('/friend-links', data)

export const updateFriendLink = (id: number, data: any) => put<any>(`/friend-links/${id}`, data)

export const deleteFriendLink = (id: number) => del<any>(`/friend-links/${id}`)

// ==================== 页面 ====================
export const getPageList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/pages', params)

export const getPage = (id: number) => get<any>(`/pages/${id}`)

export const createPage = (data: any) => post<any>('/pages', data)

export const updatePage = (id: number, data: any) => put<any>(`/pages/${id}`, data)

export const deletePage = (id: number) => del<any>(`/pages/${id}`)

// ==================== 站点配置 ====================
export const getSiteConfig = () => get<Record<string, any>>('/site/config')

export const setSiteConfig = (data: Record<string, any>) => post<any>('/site/config', data)

// ==================== 角色 ====================
export const getRoleList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/roles', params)

export const getRole = (id: number) => get<any>(`/roles/${id}`)

export const createRole = (data: any) => post<any>('/roles', data)

export const updateRole = (id: number, data: any) => put<any>(`/roles/${id}`, data)

export const deleteRole = (id: number) => del<any>(`/roles/${id}`)

export const assignRoleMenus = (id: number, data: { menu_ids: number[] }) =>
  post<any>(`/roles/${id}/menus`, data)

export const getRoleMenus = (id: number) => get<{ menu_ids: number[] }>(`/roles/${id}/menus`)

// ==================== 菜单 ====================
export const getMenuList = (params?: Record<string, any>) =>
  get<{ list: any[] }>('/menus', params)

export const getMenu = (id: number) => get<any>(`/menus/${id}`)

export const createMenu = (data: any) => post<any>('/menus', data)

export const updateMenu = (id: number, data: any) => put<any>(`/menus/${id}`, data)

export const deleteMenu = (id: number) => del<any>(`/menus/${id}`)

// ==================== 留言 ====================
export const getMessageList = (params?: Record<string, any>) =>
  get<{ list: any[]; total: number }>('/messages', params)

export const deleteMessage = (id: number) => del<any>(`/messages/${id}`)

// ==================== 仪表盘统计 ====================
export const getDashboardStats = () => get<{
  article_count: number
  user_count: number
  comment_count: number
  view_count: number
}>('/dashboard/stats')
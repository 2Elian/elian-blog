import http from './http'

// 文章
export const getArticles = (params) => http.get('/blog-api/v1/articles', { params })
export const getArticle = (id) => http.get(`/blog-api/v1/articles/${id}`)
export const searchArticles = (params) => http.get('/blog-api/v1/articles/search', { params })

// 分类 & 标签
export const getCategories = () => http.get('/blog-api/v1/categories')
export const getTags = () => http.get('/blog-api/v1/tags')

// 评论
export const getComments = (articleId, params) => http.get(`/blog-api/v1/articles/${articleId}/comments`, { params })
export const getRecentComments = (params) => http.get('/blog-api/v1/comments/recent', { params })
export const createComment = (data) => http.post('/blog-api/v1/comments', data)

// 友链
export const getFriendLinks = () => http.get('/blog-api/v1/friend-links')

// 说说/留言
export const getMessages = (params) => http.get('/blog-api/v1/messages', { params })
export const createMessage = (data) => http.post('/blog-api/v1/messages', data)

// 页面
export const getPages = () => http.get('/blog-api/v1/pages')
export const getPageBySlug = (slug) => http.get(`/blog-api/v1/pages/${slug}`)

// 站点配置
export const getSiteConfig = () => http.get('/blog-api/v1/site/config')

// 认证
export const login = (data) => http.post('/blog-api/v1/login', data)
export const register = (data) => http.post('/blog-api/v1/register', data)
export const getUserInfo = () => http.get('/blog-api/v1/user/info')
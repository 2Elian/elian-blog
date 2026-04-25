import http from './http'

// Auth
export const login = (data: { username: string; password: string }) =>
  http.post('/login', data)

export const register = (data: {
  username: string
  password: string
  email: string
  avatar?: string
  intro?: string
  website?: string
}) => http.post('/register', data)

// User
export const getUserInfo = () => http.get('/user/info')

// Articles
export const getArticles = (params?: { page?: number; page_size?: number; category_id?: number; tag_id?: number }) =>
  http.get('/articles', { params })

export const getArticle = (id: number) => http.get(`/articles/${id}`)

export const searchArticles = (params: { keyword: string; page?: number; page_size?: number }) =>
  http.get('/articles/search', { params })

// Categories & Tags
export const getCategories = () => http.get('/categories')
export const getTags = () => http.get('/tags')

// Comments
export const getComments = (articleId: number) => http.get(`/articles/${articleId}/comments`)
export const getRecentComments = () => http.get('/comments/recent')
export const postComment = (data: { article_id: number; content: string; parent_id?: number }) =>
  http.post('/comments', data)

// Products
export const getProducts = (params?: { page?: number; page_size?: number }) =>
  http.get('/products', { params })

export const getProduct = (id: number) => http.get(`/products/${id}`)

// Pages
export const getPages = () => http.get('/pages')
export const getPage = (slug: string) => http.get(`/pages/${slug}`)

// Friend Links
export const getFriendLinks = () => http.get('/friend-links')

// Messages
export const getMessages = (params?: { page?: number; page_size?: number }) =>
  http.get('/messages', { params })

export const postMessage = (data: { content: string }) =>
  http.post('/messages', data)

// Site Config
export const getSiteConfig = () => http.get('/site/config')

// About Me
export const getAboutMe = () => http.get('/site/about')

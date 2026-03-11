import http from './http.js'

export function apiListNews({ page = 1, pageSize = 9, tag = '' } = {}) {
  const params = { page, page_size: pageSize }
  if (tag && tag !== '全部') params.tag = tag
  return http.get('/news', { params })
}

export function apiGetNews(id) {
  return http.get(`/news/${id}`)
}

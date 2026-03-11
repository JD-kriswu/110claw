import http from './http.js'

export function apiListSkills({ page = 1, pageSize = 100, tag = '', sort = 'rating' } = {}) {
  const params = { page, page_size: pageSize, sort }
  if (tag) params.tag = tag
  return http.get('/skills', { params })
}

export function apiGetSkill(id) {
  return http.get(`/skills/${id}`)
}

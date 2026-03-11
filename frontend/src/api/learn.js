import http from './http.js'

export function apiListLearn() {
  return http.get('/learn')
}

export function apiGetLearnDay(day) {
  return http.get(`/learn/${day}`)
}

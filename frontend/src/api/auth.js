import http from './http.js'

export function apiLogin(username, password, rememberMe = false) {
  return http.post('/auth/login', { username, password, remember_me: rememberMe })
}

export function apiRegister(username, password, phone = '') {
  return http.post('/auth/register', { username, password, phone })
}

export function apiLogout() {
  return http.post('/auth/logout')
}

export function apiMe() {
  return http.get('/auth/me')
}

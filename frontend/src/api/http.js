import axios from 'axios'

const http = axios.create({
  baseURL: '/110claw/api/v1',
  withCredentials: true,
  timeout: 10000,
})

http.interceptors.response.use(
  res => res.data,
  err => {
    const msg = err.response?.data?.error || err.message || '请求失败'
    return Promise.reject(new Error(msg))
  }
)

export default http

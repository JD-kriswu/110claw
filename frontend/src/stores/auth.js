import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { apiLogout, apiMe } from '@/api/auth.js'

export const useAuthStore = defineStore('auth', () => {
  const user = ref(JSON.parse(localStorage.getItem('110claw_user') || 'null'))

  const isLoggedIn = computed(() => !!user.value)

  function setUser(userData) {
    user.value = userData
    localStorage.setItem('110claw_user', JSON.stringify(userData))
  }

  function clearUser() {
    user.value = null
    localStorage.removeItem('110claw_user')
  }

  async function logout() {
    try { await apiLogout() } catch (_) {}
    clearUser()
  }

  async function fetchMe() {
    try {
      const data = await apiMe()
      setUser(data)
      return data
    } catch (_) {
      clearUser()
      return null
    }
  }

  return { user, isLoggedIn, setUser, clearUser, logout, fetchMe }
})

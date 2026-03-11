<template>
  <header class="sticky top-0 z-50 bg-white/80 backdrop-blur border-b border-gray-200">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <!-- Logo -->
        <RouterLink to="/" class="flex items-center gap-2">
          <div class="w-8 h-8 bg-gradient-to-br from-blue-600 to-purple-600 rounded-lg flex items-center justify-center flex-shrink-0">
            <span class="text-white font-bold text-xs">110</span>
          </div>
          <span class="font-bold text-lg text-gray-900">110Claw</span>
        </RouterLink>

        <!-- Nav links (desktop) -->
        <nav class="hidden md:flex items-center gap-1">
          <RouterLink
            v-for="link in navLinks"
            :key="link.path"
            :to="link.path"
            class="px-4 py-2 text-sm font-medium text-gray-600 hover:text-gray-900 rounded-lg hover:bg-gray-100 transition-colors"
            active-class="text-blue-600 bg-blue-50 hover:bg-blue-50 hover:text-blue-600"
            :exact="link.path === '/'"
          >
            {{ link.label }}
          </RouterLink>
        </nav>

        <!-- Auth area -->
        <div class="flex items-center gap-2">
          <!-- Logged in: avatar + dropdown -->
          <div v-if="auth.isLoggedIn" class="relative" ref="dropdownRef">
            <button
              @click="dropdownOpen = !dropdownOpen"
              class="flex items-center gap-2 px-3 py-2 rounded-lg hover:bg-gray-100 transition-colors"
            >
              <div class="w-7 h-7 bg-gradient-to-br from-blue-500 to-purple-600 rounded-full flex items-center justify-center">
                <span class="text-white text-xs font-bold">{{ userInitial }}</span>
              </div>
              <span class="hidden sm:block text-sm font-medium text-gray-700">{{ auth.user?.username }}</span>
              <ChevronDown class="w-3.5 h-3.5 text-gray-400"/>
            </button>

            <!-- Dropdown -->
            <div
              v-if="dropdownOpen"
              class="absolute right-0 mt-1 w-48 bg-white rounded-xl shadow-lg border border-gray-200 py-1 z-50"
            >
              <div class="px-4 py-2 border-b border-gray-100">
                <p class="text-sm font-medium text-gray-900 truncate">{{ auth.user?.username }}</p>
                <p class="text-xs text-gray-500 truncate">{{ auth.user?.email }}</p>
              </div>
              <button
                @click="handleLogout"
                class="w-full flex items-center gap-2 px-4 py-2 text-sm text-red-600 hover:bg-red-50 transition-colors"
              >
                <LogOut class="w-4 h-4"/>
                退出登录
              </button>
            </div>
          </div>

          <!-- Not logged in -->
          <template v-else>
            <button
              @click="openAuth('login')"
              class="hidden sm:flex items-center gap-1 px-4 py-2 text-sm font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors"
            >
              <LogIn class="w-4 h-4"/>
              登录
            </button>
            <button
              @click="openAuth('register')"
              class="px-4 py-2 text-sm font-medium text-white bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg hover:opacity-90 transition-opacity"
            >
              注册
            </button>
          </template>

          <!-- Mobile menu button -->
          <button class="md:hidden p-2 text-gray-500 hover:text-gray-900" @click="mobileOpen = !mobileOpen">
            <Menu v-if="!mobileOpen" class="w-5 h-5"/>
            <X v-else class="w-5 h-5"/>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile menu -->
    <div v-if="mobileOpen" class="md:hidden border-t border-gray-100 bg-white">
      <div class="px-4 py-2 space-y-1">
        <RouterLink
          v-for="link in navLinks"
          :key="link.path"
          :to="link.path"
          class="block px-3 py-2 text-sm font-medium text-gray-600 hover:text-gray-900 rounded-lg hover:bg-gray-100"
          @click="mobileOpen = false"
        >
          {{ link.label }}
        </RouterLink>
        <div v-if="!auth.isLoggedIn" class="pt-2 border-t border-gray-100 flex gap-2">
          <button @click="openAuth('login'); mobileOpen = false" class="flex-1 py-2 text-sm font-medium text-gray-700 border border-gray-200 rounded-lg hover:bg-gray-50">登录</button>
          <button @click="openAuth('register'); mobileOpen = false" class="flex-1 py-2 text-sm font-medium text-white bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg">注册</button>
        </div>
      </div>
    </div>
  </header>

  <!-- Auth modal -->
  <AuthModal v-model="authModalOpen" :initial-tab="authTab"/>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { LogIn, LogOut, Menu, X, ChevronDown } from 'lucide-vue-next'
import { useAuthStore } from '@/stores/auth.js'
import AuthModal from '@/components/AuthModal.vue'

const auth = useAuthStore()
const mobileOpen = ref(false)
const dropdownOpen = ref(false)
const dropdownRef = ref(null)
const authModalOpen = ref(false)
const authTab = ref('login')

const navLinks = [
  { path: '/', label: '首页' },
  { path: '/news', label: '资讯' },
  { path: '/learn', label: '7天学习计划' },
  { path: '/skills', label: 'Skills' },
]

const userInitial = computed(() => {
  const u = auth.user?.username || auth.user?.email || '?'
  return u.charAt(0).toUpperCase()
})

function openAuth(tab) {
  authTab.value = tab
  authModalOpen.value = true
}

async function handleLogout() {
  await auth.logout()
  dropdownOpen.value = false
}

function onOutsideClick(e) {
  if (dropdownRef.value && !dropdownRef.value.contains(e.target)) {
    dropdownOpen.value = false
  }
}

onMounted(() => document.addEventListener('click', onOutsideClick))
onBeforeUnmount(() => document.removeEventListener('click', onOutsideClick))
</script>

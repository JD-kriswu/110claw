<template>
  <!-- Backdrop -->
  <Teleport to="body">
    <div
      v-if="modelValue"
      class="fixed inset-0 z-50 flex items-center justify-center p-4"
      @click.self="$emit('update:modelValue', false)"
    >
      <div class="absolute inset-0 bg-black/40 backdrop-blur-sm"></div>

      <!-- Modal -->
      <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-md overflow-hidden">
        <!-- Close -->
        <button
          @click="$emit('update:modelValue', false)"
          class="absolute top-4 right-4 p-1.5 text-gray-400 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors z-10"
        >
          <X class="w-4 h-4"/>
        </button>

        <!-- Logo header -->
        <div class="px-8 pt-8 pb-6 text-center border-b border-gray-100">
          <div class="w-12 h-12 bg-gradient-to-br from-blue-600 to-purple-600 rounded-xl flex items-center justify-center mx-auto mb-3">
            <span class="text-white font-bold text-sm">110</span>
          </div>
          <h2 class="text-xl font-bold text-gray-900">欢迎来到 110Claw</h2>
          <p class="text-sm text-gray-500 mt-1">免费安全的 Claw 社区平台</p>
        </div>

        <!-- Tabs -->
        <div class="flex border-b border-gray-100">
          <button
            @click="activeTab = 'login'"
            :class="[
              'flex-1 py-3 text-sm font-medium transition-colors',
              activeTab === 'login'
                ? 'text-blue-600 border-b-2 border-blue-600'
                : 'text-gray-500 hover:text-gray-700'
            ]"
          >登录</button>
          <button
            @click="activeTab = 'register'"
            :class="[
              'flex-1 py-3 text-sm font-medium transition-colors',
              activeTab === 'register'
                ? 'text-blue-600 border-b-2 border-blue-600'
                : 'text-gray-500 hover:text-gray-700'
            ]"
          >注册</button>
        </div>

        <!-- Form -->
        <div class="p-8">
          <!-- Error message -->
          <div v-if="errorMsg" class="mb-4 px-4 py-3 bg-red-50 border border-red-200 rounded-lg text-sm text-red-700">
            {{ errorMsg }}
          </div>

          <form @submit.prevent="submit" class="space-y-4">
            <!-- Username -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">用户名</label>
              <input
                v-model="form.username"
                type="text"
                placeholder="3-50 位字母数字"
                required
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            <!-- Phone (register only) -->
            <div v-if="activeTab === 'register'">
              <label class="block text-sm font-medium text-gray-700 mb-1">手机号（选填）</label>
              <input
                v-model="form.phone"
                type="tel"
                placeholder="请输入手机号"
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            <!-- Password -->
            <div>
              <div class="flex items-center justify-between mb-1">
                <label class="block text-sm font-medium text-gray-700">密码</label>
                <button
                  v-if="activeTab === 'login'"
                  type="button"
                  class="text-xs text-blue-600 hover:underline"
                >忘记密码？</button>
              </div>
              <div class="relative">
                <input
                  v-model="form.password"
                  :type="showPassword ? 'text' : 'password'"
                  placeholder="至少 6 位"
                  required
                  class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent pr-10"
                />
                <button
                  type="button"
                  @click="showPassword = !showPassword"
                  class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
                >
                  <Eye v-if="!showPassword" class="w-4 h-4"/>
                  <EyeOff v-else class="w-4 h-4"/>
                </button>
              </div>
            </div>

            <!-- Confirm password (register only) -->
            <div v-if="activeTab === 'register'">
              <label class="block text-sm font-medium text-gray-700 mb-1">确认密码</label>
              <input
                v-model="form.confirmPassword"
                type="password"
                placeholder="再次输入密码"
                required
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            <!-- Submit -->
            <button
              type="submit"
              :disabled="loading"
              class="w-full py-3 bg-gradient-to-r from-blue-600 to-purple-600 text-white rounded-lg font-semibold text-sm hover:opacity-90 transition-opacity disabled:opacity-60"
            >
              <span v-if="loading">处理中...</span>
              <span v-else>{{ activeTab === 'login' ? '登录' : '创建账户' }}</span>
            </button>
          </form>

          <!-- Switch tab hint -->
          <p class="mt-4 text-center text-sm text-gray-500">
            <span v-if="activeTab === 'login'">
              还没有账号？
              <button @click="activeTab = 'register'" class="text-blue-600 hover:underline font-medium">立即注册</button>
            </span>
            <span v-else>
              已有账号？
              <button @click="activeTab = 'login'" class="text-blue-600 hover:underline font-medium">立即登录</button>
            </span>
          </p>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, watch } from 'vue'
import { X, Eye, EyeOff } from 'lucide-vue-next'
import { useAuthStore } from '@/stores/auth.js'
import { apiLogin, apiRegister } from '@/api/auth.js'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  initialTab: { type: String, default: 'login' },
})
const emit = defineEmits(['update:modelValue'])

const auth = useAuthStore()
const activeTab = ref(props.initialTab)
const showPassword = ref(false)
const loading = ref(false)
const errorMsg = ref('')

const form = ref({ username: '', phone: '', password: '', confirmPassword: '' })

watch(() => props.modelValue, (val) => {
  if (val) {
    activeTab.value = props.initialTab
    form.value = { username: '', phone: '', password: '', confirmPassword: '' }
    errorMsg.value = ''
  }
})

watch(activeTab, () => {
  errorMsg.value = ''
  form.value = { username: '', phone: '', password: '', confirmPassword: '' }
})

async function submit() {
  errorMsg.value = ''

  if (activeTab.value === 'register') {
    if (form.value.password !== form.value.confirmPassword) {
      errorMsg.value = '两次输入的密码不一致'
      return
    }
    if (form.value.password.length < 6) {
      errorMsg.value = '密码至少需要 6 位'
      return
    }
  }

  loading.value = true
  try {
    if (activeTab.value === 'login') {
      const res = await apiLogin(form.value.username, form.value.password)
      auth.setUser(res.user)
    } else {
      await apiRegister(form.value.username, form.value.password, form.value.phone)
      const res = await apiLogin(form.value.username, form.value.password)
      auth.setUser(res.user)
    }
    emit('update:modelValue', false)
  } catch (err) {
    errorMsg.value = err.message
  } finally {
    loading.value = false
  }
}
</script>

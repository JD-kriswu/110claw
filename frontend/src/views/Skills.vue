<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <h1 class="text-3xl font-bold mb-2">Skills 插件</h1>
        <p class="text-gray-600">探索社区最受欢迎的插件，扩展 110Claw 的能力</p>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Search -->
      <div class="flex gap-4 mb-8">
        <div class="relative flex-1 max-w-md">
          <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-gray-400"/>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="搜索 Skills..."
            class="w-full pl-10 pr-4 py-2.5 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
          />
        </div>
        <button
          v-if="searchQuery"
          @click="searchQuery = ''"
          class="px-4 py-2 text-sm text-gray-500 hover:text-gray-700 border border-gray-200 rounded-lg hover:bg-gray-50"
        >清除</button>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div v-for="i in 6" :key="i" class="bg-white border border-gray-200 rounded-xl p-5 animate-pulse">
          <div class="flex items-start gap-3">
            <div class="w-11 h-11 bg-gray-200 rounded-xl flex-shrink-0"></div>
            <div class="flex-1 space-y-2">
              <div class="h-4 bg-gray-200 rounded w-1/2"></div>
              <div class="h-3 bg-gray-100 rounded w-full"></div>
              <div class="h-3 bg-gray-100 rounded w-4/5"></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="text-center py-20">
        <p class="text-gray-500 mb-2">加载失败</p>
        <button @click="fetchSkills" class="text-green-600 hover:underline text-sm">重试</button>
      </div>

      <template v-else>
        <!-- Stats -->
        <div class="flex items-center gap-2 mb-6 text-sm text-gray-500">
          <span>共 {{ filteredSkills.length }} 个 Skills</span>
        </div>

        <!-- Empty -->
        <div v-if="filteredSkills.length === 0" class="text-center py-20 text-gray-500">
          未找到匹配的 Skills
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
          <RouterLink
            v-for="skill in filteredSkills"
            :key="skill.id"
            :to="`/skills/${skill.id}`"
            class="group bg-white border border-gray-200 rounded-xl p-5 hover:shadow-md transition-all hover:border-green-300"
          >
            <div class="flex items-start gap-3">
              <!-- Avatar -->
              <div class="w-11 h-11 bg-gradient-to-br from-blue-100 to-purple-100 rounded-xl flex items-center justify-center flex-shrink-0">
                <span class="text-xl font-bold bg-gradient-to-br from-blue-600 to-purple-600 bg-clip-text text-transparent">
                  {{ skill.name.charAt(0) }}
                </span>
              </div>

              <div class="flex-1 min-w-0">
                <h3 class="font-semibold text-base mb-1 group-hover:text-green-700 transition-colors">{{ skill.name }}</h3>
                <p class="text-sm text-gray-600 mb-3 line-clamp-2">{{ skill.description }}</p>

                <!-- Meta -->
                <div class="flex items-center gap-3 text-xs text-gray-500">
                  <div class="flex items-center gap-1">
                    <Star class="w-3.5 h-3.5 fill-current text-yellow-400"/>
                    <span>{{ skill.rating?.toFixed(1) || '—' }}</span>
                  </div>
                  <div class="flex items-center gap-1">
                    <Download class="w-3.5 h-3.5"/>
                    <span>{{ formatNum(skill.download_count) }}</span>
                  </div>
                  <span v-if="skill.version" class="ml-auto px-2 py-0.5 bg-gray-100 rounded text-gray-600 font-mono">v{{ skill.version }}</span>
                </div>
              </div>
            </div>
          </RouterLink>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Search, Star, Download } from 'lucide-vue-next'
import { apiListSkills } from '@/api/skills.js'

const allSkills = ref([])
const searchQuery = ref('')
const loading = ref(true)
const error = ref(false)

const filteredSkills = computed(() => {
  if (!searchQuery.value.trim()) return allSkills.value
  const q = searchQuery.value.trim().toLowerCase()
  return allSkills.value.filter(s =>
    s.name.toLowerCase().includes(q) ||
    s.description.toLowerCase().includes(q) ||
    (s.tags || '').toLowerCase().includes(q)
  )
})

async function fetchSkills() {
  loading.value = true
  error.value = false
  try {
    const res = await apiListSkills({ pageSize: 100 })
    allSkills.value = res.data || []
  } catch (_) {
    error.value = true
  } finally {
    loading.value = false
  }
}

onMounted(fetchSkills)

function formatNum(n) {
  if (!n) return '0'
  if (n >= 10000) return (n / 10000).toFixed(1) + 'w'
  if (n >= 1000) return (n / 1000).toFixed(1) + 'k'
  return String(n)
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <h1 class="text-3xl font-bold mb-2">资讯</h1>
        <p class="text-gray-600">了解 Claw 的最新动态和更新</p>
      </div>
    </div>

    <!-- Tag filter -->
    <div class="bg-white border-b border-gray-200 sticky top-16 z-30">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex gap-2 overflow-x-auto py-3 scrollbar-hide">
          <button
            v-for="tag in newsTags"
            :key="tag"
            @click="selectTag(tag)"
            :class="[
              'flex-shrink-0 px-4 py-1.5 rounded-full text-sm font-medium transition-colors',
              activeTag === tag
                ? 'bg-purple-600 text-white'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]"
          >{{ tag }}</button>
        </div>
      </div>
    </div>

    <!-- News grid -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <!-- Loading skeleton -->
      <div v-if="loading && items.length === 0" class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="i in 6" :key="i" class="bg-white border border-gray-200 rounded-xl p-6 animate-pulse">
          <div class="flex gap-2 mb-3">
            <div class="h-5 w-16 bg-gray-200 rounded-full"></div>
            <div class="h-5 w-12 bg-gray-200 rounded-full ml-auto"></div>
          </div>
          <div class="h-4 bg-gray-200 rounded mb-2 w-full"></div>
          <div class="h-4 bg-gray-200 rounded mb-4 w-3/4"></div>
          <div class="h-3 bg-gray-100 rounded w-1/3"></div>
        </div>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="text-center py-20">
        <p class="text-gray-500 mb-2">加载失败</p>
        <button @click="fetchNews(true)" class="text-purple-600 hover:underline text-sm">重试</button>
      </div>

      <!-- Empty -->
      <div v-else-if="!loading && items.length === 0" class="text-center py-20 text-gray-500">
        该分类暂无资讯
      </div>

      <div v-else class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
        <RouterLink
          v-for="item in items"
          :key="item.id"
          :to="`/news/${item.id}`"
          class="group bg-white border border-gray-200 rounded-xl p-6 hover:shadow-md transition-all hover:border-purple-300"
        >
          <div class="flex items-center gap-2 mb-3">
            <span
              v-for="tag in parseTags(item.tags).slice(0, 2)"
              :key="tag"
              :class="['px-2 py-0.5 text-xs rounded-full font-medium', tagColor(tag)]"
            >{{ tag }}</span>
            <span class="ml-auto text-xs text-gray-400 flex-shrink-0">{{ formatDate(item.published_at || item.created_at) }}</span>
          </div>
          <h3 class="font-semibold text-base mb-2 line-clamp-2 group-hover:text-purple-700 transition-colors">{{ item.title }}</h3>
          <p class="text-sm text-gray-600 line-clamp-3">{{ item.summary }}</p>
        </RouterLink>
      </div>

      <!-- Load more -->
      <div v-if="items.length < total" class="mt-10 text-center">
        <button
          @click="loadMore"
          :disabled="loadingMore"
          class="px-8 py-3 border border-gray-300 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-50 transition-colors disabled:opacity-50"
        >
          {{ loadingMore ? '加载中...' : '查看更多' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { apiListNews } from '@/api/news.js'

const newsTags = ['全部', '版本发布', '安全', '教程', '生态', '技术', '社区', '隐私', '工具集成', '开源', '自动化', '移动端', '性能优化', 'ClawHub', '架构', '推荐', '本地模型', '开发效率']
const PAGE_SIZE = 9

const activeTag = ref('全部')
const items = ref([])
const total = ref(0)
const page = ref(1)
const loading = ref(false)
const loadingMore = ref(false)
const error = ref(false)

async function fetchNews(reset = false) {
  if (reset) {
    items.value = []
    page.value = 1
    total.value = 0
  }
  error.value = false
  loading.value = items.value.length === 0
  loadingMore.value = items.value.length > 0
  try {
    const res = await apiListNews({ page: page.value, pageSize: PAGE_SIZE, tag: activeTag.value })
    items.value = reset ? res.data : [...items.value, ...res.data]
    total.value = res.total
  } catch (_) {
    error.value = true
  } finally {
    loading.value = false
    loadingMore.value = false
  }
}

function selectTag(tag) {
  activeTag.value = tag
  fetchNews(true)
}

function loadMore() {
  page.value++
  fetchNews(false)
}

onMounted(() => fetchNews(true))

function parseTags(tagsStr) {
  if (!tagsStr) return []
  if (tagsStr.startsWith('[')) {
    try { return JSON.parse(tagsStr) } catch (_) {}
  }
  return tagsStr.split(',').map(t => t.trim()).filter(Boolean)
}

function formatDate(dateStr) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleDateString('zh-CN', { month: 'long', day: 'numeric' })
}

const tagColorMap = {
  '版本发布': 'bg-blue-50 text-blue-700',
  '安全': 'bg-red-50 text-red-700',
  '教程': 'bg-green-50 text-green-700',
  '生态': 'bg-purple-50 text-purple-700',
  '技术': 'bg-indigo-50 text-indigo-700',
  '社区': 'bg-orange-50 text-orange-700',
  '隐私': 'bg-teal-50 text-teal-700',
  '工具集成': 'bg-cyan-50 text-cyan-700',
  '开源': 'bg-lime-50 text-lime-700',
  '自动化': 'bg-yellow-50 text-yellow-700',
  '移动端': 'bg-pink-50 text-pink-700',
  '性能优化': 'bg-blue-50 text-blue-700',
  'ClawHub': 'bg-purple-50 text-purple-700',
  '架构': 'bg-indigo-50 text-indigo-700',
  '推荐': 'bg-orange-50 text-orange-700',
  '本地模型': 'bg-teal-50 text-teal-700',
  '开发效率': 'bg-cyan-50 text-cyan-700',
}

function tagColor(tag) {
  return tagColorMap[tag] || 'bg-gray-100 text-gray-600'
}
</script>

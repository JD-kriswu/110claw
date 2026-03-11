<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Loading -->
    <div v-if="loading" class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <div class="h-4 w-24 bg-gray-200 rounded animate-pulse mb-8"></div>
      <div class="bg-white rounded-2xl border border-gray-200 overflow-hidden animate-pulse">
        <div class="p-8 border-b border-gray-100">
          <div class="flex gap-2 mb-4">
            <div class="h-6 w-16 bg-gray-200 rounded-full"></div>
          </div>
          <div class="h-8 bg-gray-200 rounded mb-2 w-full"></div>
          <div class="h-8 bg-gray-200 rounded mb-4 w-2/3"></div>
          <div class="h-4 bg-gray-100 rounded w-1/3"></div>
        </div>
        <div class="p-8 space-y-3">
          <div class="h-4 bg-gray-100 rounded w-full"></div>
          <div class="h-4 bg-gray-100 rounded w-5/6"></div>
          <div class="h-4 bg-gray-100 rounded w-4/5"></div>
        </div>
      </div>
    </div>

    <!-- Article -->
    <div v-else-if="article" class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <!-- Back -->
      <RouterLink to="/news" class="inline-flex items-center gap-1 text-sm text-gray-500 hover:text-gray-900 mb-8 transition-colors">
        <ChevronLeft class="w-4 h-4"/>
        返回资讯列表
      </RouterLink>

      <!-- Article card -->
      <article class="bg-white rounded-2xl border border-gray-200 overflow-hidden">
        <!-- Header -->
        <div class="p-8 border-b border-gray-100">
          <div class="flex flex-wrap gap-2 mb-4">
            <span
              v-for="tag in parseTags(article.tags)"
              :key="tag"
              :class="['px-3 py-1 text-xs rounded-full font-medium', tagColor(tag)]"
            >{{ tag }}</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-bold mb-4 leading-snug">{{ article.title }}</h1>
          <div class="flex items-center gap-3 text-sm text-gray-500">
            <span>{{ formatDate(article.published_at || article.created_at) }}</span>
            <span v-if="article.view_count" class="text-gray-300">·</span>
            <span v-if="article.view_count">{{ article.view_count }} 次阅读</span>
          </div>
        </div>

        <!-- Summary -->
        <div class="px-8 py-4 bg-purple-50 border-b border-purple-100">
          <p class="text-purple-800 text-sm leading-relaxed font-medium">{{ article.summary }}</p>
        </div>

        <!-- Body (rendered markdown) -->
        <div class="p-8 prose prose-gray max-w-none" v-html="renderedContent"></div>
      </article>

      <!-- Related -->
      <div v-if="related.length > 0" class="mt-10">
        <h2 class="text-lg font-semibold mb-4">相关资讯</h2>
        <div class="grid sm:grid-cols-2 gap-4">
          <RouterLink
            v-for="item in related"
            :key="item.id"
            :to="`/news/${item.id}`"
            class="bg-white border border-gray-200 rounded-xl p-4 hover:shadow-sm hover:border-purple-300 transition-all"
          >
            <div class="flex gap-2 mb-2">
              <span
                v-for="tag in parseTags(item.tags).slice(0, 1)"
                :key="tag"
                :class="['px-2 py-0.5 text-xs rounded-full', tagColor(tag)]"
              >{{ tag }}</span>
            </div>
            <h3 class="text-sm font-medium line-clamp-2">{{ item.title }}</h3>
          </RouterLink>
        </div>
      </div>
    </div>

    <!-- 404 -->
    <div v-else class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-20 text-center">
      <p class="text-gray-500 mb-4">文章未找到</p>
      <RouterLink to="/news" class="text-purple-600 hover:underline">返回资讯列表</RouterLink>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ChevronLeft } from 'lucide-vue-next'
import { apiGetNews, apiListNews } from '@/api/news.js'

const route = useRoute()
const article = ref(null)
const related = ref([])
const loading = ref(true)

async function fetchArticle(id) {
  loading.value = true
  article.value = null
  related.value = []
  try {
    article.value = await apiGetNews(id)
    fetchRelated()
  } catch (_) {
    // article stays null → 404 shown
  } finally {
    loading.value = false
  }
}

async function fetchRelated() {
  if (!article.value) return
  const firstTag = parseTags(article.value.tags)[0]
  if (!firstTag) return
  try {
    const res = await apiListNews({ page: 1, pageSize: 5, tag: firstTag })
    related.value = (res.data || []).filter(n => n.id !== article.value.id).slice(0, 4)
  } catch (_) {}
}

watch(() => route.params.id, id => { if (id) fetchArticle(id) })
onMounted(() => fetchArticle(route.params.id))

function parseTags(tagsStr) {
  if (!tagsStr) return []
  if (tagsStr.startsWith('[')) {
    try { return JSON.parse(tagsStr) } catch (_) {}
  }
  return tagsStr.split(',').map(t => t.trim()).filter(Boolean)
}

function renderMarkdown(md) {
  if (!md) return ''
  return md
    .replace(/^### (.+)$/gm, '<h3 class="text-lg font-bold mt-6 mb-2">$1</h3>')
    .replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold mt-8 mb-3">$1</h2>')
    .replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold mt-8 mb-4">$1</h1>')
    .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
    .replace(/`([^`\n]+)`/g, '<code class="px-1.5 py-0.5 bg-gray-100 rounded text-sm font-mono text-purple-700">$1</code>')
    .replace(/```[\w]*\n([\s\S]*?)```/g, '<pre class="bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto my-4 text-sm font-mono"><code>$1</code></pre>')
    .replace(/^\|(.+)\|$/gm, (line) => {
      if (line.includes('---')) return ''
      const cells = line.split('|').filter((_, i, a) => i > 0 && i < a.length - 1)
      return '<tr>' + cells.map(c => `<td class="border border-gray-200 px-3 py-1.5 text-sm">${c.trim()}</td>`).join('') + '</tr>'
    })
    .replace(/^- (.+)$/gm, '<li class="ml-4 list-disc text-gray-700">$1</li>')
    .replace(/^\d+\. (.+)$/gm, '<li class="ml-4 list-decimal text-gray-700">$1</li>')
    .replace(/\n\n/g, '</p><p class="my-3 text-gray-700 leading-relaxed">')
    .replace(/^(?!<)(.+)$/gm, '$1')
}

const renderedContent = computed(() => {
  if (!article.value) return ''
  return `<p class="my-3 text-gray-700 leading-relaxed">${renderMarkdown(article.value.content)}</p>`
})

function formatDate(dateStr) {
  if (!dateStr) return ''
  return new Date(dateStr).toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' })
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
}

function tagColor(tag) {
  return tagColorMap[tag] || 'bg-gray-100 text-gray-600'
}
</script>

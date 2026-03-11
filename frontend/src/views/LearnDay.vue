<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Loading -->
    <div v-if="loading" class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <div class="h-4 w-24 bg-gray-200 rounded animate-pulse mb-8"></div>
      <div class="bg-white rounded-2xl border border-gray-200 overflow-hidden">
        <div class="p-8 bg-gradient-to-r from-blue-600 to-purple-600 animate-pulse">
          <div class="flex items-center gap-4">
            <div class="w-16 h-16 bg-white/20 rounded-xl flex-shrink-0"></div>
            <div class="space-y-2 flex-1">
              <div class="h-6 bg-white/30 rounded w-2/3"></div>
            </div>
          </div>
        </div>
        <div class="p-8 space-y-3 animate-pulse">
          <div class="h-4 bg-gray-100 rounded w-full"></div>
          <div class="h-4 bg-gray-100 rounded w-5/6"></div>
          <div class="h-4 bg-gray-100 rounded w-4/5"></div>
        </div>
      </div>
    </div>

    <!-- Content -->
    <div v-else-if="step" class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <!-- Back -->
      <RouterLink to="/learn" class="inline-flex items-center gap-1 text-sm text-gray-500 hover:text-gray-900 mb-8 transition-colors">
        <ChevronLeft class="w-4 h-4"/>
        返回学习路径
      </RouterLink>

      <!-- Article card -->
      <article class="bg-white rounded-2xl border border-gray-200 overflow-hidden">
        <!-- Day header -->
        <div class="p-8 bg-gradient-to-r from-blue-600 to-purple-600">
          <div class="flex items-center gap-4">
            <div class="w-16 h-16 bg-white/20 rounded-xl flex flex-col items-center justify-center flex-shrink-0">
              <span class="text-white text-xs font-medium opacity-80">DAY</span>
              <span class="text-white font-bold text-2xl leading-none">{{ step.day }}</span>
            </div>
            <div>
              <h1 class="text-2xl font-bold text-white mb-1">{{ step.title }}</h1>
            </div>
          </div>
        </div>

        <!-- Content -->
        <div class="p-8 prose prose-gray max-w-none" v-html="renderedContent"></div>

        <!-- Mark as read -->
        <div class="px-8 pb-8">
          <button
            @click="toggleComplete"
            :class="[
              'w-full py-3 rounded-xl font-semibold text-sm transition-all flex items-center justify-center gap-2',
              isCompleted
                ? 'bg-green-50 text-green-700 border-2 border-green-300 hover:bg-green-100'
                : 'bg-gradient-to-r from-blue-600 to-purple-600 text-white hover:opacity-90'
            ]"
          >
            <Check v-if="isCompleted" class="w-4 h-4"/>
            <BookOpen v-else class="w-4 h-4"/>
            {{ isCompleted ? '已标记为完成 (点击取消)' : '标记为已完成' }}
          </button>
        </div>
      </article>

      <!-- Prev / Next nav -->
      <div class="mt-6 grid grid-cols-2 gap-4">
        <RouterLink
          v-if="step.day > 1"
          :to="`/learn/day/${step.day - 1}`"
          class="flex items-center gap-2 p-4 bg-white border border-gray-200 rounded-xl hover:border-blue-300 hover:shadow-sm transition-all"
        >
          <ChevronLeft class="w-5 h-5 text-gray-400 flex-shrink-0"/>
          <div class="min-w-0">
            <div class="text-xs text-gray-400 mb-0.5">上一天</div>
            <div class="text-sm font-medium truncate">{{ prevTitle }}</div>
          </div>
        </RouterLink>
        <div v-else></div>

        <RouterLink
          v-if="step.day < 7"
          :to="`/learn/day/${step.day + 1}`"
          class="flex items-center justify-end gap-2 p-4 bg-white border border-gray-200 rounded-xl hover:border-blue-300 hover:shadow-sm transition-all text-right"
        >
          <div class="min-w-0">
            <div class="text-xs text-gray-400 mb-0.5">下一天</div>
            <div class="text-sm font-medium truncate">{{ nextTitle }}</div>
          </div>
          <ChevronRight class="w-5 h-5 text-gray-400 flex-shrink-0"/>
        </RouterLink>
        <div v-else></div>
      </div>
    </div>

    <!-- 404 -->
    <div v-else class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-20 text-center">
      <p class="text-gray-500 mb-4">该学习内容未找到</p>
      <RouterLink to="/learn" class="text-blue-600 hover:underline">返回学习路径</RouterLink>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ChevronLeft, ChevronRight, Check, BookOpen } from 'lucide-vue-next'
import { apiGetLearnDay, apiListLearn } from '@/api/learn.js'
import { useLearnProgress } from '@/composables/useLearnProgress.js'

const route = useRoute()
const step = ref(null)
const allSteps = ref([])
const loading = ref(true)

const { isCompleted: checkCompleted, toggleDay } = useLearnProgress()
const isCompleted = computed(() => step.value ? checkCompleted(step.value.day) : false)

const prevTitle = computed(() => {
  if (!step.value) return ''
  return allSteps.value.find(s => s.day === step.value.day - 1)?.title || ''
})
const nextTitle = computed(() => {
  if (!step.value) return ''
  return allSteps.value.find(s => s.day === step.value.day + 1)?.title || ''
})

function toggleComplete() {
  if (step.value) toggleDay(step.value.day)
}

async function fetchStep(day) {
  loading.value = true
  step.value = null
  try {
    const [stepRes, listRes] = await Promise.all([
      apiGetLearnDay(day),
      allSteps.value.length === 0 ? apiListLearn() : Promise.resolve({ data: allSteps.value }),
    ])
    step.value = stepRes
    if (listRes.data) allSteps.value = listRes.data
  } catch (_) {
    // step stays null → 404 shown
  } finally {
    loading.value = false
  }
}

watch(() => route.params.day, day => { if (day) fetchStep(day) })
onMounted(() => fetchStep(route.params.day))

function renderMarkdown(md) {
  if (!md) return ''
  return md
    .replace(/^### (.+)$/gm, '<h3 class="text-lg font-bold mt-6 mb-2 text-gray-900">$1</h3>')
    .replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold mt-8 mb-3 text-gray-900">$1</h2>')
    .replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold mt-4 mb-4 text-gray-900">$1</h1>')
    .replace(/\*\*(.+?)\*\*/g, '<strong class="font-semibold">$1</strong>')
    .replace(/`([^`\n]+)`/g, '<code class="px-1.5 py-0.5 bg-gray-100 rounded text-sm font-mono text-blue-700">$1</code>')
    .replace(/```[\w]*\n([\s\S]*?)```/g, '<pre class="bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto my-4 text-sm font-mono leading-relaxed"><code>$1</code></pre>')
    .replace(/^\|(.+)\|$/gm, (line) => {
      if (/^[\|\s\-:]+$/.test(line)) return '<tr class="hidden"></tr>'
      const cells = line.split('|').filter((_, i, a) => i > 0 && i < a.length - 1)
      return '<tr class="border-b border-gray-100">' + cells.map(c => `<td class="px-4 py-2 text-sm">${c.trim()}</td>`).join('') + '</tr>'
    })
    .replace(/(<tr[\s\S]+?<\/tr>)+/g, m => `<div class="overflow-x-auto my-4"><table class="w-full border border-gray-200 rounded-lg overflow-hidden text-left">${m}</table></div>`)
    .replace(/^- (.+)$/gm, '<li class="ml-5 list-disc text-gray-700 my-1">$1</li>')
    .replace(/^(\d+)\. (.+)$/gm, '<li class="ml-5 list-decimal text-gray-700 my-1">$2</li>')
    .replace(/(<li[\s\S]+?<\/li>)+/g, m => `<ul class="my-3">${m}</ul>`)
    .replace(/\n\n+/g, '</p><p class="my-3 text-gray-700 leading-relaxed">')
}

const renderedContent = computed(() => {
  if (!step.value) return ''
  return `<p class="my-3 text-gray-700 leading-relaxed">${renderMarkdown(step.value.content)}</p>`
})
</script>

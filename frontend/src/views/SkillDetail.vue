<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Loading -->
    <div v-if="loading" class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <div class="h-4 w-24 bg-gray-200 rounded animate-pulse mb-8"></div>
      <div class="bg-white rounded-2xl border border-gray-200 overflow-hidden animate-pulse">
        <div class="p-8 border-b border-gray-100">
          <div class="flex items-start gap-5">
            <div class="w-20 h-20 bg-gray-200 rounded-2xl flex-shrink-0"></div>
            <div class="flex-1 space-y-3">
              <div class="h-7 bg-gray-200 rounded w-1/2"></div>
              <div class="h-4 bg-gray-100 rounded w-full"></div>
              <div class="flex gap-2">
                <div class="h-7 w-20 bg-gray-100 rounded-lg"></div>
                <div class="h-7 w-20 bg-gray-100 rounded-lg"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Skill detail -->
    <div v-else-if="skill" class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <!-- Back -->
      <RouterLink to="/skills" class="inline-flex items-center gap-1 text-sm text-gray-500 hover:text-gray-900 mb-8 transition-colors">
        <ChevronLeft class="w-4 h-4"/>
        返回 Skills 列表
      </RouterLink>

      <!-- Main card -->
      <div class="bg-white rounded-2xl border border-gray-200 overflow-hidden">
        <!-- Header -->
        <div class="p-8 border-b border-gray-100">
          <div class="flex items-start gap-5">
            <!-- Avatar -->
            <div class="w-20 h-20 bg-gradient-to-br from-blue-100 to-purple-100 rounded-2xl flex items-center justify-center flex-shrink-0">
              <span class="text-4xl font-bold bg-gradient-to-br from-blue-600 to-purple-600 bg-clip-text text-transparent">
                {{ skill.name.charAt(0) }}
              </span>
            </div>

            <div class="flex-1 min-w-0">
              <div class="flex flex-wrap items-center gap-2 mb-2">
                <h1 class="text-2xl font-bold">{{ skill.name }}</h1>
              </div>
              <p class="text-gray-600 mb-4">{{ skill.description }}</p>

              <!-- Meta chips -->
              <div class="flex flex-wrap gap-2 text-sm">
                <span class="flex items-center gap-1 px-3 py-1 bg-gray-100 rounded-lg text-gray-700">
                  <Star class="w-3.5 h-3.5 text-yellow-500 fill-current"/>
                  {{ skill.rating?.toFixed(1) || '—' }} 评分
                </span>
                <span class="flex items-center gap-1 px-3 py-1 bg-gray-100 rounded-lg text-gray-700">
                  <Download class="w-3.5 h-3.5 text-blue-500"/>
                  {{ (skill.download_count || 0).toLocaleString() }} 安装
                </span>
                <span v-if="skill.version" class="px-3 py-1 bg-gray-100 rounded-lg text-gray-700 font-mono">v{{ skill.version }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Install command -->
        <div class="p-8 border-b border-gray-100">
          <h2 class="text-sm font-semibold text-gray-500 uppercase tracking-wide mb-3">安装命令</h2>
          <div class="flex items-center gap-2 bg-gray-900 rounded-xl px-5 py-4">
            <code class="flex-1 text-green-400 font-mono text-sm select-all">claw install {{ skill.name }}</code>
            <button
              @click="copyCmd"
              class="flex-shrink-0 p-1.5 rounded-lg hover:bg-white/10 transition-colors text-gray-400 hover:text-white"
              :title="copied ? '已复制' : '复制'"
            >
              <Check v-if="copied" class="w-4 h-4 text-green-400"/>
              <Copy v-else class="w-4 h-4"/>
            </button>
          </div>
          <p v-if="copied" class="mt-2 text-xs text-green-600">✓ 已复制到剪贴板</p>
        </div>

        <!-- Details -->
        <div class="p-8">
          <h2 class="text-sm font-semibold text-gray-500 uppercase tracking-wide mb-4">详细信息</h2>
          <dl class="grid sm:grid-cols-2 gap-4">
            <div v-if="skill.author" class="flex flex-col gap-1">
              <dt class="text-xs text-gray-400">作者</dt>
              <dd class="text-sm font-medium text-gray-900">{{ skill.author }}</dd>
            </div>
            <div v-if="skill.version" class="flex flex-col gap-1">
              <dt class="text-xs text-gray-400">版本</dt>
              <dd class="text-sm font-mono text-gray-900">v{{ skill.version }}</dd>
            </div>
          </dl>

          <!-- Tags -->
          <div v-if="parsedTags.length > 0" class="mt-6">
            <dt class="text-xs text-gray-400 mb-2">标签</dt>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="tag in parsedTags"
                :key="tag"
                class="px-3 py-1 bg-blue-50 text-blue-700 text-sm rounded-lg"
              >{{ tag }}</span>
            </div>
          </div>

          <!-- Content -->
          <div v-if="renderedContent" class="mt-8 pt-8 border-t border-gray-100 prose prose-gray max-w-none" v-html="renderedContent"></div>
        </div>
      </div>

      <!-- Related skills -->
      <div v-if="related.length > 0" class="mt-10">
        <h2 class="text-lg font-semibold mb-4">同类 Skills</h2>
        <div class="grid sm:grid-cols-2 gap-4">
          <RouterLink
            v-for="s in related"
            :key="s.id"
            :to="`/skills/${s.id}`"
            class="bg-white border border-gray-200 rounded-xl p-4 hover:shadow-sm hover:border-green-300 transition-all flex items-center gap-3"
          >
            <div class="w-10 h-10 bg-gradient-to-br from-blue-100 to-purple-100 rounded-lg flex items-center justify-center flex-shrink-0">
              <span class="font-bold text-blue-600">{{ s.name.charAt(0) }}</span>
            </div>
            <div class="min-w-0">
              <div class="text-sm font-medium truncate">{{ s.name }}</div>
              <div class="text-xs text-gray-500 truncate">{{ s.description }}</div>
            </div>
          </RouterLink>
        </div>
      </div>
    </div>

    <!-- 404 -->
    <div v-else class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-20 text-center">
      <p class="text-gray-500 mb-4">Skill 未找到</p>
      <RouterLink to="/skills" class="text-green-600 hover:underline">返回 Skills 列表</RouterLink>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { ChevronLeft, Star, Download, Copy, Check } from 'lucide-vue-next'
import { apiGetSkill, apiListSkills } from '@/api/skills.js'

const route = useRoute()
const skill = ref(null)
const related = ref([])
const loading = ref(true)
const copied = ref(false)

const parsedTags = computed(() => {
  const tags = skill.value?.tags
  if (!tags) return []
  if (tags.startsWith('[')) {
    try { return JSON.parse(tags) } catch (_) {}
  }
  return tags.split(',').map(t => t.trim()).filter(Boolean)
})

async function fetchSkill(id) {
  loading.value = true
  skill.value = null
  related.value = []
  try {
    skill.value = await apiGetSkill(id)
    fetchRelated()
  } catch (_) {
    // skill stays null → 404 shown
  } finally {
    loading.value = false
  }
}

async function fetchRelated() {
  if (!skill.value) return
  const firstTag = parsedTags.value[0]
  if (!firstTag) return
  try {
    const res = await apiListSkills({ pageSize: 5, tag: firstTag })
    related.value = (res.data || []).filter(s => s.id !== skill.value.id).slice(0, 4)
  } catch (_) {}
}

async function copyCmd() {
  if (!skill.value) return
  try {
    await navigator.clipboard.writeText(`claw install ${skill.value.name}`)
    copied.value = true
    setTimeout(() => { copied.value = false }, 2000)
  } catch (_) {}
}

function renderMarkdown(md) {
  if (!md) return ''
  return md
    .replace(/^### (.+)$/gm, '<h3 class="text-lg font-bold mt-6 mb-2">$1</h3>')
    .replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold mt-8 mb-3">$1</h2>')
    .replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold mt-8 mb-4">$1</h1>')
    .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
    .replace(/`([^`\n]+)`/g, '<code class="px-1.5 py-0.5 bg-gray-100 rounded text-sm font-mono text-green-700">$1</code>')
    .replace(/```[\w]*\n([\s\S]*?)```/g, '<pre class="bg-gray-900 text-gray-100 rounded-lg p-4 overflow-x-auto my-4 text-sm font-mono"><code>$1</code></pre>')
    .replace(/^- (.+)$/gm, '<li class="ml-4 list-disc text-gray-700">$1</li>')
    .replace(/^\d+\. (.+)$/gm, '<li class="ml-4 list-decimal text-gray-700">$1</li>')
    .replace(/\n\n/g, '</p><p class="my-3 text-gray-700 leading-relaxed">')
}

const renderedContent = computed(() => {
  if (!skill.value?.content) return ''
  return `<p class="my-3 text-gray-700 leading-relaxed">${renderMarkdown(skill.value.content)}</p>`
})

watch(() => route.params.id, id => { if (id) fetchSkill(id) })
onMounted(() => fetchSkill(route.params.id))
</script>

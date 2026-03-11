<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <h1 class="text-3xl font-bold mb-2">7 天学习路径</h1>
        <p class="text-gray-600">系统化学习，快速上手 110Claw</p>

        <!-- Progress bar -->
        <div class="mt-6 flex items-center gap-4">
          <div class="flex-1 bg-gray-200 rounded-full h-2">
            <div
              class="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full transition-all"
              :style="{ width: progressPercent + '%' }"
            ></div>
          </div>
          <span class="text-sm font-medium text-gray-600 flex-shrink-0">{{ completedDays.length }}/7 已完成</span>
        </div>
      </div>
    </div>

    <!-- Day cards -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <!-- Loading -->
      <div v-if="loading" class="grid sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-5">
        <div v-for="i in 7" :key="i" class="bg-white border border-gray-200 rounded-xl p-5 animate-pulse">
          <div class="flex items-start gap-3 mb-3">
            <div class="w-14 h-14 bg-gray-200 rounded-xl flex-shrink-0"></div>
          </div>
          <div class="h-4 bg-gray-200 rounded mb-2 w-3/4"></div>
          <div class="h-3 bg-gray-100 rounded w-full mb-1"></div>
          <div class="h-3 bg-gray-100 rounded w-2/3"></div>
        </div>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="text-center py-20">
        <p class="text-gray-500 mb-2">加载失败</p>
        <button @click="fetchSteps" class="text-blue-600 hover:underline text-sm">重试</button>
      </div>

      <div v-else class="grid sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-5">
        <RouterLink
          v-for="step in steps"
          :key="step.day"
          :to="`/learn/day/${step.day}`"
          class="group relative bg-white border border-gray-200 rounded-xl p-5 hover:shadow-lg transition-all hover:border-blue-300"
          :class="{ 'border-green-300 bg-green-50/30': isCompleted(step.day) }"
        >
          <!-- Completed badge -->
          <div v-if="isCompleted(step.day)" class="absolute top-3 right-3">
            <div class="w-6 h-6 bg-green-500 rounded-full flex items-center justify-center">
              <Check class="w-3.5 h-3.5 text-white"/>
            </div>
          </div>

          <!-- Day badge -->
          <div class="flex items-start gap-3 mb-3">
            <div
              class="w-14 h-14 rounded-xl flex flex-col items-center justify-center flex-shrink-0 transition-all"
              :class="isCompleted(step.day)
                ? 'bg-gradient-to-br from-green-400 to-green-500'
                : 'bg-gradient-to-br from-blue-500 to-blue-600 group-hover:from-blue-600 group-hover:to-purple-600'"
            >
              <span class="text-white text-xs font-medium opacity-80">DAY</span>
              <span class="text-white font-bold text-lg leading-none">{{ step.day }}</span>
            </div>
          </div>

          <h3 class="font-bold text-base mb-2 group-hover:text-blue-700 transition-colors">{{ step.title }}</h3>
          <p class="text-sm text-gray-600 mb-4 line-clamp-2">{{ step.description }}</p>

          <div class="flex items-center gap-1 text-sm text-blue-600 font-medium">
            <span>{{ isCompleted(step.day) ? '再次学习' : '开始学习' }}</span>
            <ChevronRight class="w-3.5 h-3.5"/>
          </div>
        </RouterLink>
      </div>

      <!-- Reset progress -->
      <div v-if="completedDays.length > 0" class="mt-8 text-center">
        <button @click="resetProgress" class="text-sm text-gray-400 hover:text-gray-600 transition-colors">
          重置学习进度
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ChevronRight, Check } from 'lucide-vue-next'
import { apiListLearn } from '@/api/learn.js'
import { useLearnProgress } from '@/composables/useLearnProgress.js'

const steps = ref([])
const loading = ref(true)
const error = ref(false)

const { completedDays, isCompleted, reset: resetProgress } = useLearnProgress()
const progressPercent = computed(() => Math.round(completedDays.value.length / 7 * 100))

async function fetchSteps() {
  loading.value = true
  error.value = false
  try {
    const res = await apiListLearn()
    steps.value = res.data || []
  } catch (_) {
    error.value = true
  } finally {
    loading.value = false
  }
}

onMounted(fetchSteps)
</script>

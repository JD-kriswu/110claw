<template>
  <div class="min-h-screen">
    <!-- Hero -->
    <section class="relative overflow-hidden bg-gradient-to-b from-blue-50 via-white to-white pb-16">
      <!-- Decorative blobs -->
      <div class="absolute top-0 left-0 w-full h-full overflow-hidden pointer-events-none opacity-40">
        <div class="absolute top-20 left-10 w-72 h-72 bg-blue-400 rounded-full mix-blend-multiply filter blur-3xl"></div>
        <div class="absolute top-40 right-10 w-72 h-72 bg-purple-400 rounded-full mix-blend-multiply filter blur-3xl"></div>
      </div>

      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 sm:py-28 relative">
        <div class="text-center max-w-4xl mx-auto mb-12">
          <h1 class="text-4xl sm:text-5xl md:text-6xl font-bold mb-6 bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
            免费安全的Claw社群平台
          </h1>
          <p class="text-xl sm:text-2xl text-gray-600">一起打造安全高效的数字助理</p>
        </div>

        <!-- Quick nav cards -->
        <div class="grid md:grid-cols-3 gap-4 relative z-10">
          <!-- 快速学习 -->
          <RouterLink to="/learn" class="group p-5 bg-white border border-gray-200 rounded-lg hover:shadow-lg transition-all hover:border-blue-300">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center group-hover:bg-blue-600 transition-colors flex-shrink-0">
                <BookOpen class="w-5 h-5 text-blue-600 group-hover:text-white transition-colors"/>
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="text-base font-semibold mb-1">快速学习</h3>
                <p class="text-sm text-gray-600 line-clamp-1">通过 7 天结构化学习路径，从零到熟练掌握 110Claw</p>
              </div>
              <ChevronRight class="w-4 h-4 text-gray-400 group-hover:text-blue-600 transition-colors flex-shrink-0"/>
            </div>
          </RouterLink>

          <!-- Claw 资讯 -->
          <RouterLink to="/news" class="group p-5 bg-white border border-gray-200 rounded-lg hover:shadow-lg transition-all hover:border-purple-300">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center group-hover:bg-purple-600 transition-colors flex-shrink-0">
                <Newspaper class="w-5 h-5 text-purple-600 group-hover:text-white transition-colors"/>
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="text-base font-semibold mb-1">Claw 资讯</h3>
                <p class="text-sm text-gray-600 line-clamp-1">追踪版本发布、安全更新、生态动态与社区资讯</p>
              </div>
              <ChevronRight class="w-4 h-4 text-gray-400 group-hover:text-purple-600 transition-colors flex-shrink-0"/>
            </div>
          </RouterLink>

          <!-- Skill 插件 -->
          <RouterLink to="/skills" class="group p-5 bg-white border border-gray-200 rounded-lg hover:shadow-lg transition-all hover:border-green-300">
            <div class="flex items-center gap-4">
              <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center group-hover:bg-green-600 transition-colors flex-shrink-0">
                <Puzzle class="w-5 h-5 text-green-600 group-hover:text-white transition-colors"/>
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="text-base font-semibold mb-1">Skill 插件</h3>
                <p class="text-sm text-gray-600 line-clamp-1">发现、安装和管理强大的 Skills，扩展 110Claw 能力</p>
              </div>
              <ChevronRight class="w-4 h-4 text-gray-400 group-hover:text-green-600 transition-colors flex-shrink-0"/>
            </div>
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- Latest News -->
    <section class="py-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-end justify-between mb-8">
          <div>
            <h2 class="text-3xl font-bold mb-2">最新资讯</h2>
            <p class="text-gray-600">了解 Claw 的最新动态和更新</p>
          </div>
          <RouterLink to="/news" class="hidden sm:inline-flex items-center gap-1 text-purple-600 hover:text-purple-700 font-medium">
            <span>查看全部</span>
            <ChevronRight class="w-4 h-4"/>
          </RouterLink>
        </div>

        <div class="grid md:grid-cols-3 gap-6">
          <RouterLink
            v-for="item in newsItems.slice(0,3)"
            :key="item.id"
            to="/news"
            class="p-6 bg-white border border-gray-200 rounded-lg hover:shadow-md transition-all hover:border-purple-300"
          >
            <div class="flex items-start gap-3">
              <Newspaper class="w-5 h-5 text-purple-600 flex-shrink-0 mt-1"/>
              <div>
                <div class="text-xs text-gray-500 mb-2">{{ formatDate(item.date) }}</div>
                <h3 class="font-semibold mb-2 line-clamp-2">{{ item.title }}</h3>
                <p class="text-sm text-gray-600 mb-3 line-clamp-2">{{ item.summary }}</p>
                <div class="flex flex-wrap gap-2">
                  <span
                    v-for="tag in item.tags.slice(0,2)"
                    :key="tag"
                    class="px-2 py-1 bg-purple-50 text-purple-700 text-xs rounded-full"
                  >{{ tag }}</span>
                </div>
              </div>
            </div>
          </RouterLink>
        </div>

        <div class="mt-8 text-center sm:hidden">
          <RouterLink to="/news" class="inline-flex items-center gap-1 text-purple-600 font-medium">
            <span>查看全部资讯</span><ChevronRight class="w-4 h-4"/>
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- 7-day Learning Path -->
    <section class="py-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-end justify-between mb-8">
          <div>
            <h2 class="text-3xl font-bold mb-2">7 天学习路径</h2>
            <p class="text-gray-600">系统化学习，快速上手 110Claw</p>
          </div>
          <RouterLink to="/learn" class="hidden sm:inline-flex items-center gap-1 text-blue-600 hover:text-blue-700 font-medium">
            <span>开始学习</span><ChevronRight class="w-4 h-4"/>
          </RouterLink>
        </div>

        <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-4">
          <RouterLink
            v-for="step in learningSteps"
            :key="step.day"
            :to="`/learn/day/${step.day}`"
            class="group p-5 bg-white border border-gray-200 rounded-xl hover:shadow-lg transition-all hover:border-blue-300"
          >
            <div class="flex items-start gap-3 mb-3">
              <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-lg flex items-center justify-center flex-shrink-0">
                <span class="text-white font-bold text-sm">DAY {{ step.day }}</span>
              </div>
            </div>
            <h3 class="font-bold text-base mb-2">{{ step.title }}</h3>
            <p class="text-sm text-gray-600 mb-3 line-clamp-2">{{ step.desc }}</p>
            <div class="text-sm text-blue-600 font-medium flex items-center gap-1">
              查看详情 <ChevronRight class="w-3 h-3"/>
            </div>
          </RouterLink>
          <!-- filler for 4-col layout -->
          <div class="hidden lg:block"></div>
        </div>

        <div class="mt-8 text-center sm:hidden">
          <RouterLink to="/learn" class="inline-flex items-center gap-1 text-blue-600 font-medium">
            <span>开始学习路径</span><ChevronRight class="w-4 h-4"/>
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- Hot Skills -->
    <section class="py-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-end justify-between mb-8">
          <div>
            <h2 class="text-3xl font-bold mb-2">热门 Skills</h2>
            <p class="text-gray-600">探索社区最受欢迎的插件</p>
          </div>
          <RouterLink to="/skills" class="hidden sm:inline-flex items-center gap-1 text-green-600 hover:text-green-700 font-medium">
            <span>查看全部</span><ChevronRight class="w-4 h-4"/>
          </RouterLink>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
          <RouterLink
            v-for="skill in skills"
            :key="skill.id"
            :to="`/skills/${skill.id}`"
            class="group p-4 bg-white border border-gray-200 rounded-xl hover:shadow-lg transition-all hover:border-purple-300"
          >
            <div class="flex items-start gap-3">
              <div class="w-10 h-10 bg-gradient-to-br from-blue-100 to-purple-100 rounded-lg flex items-center justify-center flex-shrink-0">
                <span class="text-lg font-bold bg-gradient-to-br from-blue-600 to-purple-600 bg-clip-text text-transparent">
                  {{ skill.name.charAt(0) }}
                </span>
              </div>
              <div class="flex-1 min-w-0">
                <div class="flex items-start justify-between gap-2 mb-1">
                  <h3 class="font-semibold text-base">{{ skill.name }}</h3>
                  <span class="px-2 py-0.5 bg-green-100 text-green-700 text-xs rounded flex-shrink-0">110Claw认证</span>
                </div>
                <p class="text-sm text-gray-600 mb-2 line-clamp-2">{{ skill.description }}</p>
                <div class="flex items-center gap-3 text-xs text-gray-500">
                  <div class="flex items-center gap-1">
                    <Star class="w-3 h-3 fill-current text-yellow-400"/>
                    <span>6.375</span>
                  </div>
                  <div class="flex items-center gap-1">
                    <ArrowRight class="w-3 h-3"/>
                    <span>295</span>
                  </div>
                  <div class="px-2 py-0.5 bg-blue-50 text-blue-700 rounded">v1.0.0</div>
                </div>
              </div>
            </div>
          </RouterLink>
        </div>

        <div class="mt-8 text-center sm:hidden">
          <RouterLink to="/skills" class="inline-flex items-center gap-1 text-green-600 font-medium">
            <span>查看全部 Skills</span><ChevronRight class="w-4 h-4"/>
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- Join CTA -->
    <section class="py-20 bg-gradient-to-r from-blue-600 to-purple-600">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <Users class="w-16 h-16 text-white mx-auto mb-6"/>
        <h2 class="text-3xl sm:text-4xl font-bold text-white mb-4">加入 110Claw 社区</h2>
        <p class="text-xl text-blue-100 mb-8">与国内开发者一起交流学习，分享经验，共同打造安全高效的数字助理</p>
        <RouterLink
          to="/learn"
          class="inline-flex items-center justify-center gap-2 px-8 py-4 bg-white text-blue-600 rounded-lg font-semibold hover:bg-blue-50 transition-colors"
        >
          <BookOpen class="w-5 h-5"/>
          <span>进一步学习</span>
        </RouterLink>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ChevronRight, BookOpen, Newspaper, Puzzle, Star, Users, ArrowRight } from 'lucide-vue-next'
import { newsItems } from '@/data/news.js'
import { learningSteps } from '@/data/learning.js'
import { skills } from '@/data/skills.js'

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' })
}
</script>

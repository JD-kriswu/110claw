<template>
  <footer class="bg-gray-900 text-gray-400 py-12">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-2 md:grid-cols-4 gap-8 mb-8">
        <!-- Brand -->
        <div class="col-span-2 md:col-span-1">
          <div class="flex items-center gap-2 mb-3">
            <svg width="28" height="28" viewBox="0 0 32 32" fill="none">
              <defs>
                <linearGradient id="footerLogoGrad" x1="0%" y1="0%" x2="100%" y2="100%">
                  <stop offset="0%" stop-color="#2563eb"/>
                  <stop offset="100%" stop-color="#7c3aed"/>
                </linearGradient>
              </defs>
              <rect width="32" height="32" rx="8" fill="url(#footerLogoGrad)"/>
              <path d="M8 16c0-4.4 3.6-8 8-8s8 3.6 8 8" stroke="white" stroke-width="2.5" stroke-linecap="round"/>
              <path d="M13 20l3 3 3-3" stroke="white" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
              <circle cx="16" cy="16" r="2.5" fill="white"/>
            </svg>
            <span class="font-bold text-white">110Claw</span>
          </div>
          <p class="text-sm">一起打造安全高效的数字助理</p>
        </div>

        <!-- 快速导航 -->
        <div>
          <h4 class="text-white font-semibold mb-3 text-sm">快速导航</h4>
          <ul class="space-y-2">
            <li v-for="link in quickLinks" :key="link.url">
              <RouterLink :to="link.url" class="text-sm hover:text-white transition-colors">{{ link.label }}</RouterLink>
            </li>
          </ul>
        </div>

        <!-- 法律条款 -->
        <div>
          <h4 class="text-white font-semibold mb-3 text-sm">法律条款</h4>
          <ul class="space-y-2">
            <li><a href="#" class="text-sm hover:text-white transition-colors">免责声明</a></li>
            <li><a href="#" class="text-sm hover:text-white transition-colors">隐私政策</a></li>
          </ul>
        </div>
      </div>

      <div class="border-t border-gray-800 pt-6 text-sm text-center">
        © 2026 110Claw. 开源项目，遵循 MIT 许可证。
      </div>
    </div>
  </footer>
</template>

<script setup>
const quickLinks = [
  { label: '7天学习计划', url: '/learn' },
  { label: '资讯', url: '/news' },
  { label: 'Skills', url: '/skills' },
]
</script>

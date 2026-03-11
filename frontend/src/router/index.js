import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/views/Layout.vue'),
    children: [
      { path: '', name: 'Home', component: () => import('@/views/Home.vue') },
      { path: 'news', name: 'News', component: () => import('@/views/News.vue') },
      { path: 'news/:id', name: 'NewsDetail', component: () => import('@/views/NewsDetail.vue') },
      { path: 'learn', name: 'Learn', component: () => import('@/views/Learn.vue') },
      { path: 'learn/day/:day', name: 'LearnDay', component: () => import('@/views/LearnDay.vue') },
      { path: 'skills', name: 'Skills', component: () => import('@/views/Skills.vue') },
      { path: 'skills/:id', name: 'SkillDetail', component: () => import('@/views/SkillDetail.vue') },
    ],
  },
]

export default createRouter({
  history: createWebHistory('/110claw/'),
  routes,
})

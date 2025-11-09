import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('@/views/Home.vue'),
    },
    {
      path: '/month',
      name: 'MonthView',
      component: () => import('@/views/MonthView.vue'),
    },
    {
      path: '/week',
      name: 'WeekView',
      component: () => import('@/views/WeekView.vue'),
    },
  ],
})

export default router


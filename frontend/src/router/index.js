import { createRouter, createWebHistory } from 'vue-router';
import Auth from '@/views/Auth.vue';
import Dashboard from '@/views/Dashboard.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/auth',
      name: 'auth',
      component: Auth,
    },
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard,
    },
  ],
});

export default router;

import { createRouter, createWebHistory } from 'vue-router';
import Auth from '@/views/Auth.vue';
import Dashboard from '@/views/Dashboard.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      meta: { layout: 'main' },
      component: Dashboard,
    },
    {
      path: '/auth',
      name: 'auth',
      meta: { layout: 'empty' },
      component: Auth,
    },
    {
      path: '/income',
      name: 'income',
      meta: { layout: 'main' },
      component: () => import('@/views/Income.vue'),
    },
    {
      path: '/expenses',
      name: 'expenses',
      meta: { layout: 'main' },
      component: () => import('@/views/Expense.vue'),
    },
    {
      path: '/deposit',
      name: 'deposit',
      meta: { layout: 'main' },
      component: () => import('@/views/Deposit.vue'),
    },
    {
      path: '/loans',
      name: 'loans',
      meta: { layout: 'main' },
      component: () => import('@/views/Loan.vue'),
    },
    {
      path: '/statistic',
      name: 'statistic',
      meta: { layout: 'main' },
      component: () => import('@/views/Statistic.vue'),
    },
  ],
});

export default router;

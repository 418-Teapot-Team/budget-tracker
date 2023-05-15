import { createRouter, createWebHistory } from 'vue-router';
import Auth from '@/views/Auth.vue';
import Dashboard from '@/views/Dashboard.vue';
import Income from '@/views/Income.vue';
import Expense from '@/views/Expense.vue';
import Deposit from '@/views/Deposit.vue';
import Loan from '@/views/Loan.vue';
import Statistic from '@/views/Statistic.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/auth',
      name: 'auth',
      meta: { layout: 'empty' },
      component: Auth,
    },
    {
      path: '/',
      name: 'dashboard',
      meta: { layout: 'main' },
      component: Dashboard,
    },
    {
      path: '/income',
      name: 'income',
      meta: { layout: 'main' },
      component: Income,
    },
    {
      path: '/expenses',
      name: 'expenses',
      meta: { layout: 'main' },
      component: Expense,
    },
    {
      path: '/deposit',
      name: 'deposit',
      meta: { layout: 'main' },
      component: Deposit,
    },
    {
      path: '/loans',
      name: 'loans',
      meta: { layout: 'main' },
      component: Loan,
    },
    {
      path: '/statistic',
      name: 'statistic',
      meta: { layout: 'main' },
      component: Statistic,
    },
  ],
});

export default router;

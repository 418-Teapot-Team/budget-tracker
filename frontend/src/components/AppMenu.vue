<template>
  <nav class="flex h-screen flex-col rounded-r-3xl bg-black w-1/5">
    <div class="w-full h-20 flex justify-center mt-6 mb-16"><teapot-icon fill="#1c1c21" /></div>
    <ul class="flex flex-col justify-start items-start self-center gap-7">
      <li v-for="link in links" :key="link.route">
        <router-link
          :to="link.route"
          class="flex flex-row justify-start gap-4 cursor-pointer"
          exact
        >
          <div
            class="h-8 w-8 p-1 rounded-full"
            :class="$route.path === link.route ? 'bg-white' : ''"
          >
            <component :is="link.icon + '-icon'" :isActive="$route.path === link.route" />
          </div>
          <span
            class="flex flex-row items-center text-lg text-white transition duration-300 border-b border-b-transparent hover:border-b-white"
            >{{ link.title }}</span
          >
        </router-link>
      </li>
    </ul>
  </nav>
</template>

<script>
import DashboardIcon from '@/components/Icons/DashboardIcon.vue';
import IncomeIcon from '@/components/Icons/IncomeIcon.vue';
import ExpensesIcon from '@/components/Icons/ExpensesIcon.vue';
import DepositIcon from '@/components/Icons/DepositIcon.vue';
import LoansIcon from '@/components/Icons/LoansIcon.vue';
import StatisticsIcon from '@/components/Icons/StatisticsIcon.vue';
import TeapotIcon from '@/components/Icons/Teapot.vue';
export default {
  name: 'AppMenu',
  components: {
    TeapotIcon,
    DashboardIcon,
    IncomeIcon,
    ExpensesIcon,
    DepositIcon,
    LoansIcon,
    StatisticsIcon,
  },
  data() {
    return {
      links: [
        {
          title: 'Dashboard',
          route: '/',
          icon: 'dashboard',
        },
        {
          title: 'Expense',
          route: '/expenses',
          icon: 'expenses',
        },
        {
          title: 'Income',
          route: '/income',
          icon: 'income',
        },

        {
          title: 'Deposits',
          route: '/deposit',
          icon: 'deposit',
        },
        {
          title: 'Loans',
          route: '/loans',
          icon: 'loans',
        },
        {
          title: 'Statistic',
          route: '/statistic',
          icon: 'statistics',
        },
      ],
    };
  },
};
</script>

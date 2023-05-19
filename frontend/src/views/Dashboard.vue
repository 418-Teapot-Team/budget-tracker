<template>
  <section id="dashboard">
    <div class="flex flex-row justify-between">
      <div class="w-8/12 stats">
        <div class="flex flex-row justify-between items-center mb-4">
          <span class="text-black text-4xl font-bold">Dashboard</span>
        </div>
        <current-month-banner :amount="currentMonthSaving" class="mb-6" />
        <div class="flex flex-row justify-between mb-6">
          <popular-category-stats
            v-for="item in topCategories"
            :key="item.category.id"
            :imageUrl="item.category.imageUrl"
            :amount="item.amount"
            :title="item.category.category"
            :style="`background-color: #${item.category.hash}`"
          />
        </div>
        <transactions-chart
          :values="data"
          :bottomLabels="labels"
          title="Incomes overview"
          class="h-full"
        />
      </div>
      <!-- -->
      <div class="w-4/12 pl-14">
        <last-transactions :data="transactions" />
      </div>
    </div>
  </section>
</template>

<script>
import CurrentMonthBanner from '@/components/CurrentMonthBanner.vue';
import PopularCategoryStats from '@/components/PopularCategoryStats.vue';
import TransactionsChart from '@/components/TransactionsChart.vue';
import LastTransactions from '@/components/LastTransactions.vue';
import useDashboardStore from '@/stores/dashboard';
import { mapActions, mapState } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'Dashboard',
  components: {
    CurrentMonthBanner,
    PopularCategoryStats,
    TransactionsChart,
    LastTransactions,
  },
  computed: {
    ...mapState(useDashboardStore, [
      'currentMonthSaving',
      'topCategories',
      'savingStats',
      'transactions',
    ]),
    labels() {
      return this.savingStats?.map((item) => item.Month);
    },
    data() {
      return this.savingStats?.map((item) => item.IncomeNet);
    },
  },

  methods: {
    ...mapActions(useDashboardStore, [
      'getTopCaregories',
      'getCurrentMonthSavings',
      'getSavingsStats',
      'getTransactions',
    ]),
    async getInitialData() {
      try {
        await this.getCurrentMonthSavings();
        await this.getTopCaregories();
        await this.getSavingsStats();
        await this.getTransactions({ take: 50, skip: 0 });
      } catch (e) {
        toast.error(e?.message);
      }
    },
  },
  mounted() {
    this.getInitialData();
  },
};
</script>

<style scoped>
.stats {
  max-height: calc(100vh - 474px);
}

@media screen and (min-width: 1280px) {
  .stats {
    max-height: calc(100vh - 538px);
  }
}
@media screen and (min-width: 1536px) {
  .stats {
    max-height: calc(100vh - 592px);
  }
}
</style>

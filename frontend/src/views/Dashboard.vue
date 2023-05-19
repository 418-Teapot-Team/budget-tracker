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
            imageUrl="https://www.svgrepo.com/show/357712/exchange-alt.svg"
            :amount="2843.43"
            title="Transfers"
            class="bg-yellow"
          />
          <popular-category-stats
            imageUrl="https://www.svgrepo.com/show/513464/house.svg"
            :amount="1287.12"
            title="Rent"
            class="bg-blue-light"
          />
          <popular-category-stats
            imageUrl="https://www.svgrepo.com/show/433318/bus2-o.svg"
            :amount="5234.21"
            title="Travelling"
            class="bg-pink"
          />
          <popular-category-stats
            imageUrl="https://www.svgrepo.com/show/443633/education-cap.svg"
            :amount="532.23"
            title="Education"
            class="bg-grey"
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
        <last-transactions />
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
export default {
  name: 'Dashboard',
  components: {
    CurrentMonthBanner,
    PopularCategoryStats,
    TransactionsChart,
    LastTransactions,
  },
  computed: {
    ...mapState(useDashboardStore, ['currentMonthSaving']),
  },
  data() {
    return {
      labels: [
        1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
        26, 27, 28, 29, 30,
      ],
      data: [
        1500, 2200, 510, 1832, 2038, 1639, 5433, 1234, 2321, 3212, 1200, 1200, 1500, 1100, 1943,
        1632, 4324, 3444, 932, 811, 7432, 5474, 2343, 2443, 1732, 1200, 1100, 1100, 1250, 6548,
      ],
    };
  },
  methods: {
    ...mapActions(useDashboardStore, ['getTopCaregories', 'getCurrentMonthSavings']),
  },
  mounted() {
    this.getTopCaregories();
    this.getCurrentMonthSavings();
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

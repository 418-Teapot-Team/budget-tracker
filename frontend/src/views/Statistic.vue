<template>
  <section>
    <!-- caption and filters -->
    <div class="flex flex-row justify-between items-end">
      <div class="flex flex-row justify-start items-end gap-4">
        <div class="flex flex-col gap-2 w-full">
          <span class="text-black text-4xl font-bold">Statistics</span>
          <div class="flex flex-row justify-start h-10 gap-3">
            <div class="h-10 w-40">
              <app-button
                text="Expenses"
                :isOutline="tab !== 'expenses'"
                @click="tab = 'expenses'"
              />
            </div>
            <div class="h-10 w-40">
              <app-button text="Incomes" :isOutline="tab !== 'incomes'" @click="tab = 'incomes'" />
            </div>
          </div>
        </div>
      </div>
      <filters @onApplyFilters="applyFilters" :withCategory="false" />
    </div>
    <!-- data -->
    <div class="pt-8 w-full flex flex-row gap-6 stats">
      <div class="flex flex-col justify-between gap-4 w-2/3">
        <balance-card class="h-1/3" />
        <transactions-chart
          :bottomLabels="labels"
          :values="data"
          :title="`${tab} overview`"
          class="h-2/3"
        />
      </div>
      <categories-chart
        class="w-1/3"
        :rightLabels="categoryLabels"
        :values="categoryDataset"
        :title="`${tab} by categories`"
      />
    </div>
  </section>
</template>

<script>
import Filters from '@/components/Filters.vue';
import BalanceCard from '@/components/BalanceCard.vue';
import CategoriesChart from '@/components/CategoriesChart.vue';
import TransactionsChart from '@/components/TransactionsChart.vue';
export default {
  name: 'Deposit',
  components: {
    Filters,
    BalanceCard,
    CategoriesChart,
    TransactionsChart,
  },
  data() {
    return {
      tab: 'expenses',
      categoryLabels: ['Transactions', 'Travelling', 'Rent', 'Education'],
      categoryDataset: [
        {
          value: 20,
          color: '#ffefe2',
        },
        {
          value: 30,
          color: '#F0F5FF',
        },
        {
          value: 45,
          color: '#FEF8FC',
        },
        {
          value: 25,
          color: '#e8e8e8',
        },
      ],
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
};
</script>

<style scoped>
.stats {
  max-height: calc(100vh - 192px);
  height: calc(100vh - 192px);
  padding-bottom: 20px;
}
</style>

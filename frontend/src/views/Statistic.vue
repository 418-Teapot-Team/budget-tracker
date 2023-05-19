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
                @click="changeTab('expenses')"
              />
            </div>
            <div class="h-10 w-40">
              <app-button
                text="Incomes"
                :isOutline="tab !== 'income'"
                @click="changeTab('income')"
              />
            </div>
          </div>
        </div>
      </div>
      <filters @onApplyFilters="applyFilters" :withCategory="false" :isStats="true" />
    </div>
    <!-- data -->
    <div class="pt-8 w-full flex flex-row gap-6 stats">
      <div class="flex flex-col justify-between gap-4 w-2/3">
        <balance-card class="h-1/3" :amount="total" />
        <transactions-chart
          :bottomLabels="labels"
          :values="data"
          :title="`${tab} overview`"
          class="h-2/3"
        />
      </div>
      <categories-chart
        class="w-1/3"
        :rightLabels="categoriesLabels"
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
import useFiltersStore from '@/stores/filters';
import useStatisticsStore from '@/stores/statistics';
import { mapState, mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

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
    };
  },
  computed: {
    ...mapState(useStatisticsStore, [
      'expensesStats',
      'expensesCategories',
      'total',
      'incomeStats',
      'incomeCategories',
    ]),
    ...mapState(useFiltersStore, ['filtersForStats']),
    categoriesLabels() {
      if (this.tab === 'income') {
        return this.incomeCategories?.map((item) => item.category?.category);
      } else {
        return this.expensesCategories?.map((item) => item.category?.category);
      }
    },
    categoryDataset() {
      if (this.tab === 'income') {
        return this.incomeCategories?.map((item) => ({
          value: item?.amount,
          hash: '#' + item?.category?.hash,
        }));
      } else {
        return this.expensesCategories?.map((item) => ({
          value: item?.amount,
          hash: '#' + item?.category.hash,
        }));
      }
    },
    labels() {
      if (this.tab === 'income') {
        return this.incomeStats?.map((item) => item.Date);
      } else {
        return this.expensesStats?.map((item) => item.Date);
      }
    },
    data() {
      if (this.tab === 'income') {
        return this.incomeStats?.map((item) => item.Value);
      } else {
        return this.expensesStats?.map((item) => item.Value);
      }
    },
  },
  methods: {
    ...mapActions(useStatisticsStore, ['getTopCaregories', 'getTotalSum', 'getStats']),
    changeTab(tabVal) {
      this.tab = tabVal;
      this.getInitialData(tabVal);
    },
    async getInitialData(tabValue) {
      const months = this.$route.query?.months ? this.$route.query?.months : 1;
      try {
        await this.getTopCaregories({ type: tabValue, months });
        await this.getTotalSum({ type: tabValue, months });
        await this.getStats({ type: tabValue, months });
        console.log(this.expensesStats);
        console.log(this.incomeStats);
      } catch (e) {
        toast.error(e?.message);
      }
    },
    applyFilters({ filter }) {
      const filterToUrl = this.filtersForStats.find((item) => filter === item.value.id);

      this.$router
        .push({
          query: {
            months: filterToUrl ? filterToUrl?.value?.months : 1,
            filterId: filterToUrl ? filterToUrl?.value?.id : 'default',
          },
        })
        .then(() => {
          this.getInitialData(this.tab);
        });
    },
  },
  mounted() {
    this.getInitialData(this.tab);
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

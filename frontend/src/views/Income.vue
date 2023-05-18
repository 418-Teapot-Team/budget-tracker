<template>
  <section>
    <!-- caption and filters -->
    <div class="flex flex-row justify-between items-center">
      <div class="flex flex-row justify-start items-center gap-4">
        <span class="text-black text-4xl font-bold">Income</span>
        <div class="h-8 w-8 cursor-pointer">
          <add-icon />
        </div>
      </div>
      <filters @onApplyFilters="applyFilters" :withCategory="true" />
    </div>

    <!-- table -->
    <div class="pt-4 w-full">
      <finances-table
        :isIncome="true"
        @onDeleteItem="deleteItem"
        @onLoadMore="loadMore"
        @onEditItem="editItem"
      />
    </div>

    <!-- popup -->
    <transaction-popup v-if="showIncomePopup" :isEdit="false" @onClose="showIncomePopup = false" />
  </section>
</template>

<script>
import FinancesTable from '@/components/FinancesTable.vue';
import AddIcon from '@/components/Icons/AddIcon.vue';
import Filters from '@/components/Filters.vue';
import TransactionPopup from '@/components/TransactionPopup.vue';
import useTransitionStore from '@/stores/transactions';
import { mapActions } from 'pinia';

export default {
  name: 'Income',
  components: {
    FinancesTable,
    TransactionPopup,
    AddIcon,
    Filters,
  },
  data() {
    return {
      showIncomePopup: false,
    };
  },
  methods: {
    ...mapActions(useTransitionStore, ['getTransactions', 'getCategories']),
    loadMore() {
      console.log('load more');
    },
    deleteItem(id) {
      console.log(id);
    },
    applyFilters({ category, filter }) {
      console.log(category, filter);
    },
    editItem(id) {
      this.showIncomePopup = true;
      console.log(id);
    },
  },
  mounted() {
    this.getCategories();
    this.getTransactions({ type: 'income' });
  },
};
</script>

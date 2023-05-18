<template>
  <section>
    <!-- caption and filters -->
    <div class="flex flex-row justify-between items-center">
      <div class="flex flex-row justify-start items-center gap-4">
        <span class="text-black text-4xl font-bold">Expense</span>
        <div class="h-8 w-8 cursor-pointer">
          <add-icon />
        </div>
      </div>
      <filters @onApplyFilters="applyFilters" :withCategory="true" />
    </div>
    <!-- table -->
    <div class="pt-4 w-full">
      <finances-table
        :isIncome="false"
        @onDeleteItem="deleteItem"
        @onLoadMore="loadMore"
        @onEditItem="editItem"
      />
    </div>

    <!-- popup -->
    <transaction-popup
      v-if="showExpsensePopup"
      :isEdit="true"
      :defaltValues="obj"
      @onClose="showExpsensePopup = false"
    />
  </section>
</template>

<script>
import FinancesTable from '@/components/FinancesTable.vue';
import AddIcon from '@/components/Icons/AddIcon.vue';
import Filters from '@/components/Filters.vue';
import TransactionPopup from '@/components/TransactionPopup.vue';
export default {
  name: 'Expense',
  components: {
    FinancesTable,
    AddIcon,
    Filters,
    TransactionPopup,
  },
  data() {
    return {
      showExpsensePopup: false,
      selectedExpense: {},
      obj: {
        category: 'bsns',
        amount: 340,
        comment: 'Comment',
      },
    };
  },
  methods: {
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
      this.showExpsensePopup = true;
      console.log(id);
    },
  },
};
</script>

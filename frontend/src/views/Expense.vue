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
      <filters @onApplyFilters="applyFilters" />
    </div>
    <!-- table -->
    <div class="pt-4 w-full">
      <finances-table :isIncome="false" @onDeleteItem="deleteItem" @onLoadMore="loadMore" />
    </div>
  </section>
</template>

<script>
import FinancesTable from '@/components/FinancesTable.vue';
import AddIcon from '@/components/Icons/AddIcon.vue';
import Filters from '@/components/Filters.vue';
export default {
  name: 'Expense',
  components: {
    FinancesTable,
    AddIcon,
    Filters,
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
  },
};
</script>

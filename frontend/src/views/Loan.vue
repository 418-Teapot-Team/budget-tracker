<template>
  <section>
    <!-- caption and filters -->
    <div class="flex flex-row justify-between items-center">
      <div class="flex flex-row justify-start items-center gap-4">
        <span class="text-black text-4xl font-bold">Loans</span>
        <div class="h-8 w-8 cursor-pointer">
          <add-icon />
        </div>
      </div>
      <filters @onApplyFilters="applyFilters" :withCategory="false" />
    </div>
    <!-- table -->
    <div class="pt-4 w-full">
      <div class="w-full flex flex-col justify-start gap-4 overflow-y-auto pb-8 loan-list">
        <loan-card />
        <loan-card />
        <loan-card />
        <loan-card />
        <loan-card />
        <loan-card />
        <div class="w-full flex h-10 flex-row justify-center items-center">
          <div class="w-36 h-10">
            <app-button text="Load more" @click="loadMore" class="h-full" />
          </div>
        </div>
      </div>
    </div>

    <!-- popup -->
  </section>
</template>

<script>
import AddIcon from '@/components/Icons/AddIcon.vue';
import Filters from '@/components/Filters.vue';
import LoanCard from '@/components/LoanCard.vue';
export default {
  name: 'Loan',
  components: {
    AddIcon,
    Filters,
    LoanCard,
  },
  data() {
    return {
      showDepositPopup: false,
      selectedDeposit: {},
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

<style scoped>
.loan-list {
  max-height: calc(100vh - 144px);
  padding-bottom: 20px;
}
.loan-list::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}

.loan-list::-webkit-scrollbar-track {
  border-radius: 100vh;
  background: #f8f9fe;
}

.loan-list::-webkit-scrollbar-thumb {
  background: #1c1c21;
  border-radius: 100vh;
  border: 3px solid #f8f9fe;
}

.loan-list::-webkit-scrollbar-thumb:hover {
  background: #1c1c21;
}
</style>

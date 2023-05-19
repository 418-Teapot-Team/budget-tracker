<template>
  <section>
    <!-- caption and filters -->
    <div class="flex flex-row justify-between items-end">
      <div class="flex flex-row justify-start items-end gap-4">
        <div class="flex flex-col gap-2">
          <span class="text-black text-4xl font-bold">Desposits</span>
          <span class="text-black text-3xl font-bold"
            >Total: <span class="font-normal">10432 / 71343</span> UAH</span
          >
        </div>
        <div class="h-8 w-8 cursor-pointer mb-1" @click="showDepositPopup = true">
          <add-icon />
        </div>
      </div>
      <filters @onApplyFilters="applyFilters" :withCategory="false" />
    </div>
    <!-- table -->
    <div class="pt-4 w-full">
      <div class="w-full flex flex-col justify-start gap-4 overflow-y-auto pb-8 deposit-list">
        <deposit-card />
        <deposit-card />
        <deposit-card />
        <deposit-card />
        <deposit-card />
        <deposit-card />
        <div class="w-full flex h-10 flex-row justify-center items-center">
          <div class="w-36 h-10">
            <app-button text="Load more" @click="loadMore" class="h-full" />
          </div>
        </div>
      </div>
    </div>

    <!-- popup -->
    <banking-popup v-if="showDepositPopup" @onClose="showDepositPopup = false" />
  </section>
</template>

<script>
import AddIcon from '@/components/Icons/AddIcon.vue';
import Filters from '@/components/Filters.vue';
import DepositCard from '@/components/DepositCard.vue';
import BankingPopup from '@/components/BankingPopup.vue';
export default {
  name: 'Deposit',
  components: {
    AddIcon,
    Filters,
    DepositCard,
    BankingPopup,
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
.deposit-list {
  max-height: calc(100vh - 192px);
  padding-bottom: 20px;
}
.deposit-list::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}

.deposit-list::-webkit-scrollbar-track {
  border-radius: 100vh;
  background: #f8f9fe;
}

.deposit-list::-webkit-scrollbar-thumb {
  background: #1c1c21;
  border-radius: 100vh;
  border: 3px solid #f8f9fe;
}

.deposit-list::-webkit-scrollbar-thumb:hover {
  background: #1c1c21;
}
</style>

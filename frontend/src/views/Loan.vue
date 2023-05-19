<template>
  <section>
    <!-- caption and filters -->
    <div class="flex flex-row justify-between items-center">
      <div class="flex flex-row justify-start items-center gap-4">
        <span class="text-black text-4xl font-bold">Loans</span>
        <div class="h-8 w-8 cursor-pointer" @click="showCerditPopup = true">
          <add-icon />
        </div>
      </div>
      <filters @onApplyFilters="applyFilters" :withCategory="false" />
    </div>
    <!-- table -->
    <div class="pt-4 w-full">
      <div class="w-full flex flex-col justify-start gap-4 overflow-y-auto pb-8 loan-list">
        <loan-card
          v-for="item in loans"
          :key="item.id"
          :data="item"
          @onDeleteLoan="deleteItem"
          @onEditLoan="editItem"
        />

        <div class="w-full flex h-10 flex-row justify-center items-center">
          <div class="w-36 h-10">
            <app-button text="Load more" @click="loadMore" class="h-full" />
          </div>
        </div>
      </div>
    </div>

    <!-- popup -->
    <banking-popup
      v-if="showCerditPopup"
      @onAdd="saveItem"
      @onEdit="updateItem"
      @onClose="closePopup"
      :isEdit="isEdit"
      :defaltValues="itemToEdit"
    />
  </section>
</template>

<script>
import AddIcon from '@/components/Icons/AddIcon.vue';
import Filters from '@/components/Filters.vue';
import LoanCard from '@/components/LoanCard.vue';
import BankingPopup from '@/components/BankingPopup.vue';
import useAccountsStore from '@/stores/accounts';
import { mapState, mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();
export default {
  name: 'Loan',
  components: {
    AddIcon,
    Filters,
    LoanCard,
    BankingPopup,
  },
  data() {
    return {
      showCerditPopup: false,
      isEdit: false,
      itemToEdit: {},
    };
  },
  computed: {
    ...mapState(useAccountsStore, ['loans']),
  },
  methods: {
    ...mapActions(useAccountsStore, [
      'getAccounts',
      'createAccount',
      'editAccount',
      'deleteAccount',
    ]),
    async saveItem(values) {
      try {
        const payload = { ...values, type: 'credit' };
        await this.createAccount(payload);
        await this.getAccounts({ type: 'credit' });
        toast.success('New loan added!');
        this.showCerditPopup = false;
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async updateItem(values) {
      try {
        const payload = { ...values, type: 'credit' };
        await this.editAccount(payload);
        await this.getAccounts({ type: 'credit' });
        toast.success('Loan changed!');
        this.showCerditPopup = false;
      } catch (e) {
        toast.error(e?.message);
      }
    },
    loadMore() {
      console.log('load more');
    },
    async deleteItem(id) {
      try {
        await this.deleteAccount({ id });
        await this.getAccounts({ type: 'deposit' });
      } catch (e) {
        toast.error(e?.message);
      }
    },
    applyFilters({ category, filter }) {
      console.log(category, filter);
    },
    editItem(id) {
      this.itemToEdit = this.loans?.find((item) => item.id === id);
      this.isEdit = true;
      this.showCerditPopup = true;
    },
    closePopup() {
      this.showCerditPopup = false;
      this.itemToEdit = {};
    },
    async getInitialData() {
      try {
        await this.getAccounts({ type: 'credit' });
        console.log(this.loans);
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

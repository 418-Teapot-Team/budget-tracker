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
      <div
        class="w-full flex flex-col justify-start gap-4 overflow-y-auto pb-8 deposit-list"
        v-if="deposits?.length"
      >
        <deposit-card
          v-for="item in deposits"
          :key="item.id"
          :data="item"
          @onDeleteDeposit="deleteItem"
          @onEditDeposit="editItem"
        />
        <div class="w-full flex h-10 flex-row justify-center items-center">
          <div class="w-36 h-10">
            <app-button text="Load more" @click="loadMore" class="h-full" />
          </div>
        </div>
      </div>
      <empty-list v-else class="mt-10" />
    </div>

    <!-- popup -->
    <banking-popup
      v-if="showDepositPopup"
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
import DepositCard from '@/components/DepositCard.vue';
import BankingPopup from '@/components/BankingPopup.vue';
import EmptyList from '@/components/EmptyList.vue';
import useAccountsStore from '@/stores/accounts';
import { mapState, mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'Deposit',
  components: {
    AddIcon,
    Filters,
    DepositCard,
    BankingPopup,
    EmptyList,
  },
  data() {
    return {
      showDepositPopup: false,
      isEdit: false,
      itemToEdit: {},
    };
  },
  computed: {
    ...mapState(useAccountsStore, ['deposits']),
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
        const payload = { ...values, type: 'deposit' };
        await this.createAccount(payload);
        await this.getAccounts({ type: 'deposit' });
        toast.success('New deposit added!');
        this.showDepositPopup = false;
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async updateItem(values) {
      try {
        const payload = { ...values, type: 'deposit' };
        await this.editAccount(payload);
        await this.getAccounts({ type: 'deposit' });
        toast.success('Deposit changed!');
        this.showDepositPopup = false;
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
      this.itemToEdit = this.deposits?.find((item) => item.id === id);
      this.isEdit = true;
      this.showDepositPopup = true;
    },
    closePopup() {
      this.showDepositPopup = false;
      this.itemToEdit = {};
    },
    async getInitialData() {
      try {
        await this.getAccounts({ type: 'deposit' });
        console.log(this.deposits);
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

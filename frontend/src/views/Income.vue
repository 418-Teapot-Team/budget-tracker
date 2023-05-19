<template>
  <section>
    <!-- caption and filters -->
    <div class="flex flex-row justify-between items-center">
      <div class="flex flex-row justify-start items-center gap-4">
        <span class="text-black text-4xl font-bold">Income</span>
        <div class="h-8 w-8 cursor-pointer" @click="showIncomePopup = true">
          <add-icon />
        </div>
      </div>
      <filters @onApplyFilters="applyFilters" :withCategory="true" :categories="incomesCtegories" />
    </div>

    <!-- table -->
    <div class="pt-4 w-full">
      <finances-table
        v-if="incomes?.length"
        :data="incomes"
        :isIncome="true"
        @onDeleteItem="deleteItem"
        @onLoadMore="loadMore"
        @onEditItem="editItem"
      />
      <empty-list v-else class="mt-10" />
    </div>

    <!-- popup -->
    <transaction-popup
      v-if="showIncomePopup"
      :isEdit="isEdit"
      @onClose="closePopup"
      :categories="incomesCtegories"
      :defaltValues="itemToEdit"
      @onAdd="saveItem"
      @onEdit="updateItem"
    />
  </section>
</template>

<script>
import FinancesTable from '@/components/FinancesTable.vue';
import EmptyList from '@/components/EmptyList.vue';
import AddIcon from '@/components/Icons/AddIcon.vue';
import Filters from '@/components/Filters.vue';
import TransactionPopup from '@/components/TransactionPopup.vue';
import useFiltersStore from '@/stores/filters';
import useTransitionStore from '@/stores/transactions';
import { mapActions, mapState } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'Income',
  components: {
    FinancesTable,
    TransactionPopup,
    AddIcon,
    Filters,
    EmptyList,
  },
  data() {
    return {
      showIncomePopup: false,
      itemToEdit: {},
      isEdit: false,
    };
  },
  computed: {
    ...mapState(useTransitionStore, ['incomesCtegories', 'incomes']),
    ...mapState(useFiltersStore, ['filters']),
  },
  methods: {
    ...mapActions(useTransitionStore, [
      'getTransactions',
      'getCategories',
      'editTransaction',
      'createTransaction',
      'deleteTransaction',
    ]),
    loadMore() {
      console.log('load more');
    },
    async deleteItem(id) {
      try {
        await this.deleteTransaction({ id });
        await this.getTransactions({ type: 'income' });
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async saveItem(values) {
      try {
        const payload = { ...values, type: 'income' };
        await this.createTransaction(payload);
        await this.getTransactions({ type: 'income' });
        toast.success('New income added!');
        this.showIncomePopup = false;
      } catch (e) {
        toast.error(e?.message);
      }
    },
    applyFilters({ category, filter }) {
      const categoryToUrl = this.incomesCtegories?.find((item) => category === item.id);
      const filterToUrl = this.filters.find((item) => filter === item.value.id);

      this.$router
        .push({
          query: {
            category: categoryToUrl ? categoryToUrl?.category : 'All',
            categoryId: categoryToUrl ? categoryToUrl.id : 'all',
            orderBy: filterToUrl ? filterToUrl?.value?.orderBy : 'default',
            reverse: filterToUrl ? filterToUrl?.value?.reverse : false,
            filterId: filterToUrl ? filterToUrl?.value?.id : 'default',
          },
        })
        .then(() => this.getInitialData());
    },
    editItem(id) {
      this.itemToEdit = this.incomes?.find((item) => item.id === id);
      this.isEdit = true;
      this.showIncomePopup = true;
    },
    async updateItem(values) {
      try {
        const payload = { ...values, type: 'income' };
        await this.editTransaction(payload);
        await this.getTransactions({ type: 'income' });
        toast.success('Income changed!');
        this.showIncomePopup = false;
      } catch (e) {
        toast.error(e?.message);
      }
    },
    closePopup() {
      this.showIncomePopup = false;
      this.itemToEdit = {};
    },
    async getInitialData() {
      const orderBy =
        this.$route.query?.orderBy && this.$route.query?.orderBy !== 'default'
          ? this.$route.query?.orderBy
          : '';
      const reverse = this.$route.query?.reverse === 'true' ? true : false;
      const category =
        this.$route.query?.categoryId && this.$route.query?.categoryId !== 'all'
          ? this.$route.query?.categoryId
          : '';
      try {
        await this.getCategories();
        await this.getTransactions({ type: 'income', orderBy, reverse, category });
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

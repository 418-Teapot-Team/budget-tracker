<template>
  <section>
    <!-- caption and filters -->
    <div class="flex flex-row justify-between items-center">
      <div class="flex flex-row justify-start items-center gap-4">
        <span class="text-black text-4xl font-bold">Expense</span>
        <div class="h-8 w-8 cursor-pointer" @click="showExpensePopup = true">
          <add-icon />
        </div>
      </div>
      <filters
        @onApplyFilters="applyFilters"
        :withCategory="true"
        :categories="expensesCategories"
      />
    </div>
    <!-- table -->
    <div class="pt-4 w-full">
      <finances-table
        v-if="expenses?.length"
        :data="expenses"
        :isIncome="false"
        @onDeleteItem="deleteItem"
        @onLoadMore="loadMore"
        @onEditItem="editItem"
      />
      <empty-list v-else class="mt-10" />
    </div>

    <!-- popup -->
    <transaction-popup
      v-if="showExpensePopup"
      :isEdit="isEdit"
      @onClose="closePopup"
      :defaltValues="itemToEdit"
      :categories="expensesCategories"
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
import useTransitionStore from '@/stores/transactions';
import useFiltersStore from '@/stores/filters';
import { mapActions, mapState } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'Expense',
  components: {
    FinancesTable,
    AddIcon,
    Filters,
    TransactionPopup,
    EmptyList,
  },
  data() {
    return {
      showExpensePopup: false,
      itemToEdit: {},
      isEdit: false,
    };
  },

  loadMore() {
    console.log('load more');
  },
  computed: {
    ...mapState(useTransitionStore, ['expensesCategories', 'expenses']),
    ...mapState(useFiltersStore, ['filters']),
  },
  methods: {
    ...mapActions(useTransitionStore, [
      'getTransactions',
      'getCategories',
      'createTransaction',
      'editTransaction',
      'deleteTransaction',
    ]),
    loadMore() {
      console.log('load more');
    },
    async deleteItem(id) {
      try {
        await this.deleteTransaction({ id });
        await this.getTransactions({ type: 'expenses' });
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async updateItem(values) {
      try {
        const payload = { ...values, type: 'expenses' };
        await this.editTransaction(payload);
        await this.getTransactions({ type: 'expenses' });
        toast.success('Expense changed!');
        this.showExpensePopup = false;
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async saveItem(values) {
      try {
        const payload = { ...values, type: 'expenses' };
        await this.createTransaction(payload);
        await this.getTransactions({ type: 'expenses' });
        toast.success('New expense added!');
        this.showExpensePopup = false;
      } catch (e) {
        toast.error(e?.message);
      }
    },
    applyFilters({ category, filter }) {
      const categoryToUrl = this.expensesCategories?.find((item) => category === item.id);
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
      this.itemToEdit = this.expenses?.find((item) => item.id === id);
      this.isEdit = true;
      this.showExpensePopup = true;
    },
    closePopup() {
      this.showExpensePopup = false;
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
        this.getCategories();
        this.getTransactions({ type: 'expenses', orderBy, reverse, category });
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

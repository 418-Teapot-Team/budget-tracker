import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('transactions', {
  state: () => ({
    incomes: [],
    expenses: [],
    incomesCtegories: [],
    expensesCategories: [],
  }),

  actions: {
    async getTransactions({ type, orderBy = '', reverse = false, category = '' }) {
      const { data } = await HttpClient.get(
        `api/lists/${type}?orderBy=${orderBy}&reverse=${reverse}&category=${category}`
      );
      if (type === 'income') {
        this.incomes = data.result;
      } else {
        this.expenses = data.result;
      }
    },
    async getCategories() {
      const { data } = await HttpClient.get(`api/get-categories`);
      this.incomesCtegories = [];
      this.expensesCategories = [];
      data?.result?.forEach((item) => {
        if (item.type === 'income') {
          this.incomesCtegories.push(item);
        } else {
          this.expensesCategories.push(item);
        }
      });
    },
    async createTransaction(payload) {
      await HttpClient.post('api/lists/create', {
        type: payload.type,
        category: payload.category,
        amount: Number(payload.amount),
        comment: payload.comment,
      });
    },
    async editTransaction(payload) {
      await HttpClient.put('api/lists/update', {
        id: payload.id,
        category: payload.category,
        amount: Number(payload.amount),
        comment: payload.comment,
        type: payload.type,
        createdAt: payload.createdAt,
      });
    },
    async deleteTransaction(payload) {
      await HttpClient.delete('api/lists/delete', {
        data: { id: payload.id },
      });
    },
  },
});

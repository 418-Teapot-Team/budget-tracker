import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('transactions', {
  state: () => ({
    incomes: [],
    expenses: [],
    categories: [],
  }),

  actions: {
    async getTransactions({ type }) {
      const { data } = await HttpClient.get(`api/lists/${type}`);
      console.log(data);
      if (type === 'incomes') {
        this.incomes = data;
      } else {
        this.expenses = data;
      }
    },
    async getCategories() {
      const { data } = await HttpClient.get(`api/categories `);
      console.log(data);
      this.categories = data;
    },
  },
});

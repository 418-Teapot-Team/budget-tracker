import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('dashboard', {
  state: () => ({
    expensesStats: [],
    expensesCategories: [],
    total: 0,
    incomeStats: [],
    incomeCategories: [],
  }),

  actions: {
    async getTopCaregories({ type, months = 1 }) {
      const { data } = await HttpClient.get(
        `api/lists/get-top-categories?type=${type}&monts=${months}`
      );
      console.log(data);
      if (type === 'income') {
        this.incomeCategories = data.result;
      } else {
        this.expensesCategories = data.result;
      }
    },
    async getTotalSum({ type, months = 1 }) {
      const { data } = await HttpClient.get(
        `api/lists/get-total-amount?type=${type}&months=${months}`
      );
      this.total = data.result;
    },
    async getStats() {
      const { data } = await HttpClient.get('api/lists/get-saving-stats');
      this.savingStats = data.result;
    },
  },
});

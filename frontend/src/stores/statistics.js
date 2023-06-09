import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('statistics', {
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
        `api/lists/get-top-categories?type=${type}&months=${months}`
      );

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
    async getStats({ type, months = 1 }) {
      const { data } = await HttpClient.get(`api/lists/get-stats?type=${type}&months=${months}`);

      if (type === 'income') {
        this.incomeStats = data.result;
      } else {
        this.expensesStats = data.result;
      }
    },
  },
});
